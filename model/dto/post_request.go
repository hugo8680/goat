package dto

// SavePostRequest 保存岗位
type SavePostRequest struct {
	PostId   int    `json:"postId"`
	PostCode string `json:"postCode"`
	PostName string `json:"postName"`
	PostSort int    `json:"postSort"`
	Status   string `json:"status"`
	CreateBy string `json:"createBy"`
	UpdateBy string `json:"updateBy"`
	Remark   string `json:"remark"`
}

// PostListRequest 岗位列表
type PostListRequest struct {
	PageRequest
	PostCode string `query:"postCode" form:"postCode"`
	PostName string `query:"postName" form:"postName"`
	Status   string `query:"status" form:"status"`
}

// CreatePostRequest 新增岗位
type CreatePostRequest struct {
	PostCode string `json:"postCode"`
	PostName string `json:"postName"`
	PostSort int    `json:"postSort"`
	Status   string `json:"status"`
	Remark   string `json:"remark"`
}

// UpdatePostRequest 更新岗位
type UpdatePostRequest struct {
	PostId   int    `json:"postId"`
	PostCode string `json:"postCode"`
	PostName string `json:"postName"`
	PostSort int    `json:"postSort"`
	Status   string `json:"status"`
	Remark   string `json:"remark"`
}
