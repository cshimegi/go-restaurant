package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"alan/restaurant/appetizers/services"
)

type IAppetizerController interface {
	ListAll(c *gin.Context)
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

// ListAll lists all appetizers
func (a *AppetizerController) ListAll(c *gin.Context) {
	appetizers, err := a.svc.ListAll(c)

	switch err.(type) {
	case nil:
		c.JSON(http.StatusOK, gin.H{"data": appetizers})
	default:
		a.log.WithField("error", err.Error()).Error("Error listing appetizers")
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
}
