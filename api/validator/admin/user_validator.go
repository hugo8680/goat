package admin

import (
	"errors"
	"github.com/hugo8680/goat/common/constant/regexp"
	"github.com/hugo8680/goat/common/utils"
	"github.com/hugo8680/goat/model/dto"
)

// UpdateProfileValidator 更新个人资料验证
func UpdateProfileValidator(param dto.UpdateProfileRequest) error {
	if param.NickName == "" {
		return errors.New("请输入用户昵称")
	}
	if !utils.CheckRegex(regexp.EMAIL, param.Email) {
		return errors.New("邮箱格式错误")
	}
	if !utils.CheckRegex(regexp.PHONE, param.PhoneNumber) {
		return errors.New("手机号格式错误")
	}
	return nil
}

// UserProfileUpdatePwdValidator 更新个人密码验证
func UserProfileUpdatePwdValidator(param dto.UserProfileUpdatePwdRequest) error {
	if param.OldPassword == "" {
		return errors.New("请输入旧密码")
	}
	if param.NewPassword == "" {
		return errors.New("请输入新密码")
	}
	return nil
}

// CreateUserValidator 添加用户验证
func CreateUserValidator(param dto.CreateUserRequest) error {
	if param.NickName == "" {
		return errors.New("请输入用户昵称")
	}
	if param.UserName == "" {
		return errors.New("请输入用户名称")
	}
	if param.Password == "" {
		return errors.New("请输入用户密码")
	}
	if param.PhoneNumber != "" && !utils.CheckRegex(regexp.PHONE, param.PhoneNumber) {
		return errors.New("手机号码格式错误")
	}
	if param.Email != "" && !utils.CheckRegex(regexp.EMAIL, param.Email) {
		return errors.New("邮箱账号格式错误")
	}
	return nil
}

// UpdateUserValidator 更新用户验证
func UpdateUserValidator(param dto.UpdateUserRequest) error {
	if param.UserId <= 0 {
		return errors.New("参数错误")
	}
	if param.NickName == "" {
		return errors.New("请输入用户昵称")
	}
	if param.PhoneNumber != "" && !utils.CheckRegex(regexp.PHONE, param.PhoneNumber) {
		return errors.New("手机号码格式错误")
	}
	if param.Email != "" && !utils.CheckRegex(regexp.EMAIL, param.Email) {
		return errors.New("邮箱账号格式错误")
	}
	return nil
}

// RemoveUserValidator 删除用户验证
func RemoveUserValidator(userIds []int, authUserId int) error {
	if utils.Contains(userIds, 1) {
		return errors.New("超级管理员无法删除")
	}
	if utils.Contains(userIds, authUserId) {
		return errors.New("当前用户无法删除")
	}
	return nil
}

// ChangeUserStatusValidator 修改用户状态验证
func ChangeUserStatusValidator(param dto.UpdateUserRequest) error {
	if param.UserId <= 0 {
		return errors.New("参数错误")
	}
	if param.Status == "" {
		return errors.New("请选择状态")
	}
	return nil
}

// ResetUserPwdValidator 重置用户密码验证
func ResetUserPwdValidator(param dto.UpdateUserRequest) error {
	if param.UserId <= 0 {
		return errors.New("参数错误")
	}
	if param.Password == "" {
		return errors.New("请输入用户密码")
	}
	return nil
}

// ImportUserValidator 导入用户验证
func ImportUserValidator(param dto.CreateUserRequest) error {
	if param.NickName == "" {
		return errors.New("请输入用户昵称")
	}
	if param.UserName == "" {
		return errors.New("请输入用户名称")
	}
	if param.PhoneNumber != "" && !utils.CheckRegex(regexp.PHONE, param.PhoneNumber) {
		return errors.New("手机号码格式错误")
	}
	if param.Email != "" && !utils.CheckRegex(regexp.EMAIL, param.Email) {
		return errors.New("邮箱账号格式错误")
	}
	return nil
}
