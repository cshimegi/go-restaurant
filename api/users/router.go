package main

import (
	"users/controllers"

	"github.com/gin-gonic/gin"
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

func initRouters(engine *gin.Engine) {
	group := engine.Group(apiPathPrefix)
	{
		group.GET("", userAPI.ListAll)
	}
}
