package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"alan/blog/users/shared/domain"
)

// IUserService is the interface for user API
type IUserService interface {
	ListAll(c *gin.Context) ([]domain.User, error)
}

// UserController defines user's controller
type UserController struct {
	svc IUserService
}

// NewUserController defines api paths to user service
func NewUserController(svc IUserService) UserController {
	return UserController{
		svc: svc,
	}
}

// ListAll lists all users
func (u *UserController) ListAll(c *gin.Context) {
	users, err := u.svc.ListAll(c)

	if err != nil {
		// TODO: use logrus
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
