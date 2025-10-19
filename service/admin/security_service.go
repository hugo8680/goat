package admin

import (
	"errors"
	"github.com/hugo8680/goat/model/dto"

	"github.com/gin-gonic/gin"
)

type SecurityService struct {
}

// GetCurrentUserId 获取当前用户id
func (s *SecurityService) GetCurrentUserId(ctx *gin.Context) (int, error) {
	tokenService := NewTokenService()
	authUser, err := tokenService.Parse(ctx)
	if err != nil {
		return 0, err
	}
	if authUser == nil {
		return 0, errors.New("用户不存在")
	}
	return authUser.UserId, nil
}

// GetCurrentUserDeptId 获取当前用户部门id
func (s *SecurityService) GetCurrentUserDeptId(ctx *gin.Context) (int, error) {
	tokenService := NewTokenService()
	authUser, err := tokenService.Parse(ctx)
	if err != nil {
		return 0, err
	}
	if authUser == nil {
		return 0, errors.New("用户不存在")
	}
	return authUser.DeptId, nil
}

// GetCurrentUserName 获取当前账户名
func (s *SecurityService) GetCurrentUserName(ctx *gin.Context) (string, error) {
	tokenService := NewTokenService()
	authUser, err := tokenService.Parse(ctx)
	if err != nil {
		return "", err
	}
	if authUser == nil {
		return "", errors.New("用户不存在")
	}
	return authUser.UserName, nil
}

// GetCurrentUser 获取当前账户
func (s *SecurityService) GetCurrentUser(ctx *gin.Context) (*dto.UserTokenResponse, error) {
	tokenService := NewTokenService()
	authUser, err := tokenService.Parse(ctx)
	if err != nil {
		return nil, err
	}
	if authUser == nil {
		return nil, errors.New("用户不存在")
	}
	return authUser, nil
}

// HasPerm 验证用户是否具备某权限
func (s *SecurityService) HasPerm(userId int, perm string) bool {
	return (&UserService{}).HasPerms(userId, []string{perm})
}

// LackPerm 验证用户是否不具备某权限
func (s *SecurityService) LackPerm(userId int, perm string) bool {
	return !(&UserService{}).HasPerms(userId, []string{perm})
}

// HasAnyPerms 验证用户是否具有以下任意一个权限
func (s *SecurityService) HasAnyPerms(userId int, perms []string) bool {
	return (&UserService{}).HasPerms(userId, perms)
}

// HasRole 验证用户是否拥有某个角色
func (s *SecurityService) HasRole(userId int, roleKey string) bool {
	return (&UserService{}).HasRoles(userId, []string{roleKey})
}

// LackRole 验证用户是否不具备某个角色
func (s *SecurityService) LackRole(userId int, roleKey string) bool {
	return !(&UserService{}).HasRoles(userId, []string{roleKey})
}

// HasAnyRoles 验证用户是否具有以下任意一个角色
func (s *SecurityService) HasAnyRoles(userId int, roleKey []string) bool {
	return (&UserService{}).HasRoles(userId, roleKey)
}
