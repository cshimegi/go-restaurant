package main

import (
	"posts/controllers"

	"github.com/gin-gonic/gin"
)

var (
	postAPI controllers.PostAPI
)

const (
	apiPathPrefix = "/api/posts"
)

func init() {
	postAPI = controllers.NewPostAPI()
}

func initRouters(engine *gin.Engine) {
	group := engine.Group(apiPathPrefix)
	{
		group.GET("", postAPI.ListAll)
	}
}
