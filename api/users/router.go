package users

import (
	"github.com/gin-gonic/gin"

	"alan/blog/users/controllers"
	"alan/blog/users/dao"
	"alan/blog/users/services"
	"alan/blog/users/shared/logger"
)

var (
	controller controllers.UserController
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
