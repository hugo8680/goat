package connector

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func InitializeLogger(server *gin.Engine) {
	gin.ForceConsoleColor()
	server.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[syslog] %s - [%s] \"%s %s %s\" %d %s \"%s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.DateTime),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
		)
	}))
}
