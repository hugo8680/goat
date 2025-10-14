package middleware

import (
	"forum-service/framework/response"
	"forum-service/service"

	"github.com/gin-gonic/gin"
)

// PermissionCheckMiddleware 验证用户是否具备某权限
func PermissionCheckMiddleware(perm string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		securityService := &service.SecurityService{}
		authUserId, _ := securityService.GetCurrentUserId(ctx)
		if authUserId == 1 {
			ctx.Next()
			return
		}
		if hasPerm := securityService.HasPerm(authUserId, perm); !hasPerm {
			response.Error(ctx).SetCode(601).SetMsg("权限不足").Json()
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
