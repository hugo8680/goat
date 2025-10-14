package response_writer

import (
	"bytes"

	"github.com/gin-gonic/gin"
)

// ResponseWriter 重写gin的ResponseWriter，用户接收响应体
type ResponseWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func (r *ResponseWriter) Write(b []byte) (int, error) {
	r.Body.Write(b)
	return r.ResponseWriter.Write(b)
}

func (r *ResponseWriter) WriteString(s string) (int, error) {
	r.Body.WriteString(s)
	return r.ResponseWriter.WriteString(s)
}
