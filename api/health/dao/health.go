package dao

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"alan/blog/health/domain"
)

// IHealthStore is interface of health store
type IHealthStore interface {
	Retrieve(c *gin.Context) (domain.ApiInfo, error)
	ListAll(c *gin.Context) ([]domain.ApiInfo, error)
}

// HealthStore is implementation of user store
type HealthStore struct {
	DB *gorm.DB
}

func (h *HealthStore) Retrieve(c *gin.Context) (domain.ApiInfo, error) {
	var apiInfo domain.ApiInfo
	result := h.DB.Order("release_at desc").First(&apiInfo)
	if result.Error != nil {
		return apiInfo, result.Error
	}
	return apiInfo, nil
}

func (h *HealthStore) ListAll(c *gin.Context) ([]domain.ApiInfo, error) {
	var apiInfos []domain.ApiInfo
	result := h.DB.Find(&apiInfos)
	if result.Error != nil {
		return apiInfos, result.Error
	}
	return apiInfos, nil
}
