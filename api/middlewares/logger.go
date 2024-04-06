package middlwares

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
)

func JsonLoggerMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(
		func(params gin.LogFormatterParams) string {
			messages := make(map[string]any)

			messages["method"] = params.Method
			messages["status"] = params.StatusCode
			messages["path"] = params.Path
			messages["time"] = time.Now().Format("2006-01-02 15:04:05")
			messages["response_time"] = params.Latency.String()
			messages["remote_addr"] = params.ClientIP
			messages["user_agent"] = params.Request.UserAgent()
			messages["error"] = params.ErrorMessage
			messages["referer"] = params.Request.Referer()
			messages["proto"] = params.Request.Proto

			str, _ := json.Marshal(messages)
			return string(str) + "\n"
		},
	)
}
