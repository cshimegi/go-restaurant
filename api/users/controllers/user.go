package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"alan/blog/users/dao"
	"alan/blog/users/services"
)

var (
	userService *services.UserService
)

// UserAPI is the interface for user API
type UserAPI interface {
	ListAll(c *gin.Context)
}

// UserController defines user's controller
type UserController struct{}

func init() {
	userStore := dao.NewUserStore()
	userService = services.NewUserService(userStore)
}

// NewUserAPI defines api paths to user service
func NewUserAPI() *UserController {
	return &UserController{}
}

// ListAll lists all users
func (u UserController) ListAll(c *gin.Context) {
	users, err := userService.ListAll(c)

	if err != nil {
		// TODO: use logrus
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
