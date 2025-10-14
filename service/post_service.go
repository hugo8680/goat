package service

import (
	"errors"
	"forum-service/framework/connector"
	"forum-service/model"
	"forum-service/model/dto"
)

type PostService struct {
}

// Create 创建岗位
func (s *PostService) Create(param dto.SavePostRequest) error {
	if post := s.GetByPostName(param.PostName); post.PostId > 0 {
		return errors.New("新增岗位" + param.PostName + "失败，岗位名称已存在")
	}
	if post := s.GetByPostCode(param.PostCode); post.PostId > 0 {
		return errors.New("新增岗位" + param.PostName + "失败，岗位编码已存在")
	}
	return connector.GetDB().Model(model.SysPost{}).Create(&model.SysPost{
		PostCode: param.PostCode,
		PostName: param.PostName,
		PostSort: param.PostSort,
		Status:   param.Status,
		Remark:   param.Remark,
		CreateBy: param.CreateBy,
	}).Error
}

// Update 更新岗位
func (s *PostService) Update(param dto.SavePostRequest) error {
	if post := s.GetByPostName(param.PostName); post.PostId > 0 && post.PostId != param.PostId {
		return errors.New("修改岗位" + param.PostName + "失败，岗位名称已存在")
	}
	if post := s.GetByPostCode(param.PostCode); post.PostId > 0 && post.PostId != param.PostId {
		return errors.New("修改岗位" + param.PostName + "失败，岗位编码已存在")
	}
	return connector.GetDB().Model(model.SysPost{}).Where("post_id = ?", param.PostId).Updates(&model.SysPost{
		PostCode: param.PostCode,
		PostName: param.PostName,
		PostSort: param.PostSort,
		Status:   param.Status,
		Remark:   param.Remark,
		UpdateBy: param.UpdateBy,
	}).Error
}

// Delete 删除岗位
func (s *PostService) Delete(postIds []int) error {
	return connector.GetDB().Model(model.SysPost{}).Where("post_id IN ?", postIds).Delete(&model.SysPost{}).Error
}

// List 岗位列表
func (s *PostService) List(param dto.PostListRequest, isPaging bool) ([]dto.PostListResponse, int) {
	var count int64
	posts := make([]dto.PostListResponse, 0)
	query := connector.GetDB().Model(model.SysPost{}).Order("post_sort, post_id")
	if param.PostCode != "" {
		query.Where("post_code LIKE ?", "%"+param.PostCode+"%")
	}
	if param.PostName != "" {
		query.Where("post_name LIKE ?", "%"+param.PostName+"%")
	}
	if param.Status != "" {
		query.Where("status = ?", param.Status)
	}
	if isPaging {
		query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize)
	}
	query.Find(&posts)
	return posts, int(count)
}

// Get 根据岗位id获取岗位详情
func (s *PostService) Get(postId int) dto.PostDetailResponse {
	var post dto.PostDetailResponse
	connector.GetDB().Model(model.SysPost{}).Where("post_id = ?", postId).Last(&post)
	return post
}

// GetByPostName 根据岗位名称获取岗位详情
func (s *PostService) GetByPostName(postName string) dto.PostDetailResponse {
	var post dto.PostDetailResponse
	connector.GetDB().Model(model.SysPost{}).Where("post_name = ?", postName).Last(&post)
	return post
}

// GetByPostCode 根据岗位编码获取岗位详情
func (s *PostService) GetByPostCode(postCode string) dto.PostDetailResponse {
	var post dto.PostDetailResponse
	connector.GetDB().Model(model.SysPost{}).Where("post_code = ?", postCode).Last(&post)
	return post
}

// ListIdsByUserId 根据用户id查询岗位id集合
func (s *PostService) ListIdsByUserId(userId int) []int {
	var postIds []int
	connector.GetDB().Model(model.SysPost{}).
		Joins("JOIN sys_user_post ON sys_user_post.post_id = sys_post.post_id").
		Where("sys_user_post.user_id = ? AND sys_post.status = 0", userId).
		Pluck("sys_post.post_id", &postIds)
	return postIds
}

// ListNamesByUserId 根据用户id查询角色名
func (s *PostService) ListNamesByUserId(userId int) []string {
	var postNames []string
	connector.GetDB().Model(model.SysPost{}).
		Joins("JOIN sys_user_post ON sys_user_post.post_id = sys_post.post_id").
		Where("sys_user_post.user_id = ? AND sys_post.status = 0", userId).
		Pluck("sys_post.post_name", &postNames)
	return postNames
}
