package services

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"alan/restaurant/appetizers/shared/domain"
)

// IAppetizerStore is interface of appetizer store
type IAppetizerStore interface {
	ListAll(c *gin.Context) ([]domain.Appetizer, error)
}

// AppetizerService is implementation of post service
type AppetizerService struct {
	store IAppetizerStore
	log   *logrus.Entry
}

// NewAppetizerService is used to instantiate appetizerService
func NewAppetizerService(store IAppetizerStore, log *logrus.Entry) *AppetizerService {
	return &AppetizerService{
		store: store,
		log:   log,
	}
}

// ListAll lists all appetizers
func (a *AppetizerService) ListAll(c *gin.Context) ([]domain.Appetizer, error) {
	appetizers, err := a.store.ListAll(c)
	if err != nil {
		a.log.WithField("error", err).Error("failed to list all appetizers")
		return nil, err
	}
	return appetizers, nil
}
