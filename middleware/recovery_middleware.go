package middleware

import (
	"github.com/hugo8680/goat/framework/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RecoveryMiddleware 捕获panic恢复程序的中间件
func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("recovered: %v", err)
				response.Error(ctx).SetCode(http.StatusInternalServerError).SetMsg("服务器错误").Json()
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
