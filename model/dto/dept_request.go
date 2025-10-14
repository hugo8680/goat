package dto

// SaveDeptRequest 保存部门
type SaveDeptRequest struct {
	DeptId    int    `json:"deptId"`
	ParentId  int    `json:"parentId"`
	Ancestors string `json:"ancestors"`
	DeptName  string `json:"deptName"`
	OrderNum  int    `json:"orderNum"`
	Leader    string `json:"leader"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	CreateBy  string `json:"createBy"`
	UpdateBy  string `json:"updateBy"`
}

// DeptListRequest 部门列表
type DeptListRequest struct {
	DeptName string `query:"deptName" form:"deptName"`
	Status   string `query:"status" form:"status"`
}

// CreateDeptRequest 新增部门
type CreateDeptRequest struct {
	ParentId int    `json:"parentId"`
	DeptName string `json:"deptName"`
	OrderNum int    `json:"orderNum"`
	Leader   string `json:"leader"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Status   string `json:"status"`
}

// UpdateDeptRequest 更新部门
type UpdateDeptRequest struct {
	DeptId    int    `json:"deptId"`
	ParentId  int    `json:"parentId"`
	Ancestors string `json:"ancestors"`
	DeptName  string `json:"deptName"`
	OrderNum  int    `json:"orderNum"`
	Leader    string `json:"leader"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Status    string `json:"status"`
}
