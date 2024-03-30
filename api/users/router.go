package users

import (
	"github.com/gin-gonic/gin"

	"alan/blog/users/controllers"
)

var (
	userAPI controllers.UserAPI
)

const (
	apiPathPrefix = "/api/users"
)

func init() {
	userAPI = controllers.NewUserAPI()
}

func InitRouters(engine *gin.Engine) {
	group := engine.Group(apiPathPrefix)
	{
		group.GET("", userAPI.ListAll)
	}
}
