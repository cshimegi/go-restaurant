package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"alan/restaurant/users/shared/domain"
)

// IUserService is the interface for user API
type IUserService interface {
	ListAll(c *gin.Context) ([]domain.User, error)
}

// UserController defines user's controller
type UserController struct {
	svc IUserService
	log *logrus.Entry
}

// NewUserController defines api paths to user service
func NewUserController(svc IUserService, log *logrus.Entry) UserController {
	return UserController{
		svc: svc,
		log: log,
	}
}

// ListAll lists all users
func (u *UserController) ListAll(c *gin.Context) {
	users, err := u.svc.ListAll(c)

	if err != nil {
		u.log.WithField("error", err).Error("Error listing users")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
