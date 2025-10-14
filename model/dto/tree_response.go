package dto

// TreeResponse 树状图
type TreeResponse struct {
	Id       int            `json:"id"`
	Label    string         `json:"label"`
	Children []TreeResponse `json:"children" gorm:"-"`
	ParentId int            `json:"-"`
}
