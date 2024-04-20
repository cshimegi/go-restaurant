package appetizers

import (
	"github.com/gin-gonic/gin"

	"alan/restaurant/appetizers/controllers"
	"alan/restaurant/appetizers/dao"
	"alan/restaurant/appetizers/services"
	"alan/restaurant/users/shared/logger"
)

var (
	controller controllers.IAppetizerController
)

const (
	apiPathPrefix = "/api/appetizers"
)

func init() {
	log := logger.NewUserLogger()
	store := dao.NewAppetizerStore(log)
	service := services.NewAppetizerService(store, log)
	controller = controllers.NewAppetizerController(service, log)
}

func InitRouters(engine *gin.Engine) {
	group := engine.Group(apiPathPrefix)
	{
		group.POST("", controller.Create)
		group.GET("", controller.ListAll)
		group.GET("/:id", controller.GetById)
		group.PATCH("/:id", controller.PatchById)
		group.DELETE("/:id", controller.DeleteById)
	}
}
