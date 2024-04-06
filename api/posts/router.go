package posts

import (
	"github.com/gin-gonic/gin"

	"alan/blog/posts/controllers"
	"alan/blog/posts/dao"
	"alan/blog/posts/services"
	"alan/blog/users/shared/logger"
)

var (
	controller controllers.PostController
)

const (
	apiPathPrefix = "/api/posts"
)

func init() {
	log := logger.NewUserLogger()
	store := dao.NewPostStore(log)
	service := services.NewPostService(store, log)
	controller = controllers.NewPostController(service, log)
}

func InitRouters(engine *gin.Engine) {
	group := engine.Group(apiPathPrefix)
	{
		group.GET("", controller.ListAll)
	}
}
