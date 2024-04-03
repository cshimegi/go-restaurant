package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"alan/blog/health/dao"
	"alan/blog/health/services"
)

var (
	HealthService *services.HealthService
)

// HealthAPI is the interface for Health API
type HealthAPI interface {
	Retrieve(c *gin.Context)
	ListAll(c *gin.Context)
}

// HealthController defines Health's controller
type HealthController struct{}

func init() {
	HealthStore := dao.NewHealthStore()
	HealthService = services.NewHealthService(HealthStore)
}

// NewHealthAPI defines api paths to Health service
func NewHealthAPI() *HealthController {
	return &HealthController{}
}

// Retrieve Api Info
func (u HealthController) Retrieve(c *gin.Context) {
	apiInfo, err := HealthService.Retrieve(c)

	if err != nil {
		// TODO: use logrus
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": apiInfo,
	})
}

func (u HealthController) ListAll(c *gin.Context) {
	apiInfos, err := HealthService.ListAll(c)

	if err != nil {
		// TODO: use logrus
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": apiInfos,
	})
}
