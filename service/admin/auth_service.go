package admin

import (
	"errors"
	"github.com/hugo8680/goat/common/constant/redis_key"
	"github.com/hugo8680/goat/common/password"
	"github.com/hugo8680/goat/common/serializer/datetime"
	"github.com/hugo8680/goat/framework/config"
	"github.com/hugo8680/goat/framework/connector"
	"github.com/hugo8680/goat/model/dto"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthService struct {
}

func (s *AuthService) GetCaptchaImage() dto.CaptchaResponse {
	captchaService := NewCaptchaService()
	id, b64s := captchaService.Generate()
	b64s = strings.Replace(b64s, "data:image/png;base64,", "", 1)
	conf := (&ConfigService{}).GetCacheByConfigKey("sys.account.captchaEnabled")
	return dto.CaptchaResponse{
		Uuid:           id,
		Img:            b64s,
		CaptchaEnabled: conf.ConfigValue == "true",
	}
}

func (s *AuthService) Register(param *dto.RegisterRequest) error {
	configService := &ConfigService{}
	userService := &UserService{}
	captchaService := NewCaptchaService()
	if conf := configService.GetCacheByConfigKey("sys.account.registerUser"); conf.ConfigValue != "true" {
		return errors.New("当前系统没有开启注册功能")
	}
	if conf := configService.GetCacheByConfigKey("sys.account.captchaEnabled"); conf.ConfigValue == "true" {
		if !captchaService.Verify(param.Uuid, param.Code) {
			return errors.New("验证码错误")
		}
	}
	if user := userService.GetByUserName(param.Username); user.UserId > 0 {
		return errors.New("注册账号已存在")
	}
	if err := userService.Create(dto.SaveUserRequest{
		UserName: param.Username,
		NickName: param.Username,
		Password: password.Generate(param.Password),
		Status:   "0",
		Remark:   "注册用户",
		CreateBy: "注册用户",
	}, nil, nil); err != nil {
		return err
	}
	return nil
}

func (s *AuthService) Login(param *dto.LoginRequest, ctx *gin.Context) (string, error) {
	configService := &ConfigService{}
	userService := &UserService{}
	tokenService := NewTokenService()
	captchaService := NewCaptchaService()
	setting := config.GetSetting()
	cache := connector.GetCache()
	if conf := configService.GetCacheByConfigKey("sys.account.captchaEnabled"); conf.ConfigValue == "true" {
		if !captchaService.Verify(param.Uuid, param.Code) {
			return "", errors.New("验证码错误")
		}
	}
	user := userService.GetByUserName(param.Username)
	if user.UserId <= 0 || user.Status != "0" {
		return "", errors.New("用户不存在或被禁用")
	}
	// 登陆密码错误次数超过限制，锁定账号10分钟
	count, _ := cache.Get(ctx.Request.Context(), redis_key.LoginPasswordErrorKey+param.Username).Int()
	if count >= setting.Auth.Password.MaxRetryCount {
		return "", errors.New("密码错误次数超过限制，请" + strconv.Itoa(setting.Auth.Password.LockTime) + "分钟后重试")
	}
	if !password.Verify(user.Password, param.Password) {
		// 密码错误次数加1，并设置缓存过期时间为锁定时间
		cache.Set(ctx.Request.Context(), redis_key.LoginPasswordErrorKey+param.Username, count+1, time.Minute*time.Duration(setting.Auth.Password.LockTime))
		return "", errors.New("密码错误")
	}
	// 登录成功，删除错误次数
	cache.Del(ctx.Request.Context(), redis_key.LoginPasswordErrorKey+param.Username)
	token, err := tokenService.Create(&user)
	if err != nil {
		return "", err
	}
	// 更新登录的ip和时间
	err = userService.Update(dto.SaveUserRequest{
		UserId:    user.UserId,
		LoginIP:   ctx.ClientIP(),
		LoginDate: datetime.Datetime{Time: time.Now()},
	}, nil, nil)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *AuthService) GetAuthInfo(ctx *gin.Context) dto.AuthInfoResponse {
	securityService := &SecurityService{}
	userService := &UserService{}
	deptService := &DeptService{}
	roleService := &RoleService{}
	menuService := &MenuService{}
	userId, _ := securityService.GetCurrentUserId(ctx)
	user := userService.Get(userId)
	user.Admin = user.UserId == 1
	dept := deptService.Get(user.DeptId)
	roles := roleService.ListByUserId(user.UserId)
	data := dto.AuthUserInfoResponse{
		UserDetailResponse: user,
		Dept:               dept,
		Roles:              roles,
	}
	roleKeys := roleService.ListKeyByUserId(user.UserId)
	perms := menuService.ListPermsByUserId(user.UserId)
	return dto.AuthInfoResponse{
		User:        data,
		Roles:       roleKeys,
		Permissions: perms,
	}
}

func (s *AuthService) GetRouters(ctx *gin.Context) []dto.MenuMetaTreeResponse {
	securityService := &SecurityService{}
	menuService := &MenuService{}
	userId, _ := securityService.GetCurrentUserId(ctx)
	menus := menuService.GetMCListByUserId(userId)
	tree := menuService.MCListToTree(menus, 0)
	return menuService.BuildRouterMenus(tree)
}

func (s *AuthService) Logout(ctx *gin.Context) error {
	tokenService := NewTokenService()
	return tokenService.Delete(ctx)
}
