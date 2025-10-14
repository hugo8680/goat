package service

import (
	"errors"
	"forum-service/common/constant/menu_key"
	"forum-service/framework/connector"
	"forum-service/model"
	"forum-service/model/dto"
	"strings"
)

type MenuService struct {
}

// Create 新增菜单
func (s *MenuService) Create(param dto.SaveMenuRequest) error {
	if menu := s.GetByMenuName(param.MenuName); menu.MenuId > 0 {
		return errors.New("新增菜单" + param.MenuName + "失败，菜单名称已存在")
	}
	return connector.GetDB().Model(model.SysMenu{}).Create(&model.SysMenu{
		MenuName:  param.MenuName,
		ParentId:  param.ParentId,
		OrderNum:  param.OrderNum,
		Path:      param.Path,
		Component: param.Component,
		Query:     param.Query,
		RouteName: param.RouteName,
		IsFrame:   param.IsFrame,
		IsCache:   param.IsCache,
		MenuType:  param.MenuType,
		Visible:   param.Visible,
		Perms:     param.Perms,
		Icon:      param.Icon,
		Status:    param.Status,
		Remark:    param.Remark,
		CreateBy:  param.CreateBy,
	}).Error
}

// Update 更新菜单
func (s *MenuService) Update(param dto.SaveMenuRequest) error {
	if menu := s.GetByMenuName(param.MenuName); menu.MenuId > 0 && menu.MenuId != param.MenuId {
		return errors.New("修改菜单" + param.MenuName + "失败，菜单名称已存在")
	}
	return connector.GetDB().Model(model.SysMenu{}).Where("menu_id = ?", param.MenuId).Updates(&model.SysMenu{
		MenuName:  param.MenuName,
		ParentId:  param.ParentId,
		OrderNum:  param.OrderNum,
		Path:      param.Path,
		Component: param.Component,
		Query:     param.Query,
		RouteName: param.RouteName,
		IsFrame:   param.IsFrame,
		IsCache:   param.IsCache,
		MenuType:  param.MenuType,
		Visible:   param.Visible,
		Perms:     param.Perms,
		Icon:      param.Icon,
		Status:    param.Status,
		UpdateBy:  param.UpdateBy,
		Remark:    param.Remark,
	}).Error
}

// Delete 删除菜单
func (s *MenuService) Delete(menuId int) error {
	if s.HasChildren(menuId) {
		return errors.New("存在子菜单，不允许删除")
	}
	if s.HasAssigned(menuId) {
		return errors.New("菜单已分配，不允许删除")
	}
	return connector.GetDB().Model(model.SysMenu{}).Where("menu_id = ?", menuId).Delete(&model.SysMenu{}).Error
}

// List 菜单列表
func (s *MenuService) List(param dto.MenuListRequest) []dto.MenuListResponse {
	menus := make([]dto.MenuListResponse, 0)
	query := connector.GetDB().Model(model.SysMenu{}).Order("sys_menu.parent_id, sys_menu.order_num, sys_menu.menu_id")
	if param.MenuName != "" {
		query.Where("menu_name LIKE ?", "%"+param.MenuName+"%")
	}
	if param.Status != "" {
		query.Where("status = ?", param.Status)
	}
	query.Find(&menus)
	return menus
}

// Get 根据菜单id查询菜单
func (s *MenuService) Get(menuId int) dto.MenuDetailResponse {
	var menu dto.MenuDetailResponse
	connector.GetDB().Model(model.SysMenu{}).Where("menu_id = ?", menuId).Last(&menu)
	return menu
}

// GetByMenuName 根据菜单名称查询菜单
func (s *MenuService) GetByMenuName(menuName string) dto.MenuDetailResponse {
	var menu dto.MenuDetailResponse
	connector.GetDB().Model(model.SysMenu{}).Where("menu_name = ?", menuName).Last(&menu)
	return menu
}

// HasChildren 查询是否存在下级菜单
func (s *MenuService) HasChildren(menuId int) bool {
	var count int64
	connector.GetDB().Model(model.SysMenu{}).Where("parent_id = ?", menuId).Count(&count)
	return count > 0
}

// HasAssigned 查询菜单是否已分配到权限
func (s *MenuService) HasAssigned(menuId int) bool {
	var count int64
	connector.GetDB().Model(model.SysRoleMenu{}).Where("menu_id = ?", menuId).Count(&count)
	return count > 0
}

