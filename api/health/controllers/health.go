package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"alan/restaurant/health/services"
)

type IHealthController interface {
	GetLatest(c *gin.Context)
	ListAll(c *gin.Context)
}

// HealthController defines Health's controller
type HealthController struct {
	svc services.IHealthService
	log *logrus.Entry
}

// NewHealthController defines api paths to Health service
func NewHealthController(svc services.IHealthService, log *logrus.Entry) IHealthController {
	return &HealthController{
		svc: svc,
		log: log,
	}
}

// GetLatest Api Info
func (h *HealthController) GetLatest(c *gin.Context) {
	apiInfo, err := h.svc.GetLatest(c)

	switch {
	case err == nil:
		c.JSON(http.StatusOK, gin.H{"data": apiInfo})
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func (h *HealthController) ListAll(c *gin.Context) {
	apiInfos, err := h.svc.ListAll(c)

	switch {
	case err == nil:
		c.JSON(http.StatusOK, gin.H{"data": apiInfos})
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
