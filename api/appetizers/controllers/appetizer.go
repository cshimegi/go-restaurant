package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"alan/restaurant/appetizers/services"
	"alan/restaurant/appetizers/shared/domain"
	sc "alan/restaurant/shared/consts"
)

type IAppetizerController interface {
	Create(c *gin.Context)
	ListAll(c *gin.Context)
	GetById(c *gin.Context)
	PatchById(c *gin.Context)
	DeleteById(c *gin.Context)
}

// AppetizerController defines type of post controller
type AppetizerController struct {
	svc services.IAppetizerService
	log *logrus.Entry
}

// NewAppetizerController defines api paths to post service
func NewAppetizerController(svc services.IAppetizerService, log *logrus.Entry) IAppetizerController {
	return &AppetizerController{
		svc: svc,
		log: log,
	}
}

func (a *AppetizerController) Create(c *gin.Context) {
	appetizer := domain.Appetizer{}
	if err := c.ShouldBindJSON(&appetizer); err != nil {
		a.log.WithField("error", err.Error()).Error("Bad request when creating appetizer")
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	newAppetizer, err := a.svc.Create(c, appetizer)
	switch {
	case err == nil:
		c.JSON(http.StatusCreated, gin.H{"data": newAppetizer})
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
}

// ListAll lists all appetizers
func (a *AppetizerController) ListAll(c *gin.Context) {
	appetizers, err := a.svc.ListAll(c)

	switch err.(type) {
	case nil:
		c.JSON(http.StatusOK, gin.H{"data": appetizers})
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
}

func (a *AppetizerController) GetById(c *gin.Context) {
	uid, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		a.log.WithField("error", err.Error()).Error("Bad request when getting appetizer by id")
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	appetizer, err := a.svc.GetById(c, uint(uid))
	switch {
	case err == nil:
		c.JSON(http.StatusOK, gin.H{"data": appetizer})
	case errors.Is(err, sc.ErrNotFound):
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
}

func (a *AppetizerController) PatchById(c *gin.Context) {
	uid, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		a.log.WithField("error", err.Error()).Error("Bad request when getting appetizer by id")
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	var appetizer domain.Appetizer
	if err := c.ShouldBindJSON(&appetizer); err != nil {
		a.log.WithField("error", err.Error()).Error("Bad requested json body")
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	updatedAppetizer, err := a.svc.PatchById(c, uint(uid), appetizer)
	switch {
	case err == nil:
		c.JSON(http.StatusOK, gin.H{"data": updatedAppetizer})
	case errors.Is(err, sc.ErrNotFound):
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
}

func (a *AppetizerController) DeleteById(c *gin.Context) {
	uid, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		a.log.WithField("error", err.Error()).Error("Bad requested appetizer id")
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	err = a.svc.DeleteById(c, uint(uid))
	switch {
	case err == nil:
		c.Status(http.StatusNoContent)
	case errors.Is(err, sc.ErrNotFound):
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
}