// ListPermsByUserId 根据用户id查询菜单权限perms
func (s *MenuService) ListPermsByUserId(userId int) []string {
	perms := make([]string, 0)
	// 超级管理员拥有所有权限
	if userId == 1 {
		perms = append(perms, "*:*:*")
	} else {
		connector.GetDB().Model(model.SysMenu{}).
			Joins("JOIN sys_role_menu ON sys_menu.menu_id = sys_role_menu.menu_id").
			Joins("JOIN sys_role ON sys_role_menu.role_id = sys_role.role_id").
			Joins("JOIN sys_user_role ON sys_role.role_id = sys_user_role.role_id").
			Where("sys_user_role.user_id = ? AND sys_menu.status = 0", userId).
			Pluck("sys_menu.perms", &perms)
	}
	return perms
}

// ListIdsByRoleId 根据角色id查询拥有的菜单id集合
func (s *MenuService) ListIdsByRoleId(roleId int) []int {
	menuIds := make([]int, 0)
	connector.GetDB().Model(model.SysRoleMenu{}).
		Joins("JOIN sys_menu ON sys_menu.menu_id = sys_role_menu.menu_id").
		Where("sys_menu.status = 0 AND sys_role_menu.role_id = ?", roleId).
		Pluck("sys_menu.menu_id", &menuIds)
	return menuIds
}

// Tree 菜单下拉树列表
func (s *MenuService) Tree() []dto.TreeResponse {
	menus := make([]dto.TreeResponse, 0)
	connector.GetDB().Model(model.SysMenu{}).Order("order_num, menu_id").
		Select("menu_id as id", "menu_name as label", "parent_id").
		Where("status = 0").
		Find(&menus)
	return menus
}

// RemakeTree 菜单下拉列表转树形结构
func (s *MenuService) RemakeTree(menus []dto.TreeResponse, parentId int) []dto.TreeResponse {
	tree := make([]dto.TreeResponse, 0)
	for _, menu := range menus {
		if menu.ParentId == parentId {
			tree = append(tree, dto.TreeResponse{
				Id:       menu.Id,
				Label:    menu.Label,
				ParentId: menu.ParentId,
				Children: s.RemakeTree(menus, menu.Id),
			})
		}
	}
	return tree
}

// GetMCListByUserId 根据用户id查询拥有的菜单权限
//
// （M-目录；C-菜单；F-按钮）
func (s *MenuService) GetMCListByUserId(userId int) []dto.MenuListResponse {
	menus := make([]dto.MenuListResponse, 0)
	query := connector.GetDB().Model(model.SysMenu{}).
		Distinct("sys_menu.*").
		Order("sys_menu.parent_id, sys_menu.order_num").
		Joins("LEFT JOIN sys_role_menu ON sys_menu.menu_id = sys_role_menu.menu_id").
		Joins("LEFT JOIN sys_role ON sys_role_menu.role_id = sys_role.role_id").
		Joins("LEFT JOIN sys_user_role ON sys_role.role_id = sys_user_role.role_id").
		Where("sys_menu.status = 0 AND sys_menu.menu_type IN ?", []string{menu_key.MENU_TYPE_DIRECTORY, menu_key.MENU_TYPE_MENU})
	if userId > 1 {
		query = query.Where("sys_user_role.user_id = ? AND sys_role.status = 0", userId)
	}
	query.Find(&menus)
	return menus
}

// MCListToTree 菜单权限列表转树形结构
func (s *MenuService) MCListToTree(menus []dto.MenuListResponse, parentId int) []dto.MenuListTreeResponse {
	tree := make([]dto.MenuListTreeResponse, 0)
	for _, menu := range menus {
		if menu.ParentId == parentId {
			tree = append(tree, dto.MenuListTreeResponse{
				MenuListResponse: menu,
				Children:         s.MCListToTree(menus, menu.MenuId),
			})
		}
	}
	return tree
}

