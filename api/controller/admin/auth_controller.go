package admin

import (
	"github.com/hugo8680/goat/api/validator/admin"
	"github.com/hugo8680/goat/framework/response"
	"github.com/hugo8680/goat/model/dto"
	adminService "github.com/hugo8680/goat/service/admin"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *adminService.AuthService
}

// NewAuthController 构造函数
func NewAuthController() *AuthController {
	return &AuthController{
		authService: &adminService.AuthService{},
	}
}

// GetCaptchaImage 获取验证码
func (c *AuthController) GetCaptchaImage(ctx *gin.Context) {
	captchaResponse := c.authService.GetCaptchaImage()
	response.Success(ctx).SetData("uuid", captchaResponse.Uuid).SetData("img", captchaResponse.Img).SetData("captchaEnabled", captchaResponse.CaptchaEnabled).Json()
}

// Register 注册
func (c *AuthController) Register(ctx *gin.Context) {
	var param dto.RegisterRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err := admin.RegisterValidator(param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err := c.authService.Register(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// Login 登录
func (c *AuthController) Login(ctx *gin.Context) {
	var param dto.LoginRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetCode(400).SetMsg(err.Error()).Json()
		return
	}
	if err := admin.LoginValidator(param); err != nil {
		response.Error(ctx).SetCode(400).SetMsg(err.Error()).Json()
		return
	}
	token, err := c.authService.Login(&param, ctx)
	if err != nil {
		response.Error(ctx).SetCode(400).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).SetData("token", token).Json()
}

// GetInfo 获取授权信息
func (c *AuthController) GetInfo(ctx *gin.Context) {
	authInfo := c.authService.GetAuthInfo(ctx)
	response.Success(ctx).SetData("user", authInfo.User).SetData("roles", authInfo.Roles).SetData("permissions", authInfo.Permissions).Json()
}

// GetRouters 获取当前用户的路由
func (c *AuthController) GetRouters(ctx *gin.Context) {
	routers := c.authService.GetRouters(ctx)
	response.Success(ctx).SetData("data", routers).Json()
}

// Logout 退出登录
func (c *AuthController) Logout(ctx *gin.Context) {
	err := c.authService.Logout(ctx)
	if err != nil {
		response.Error(ctx).SetCode(400).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}
