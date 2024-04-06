package health

import (
	"alan/blog/health/shared/logger"
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
	log := logger.NewHealthLogger()
	store := dao.NewHealthStore(log)
	service := services.NewHealthService(store, log)
	controller = controllers.NewHealthController(service, log)
}

func InitRouters(engine *gin.Engine) {
	group := engine.Group(apiPathPrefix)
	{
		group.GET("", controller.Retrieve)
		group.GET("/all", controller.ListAll)
	}
}
