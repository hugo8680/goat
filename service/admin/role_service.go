package admin

import (
	"errors"
	"github.com/hugo8680/goat/framework/connector"
	"github.com/hugo8680/goat/model"
	"github.com/hugo8680/goat/model/dto"
)

type RoleService struct {
}

// Create 新增角色
func (s *RoleService) Create(param dto.SaveRoleRequest, menuIds []int) error {
	db := connector.GetDB()
	if role := s.GetByRoleName(param.RoleName); role.RoleId > 0 {
		return errors.New("新增角色" + param.RoleName + "失败，角色名已存在")
	}
	if role := s.GetByRoleKey(param.RoleKey); role.RoleId > 0 {
		return errors.New("新增角色" + param.RoleName + "失败，权限字符已存在")
	}
	tx := db.Begin()
	role := model.SysRole{
		RoleName:          param.RoleName,
		RoleKey:           param.RoleKey,
		RoleSort:          param.RoleSort,
		MenuCheckStrictly: param.MenuCheckStrictly,
		DeptCheckStrictly: param.DeptCheckStrictly,
		Status:            param.Status,
		CreateBy:          param.CreateBy,
		Remark:            param.Remark,
	}
	if err := tx.Model(model.SysRole{}).Create(&role).Error; err != nil {
		tx.Rollback()
		return err
	}
	if len(menuIds) > 0 {
		for _, menuId := range menuIds {
			if err := tx.Model(model.SysRoleMenu{}).Create(&model.SysRoleMenu{
				RoleId: role.RoleId,
				MenuId: menuId,
			}).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	return tx.Commit().Error
}

// Update 更新角色
func (s *RoleService) Update(param dto.SaveRoleRequest, menuIds, deptIds []int) error {
	db := connector.GetDB()
	if role := s.GetByRoleName(param.RoleName); role.RoleId > 0 && role.RoleId != param.RoleId {
		return errors.New("修改角色" + param.RoleName + "失败，角色名已存在")
	}
	if role := s.GetByRoleKey(param.RoleKey); role.RoleId > 0 && role.RoleId != param.RoleId {
		return errors.New("修改角色" + param.RoleName + "失败，权限字符已存在")
	}
	tx := db.Begin()
	if err := tx.Model(model.SysRole{}).Where("role_id = ?", param.RoleId).Updates(&model.SysRole{
		RoleName:          param.RoleName,
		RoleKey:           param.RoleKey,
		RoleSort:          param.RoleSort,
		DataScope:         param.DataScope,
		MenuCheckStrictly: param.MenuCheckStrictly,
		DeptCheckStrictly: param.DeptCheckStrictly,
		Status:            param.Status,
		UpdateBy:          param.UpdateBy,
		Remark:            param.Remark,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if menuIds != nil {
		if err := tx.Model(model.SysRoleMenu{}).Where("role_id = ?", param.RoleId).Delete(&model.SysRoleMenu{}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	if len(menuIds) > 0 {
		for _, menuId := range menuIds {
			if err := tx.Model(model.SysRoleMenu{}).Create(&model.SysRoleMenu{
				RoleId: param.RoleId,
				MenuId: menuId,
			}).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	if deptIds != nil {
		if err := tx.Model(model.SysRoleDept{}).Where("role_id = ?", param.RoleId).Delete(&model.SysRoleDept{}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	if len(deptIds) > 0 {
		for _, deptId := range deptIds {
			if err := tx.Model(model.SysRoleDept{}).Create(&model.SysRoleDept{
				RoleId: param.RoleId,
				DeptId: deptId,
			}).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	return tx.Commit().Error
}

// Delete 删除角色
func (s *RoleService) Delete(roleIds []int) error {
	db := connector.GetDB()
	tx := db.Begin()
	if err := tx.Model(model.SysRole{}).Where("role_id IN ?", roleIds).Delete(&model.SysRole{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(model.SysRoleMenu{}).Where("role_id IN ?", roleIds).Delete(&model.SysRoleMenu{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(model.SysRoleDept{}).Where("role_id IN ?", roleIds).Delete(&model.SysRoleDept{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// List 获取角色列表
func (s *RoleService) List(param dto.RoleListRequest, isPaging bool) ([]dto.RoleListResponse, int) {
	var count int64
	roles := make([]dto.RoleListResponse, 0)
	query := connector.GetDB().Model(model.SysRole{}).Order("role_sort, role_id")
	if param.RoleName != "" {
		query.Where("role_name LIKE ?", "%"+param.RoleName+"%")
	}
	if param.RoleKey != "" {
		query.Where("role_key LIKE ?", "%"+param.RoleKey+"%")
	}
	if param.Status != "" {
		query.Where("status = ?", param.Status)
	}
	if param.BeginTime != "" && param.EndTime != "" {
		query = query.Where("sys_user.create_time BETWEEN ? AND ?", param.BeginTime, param.EndTime)
	}
	if isPaging {
		query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize)
	}
	query.Find(&roles)
	return roles, int(count)
}

// Get 获取角色详情
func (s *RoleService) Get(roleId int) dto.RoleDetailResponse {
	var role dto.RoleDetailResponse
	connector.GetDB().Model(model.SysRole{}).Where("role_id = ?", roleId).Last(&role)
	return role
}

// AuthUsers 批量授权用户
func (s *RoleService) AuthUsers(roleId int, userIds []int) error {
	db := connector.GetDB()
	tx := db.Begin()
	for _, userId := range userIds {
		if err := tx.Model(model.SysUserRole{}).Create(&model.SysUserRole{
			UserId: userId,
			RoleId: roleId,
		}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

// UnAuthUsers 批量取消授权用户
func (s *RoleService) UnAuthUsers(roleId int, userIds []int) error {
	return connector.GetDB().Model(model.SysUserRole{}).Where("role_id = ? AND user_id in ?", roleId, userIds).Delete(&model.SysUserRole{}).Error
}

// ListByUserId 根据用户id查询角色列表
func (s *RoleService) ListByUserId(userId int) []dto.RoleListResponse {
	roles := make([]dto.RoleListResponse, 0)
	connector.GetDB().Model(model.SysRole{}).Select("sys_role.*").
		Joins("JOIN sys_user_role ON sys_role.role_id = sys_user_role.role_id").
		Where("sys_user_role.user_id = ? AND sys_role.status = 0", userId).
		Find(&roles)
	return roles
}

// ListKeyByUserId 根据用户id查询角色key
func (s *RoleService) ListKeyByUserId(userId int) []string {
	roleKeys := make([]string, 0)
	connector.GetDB().Model(model.SysRole{}).
		Joins("JOIN sys_user_role ON sys_user_role.role_id = sys_role.role_id").
		Where("sys_user_role.user_id = ? AND sys_role.status = 0", userId).
		Pluck("sys_role.role_key", &roleKeys)
	return roleKeys
}

// ListNameByUserId 根据用户id查询角色名
func (s *RoleService) ListNameByUserId(userId int) []string {
	var roleNames []string
	connector.GetDB().Model(model.SysRole{}).
		Joins("JOIN sys_user_role ON sys_user_role.role_id = sys_role.role_id").
		Where("sys_user_role.user_id = ? AND sys_role.status = 0", userId).
		Pluck("sys_role.role_name", &roleNames)
	return roleNames
}

// GetByRoleName 根据角色名称查询角色
func (s *RoleService) GetByRoleName(roleName string) dto.RoleDetailResponse {
	var role dto.RoleDetailResponse
	connector.GetDB().Model(model.SysRole{}).Where("role_name = ?", roleName).Last(&role)
	return role
}

// GetByRoleKey 根据角色key称查询角色
func (s *RoleService) GetByRoleKey(roleKey string) dto.RoleDetailResponse {
	var role dto.RoleDetailResponse
	connector.GetDB().Model(model.SysRole{}).Where("role_key = ?", roleKey).Last(&role)
	return role
}
