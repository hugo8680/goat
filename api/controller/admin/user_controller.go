package admin

import (
	"github.com/hugo8680/goat/api/validator/admin"
	"github.com/hugo8680/goat/common/constant/auth"
	"github.com/hugo8680/goat/common/password"
	"github.com/hugo8680/goat/common/serializer/datetime"
	"github.com/hugo8680/goat/common/uploader"
	"github.com/hugo8680/goat/common/utils"
	"github.com/hugo8680/goat/framework/config"
	"github.com/hugo8680/goat/framework/response"
	"github.com/hugo8680/goat/model/dto"
	adminService "github.com/hugo8680/goat/service/admin"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"gitee.com/hanshuangjianke/go-excel/excel"
	"github.com/gin-gonic/gin"
	excelize "github.com/xuri/excelize/v2"
)

type UserController struct {
	userService   *adminService.UserService
	deptService   *adminService.DeptService
	roleService   *adminService.RoleService
	postService   *adminService.PostService
	configService *adminService.ConfigService
	setting       *config.Setting
}

func NewUserController() *UserController {
	return &UserController{
		userService:   &adminService.UserService{},
		deptService:   &adminService.DeptService{},
		roleService:   &adminService.RoleService{},
		postService:   &adminService.PostService{},
		configService: &adminService.ConfigService{},
		setting:       config.GetSetting(),
	}
}

// DeptTree 获取部门树
func (c *UserController) DeptTree(ctx *gin.Context) {
	user, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	deptList := c.deptService.TreeByUserId(user.(*dto.UserTokenResponse).UserId)
	tree := c.userService.RemakeTreeByUserId(deptList, 0)
	response.Success(ctx).SetData("data", tree).Json()
}

// List 获取用户列表
func (c *UserController) List(ctx *gin.Context) {
	var param dto.UserListRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	user, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	users, total := c.userService.List(param, user.(*dto.UserTokenResponse).UserId, true)
	for key, user := range users {
		users[key].Dept.DeptName = user.DeptName
		users[key].Dept.Leader = user.Leader
	}
	response.Success(ctx).SetPageData(users, total).Json()
}

// Get 用户详情
func (c *UserController) Get(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("userId"))
	r := response.Success(ctx)
	if userId > 0 {
		user := c.userService.Get(userId)
		user.Admin = user.UserId == 1
		dept := c.deptService.Get(user.DeptId)
		roles := c.roleService.ListByUserId(user.UserId)
		r.SetData("data", dto.AuthUserInfoResponse{
			UserDetailResponse: user,
			Dept:               dept,
			Roles:              roles,
		})
		roleIds := make([]int, 0)
		for _, role := range roles {
			roleIds = append(roleIds, role.RoleId)
		}
		r.SetData("roleIds", roleIds)
		postIds := c.postService.ListIdsByUserId(user.UserId)
		r.SetData("postIds", postIds)
	}
	roles, _ := c.roleService.List(dto.RoleListRequest{}, false)
	if userId != 1 {
		roles = utils.Filter(roles, func(role dto.RoleListResponse) bool {
			return role.RoleId != 1
		})
	}
	r.SetData("roles", roles)
	posts, _ := c.postService.List(dto.PostListRequest{}, false)
	r.SetData("posts", posts)
	r.Json()
}

