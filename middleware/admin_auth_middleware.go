package middleware

import (
	"github.com/hugo8680/goat/common/constant/auth"
	"github.com/hugo8680/goat/framework/response"
	"github.com/hugo8680/goat/service/admin"
	"time"

	"github.com/gin-gonic/gin"
)

// AdminAuthMiddleware 认证中间件
func AdminAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		securityService := &admin.SecurityService{}
		tokenService := admin.NewTokenService()
		authUser, err := securityService.GetCurrentUser(ctx)
		if err != nil {
			response.Error(ctx).SetCode(401).SetMsg("未登录").Json()
			ctx.Abort()
			return
		}
		// 判断token临期，小于20分钟刷新
		if authUser.ExpireTime.Time.Before(time.Now().Add(time.Minute * 20)) {
			tokenService.Refresh(ctx, authUser)
		}
		if authUser.Status != "0" {
			response.Error(ctx).SetCode(601).SetMsg("用户被禁用").Json()
			ctx.Abort()
			return
		}
		ctx.Set(auth.CONTEXT_USER_KEY, authUser)
		ctx.Next()
	}
}
