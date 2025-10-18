package middleware

import (
	"bytes"
	"encoding/json"
	"forum-service/common/ip"
	"forum-service/common/response_writer"
	"forum-service/common/serializer/datetime"
	"forum-service/framework/response"
	"forum-service/model/dto"
	"forum-service/service"
	"io"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// OperLogMiddleware 操作日志中间件
//
// title 操作模块标题
// businessType 操作类型 constant.REQUEST_BUSINESS_TYPE_*
func OperLogMiddleware(title string, businessType int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		securityService := &service.SecurityService{}
		operLogService := &service.OperLogService{}
		var operName, deptName string
		if authUser, _ := securityService.GetCurrentUser(ctx); authUser != nil {
			operName = authUser.NickName
			deptName = authUser.DeptName
		}
		// 记录请求时间，用于获取请求耗时
		requestStartTime := time.Now()
		// 因读取请求体后，请求体的数据流会被消耗完毕，未避免EOF错误，需要缓存请求体，并且每次使用后需要重新赋值给ctx.Request.Body
		bodyBytes, _ := ctx.GetRawData()
		// 将缓存的请求体重新赋值给ctx.Request.Body，供下方ctx.ShouldBind使用
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		rw := &response_writer.ResponseWriter{
			ResponseWriter: ctx.Writer,
			Body:           bytes.NewBufferString(""),
		}
		param := make(map[string]interface{})
		err := ctx.ShouldBind(&param)
		if err != nil {
			err = json.Unmarshal(bodyBytes, &param)
			if err != nil {
				log.Println(err)
			}
		} else {
			// 因ctx.ShouldBind后，请求体的数据流会被消耗完毕，需要将缓存的请求体重新赋值给ctx.Request.Body
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}
		// 将query参数转为map并添加到请求参数中，用query-key的形式以便区分
		for key, value := range ctx.Request.URL.Query() {
			param[key] = value
		}
		operParam, _ := json.Marshal(&param)
		ipAddress := ip.GetAddress(ctx.ClientIP(), ctx.Request.UserAgent())
		sysOperLog := dto.SaveOperLogRequest{
			Title:         title,
			BusinessType:  businessType,
			Method:        ctx.HandlerName(),
			RequestMethod: ctx.Request.Method,
			OperName:      operName,
			DeptName:      deptName,
			OperUrl:       ctx.Request.URL.Path,
			OperIp:        ipAddress.Ip,
			OperLocation:  ipAddress.Addr,
			OperParam:     string(operParam),
			JsonResult:    "",
			Status:        "0",
			ErrorMsg:      "",
			OperTime:      datetime.Datetime{Time: time.Now()},
			CostTime:      0,
		}
		ctx.Writer = rw
		ctx.Next()
		// 解析响应
		var body response.Response
		if ctx.Request.Header.Get("Content-Type") == "application/json" {
			sysOperLog.JsonResult = rw.Body.String()
			err = json.Unmarshal(rw.Body.Bytes(), &body)
			if err != nil || body.Code != 200 {
				sysOperLog.Status = "1"
				sysOperLog.ErrorMsg = body.Msg
			}
		} else {
			sysOperLog.Status = "0"
			sysOperLog.ErrorMsg = "OK"
		}
		duration := time.Since(requestStartTime)
		sysOperLog.CostTime = int(duration.Milliseconds())
		err = operLogService.Create(sysOperLog)
		if err != nil {
			response.Error(ctx).SetCode(500).SetMsg(err.Error()).Json()
		}
	}
}
