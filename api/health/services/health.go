package services

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"alan/restaurant/health/dao"
	"alan/restaurant/health/shared/domain"
)

// IHealthService is the interface for Health service
type IHealthService interface {
	GetLatest(c *gin.Context) (*domain.ApiInfo, error)
	ListAll(c *gin.Context) ([]domain.ApiInfo, error)
}

// HealthService is implementation of user service
type HealthService struct {
	store dao.IHealthStore
	log   *logrus.Entry
}

// NewHealthService is implementation of user service
func NewHealthService(store dao.IHealthStore, log *logrus.Entry) IHealthService {
	return &HealthService{
		store: store,
		log:   log,
	}
}

// GetLatest Api Information
func (h *HealthService) GetLatest(c *gin.Context) (*domain.ApiInfo, error) {
	apiInfo, err := h.store.GetLatest(c)

	switch {
	case err == nil:
		return apiInfo, nil
	default:
		h.log.WithField("error", err.Error()).Error("Error getting latest api info")
		return apiInfo, err
	}
}

func (h *HealthService) ListAll(c *gin.Context) ([]domain.ApiInfo, error) {
	apiInfo, err := h.store.ListAll(c)

	switch {
	case err == nil:
		return apiInfo, nil
	default:
		h.log.WithField("error", err.Error()).Error("Error listing all api infos")
		return apiInfo, err
	}
}
