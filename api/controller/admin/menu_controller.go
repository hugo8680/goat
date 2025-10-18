package admin

import (
	"forum-service/api/validator/admin"
	"forum-service/common/constant/auth"
	"forum-service/framework/response"
	"forum-service/model/dto"
	"forum-service/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MenuController struct {
	menuService *service.MenuService
}

func NewMenuController() *MenuController {
	return &MenuController{
		menuService: &service.MenuService{},
	}
}

// List 菜单列表
func (c *MenuController) List(ctx *gin.Context) {
	var param dto.MenuListRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	menus := c.menuService.List(param)
	response.Success(ctx).SetData("data", menus).Json()
}

// Get 菜单详情
func (c *MenuController) Get(ctx *gin.Context) {
	menuId, _ := strconv.Atoi(ctx.Param("menuId"))
	menu := c.menuService.Get(menuId)
	response.Success(ctx).SetData("data", menu).Json()
}

// Tree 获取菜单下拉树列表
func (c *MenuController) Tree(ctx *gin.Context) {
	menus := c.menuService.Tree()
	tree := c.menuService.RemakeTree(menus, 0)
	response.Success(ctx).SetData("data", tree).Json()
}

// RoleMenuTree 加载对应角色菜单列表树
func (c *MenuController) RoleMenuTree(ctx *gin.Context) {
	roleId, _ := strconv.Atoi(ctx.Param("roleId"))
	roleHasMenuIds := c.menuService.ListIdsByRoleId(roleId)
	menus := c.menuService.Tree()
	tree := c.menuService.RemakeTree(menus, 0)
	response.Success(ctx).SetData("menus", tree).SetData("checkedKeys", roleHasMenuIds).Json()
}

// Create 新增菜单
func (c *MenuController) Create(ctx *gin.Context) {
	var param dto.CreateMenuRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err := admin.CreateMenuValidator(param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	user, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	if err := c.menuService.Create(dto.SaveMenuRequest{
		MenuName:  param.MenuName,
		ParentId:  param.ParentId,
		OrderNum:  param.OrderNum,
		Path:      param.Path,
		Component: param.Component,
		Query:     param.Query,
		RouteName: param.RouteName,
		IsFrame:   &param.IsFrame,
		IsCache:   &param.IsCache,
		MenuType:  param.MenuType,
		Visible:   param.Visible,
		Perms:     param.Perms,
		Icon:      param.Icon,
		Status:    param.Status,
		CreateBy:  user.(*dto.UserTokenResponse).UserName,
	}); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// Update 更新菜单
func (c *MenuController) Update(ctx *gin.Context) {
	var param dto.UpdateMenuRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err := admin.UpdateMenuValidator(param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	user, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	if err := c.menuService.Update(dto.SaveMenuRequest{
		MenuId:    param.MenuId,
		MenuName:  param.MenuName,
		ParentId:  param.ParentId,
		OrderNum:  param.OrderNum,
		Path:      param.Path,
		Component: param.Component,
		Query:     param.Query,
		RouteName: param.RouteName,
		IsFrame:   &param.IsFrame,
		IsCache:   &param.IsCache,
		MenuType:  param.MenuType,
		Visible:   param.Visible,
		Perms:     param.Perms,
		Icon:      param.Icon,
		Status:    param.Status,
		UpdateBy:  user.(*dto.UserTokenResponse).UserName,
	}); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// Delete 删除菜单
func (c *MenuController) Delete(ctx *gin.Context) {
	menuId, _ := strconv.Atoi(ctx.Param("menuId"))
	if err := c.menuService.Delete(menuId); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}
