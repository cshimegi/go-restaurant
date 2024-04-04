package services

import (
	"github.com/gin-gonic/gin"

	"alan/blog/health/domain"
)

// IHealthStore is interface of health store
type IHealthStore interface {
	Retrieve(c *gin.Context) (domain.ApiInfo, error)
	ListAll(c *gin.Context) ([]domain.ApiInfo, error)
}

// HealthService is implementation of user service
type HealthService struct {
	store IHealthStore
}

// NewHealthService is implementation of user service
func NewHealthService(store IHealthStore) *HealthService {
	return &HealthService{
		store: store,
	}
}

// Retrieve Api Information
func (h *HealthService) Retrieve(c *gin.Context) (domain.ApiInfo, error) {
	apiInfo, err := h.store.Retrieve(c)
	if err != nil {
		return apiInfo, err
	}
	return apiInfo, nil
}

func (h *HealthService) ListAll(c *gin.Context) ([]domain.ApiInfo, error) {
	apiInfo, err := h.store.ListAll(c)
	if err != nil {
		return apiInfo, err
	}
	return apiInfo, nil
}
