package users

import (
	"github.com/gin-gonic/gin"

	"alan/restaurant/users/controllers"
	"alan/restaurant/users/dao"
	"alan/restaurant/users/services"
	"alan/restaurant/users/shared/logger"
)

var (
	controller controllers.IUserController
)

const (
	apiPathPrefix = "/api/users"
)

func init() {
	log := logger.NewUserLogger()
	store := dao.NewUserStore(log)
	service := services.NewUserService(store, log)
	controller = controllers.NewUserController(service, log)
}

func InitRouters(engine *gin.Engine) {
	group := engine.Group(apiPathPrefix)
	{
		group.GET("", controller.ListAll)
	}
}
