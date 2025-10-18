package connector

import (
	"fmt"
	"forum-service/common/serializer/datetime"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	GREEN   = "\033[97;42m"
	WHITE   = "\033[90;47m"
	YELLOW  = "\033[90;43m"
	RED     = "\033[97;41m"
	BLUE    = "\033[97;44m"
	MAGENTA = "\033[97;45m"
	CYAN    = "\033[97;46m"
	RESET   = "\033[0m"
)

func InitializeLogger(server *gin.Engine) {
	gin.ForceConsoleColor()
	server.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		var statusColor, methodColor, resetColor string
		if param.IsOutputColor() {
			statusColor = param.StatusCodeColor()
			methodColor = param.MethodColor()
			resetColor = param.ResetColor()
		}
		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}
		return fmt.Sprintf("[GIN-%s] %v |%s %3d %s| %13v | %12s |%s %-7s %s %#v\n%s",
			gin.Mode(),
			param.TimeStamp.Format(datetime.DATETIME_FORMAT0),
			statusColor,
			param.StatusCode,
			resetColor,
			param.Latency,
			param.ClientIP,
			methodColor,
			param.Method,
			resetColor,
			param.Path,
			param.ErrorMessage,
		)
	}))
}
