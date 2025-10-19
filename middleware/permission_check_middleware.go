package middleware

import (
	"github.com/hugo8680/goat/framework/response"
	"github.com/hugo8680/goat/service/admin"

	"github.com/gin-gonic/gin"
)

// PermissionCheckMiddleware 验证用户是否具备某权限
func PermissionCheckMiddleware(perm string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		securityService := &admin.SecurityService{}
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