// Create 新增用户
func (c *UserController) Create(ctx *gin.Context) {
	var param dto.CreateUserRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err := admin.CreateUserValidator(param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	user, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	if err := c.userService.Create(dto.SaveUserRequest{
		DeptId:      param.DeptId,
		UserName:    param.UserName,
		NickName:    param.NickName,
		Email:       param.Email,
		PhoneNumber: param.PhoneNumber,
		Sex:         param.Sex,
		Password:    password.Generate(param.Password),
		Status:      param.Status,
		Remark:      param.Remark,
		CreateBy:    user.(*dto.UserTokenResponse).UserName,
	}, param.RoleIds, param.PostIds); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// Update 更新用户
func (c *UserController) Update(ctx *gin.Context) {
	var param dto.UpdateUserRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err := admin.UpdateUserValidator(param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	user, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	if err := c.userService.Update(dto.SaveUserRequest{
		UserId:      param.UserId,
		DeptId:      param.DeptId,
		NickName:    param.NickName,
		Email:       param.Email,
		PhoneNumber: param.PhoneNumber,
		Sex:         param.Sex,
		Status:      param.Status,
		Remark:      param.Remark,
		UpdateBy:    user.(*dto.UserTokenResponse).UserName,
	}, param.RoleIds, param.PostIds); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// Delete 删除用户
func (c *UserController) Delete(ctx *gin.Context) {
	userIds, err := utils.StringToIntSlice(ctx.Param("userIds"), ",")
	if err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	user, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	if err = admin.RemoveUserValidator(userIds, user.(*dto.UserTokenResponse).UserId); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err = c.userService.Delete(userIds); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// ChangeStatus 更改用户状态
func (c *UserController) ChangeStatus(ctx *gin.Context) {
	var param dto.UpdateUserRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err := admin.ChangeUserStatusValidator(param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	user, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	if err := c.userService.Update(dto.SaveUserRequest{
		UserId:   param.UserId,
		Status:   param.Status,
		UpdateBy: user.(*dto.UserTokenResponse).UserName,
	}, nil, nil); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// ResetPassword 重置用户密码
func (c *UserController) ResetPassword(ctx *gin.Context) {
	var param dto.UpdateUserRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err := admin.ResetUserPwdValidator(param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	user, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	if err := c.userService.Update(dto.SaveUserRequest{
		UserId:   param.UserId,
		Password: password.Generate(param.Password),
		UpdateBy: user.(*dto.UserTokenResponse).UserName,
	}, nil, nil); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// ListRoleByUserId 根据用户编号获取授权角色
func (c *UserController) ListRoleByUserId(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("userId"))
	r := response.Success(ctx)
	var userHasRoleIds []int
	if userId > 0 {
		user := c.userService.Get(userId)
		user.Admin = user.UserId == 1
		dept := c.deptService.Get(user.DeptId)
		roles := c.roleService.ListByUserId(user.UserId)
		for _, role := range roles {
			userHasRoleIds = append(userHasRoleIds, role.RoleId)
		}
		r.SetData("user", dto.AuthUserInfoResponse{
			UserDetailResponse: user,
			Dept:               dept,
			Roles:              roles,
		})
	}
	roles, _ := c.roleService.List(dto.RoleListRequest{}, false)
	if userId != 1 {
		roles = utils.Filter(roles, func(role dto.RoleListResponse) bool {
			return role.RoleId != 1
		})
		// 设置角色选中标识，如果角色在用户所拥有的角色列表中设置标识为true
		for key, role := range roles {
			if utils.Contains(userHasRoleIds, role.RoleId) {
				roles[key].Flag = true
			}
		}
	}
	r.SetData("roles", roles)
	r.Json()
}

// AuthRoles 用户授权角色
func (c *UserController) AuthRoles(ctx *gin.Context) {
	var param dto.AddUserAuthRoleRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	roleIds, err := utils.StringToIntSlice(param.RoleIds, ",")
	if err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err := c.userService.AuthRoles(param.UserId, roleIds); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// DownloadImportTemplate 下载导入用户模板
func (c *UserController) DownloadImportTemplate(ctx *gin.Context) {
	list := make([]dto.UserImportRequest, 0)
	list = append(list, dto.UserImportRequest{
		DeptId:      1,
		UserName:    "example",
		NickName:    "模板",
		Email:       "example@example.com",
		PhoneNumber: "12345678901",
		Sex:         "1",
		Status:      "0",
	})
	file, err := excel.NormalDynamicExport("Sheet1", "", "", false, false, list, nil)
	if err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	excel.DownLoadExcel("user_template_"+time.Now().Format(datetime.DATETIME_FORMAT2), ctx.Writer, file)
}

// Import 导入用户数据
func (c *UserController) Import(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	fileName := c.setting.System.UploadPath + file.Filename
	// 临时保存文件
	err = ctx.SaveUploadedFile(file, fileName)
	if err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			response.Error(ctx).SetMsg(err.Error()).Json()
		}
	}(fileName)
	// 是否更新已经存在的用户数据
	updateSupport, _ := strconv.ParseBool(ctx.Query("updateSupport"))
	excelFile, err := excelize.OpenFile(fileName)
	if err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	list := make([]dto.UserImportRequest, 0)
	if err = excel.ImportExcel(excelFile, &list, 0, 1); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if len(list) <= 0 {
		response.Error(ctx).SetMsg("导入用户数据不能为空").Json()
		return
	}
	var successNum, failNum int
	var failMsg []string
	user, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	authUserName := user.(*dto.UserTokenResponse).UserName
	for _, item := range list {
		user := c.userService.GetByUserName(item.UserName)
		// 插入新用户
		if user.UserId <= 0 {
			if err = admin.ImportUserValidator(dto.CreateUserRequest{
				DeptId:      item.DeptId,
				UserName:    item.UserName,
				NickName:    item.NickName,
				Email:       item.Email,
				PhoneNumber: item.PhoneNumber,
				Sex:         item.Sex,
				Status:      item.Status,
			}); err != nil {
				failNum = failNum + 1
				failMsg = append(failMsg, strconv.Itoa(failNum)+"、账号 "+item.UserName+" 新增失败："+err.Error())
				continue
			}
			if err = c.userService.Create(dto.SaveUserRequest{
				DeptId:      item.DeptId,
				UserName:    item.UserName,
				NickName:    item.NickName,
				Email:       item.Email,
				PhoneNumber: item.PhoneNumber,
				Sex:         item.Sex,
				Password:    password.Generate(c.configService.GetCacheByConfigKey("sys.user.initPassword").ConfigValue),
				Status:      item.Status,
				CreateBy:    authUserName,
			}, nil, nil); err != nil {
				failNum = failNum + 1
				failMsg = append(failMsg, strconv.Itoa(failNum)+"、账号 "+item.UserName+" 新增失败："+err.Error())
				continue
			}
			successNum = successNum + 1
			continue
		} else if updateSupport {
			if err = admin.UpdateUserValidator(dto.UpdateUserRequest{
				UserId:      user.UserId,
				DeptId:      item.DeptId,
				NickName:    item.NickName,
				Email:       item.Email,
				PhoneNumber: item.PhoneNumber,
				Sex:         item.Sex,
				Status:      item.Status,
			}); err != nil {
				failNum = failNum + 1
				failMsg = append(failMsg, strconv.Itoa(failNum)+"、账号 "+item.UserName+" 更新失败："+err.Error())
				continue
			}
			// 更新已经存在的用户
			if err = c.userService.Update(dto.SaveUserRequest{
				UserId:      user.UserId,
				DeptId:      item.DeptId,
				NickName:    item.NickName,
				Email:       item.Email,
				PhoneNumber: item.PhoneNumber,
				Sex:         item.Sex,
				Status:      item.Status,
				UpdateBy:    authUserName,
			}, nil, nil); err != nil {
				failNum = failNum + 1
				failMsg = append(failMsg, strconv.Itoa(failNum)+"、账号 "+item.UserName+" 更新失败："+err.Error())
				continue
			}
			successNum = successNum + 1
			// successMsg = append(successMsg, strconv.Itoa(successNum)+"、账号 "+item.UserName+" 更新成功")
			continue
		} else {
			failNum = failNum + 1
			failMsg = append(failMsg, strconv.Itoa(failNum)+"、账号 "+item.UserName+" 已存在")
		}
	}
	if failNum > 0 {
		response.Error(ctx).SetMsg("导入失败，共 " + strconv.Itoa(failNum) + " 条数据错误，错误如下：" + strings.Join(failMsg, "<br/>")).Json()
		return
	}
	response.Success(ctx).SetMsg("导入成功，共 " + strconv.Itoa(successNum) + " 条数据").Json()
}

// Export 导出用户数据
func (c *UserController) Export(ctx *gin.Context) {
	var param dto.UserListRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	list := make([]dto.UserExportResponse, 0)
	user, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	users, _ := c.userService.List(param, user.(*dto.UserTokenResponse).UserId, false)
	for _, user := range users {
		loginDate := user.LoginDate.Format("2006-01-02 15:04:05")
		if user.LoginDate.IsZero() {
			loginDate = ""
		}
		list = append(list, dto.UserExportResponse{
			UserId:      user.UserId,
			UserName:    user.UserName,
			NickName:    user.NickName,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			Sex:         user.Sex,
			Status:      user.Status,
			LoginIp:     user.LoginIp,
			LoginDate:   loginDate,
			DeptName:    user.DeptName,
			DeptLeader:  user.Leader,
		})
	}
	file, err := excel.NormalDynamicExport("Sheet1", "", "", false, false, list, nil)
	if err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	excel.DownLoadExcel("user_"+time.Now().Format(datetime.DATETIME_FORMAT2), ctx.Writer, file)
}

// GetProfile 个人信息
func (c *UserController) GetProfile(ctx *gin.Context) {
	curUser, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	user := c.userService.Get(curUser.(*dto.UserTokenResponse).UserId)
	user.Admin = user.UserId == 1
	dept := c.deptService.Get(user.DeptId)
	roles := c.roleService.ListByUserId(user.UserId)
	data := dto.AuthUserInfoResponse{
		UserDetailResponse: user,
		Dept:               dept,
		Roles:              roles,
	}
	// 获取角色组
	roleGroup := c.roleService.ListNameByUserId(user.UserId)
	// 获取岗位组
	postGroup := c.postService.ListNamesByUserId(user.UserId)
	response.Success(ctx).SetData("data", data).SetData("roleGroup", strings.Join(roleGroup, ",")).SetData("postGroup", strings.Join(postGroup, ",")).Json()
}

// UpdateProfile 修改个人信息
func (c *UserController) UpdateProfile(ctx *gin.Context) {
	var param dto.UpdateProfileRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err := admin.UpdateProfileValidator(param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	user, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	if err := c.userService.Update(dto.SaveUserRequest{
		UserId:      user.(*dto.UserTokenResponse).UserId,
		NickName:    param.NickName,
		Email:       param.Email,
		PhoneNumber: param.PhoneNumber,
		Sex:         param.Sex,
	}, nil, nil); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// UpdatePassword 修改个人密码
func (c *UserController) UpdatePassword(ctx *gin.Context) {
	var param dto.UserProfileUpdatePwdRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err := admin.UserProfileUpdatePwdValidator(param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	curUser, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	user := c.userService.Get(curUser.(*dto.UserTokenResponse).UserId)
	if !password.Verify(user.Password, param.OldPassword) {
		response.Error(ctx).SetMsg("旧密码输入错误").Json()
		return
	}
	if err := c.userService.Update(dto.SaveUserRequest{
		UserId:   user.UserId,
		Password: password.Generate(param.NewPassword),
	}, nil, nil); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// UpdateAvatar 上传头像
func (c *UserController) UpdateAvatar(ctx *gin.Context) {
	fileHeader, err := ctx.FormFile("avatarfile")
	if err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	fileContent, err := io.ReadAll(file)
	if err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	fileResult, err := uploader.NewUploader(
		uploader.SetLimitType([]string{
			"image/jpeg",
			"image/png",
			"image/svg+xml",
		}),
	).SetFile(&uploader.File{
		FileName:    fileHeader.Filename,
		FileType:    fileHeader.Header.Get("Content-Type"),
		FileHeader:  fileHeader.Header,
		FileContent: fileContent,
	}).Save()
	if err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	imgUrl := "/" + fileResult.UrlPath + fileResult.FileName
	user, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	if err = c.userService.Update(dto.SaveUserRequest{
		UserId: user.(*dto.UserTokenResponse).UserId,
		Avatar: imgUrl,
	}, nil, nil); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).SetData("imgUrl", imgUrl).Json()
}
