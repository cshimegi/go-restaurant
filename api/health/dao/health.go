package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"alan/blog/health/shared/domain"
)

// HealthStore is implementation of user store
type HealthStore struct {
	DB  *gorm.DB
	log *logrus.Entry
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
