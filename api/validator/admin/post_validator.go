package admin

import (
	"errors"
	"forum-service/model/dto"
)

// CreatePostValidator 添加岗位验证
func CreatePostValidator(param dto.CreatePostRequest) error {
	if param.PostCode == "" {
		return errors.New("请输入岗位编码")
	}
	if param.PostName == "" {
		return errors.New("请输入岗位名称")
	}
	return nil
}

// UpdatePostValidator 更新岗位验证
func UpdatePostValidator(param dto.UpdatePostRequest) error {
	if param.PostId <= 0 {
		return errors.New("参数错误")
	}
	if param.PostCode == "" {
		return errors.New("请输入岗位编码")
	}
	if param.PostName == "" {
		return errors.New("请输入岗位名称")
	}
	return nil
}
