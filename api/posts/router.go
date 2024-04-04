package posts

import (
	"alan/blog/posts/dao"
	"alan/blog/posts/services"
	"github.com/gin-gonic/gin"

	"alan/blog/posts/controllers"
)

var (
	controller controllers.PostController
)

const (
	apiPathPrefix = "/api/posts"
)

func init() {
	store := dao.NewPostStore()
	service := services.NewPostService(store)
	controller = controllers.NewPostController(service)
}

func InitRouters(engine *gin.Engine) {
	group := engine.Group(apiPathPrefix)
	{
		group.GET("", controller.ListAll)
	}
}
