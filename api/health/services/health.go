package services

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	
	"alan/blog/health/shared/domain"
)

// IHealthStore is interface of health store
type IHealthStore interface {
	Retrieve(c *gin.Context) (domain.ApiInfo, error)
	ListAll(c *gin.Context) ([]domain.ApiInfo, error)
}

// HealthService is implementation of user service
type HealthService struct {
	store IHealthStore
	log   *logrus.Entry
}

// NewHealthService is implementation of user service
func NewHealthService(store IHealthStore, log *logrus.Entry) *HealthService {
	return &HealthService{
		store: store,
		log:   log,
	}
}

// Retrieve Api Information
func (h *HealthService) Retrieve(c *gin.Context) (domain.ApiInfo, error) {
	apiInfo, err := h.store.Retrieve(c)
	if err != nil {
		h.log.WithField("error", err).Error("Error retrieving api info")
		return apiInfo, err
	}
	return apiInfo, nil
}

func (h *HealthService) ListAll(c *gin.Context) ([]domain.ApiInfo, error) {
	apiInfo, err := h.store.ListAll(c)
	if err != nil {
		h.log.WithField("error", err).Error("Error retrieving api info")
		return apiInfo, err
	}
	return apiInfo, nil
}
