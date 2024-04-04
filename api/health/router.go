package health

import (
	"github.com/gin-gonic/gin"

	"alan/blog/health/controllers"
	"alan/blog/health/dao"
	"alan/blog/health/services"
)

var (
	controller controllers.HealthController
)

const (
	apiPathPrefix = "/api/health"
)

func init() {
	store := dao.NewHealthStore()
	service := services.NewHealthService(store)
	controller = controllers.NewHealthController(service)
}

func InitRouters(engine *gin.Engine) {
	group := engine.Group(apiPathPrefix)
	{
		group.GET("", controller.Retrieve)
		group.GET("/all", controller.ListAll)
	}
}
