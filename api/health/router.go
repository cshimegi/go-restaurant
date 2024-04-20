package health

import (
	"github.com/gin-gonic/gin"

	"alan/restaurant/health/controllers"
	"alan/restaurant/health/dao"
	"alan/restaurant/health/services"
	"alan/restaurant/health/shared/logger"
)

var (
	controller controllers.IHealthController
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
		group.GET("", controller.GetLatest)
		group.GET("/all", controller.ListAll)
	}
}
