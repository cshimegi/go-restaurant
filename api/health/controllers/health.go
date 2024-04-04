package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"alan/blog/health/domain"
)

// IHealthService is the interface for Health service
type IHealthService interface {
	Retrieve(c *gin.Context) (domain.ApiInfo, error)
	ListAll(c *gin.Context) ([]domain.ApiInfo, error)
}

// HealthController defines Health's controller
type HealthController struct {
	svc IHealthService
}

// NewHealthController defines api paths to Health service
func NewHealthController(svc IHealthService) HealthController {
	return HealthController{
		svc: svc,
	}
}

// Retrieve Api Info
func (h *HealthController) Retrieve(c *gin.Context) {
	apiInfo, err := h.svc.Retrieve(c)

	if err != nil {
		// TODO: use logrus
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": apiInfo,
	})
}

func (h *HealthController) ListAll(c *gin.Context) {
	apiInfos, err := h.svc.ListAll(c)

	if err != nil {
		// TODO: use logrus
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": apiInfos,
	})
}
