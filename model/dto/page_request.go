package dto

// PageRequest 分页请求
type PageRequest struct {
	PageNum  int `query:"pageNum" form:"pageNum"`
	PageSize int `query:"pageSize" form:"pageSize"`
}
