package admin

import (
	"forum-service/api/validator/admin"
	"forum-service/common/constant/auth"
	"forum-service/common/serializer/datetime"
	"forum-service/common/utils"
	"forum-service/framework/response"
	"forum-service/model/dto"
	"forum-service/service"
	"strconv"
	"time"

	"gitee.com/hanshuangjianke/go-excel/excel"
	"github.com/gin-gonic/gin"
)

type RoleController struct {
	roleService *service.RoleService
	deptService *service.DeptService
	userService *service.UserService
}

func NewRoleController() *RoleController {
	return &RoleController{
		roleService: &service.RoleService{},
		deptService: &service.DeptService{},
		userService: &service.UserService{},
	}
}

// List 角色列表
func (c *RoleController) List(ctx *gin.Context) {
	var param dto.RoleListRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	roles, total := c.roleService.List(param, true)
	response.Success(ctx).SetPageData(roles, total).Json()
}

// Get 角色详情
func (c *RoleController) Get(ctx *gin.Context) {
	roleId, _ := strconv.Atoi(ctx.Param("roleId"))
	role := c.roleService.Get(roleId)
	response.Success(ctx).SetData("data", role).Json()
}

// Create 新增角色
func (c *RoleController) Create(ctx *gin.Context) {
	var param dto.CreateRoleRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err := admin.CreateRoleValidator(param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	user, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	menuCheckStrictly, deptCheckStrictly := 0, 0
	if param.MenuCheckStrictly {
		menuCheckStrictly = 1
	}
	if param.DeptCheckStrictly {
		deptCheckStrictly = 1
	}
	if err := c.roleService.Create(dto.SaveRoleRequest{
		RoleName:          param.RoleName,
		RoleKey:           param.RoleKey,
		RoleSort:          param.RoleSort,
		MenuCheckStrictly: &menuCheckStrictly,
		DeptCheckStrictly: &deptCheckStrictly,
		Status:            param.Status,
		CreateBy:          user.(*dto.UserTokenResponse).UserName,
		Remark:            param.Remark,
	}, param.MenuIds); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// Update 更新角色
func (c *RoleController) Update(ctx *gin.Context) {
	var param dto.UpdateRoleRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err := admin.UpdateRoleValidator(param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	menuCheckStrictly, deptCheckStrictly := 0, 0
	if param.MenuCheckStrictly {
		menuCheckStrictly = 1
	}
	if param.DeptCheckStrictly {
		deptCheckStrictly = 1
	}
	user, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	if err := c.roleService.Update(dto.SaveRoleRequest{
		RoleId:            param.RoleId,
		RoleName:          param.RoleName,
		RoleKey:           param.RoleKey,
		RoleSort:          param.RoleSort,
		MenuCheckStrictly: &menuCheckStrictly,
		DeptCheckStrictly: &deptCheckStrictly,
		Status:            param.Status,
		UpdateBy:          user.(*dto.UserTokenResponse).UserName,
		Remark:            param.Remark,
	}, param.MenuIds, nil); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// Delete 删除角色
func (c *RoleController) Delete(ctx *gin.Context) {
	roleIds, err := utils.StringToIntSlice(ctx.Param("roleIds"), ",")
	if err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	user, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	roles := c.roleService.ListByUserId(user.(*dto.UserTokenResponse).UserId)
	for _, role := range roles {
		if err = admin.RemoveRoleValidator(roleIds, role.RoleId, role.RoleName); err != nil {
			response.Error(ctx).SetMsg(err.Error()).Json()
			return
		}
	}
	if err = c.roleService.Delete(roleIds); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// ChangeStatus 修改角色状态
func (c *RoleController) ChangeStatus(ctx *gin.Context) {
	var param dto.UpdateRoleRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err := admin.ChangeRoleStatusValidator(param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	user, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	if err := c.roleService.Update(dto.SaveRoleRequest{
		RoleId:   param.RoleId,
		Status:   param.Status,
		UpdateBy: user.(*dto.UserTokenResponse).UserName,
	}, nil, nil); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// DeptTree 部门树
func (c *RoleController) DeptTree(ctx *gin.Context) {
	roleId, _ := strconv.Atoi(ctx.Param("roleId"))
	roleHasDeptIds := c.deptService.ListIdsByRoleId(roleId)
	depts := c.deptService.Tree()
	tree := c.deptService.RemakeTree(depts, 0)
	response.Success(ctx).SetData("depts", tree).SetData("checkedKeys", roleHasDeptIds).Json()
}

// AssignDataScope 分配数据权限
func (c *RoleController) AssignDataScope(ctx *gin.Context) {
	var param dto.UpdateRoleRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Success(ctx).SetMsg(err.Error()).Json()
		return
	}
	deptCheckStrictly := 0
	if param.DeptCheckStrictly {
		deptCheckStrictly = 1
	}
	user, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	if err := c.roleService.Update(dto.SaveRoleRequest{
		RoleId:            param.RoleId,
		DataScope:         param.DataScope,
		DeptCheckStrictly: &deptCheckStrictly,
		UpdateBy:          user.(*dto.UserTokenResponse).UserName,
	}, nil, param.DeptIds); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// RoleUsersAllocated 查询已分配用户角色列表
func (c *RoleController) RoleUsersAllocated(ctx *gin.Context) {
	var param dto.RoleAuthUserAllocatedListRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	user, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	users, total := c.userService.ListByRoleId(param, user.(*dto.UserTokenResponse).UserId, true)
	response.Success(ctx).SetPageData(users, total).Json()
}

// RoleUsersUnAllocated 查询未分配用户角色列表
func (c *RoleController) RoleUsersUnAllocated(ctx *gin.Context) {
	var param dto.RoleAuthUserAllocatedListRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	user, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	users, total := c.userService.ListByRoleId(param, user.(*dto.UserTokenResponse).UserId, false)
	response.Success(ctx).SetPageData(users, total).Json()
}

// AuthUsers 批量选择用户授权
func (c *RoleController) AuthUsers(ctx *gin.Context) {
	var param dto.RoleAuthUserSelectAllRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	userIds, err := utils.StringToIntSlice(param.UserIds, ",")
	if err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err = c.roleService.AuthUsers(param.RoleId, userIds); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// UnAuthUser 取消授权用户
func (c *RoleController) UnAuthUser(ctx *gin.Context) {
	var param dto.RoleAuthUserCancelRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err := c.roleService.UnAuthUsers(param.RoleId, []int{param.UserId}); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// UnAuthUsers 批量取消授权用户
func (c *RoleController) UnAuthUsers(ctx *gin.Context) {
	var param dto.RoleAuthUserCancelAllRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	userIds, err := utils.StringToIntSlice(param.UserIds, ",")
	if err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err := c.roleService.UnAuthUsers(param.RoleId, userIds); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// Export 数据导出
func (c *RoleController) Export(ctx *gin.Context) {
	var param dto.RoleListRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	list := make([]dto.RoleExportResponse, 0)
	roles, _ := c.roleService.List(param, false)
	for _, role := range roles {
		list = append(list, dto.RoleExportResponse{
			RoleId:    role.RoleId,
			RoleName:  role.RoleName,
			RoleKey:   role.RoleKey,
			RoleSort:  role.RoleSort,
			DataScope: role.DataScope,
			Status:    role.Status,
		})
	}
	file, err := excel.NormalDynamicExport("Sheet1", "", "", false, false, list, nil)
	if err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	excel.DownLoadExcel("role_"+time.Now().Format(datetime.DATETIME_FORMAT2), ctx.Writer, file)
}
