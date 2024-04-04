package users

import (
	"github.com/gin-gonic/gin"

	"alan/blog/users/controllers"
	"alan/blog/users/dao"
	"alan/blog/users/services"
)

var (
	controller controllers.UserController
)

const (
	apiPathPrefix = "/api/users"
)

func init() {
	store := dao.NewUserStore()
	service := services.NewUserService(store)
	controller = controllers.NewUserController(service)
}

func InitRouters(engine *gin.Engine) {
	group := engine.Group(apiPathPrefix)
	{
		group.GET("", controller.ListAll)
	}
}
