package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"alan/restaurant/users/services"
)

// IUserController is the interface for user controller
type IUserController interface {
	ListAll(c *gin.Context)
}

// UserController defines user's controller
type UserController struct {
	svc services.IUserService
	log *logrus.Entry
}

// NewUserController defines api paths to user service
func NewUserController(svc services.IUserService, log *logrus.Entry) IUserController {
	return &UserController{
		svc: svc,
		log: log,
	}
}

// ListAll lists all users
func (u *UserController) ListAll(c *gin.Context) {
	users, err := u.svc.ListAll(c)

	switch err.(type) {
	case nil:
		c.JSON(http.StatusOK, gin.H{"data": users})
	default:
		u.log.WithField("error", err.Error()).Error("Error listing users")
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
}
