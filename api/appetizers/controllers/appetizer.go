package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"alan/restaurant/appetizers/shared/domain"
)

// IAppetizerService is interface of post service
type IAppetizerService interface {
	ListAll(c *gin.Context) ([]domain.Appetizer, error)
}

// AppetizerController defines type of post controller
type AppetizerController struct {
	svc IAppetizerService
	log *logrus.Entry
}

// NewAppetizerController defines api paths to post service
func NewAppetizerController(svc IAppetizerService, log *logrus.Entry) AppetizerController {
	return AppetizerController{
		svc: svc,
		log: log,
	}
}

// ListAll lists all appetizers
func (a *AppetizerController) ListAll(c *gin.Context) {
	appetizers, err := a.svc.ListAll(c)

	if err != nil {
		a.log.WithField("error", err).Error("Error listing appetizers")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": appetizers,
	})
}
