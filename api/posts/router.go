package posts

import (
	"github.com/gin-gonic/gin"

	"alan/blog/posts/controllers"
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

func InitRouters(engine *gin.Engine) {
	group := engine.Group(apiPathPrefix)
	{
		group.GET("", postAPI.ListAll)
	}
}
