package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/hugo8680/goat/common/ip"
	"github.com/hugo8680/goat/common/response_writer"
	"github.com/hugo8680/goat/common/serializer/datetime"
	"github.com/hugo8680/goat/framework/response"
	"github.com/hugo8680/goat/model/dto"
	"github.com/hugo8680/goat/service/admin"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

// LoginLogMiddleware 登录信息记录
func LoginLogMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		loginLogService := &admin.LoginLogService{}
		// 因读取请求体后，请求体的数据流会被消耗完毕，未避免EOF错误，需要缓存请求体，并且每次使用后需要重新赋值给ctx.Request.Body
		bodyBytes, _ := ctx.GetRawData()
		// 将缓存的请求体重新赋值给ctx.Request.Body，供下方ctx.ShouldBind使用
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		rw := &response_writer.ResponseWriter{
			ResponseWriter: ctx.Writer,
			Body:           bytes.NewBufferString(""),
		}
		var param dto.LoginRequest
		if err := ctx.ShouldBind(&param); err != nil {
			response.Error(ctx).SetCode(400).SetMsg(err.Error()).Json()
			ctx.Abort()
			return
		}
		// 因ctx.ShouldBind后，请求体的数据流会被消耗完毕，需要将缓存的请求体重新赋值给ctx.Request.Body
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		ipAddr := ip.GetAddress(ctx.ClientIP(), ctx.Request.UserAgent())
		loginLog := dto.SaveLoginLogRequest{
			UserName:      param.Username,
			Ipaddr:        ipAddr.Ip,
			LoginLocation: ipAddr.Addr,
			Browser:       ipAddr.Browser,
			Os:            ipAddr.Os,
			Status:        "0",
			LoginTime:     datetime.Datetime{Time: time.Now()},
		}
		ctx.Writer = rw
		ctx.Next()
		// 解析响应
		var body response.Response
		err := json.Unmarshal(rw.Body.Bytes(), &body)
		if err != nil || body.Code != 200 {
			loginLog.Status = "1"
		}
		loginLog.Msg = body.Msg
		err = loginLogService.Create(loginLog)
		if err != nil {
			response.Error(ctx).SetCode(500).SetMsg(err.Error()).Json()
		}
	}
}
