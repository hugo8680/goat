package response

import (
	"github.com/gin-gonic/gin"
)

// Response 响应
type Response struct {
	ctx    *gin.Context
	Status int
	Code   int
	Msg    string
	Data   map[string]interface{}
}

// New 默认响应
func New(ctx *gin.Context) *Response {
	return &Response{
		Status: 200,
		Code:   200,
		Msg:    "成功",
		Data:   make(map[string]interface{}),
		ctx:    ctx,
	}
}

// Success 初始化成功响应
func Success(ctx *gin.Context) *Response {
	return &Response{
		Status: 200,
		Code:   200,
		Msg:    "成功",
		Data:   make(map[string]interface{}),
		ctx:    ctx,
	}
}

// Error 初始化失败响应
func Error(ctx *gin.Context) *Response {
	return &Response{
		Status: 200,
		Code:   500,
		Msg:    "失败",
		Data:   make(map[string]interface{}),
		ctx:    ctx,
	}
}

// SetStatus 设置状态码
func (r *Response) SetStatus(status int) *Response {
	r.Status = status
	return r
}

// SetCode 设置响应码
func (r *Response) SetCode(code int) *Response {
	r.Code = code
	return r
}

// SetMsg 设置响应信息
func (r *Response) SetMsg(msg string) *Response {
	r.Msg = msg
	return r
}

// SetData 设置响应数据
func (r *Response) SetData(key string, value interface{}) *Response {
	if key == "code" || key == "msg" {
		return r
	}
	r.Data[key] = value
	return r
}

// SetPageData 设置分页响应数据
func (r *Response) SetPageData(rows interface{}, total int) *Response {
	r.Data["rows"] = rows
	r.Data["total"] = total
	return r
}

// SetDataMap 设置响应数据
func (r *Response) SetDataMap(data map[string]interface{}) *Response {
	for key, value := range data {
		if key == "code" || key == "msg" {
			continue
		}
		r.Data[key] = value
	}
	return r
}

// Json 序列化返回
func (r *Response) Json() {
	response := gin.H{
		"code": r.Code,
		"msg":  r.Msg,
	}
	for key, value := range r.Data {
		response[key] = value
	}
	r.ctx.JSON(r.Status, response)
}
