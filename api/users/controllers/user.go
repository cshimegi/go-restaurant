package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	sc "alan/restaurant/shared/consts"
	"alan/restaurant/users/services"
	"alan/restaurant/users/shared/domain"
)

// IUserController is the interface for user controller
type IUserController interface {
	Create(c *gin.Context)
	ListAll(c *gin.Context)
	GetById(c *gin.Context)
	PatchById(c *gin.Context)
	DeleteById(c *gin.Context)
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

func (u *UserController) Create(c *gin.Context) {
	user := domain.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		u.log.WithField("error", err.Error()).Error("Bad request when creating user")
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	newUser, err := u.svc.Create(c, user)
	switch {
	case err == nil:
		c.JSON(http.StatusCreated, gin.H{"data": newUser})
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
}

// ListAll lists all users
func (u *UserController) ListAll(c *gin.Context) {
	users, err := u.svc.ListAll(c)

	switch {
	case err == nil:
		c.JSON(http.StatusOK, gin.H{"data": users})
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
}

func (u *UserController) GetById(c *gin.Context) {
	uid, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		u.log.WithField("error", err.Error()).Error("Bad request when getting user by id")
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := u.svc.GetById(c, uint(uid))
	switch {
	case err == nil:
		c.JSON(http.StatusOK, gin.H{"data": user})
	case errors.Is(err, sc.ErrNotFound):
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
}

func (u *UserController) PatchById(c *gin.Context) {
	uid, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		u.log.WithField("error", err.Error()).Error("Bad requested user id")
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		u.log.WithField("error", err.Error()).Error("Bad requested json body")
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	patchedUser, err := u.svc.PatchById(c, uint(uid), user)
	switch {
	case err == nil:
		c.JSON(http.StatusOK, gin.H{"data": patchedUser})
	case errors.Is(err, sc.ErrNotFound):
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
}

func (u *UserController) DeleteById(c *gin.Context) {
	uid, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		u.log.WithField("error", err.Error()).Error("Bad requested user id")
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	err = u.svc.DeleteById(c, uint(uid))
	switch {
	case err == nil:
		c.Status(http.StatusNoContent)
	case errors.Is(err, sc.ErrNotFound):
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
}
