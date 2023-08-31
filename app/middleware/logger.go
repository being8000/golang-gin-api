package middleware

import (
	"fmt"
	"time"
	"zehan/gin/app/pkg"

	"github.com/gin-gonic/gin"
)

func logFormatter(param gin.LogFormatterParams) string {
	requestID := param.Request.Header.Get(RequestIDKey)

	var statusColor, methodColor, resetColor string
	if param.IsOutputColor() {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}

	if param.Latency > time.Minute {
		param.Latency = param.Latency.Truncate(time.Second)
	}
	return fmt.Sprintf("%v %s |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
		param.TimeStamp.Format(time.RFC3339),
		requestID,
		statusColor, param.StatusCode, resetColor,
		param.Latency,
		param.ClientIP,
		methodColor, param.Method, resetColor,
		param.Path,
		param.ErrorMessage,
	)

	// 	// your custom format
	// 	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	// 	param.ClientIP,
	// 	param.TimeStamp.Format(time.RFC1123),
	// 	param.Method,
	// 	param.Path,
	// 	param.Request.Proto,
	// 	param.StatusCode,
	// 	param.Latency,
	// 	param.Request.UserAgent(),
	// 	param.ErrorMessage,
	// )
}

func Logger() gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Output:    pkg.GetLogWriter(),
		Formatter: logFormatter,
		SkipPaths: []string{},
	})
}
