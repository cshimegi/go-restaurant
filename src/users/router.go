package main

import (
	ctrl "users/controllers"

	"github.com/gin-gonic/gin"
)

const (
	apiPathPrefix = "/api/users"
)

func initRouters(engine *gin.Engine) {
	group := engine.Group(apiPathPrefix)
	{
		group.GET("", ctrl.ListUsers)
	}
}
