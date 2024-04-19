package services

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"alan/restaurant/appetizers/dao"
	"alan/restaurant/appetizers/shared/domain"
)

// IAppetizerService is interface of post service
type IAppetizerService interface {
	ListAll(c *gin.Context) ([]domain.Appetizer, error)
}

// AppetizerService is implementation of post service
type AppetizerService struct {
	store dao.IAppetizerStore
	log   *logrus.Entry
}

// NewAppetizerService is used to instantiate appetizerService
func NewAppetizerService(store dao.IAppetizerStore, log *logrus.Entry) IAppetizerService {
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
