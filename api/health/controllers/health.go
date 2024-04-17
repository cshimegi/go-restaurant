package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"alan/restaurant/health/shared/domain"
)

// IHealthService is the interface for Health service
type IHealthService interface {
	Retrieve(c *gin.Context) (domain.ApiInfo, error)
	ListAll(c *gin.Context) ([]domain.ApiInfo, error)
}

// HealthController defines Health's controller
type HealthController struct {
	svc IHealthService
	log *logrus.Entry
}

// NewHealthController defines api paths to Health service
func NewHealthController(svc IHealthService, log *logrus.Entry) HealthController {
	return HealthController{
		svc: svc,
		log: log,
	}
}

// Retrieve Api Info
func (h *HealthController) Retrieve(c *gin.Context) {
	apiInfo, err := h.svc.Retrieve(c)

	if err != nil {
		h.log.WithField("error", err).Error("Error retrieving api info")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": apiInfo,
	})
}

func (h *HealthController) ListAll(c *gin.Context) {
	apiInfos, err := h.svc.ListAll(c)

	if err != nil {
		h.log.WithField("error", err).Error("Error listing api infos")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": apiInfos,
	})
}