// BuildRouterMenus 构建前端路由所需要的菜单
func (s *MenuService) BuildRouterMenus(menus []dto.MenuListTreeResponse) []dto.MenuMetaTreeResponse {
	routers := make([]dto.MenuMetaTreeResponse, 0)
	for _, menu := range menus {
		router := dto.MenuMetaTreeResponse{
			Name:      s.GetRouteName(menu),
			Path:      s.GetRoutePath(menu),
			Component: s.GetComponent(menu),
			Hidden:    menu.Visible == "1",
			Meta: dto.MenuMetaResponse{
				Title:   menu.MenuName,
				Icon:    menu.Icon,
				NoCache: menu.IsCache == 1,
			},
		}
		if len(menu.Children) > 0 && menu.MenuType == menu_key.MENU_TYPE_DIRECTORY {
			router.AlwaysShow = true
			router.Redirect = "noRedirect"
			router.Children = s.BuildRouterMenus(menu.Children)
		} else if s.IsMenuFrame(menu) {
			children := dto.MenuMetaTreeResponse{
				Path:      menu.Path,
				Component: menu.Component,
				Name:      s.GetRouteNameOrDefault(menu.RouteName, menu.Path),
				Meta: dto.MenuMetaResponse{
					Title:   menu.MenuName,
					Icon:    menu.Icon,
					NoCache: menu.IsCache == 1,
				},
				Query: menu.Query,
			}
			router.Children = append(router.Children, children)
		} else if menu.ParentId == 0 && s.IsInnerLink(menu) {
			router.Meta = dto.MenuMetaResponse{
				Title: menu.MenuName,
				Icon:  menu.Icon,
			}
			router.Path = "/"
			children := dto.MenuMetaTreeResponse{
				Path:      s.InnerLinkReplacePach(menu.Path),
				Component: menu_key.INNER_LINK_COMPONENT,
				Name:      s.GetRouteNameOrDefault(menu.RouteName, menu.Path),
				Meta: dto.MenuMetaResponse{
					Title: menu.MenuName,
					Icon:  menu.Icon,
					Link:  menu.Path,
				},
			}
			router.Children = append(router.Children, children)
		}
		routers = append(routers, router)
	}
	return routers
}

// GetRouteName 获取路由名称
func (s *MenuService) GetRouteName(menu dto.MenuListTreeResponse) string {
	if s.IsMenuFrame(menu) {
		return ""
	}
	return s.GetRouteNameOrDefault(menu.RouteName, menu.Path)
}

// GetRouteNameOrDefault 获取路由名称，如没有配置路由名称则取路由地址
func (s *MenuService) GetRouteNameOrDefault(name, path string) string {
	if name == "" {
		name = path
	}
	return strings.ToUpper(string(name[0])) + name[1:]
}

// GetRoutePath 获取路由地址
func (s *MenuService) GetRoutePath(menu dto.MenuListTreeResponse) string {
	routePath := menu.Path
	// 内链打开外网方式
	if menu.ParentId != 0 && !s.IsInnerLink(menu) {
		routePath = s.InnerLinkReplacePach(routePath)
	}
	// 非外链并且是一级目录（类型为目录）
	if menu.ParentId == 0 && menu.MenuType == menu_key.MENU_TYPE_DIRECTORY && menu.IsFrame == menu_key.MENU_NO_FRAME {
		routePath = "/" + routePath
	} else if s.IsMenuFrame(menu) {
		// 非外链并且是一级目录（类型为菜单）
		routePath = "/"
	}
	return routePath
}

// GetComponent 获取组件信息
func (s *MenuService) GetComponent(menu dto.MenuListTreeResponse) string {
	component := menu_key.LAYOUT_COMPONENT
	if menu.Component != "" && !s.IsMenuFrame(menu) {
		component = menu.Component
	} else if menu.Component == "" && menu.ParentId != 0 && s.IsInnerLink(menu) {
		component = menu_key.INNER_LINK_COMPONENT
	} else if menu.Component == "" && s.IsParentView(menu) {
		component = menu_key.PARENT_VIEW_COMPONENT
	}
	return component
}

// IsMenuFrame 是否为菜单内部跳转
func (s *MenuService) IsMenuFrame(menu dto.MenuListTreeResponse) bool {
	return menu.ParentId == 0 && menu_key.MENU_TYPE_MENU == menu.MenuType && menu.IsFrame == menu_key.MENU_NO_FRAME
}

// IsInnerLink 是否为内链组件
func (s *MenuService) IsInnerLink(menu dto.MenuListTreeResponse) bool {
	return menu.IsFrame == menu_key.MENU_NO_FRAME && strings.HasPrefix(menu.Path, "http")
}

// IsParentView 是否为parent_view组件
func (s *MenuService) IsParentView(menu dto.MenuListTreeResponse) bool {
	return menu.ParentId != 0 && menu.MenuType == menu_key.MENU_TYPE_DIRECTORY
}

// InnerLinkReplacePach 内链域名特殊字符替换
func (s *MenuService) InnerLinkReplacePach(path string) string {
	// 去掉 http:// 和 https://
	path = strings.ReplaceAll(path, "http://", "")
	path = strings.ReplaceAll(path, "https://", "")
	path = strings.ReplaceAll(path, "www.", "")
	// 将 . 替换为 /
	path = strings.ReplaceAll(path, ".", "/")
	// 将 : 替换为 /
	path = strings.ReplaceAll(path, ":", "/")
	return path
}
