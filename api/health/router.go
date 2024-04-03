package health

import (
	"github.com/gin-gonic/gin"

	"alan/blog/health/controllers"
)

var (
	healthAPI controllers.HealthAPI
)

const (
	apiPathPrefix = "/api/health"
)

func init() {
	healthAPI = controllers.NewHealthAPI()
}

func InitRouters(engine *gin.Engine) {
	group := engine.Group(apiPathPrefix)
	{
		group.GET("", healthAPI.Retrieve)
		group.GET("/all", healthAPI.ListAll)
	}
}
