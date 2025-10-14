package service

import (
	"errors"
	"forum-service/framework/connector"
	"forum-service/model"
	"forum-service/model/dto"
)

type UserService struct {
}

// Create 新增用户
func (s *UserService) Create(param dto.SaveUserRequest, roleIds, postIds []int) error {
	if user := s.GetByUserName(param.UserName); user.UserId > 0 {
		return errors.New("新增用户" + param.UserName + "失败，用户名已存在")
	}
	if param.Email != "" {
		return errors.New("新增用户" + param.UserName + "失败，邮箱已存在")
	}
	if param.PhoneNumber != "" {
		if user := s.GetByPhoneNumber(param.PhoneNumber); user.UserId > 0 {
			return errors.New("新增用户" + param.UserName + "失败，手机号已存在")
		}
	}
	db := connector.GetDB()
	tx := db.Begin()
	user := model.SysUser{
		DeptId:      param.DeptId,
		UserName:    param.UserName,
		NickName:    param.NickName,
		UserType:    param.UserType,
		Email:       param.Email,
		PhoneNumber: param.PhoneNumber,
		Sex:         param.Sex,
		Avatar:      param.Avatar,
		Password:    param.Password,
		LoginIP:     param.LoginIP,
		LoginDate:   param.LoginDate,
		Status:      param.Status,
		CreateBy:    param.CreateBy,
		Remark:      param.Remark,
	}
	if err := tx.Model(model.SysUser{}).Create(&user).Error; err != nil {
		tx.Rollback()
		return err
	}
	if len(roleIds) > 0 {
		for _, roleId := range roleIds {
			if err := tx.Model(model.SysUserRole{}).Create(&model.SysUserRole{
				UserId: user.UserId,
				RoleId: roleId,
			}).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	if len(postIds) > 0 {
		for _, postId := range postIds {
			if err := tx.Model(model.SysUserPost{}).Create(&model.SysUserPost{
				UserId: user.UserId,
				PostId: postId,
			}).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	return tx.Commit().Error
}

// Update 更新用户
func (s *UserService) Update(param dto.SaveUserRequest, roleIds, postIds []int) error {
	if param.Email != "" {
		if user := s.GetByEmail(param.Email); user.UserId > 0 && user.UserId != param.UserId {
			return errors.New("修改用户" + param.UserName + "失败，邮箱已存在")
		}
	}
	if param.PhoneNumber != "" {
		if user := s.GetByPhoneNumber(param.PhoneNumber); user.UserId > 0 && user.UserId != param.UserId {
			return errors.New("修改用户" + param.UserName + "失败，手机号已存在")
		}
	}
	db := connector.GetDB()
	tx := db.Begin()
	if err := tx.Model(model.SysUser{}).Where("user_id = ?", param.UserId).Updates(&model.SysUser{
		DeptId:      param.DeptId,
		NickName:    param.NickName,
		UserType:    param.UserType,
		Email:       param.Email,
		PhoneNumber: param.PhoneNumber,
		Sex:         param.Sex,
		Avatar:      param.Avatar,
		Password:    param.Password,
		LoginIP:     param.LoginIP,
		LoginDate:   param.LoginDate,
		Status:      param.Status,
		UpdateBy:    param.UpdateBy,
		Remark:      param.Remark,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if roleIds != nil {
		if err := tx.Model(model.SysUserRole{}).Where("user_id = ?", param.UserId).Delete(&model.SysUserRole{}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	if len(roleIds) > 0 {
		for _, roleId := range roleIds {
			if err := tx.Model(model.SysUserRole{}).Create(&model.SysUserRole{
				UserId: param.UserId,
				RoleId: roleId,
			}).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	if postIds != nil {
		if err := tx.Model(model.SysUserPost{}).Where("user_id = ?", param.UserId).Delete(&model.SysUserPost{}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	if len(postIds) > 0 {
		for _, postId := range postIds {
			if err := tx.Model(model.SysUserPost{}).Create(&model.SysUserPost{
				UserId: param.UserId,
				PostId: postId,
			}).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	return tx.Commit().Error
}

// Delete 删除用户
func (s *UserService) Delete(userIds []int) error {
	tx := connector.GetDB().Begin()
	if err := tx.Model(model.SysUser{}).Where("user_id IN ?", userIds).Delete(&model.SysUser{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(model.SysUserRole{}).Where("user_id IN ?", userIds).Delete(&model.SysUserRole{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(model.SysUserPost{}).Where("user_id IN ?", userIds).Delete(&model.SysUserPost{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// AuthRoles 用户授权角色
func (s *UserService) AuthRoles(userId int, roleIds []int) error {
	db := connector.GetDB()
	tx := db.Begin()
	// 清理用户角色
	if err := tx.Model(model.SysUserRole{}).Where("user_id = ?", userId).Delete(&model.SysUserRole{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 重新插入分配的角色
	if len(roleIds) > 0 {
		for _, roleId := range roleIds {
			if err := tx.Model(model.SysUserRole{}).Create(&model.SysUserRole{
				UserId: userId,
				RoleId: roleId,
			}).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	return tx.Commit().Error
}

// List 获取用户列表
func (s *UserService) List(param dto.UserListRequest, userId int, isPaging bool) ([]dto.UserListResponse, int) {
	var count int64
	users := make([]dto.UserListResponse, 0)
	query := connector.GetDB().Model(model.SysUser{}).
		Select("sys_user.*", "sys_dept.dept_name", "sys_dept.leader").
		Joins("LEFT JOIN sys_dept ON sys_user.dept_id = sys_dept.dept_id").
		Scopes((&DataScopeService{}).GetDataScope("sys_dept", userId, "sys_user"))
	if param.UserName != "" {
		query = query.Where("sys_user.user_name LIKE ?", "%"+param.UserName+"%")
	}
	if param.PhoneNumber != "" {
		query = query.Where("sys_user.phone_number LIKE ?", "%"+param.PhoneNumber+"%")
	}
	if param.Status != "" {
		query = query.Where("sys_user.status = ?", param.Status)
	}
	if param.DeptId != 0 {
		query = query.Where("sys_user.dept_id = ?", param.DeptId)
	}
	if param.BeginTime != "" && param.EndTime != "" {
		query = query.Where("sys_user.create_time BETWEEN ? AND ?", param.BeginTime, param.EndTime)
	}
	if isPaging {
		query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize)
	}
	query.Find(&users)
	return users, int(count)
}

// Get 根据用户id查询用户信息
func (s *UserService) Get(userId int) dto.UserDetailResponse {
	var user dto.UserDetailResponse
	connector.GetDB().Model(model.SysUser{}).Where("user_id = ?", userId).Last(&user)
	return user
}

// 根据用户名查询用户信息
func (s *UserService) GetByUserName(userName string) dto.UserTokenResponse {
	var user dto.UserTokenResponse
	connector.GetDB().Model(model.SysUser{}).
		Select(
			"sys_user.user_id",
			"sys_user.dept_id",
			"sys_user.user_name",
			"sys_user.nick_name",
			"sys_user.user_type",
			"sys_user.password",
			"sys_user.status",
			"sys_dept.dept_name",
		).
		Joins("LEFT JOIN sys_dept ON sys_user.dept_id = sys_dept.dept_id").
		Where("sys_user.user_name = ?", userName).
		Last(&user)
	return user
}

// GetByEmail 根据邮箱查询用户信息
func (s *UserService) GetByEmail(email string) dto.UserTokenResponse {
	var user dto.UserTokenResponse
	connector.GetDB().Model(model.SysUser{}).
		Select(
			"sys_user.user_id",
			"sys_user.dept_id",
			"sys_user.user_name",
			"sys_user.nick_name",
			"sys_user.user_type",
			"sys_user.password",
			"sys_user.status",
			"sys_dept.dept_name",
		).
		Joins("LEFT JOIN sys_dept ON sys_user.dept_id = sys_dept.dept_id").
		Where("sys_user.email = ?", email).
		Last(&user)
	return user
}

// GetByPhoneNumber 根据手机号码查询用户信息
func (s *UserService) GetByPhoneNumber(phoneNumber string) dto.UserTokenResponse {
	var user dto.UserTokenResponse
	connector.GetDB().Model(model.SysUser{}).
		Select(
			"sys_user.user_id",
			"sys_user.dept_id",
			"sys_user.user_name",
			"sys_user.nick_name",
			"sys_user.user_type",
			"sys_user.password",
			"sys_user.status",
			"sys_dept.dept_name",
		).
		Joins("LEFT JOIN sys_dept ON sys_user.dept_id = sys_dept.dept_id").
		Where("sys_user.phone_number = ?", phoneNumber).
		Last(&user)
	return user
}

// ListByRoleId 根据角色id查询已分配角色的用户列表
//
// isAllocation：true-已分配；false-未分配
func (s *UserService) ListByRoleId(param dto.RoleAuthUserAllocatedListRequest, userId int, isAllocation bool) ([]dto.UserListResponse, int) {
	var count int64
	users := make([]dto.UserListResponse, 0)
	query := connector.GetDB().Model(model.SysUser{}).
		Select("sys_user.*", "sys_dept.dept_name", "sys_dept.leader").
		Joins("LEFT JOIN sys_dept ON sys_user.dept_id = sys_dept.dept_id").
		Scopes((&DataScopeService{}).GetDataScope("sys_dept", userId, "sys_user"))
	if isAllocation {
		query.Joins("JOIN sys_user_role ON sys_user_role.user_id = sys_user.user_id").
			Where("sys_user_role.role_id = ?", param.RoleId)
	} else {
		query.Joins("LEFT JOIN sys_user_role ON sys_user_role.user_id = sys_user.user_id").
			Where("sys_user_role.user_id IS NULL")
	}
	if param.UserName != "" {
		query = query.Where("sys_user.user_name LIKE ?", "%"+param.UserName+"%")
	}
	if param.PhoneNumber != "" {
		query = query.Where("sys_user.phone_number LIKE ?", "%"+param.PhoneNumber+"%")
	}
	query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize).Find(&users)
	return users, int(count)
}

// HasPerms 查询用户是否拥有某权限，拥有返回true
func (s *UserService) HasPerms(userId int, perms []string) bool {
	var count int64
	connector.GetDB().Model(model.SysUserRole{}).
		Joins("JOIN sys_role ON sys_user_role.role_id = sys_role.role_id AND sys_role.status = 0").
		Joins("JOIN sys_role_menu ON sys_role_menu.role_id = sys_role.role_id").
		Joins("JOIN sys_menu ON sys_menu.menu_id = sys_role_menu.menu_id AND sys_menu.status = 0").
		Where("sys_role.delete_time IS NULL AND sys_menu.delete_time IS NULL").
		Where("sys_user_role.user_id = ? AND sys_menu.perms IN ?", userId, perms).
		Count(&count)
	return count > 0
}

// HasRoles 查询用户是否拥有某角色，拥有返回true
func (s *UserService) HasRoles(userId int, roles []string) bool {
	var count int64
	connector.GetDB().Model(model.SysUserRole{}).
		Joins("JOIN sys_role ON sys_user_role.role_id = sys_role.role_id AND sys_role.status = 0").
		Where("sys_role.delete_time IS NULL").
		Where("sys_user_role.user_id = ? AND sys_role.role_key IN ?", userId, roles).
		Count(&count)
	return count > 0
}
