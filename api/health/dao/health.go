package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"alan/restaurant/health/shared/domain"
)

// HealthStore is implementation of user store
type HealthStore struct {
	DB  *gorm.DB
	log *logrus.Entry
}

func (h *HealthStore) GetLatest(c *gin.Context) (*domain.ApiInfo, error) {
	var apiInfo *domain.ApiInfo
	result := h.DB.Last(&apiInfo)

	switch {
	case result.Error == nil:
		return apiInfo, nil
	default:
		return nil, result.Error
	}
}

func (h *HealthStore) ListAll(c *gin.Context) ([]domain.ApiInfo, error) {
	var apiInfos []domain.ApiInfo
	result := h.DB.Find(&apiInfos)

	switch {
	case result.Error == nil:
		return apiInfos, nil
	default:
		return nil, result.Error
	}
}
