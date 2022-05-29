package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"users/services"
)

var (
	userService services.UserService
)

// UserAPI is the interface for user API
type UserAPI interface {
	ListAll(c *gin.Context)
}

// UserController defines user's controller
type UserController struct{}

func init() {
	userService = services.NewUserService()
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
