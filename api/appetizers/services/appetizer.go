package services

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"alan/restaurant/appetizers/dao"
	"alan/restaurant/appetizers/shared/domain"
)

// IAppetizerService is interface of post service
type IAppetizerService interface {
	Create(c *gin.Context, appetizer domain.Appetizer) (*domain.Appetizer, error)
	ListAll(c *gin.Context) ([]domain.Appetizer, error)
	GetById(c *gin.Context, id uint) (*domain.Appetizer, error)
	PatchById(c *gin.Context, id uint, appetizer domain.Appetizer) (*domain.Appetizer, error)
	DeleteById(c *gin.Context, id uint) error
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

func (a *AppetizerService) Create(c *gin.Context, appetizer domain.Appetizer) (*domain.Appetizer, error) {
	newAppetizer, err := a.store.Create(c, appetizer)
	switch {
	case err == nil:
		return newAppetizer, nil
	default:
		a.log.WithField("error", err.Error()).Error("Error creating appetizer")
		return nil, err
	}
}

// ListAll lists all appetizers
func (a *AppetizerService) ListAll(c *gin.Context) ([]domain.Appetizer, error) {
	appetizers, err := a.store.ListAll(c)

	switch {
	case err == nil:
		return appetizers, nil
	default:
		a.log.WithField("error", err.Error()).Error("Error listing all appetizers")
		return nil, err
	}
}

func (a *AppetizerService) GetById(c *gin.Context, id uint) (*domain.Appetizer, error) {
	appetizer, err := a.store.GetById(c, id)

	switch {
	case err == nil:
		return appetizer, nil
	default:
		a.log.WithField("error", err.Error()).Error("Error getting appetizer by Id")
		return nil, err
	}
}

func (a *AppetizerService) PatchById(c *gin.Context, id uint, appetizer domain.Appetizer) (*domain.Appetizer, error) {
	updatedAppetizer, err := a.store.PatchById(c, id, appetizer)

	switch {
	case err == nil:
		return updatedAppetizer, nil
	default:
		a.log.WithField("error", err.Error()).Error("Error patching appetizer by Id")
		return nil, err
	}
}

func (a *AppetizerService) DeleteById(c *gin.Context, id uint) error {
	err := a.store.DeleteById(c, id)

	switch {
	case err == nil:
		return nil
	default:
		a.log.WithField("error", err.Error()).Error("Error deleting appetizer by Id")
		return err
	}
}
