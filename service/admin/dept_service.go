package admin

import (
	"errors"
	"github.com/hugo8680/goat/framework/connector"
	"github.com/hugo8680/goat/model"
	"github.com/hugo8680/goat/model/dto"
	"strconv"
)

type DeptService struct {
}

// Create 创建部门
func (s *DeptService) Create(param dto.SaveDeptRequest) error {
	if dept := s.GetByDeptName(param.DeptName); dept.DeptId > 0 {
		return errors.New("新增部门" + param.DeptName + "失败，部门名称已存在")
	}
	// 拼接ancestors，获取上级的祖级列表
	parentDept := s.Get(param.ParentId)
	if parentDept.Status == "1" {
		return errors.New("部门停用，不允许新增")
	}
	ancestors := parentDept.Ancestors + "," + strconv.Itoa(parentDept.DeptId)
	return connector.GetDB().Model(model.SysDept{}).Create(&model.SysDept{
		ParentId:  param.ParentId,
		Ancestors: ancestors,
		DeptName:  param.DeptName,
		OrderNum:  param.OrderNum,
		Leader:    param.Leader,
		Phone:     param.Phone,
		Email:     param.Email,
		Status:    param.Status,
		CreateBy:  param.CreateBy,
	}).Error
}

// Update 更新部门
func (s *DeptService) Update(param dto.SaveDeptRequest) error {
	if dept := s.GetByDeptName(param.DeptName); dept.DeptId > 0 && dept.DeptId != param.DeptId {
		return errors.New("修改部门" + param.DeptName + "失败，部门名称已存在")
	}
	if dept := s.Get(param.DeptId); dept.ParentId != param.ParentId && s.HasChildren(param.DeptId) {
		return errors.New("存在子级部门，无法直接修改所属部门")
	}
	// 拼接ancestors，获取上级的祖级列表
	parentDept := s.Get(param.ParentId)
	if parentDept.Status == "1" {
		return errors.New("部门停用，不允许新增")
	}
	ancestors := parentDept.Ancestors + "," + strconv.Itoa(parentDept.DeptId)
	return connector.GetDB().Model(model.SysDept{}).Where("dept_id = ?", param.DeptId).Updates(&model.SysDept{
		ParentId:  param.ParentId,
		Ancestors: ancestors,
		DeptName:  param.DeptName,
		OrderNum:  param.OrderNum,
		Leader:    param.Leader,
		Phone:     param.Phone,
		Email:     param.Email,
		Status:    param.Status,
		UpdateBy:  param.UpdateBy,
	}).Error
}

// Delete 删除部门
func (s *DeptService) Delete(deptId int) error {
	if s.HasChildren(deptId) {
		return errors.New("存在下级部门，不允许删除")
	}
	if (&UserService{}).HasUser(deptId) {
		return errors.New("部门存在用户，不允许删除")
	}
	return connector.GetDB().Model(model.SysDept{}).Where("dept_id = ?", deptId).Delete(&model.SysDept{}).Error
}

// List 获取部门列表
func (s *DeptService) List(param dto.DeptListRequest, userId int) []dto.DeptListResponse {
	deptList := make([]dto.DeptListResponse, 0)
	query := connector.GetDB().Model(model.SysDept{}).Order("order_num, dept_id").Scopes((&DataScopeService{}).GetDataScope("sys_dept", userId, ""))
	if param.DeptName != "" {
		query.Where("dept_name LIKE ?", "%"+param.DeptName+"%")
	}
	if param.Status != "" {
		query.Where("status = ?", param.Status)
	}
	query.Find(&deptList)
	return deptList
}

// Get 根据部门id查询部门信息
func (s *DeptService) Get(deptId int) dto.DeptDetailResponse {
	var dept dto.DeptDetailResponse
	connector.GetDB().Model(model.SysDept{}).Where("dept_id = ?", deptId).Last(&dept)
	return dept
}

// GetByDeptName 根据部门名称查询部门信息
func (s *DeptService) GetByDeptName(deptName string) dto.DeptDetailResponse {
	var dept dto.DeptDetailResponse
	connector.GetDB().Model(model.SysDept{}).Where("dept_name = ?", deptName).Last(&dept)
	return dept
}

// ListIdsByRoleId 根据角色id获取部门id集合
func (s *DeptService) ListIdsByRoleId(roleId int) []int {
	deptIds := make([]int, 0)
	connector.GetDB().Model(model.SysRoleDept{}).
		Joins("JOIN sys_dept ON sys_dept.dept_id = sys_role_dept.dept_id").
		Where("sys_dept.status = 0 AND sys_role_dept.role_id = ?", roleId).
		Pluck("sys_dept.dept_id", &deptIds)
	return deptIds
}

// Tree 部门下拉树列表
func (s *DeptService) Tree() []dto.TreeResponse {
	deptList := make([]dto.TreeResponse, 0)
	connector.GetDB().Model(model.SysDept{}).Order("order_num, dept_id").
		Select("dept_id as id", "dept_name as label", "parent_id").
		Where("status = 0").
		Find(&deptList)
	return deptList
}

// RemakeTree 重构部门树
func (s *DeptService) RemakeTree(depts []dto.TreeResponse, parentId int) []dto.TreeResponse {
	tree := make([]dto.TreeResponse, 0)
	for _, dept := range depts {
		if dept.ParentId == parentId {
			tree = append(tree, dto.TreeResponse{
				Id:       dept.Id,
				Label:    dept.Label,
				ParentId: dept.ParentId,
				Children: s.RemakeTree(depts, dept.Id),
			})
		}
	}
	return tree
}

// TreeByUserId 获取用户所属的部门树
func (s *DeptService) TreeByUserId(userId int) []dto.DeptTreeResponse {
	depts := make([]dto.DeptTreeResponse, 0)
	connector.GetDB().Model(model.SysDept{}).
		Select(
			"dept_id as id",
			"dept_name as label",
			"parent_id",
		).
		Order("order_num, dept_id").
		Where("status = 0").
		Scopes((&DataScopeService{}).GetDataScope("sys_dept", userId, "")).
		Find(&depts)
	return depts
}

// RemakeTreeByUserId 部门列表转树形
func (s *UserService) RemakeTreeByUserId(depts []dto.DeptTreeResponse, parentId int) []dto.DeptTreeResponse {
	tree := make([]dto.DeptTreeResponse, 0)
	// 构建树形结构
	for _, dept := range depts {
		if dept.ParentId == parentId {
			dept.Children = s.RemakeTreeByUserId(depts, dept.Id)
			tree = append(tree, dept)
		}
	}
	return tree
}

// HasChildren 查询部门是否存在下级
func (s *DeptService) HasChildren(deptId int) bool {
	var count int64
	connector.GetDB().Model(model.SysDept{}).Where("parent_id = ?", deptId).Count(&count)
	return count > 0
}

// HasUser 根据部门id查询是否存在用户
func (s *UserService) HasUser(deptId int) bool {
	var count int64
	connector.GetDB().Model(model.SysUser{}).Where("dept_id = ?", deptId).Count(&count)
	return count > 0
}
