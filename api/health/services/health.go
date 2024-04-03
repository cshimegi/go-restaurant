package services

import (
	"github.com/gin-gonic/gin"

	"alan/blog/health/dao"
	"alan/blog/health/domain"
)

// HealthService is implementation of user service
type HealthService struct {
	store dao.IHealthStore
}

// NewHealthService is implementation of user service
func NewHealthService(store dao.IHealthStore) *HealthService {
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
