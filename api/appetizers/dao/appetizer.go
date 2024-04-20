package dao

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"alan/restaurant/appetizers/shared/domain"
	sc "alan/restaurant/shared/consts"
)

// AppetizerStore is implementation of post store
type AppetizerStore struct {
	DB  *gorm.DB
	log *logrus.Entry
}

func (a *AppetizerStore) Create(c *gin.Context, appetizer domain.Appetizer) (*domain.Appetizer, error) {
	result := a.DB.Create(&appetizer)

	switch {
	case result.Error == nil:
		return a.GetById(c, appetizer.ID)
	default:
		a.log.WithField("error", result.Error.Error()).Error("Error creating appetizer")
		return nil, sc.ErrInternal
	}
}

// ListAll lists all appetizers
func (a *AppetizerStore) ListAll(c *gin.Context) ([]domain.Appetizer, error) {
	var appetizers []domain.Appetizer
	result := a.DB.Find(&appetizers)

	switch {
	case result.Error == nil:
		return appetizers, nil
	default:
		a.log.WithField(
			"error", result.Error.Error(),
		).Error("Error listing appetizers from database")
		return nil, sc.ErrInternal
	}
}

func (a *AppetizerStore) GetById(c *gin.Context, id uint) (*domain.Appetizer, error) {
	var appetizer domain.Appetizer
	result := a.DB.First(&appetizer, id)

	switch {
	case result.Error == nil:
		return &appetizer, nil
	case errors.Is(result.Error, gorm.ErrRecordNotFound):
		return nil, sc.ErrNotFound
	default:
		a.log.WithFields(
			logrus.Fields{"error": result.Error.Error(), "id": id},
		).Error("Error getting appetizer by id")
		return nil, sc.ErrInternal
	}
}

func (a *AppetizerStore) PatchById(c *gin.Context, id uint, appetizer domain.Appetizer) (*domain.Appetizer, error) {
	var updatedAppetizer domain.Appetizer
	result := a.DB.Model(&updatedAppetizer).Where("id = ?", id).Updates(appetizer)

	switch {
	case result.Error == nil:
		return &updatedAppetizer, nil
	case errors.Is(result.Error, gorm.ErrRecordNotFound):
		return nil, sc.ErrNotFound
	default:
		a.log.WithFields(
			logrus.Fields{"error": result.Error.Error(), "id": id},
		).Error("Error patching appetizer by id")
		return nil, sc.ErrInternal
	}
}

func (a *AppetizerStore) DeleteById(c *gin.Context, id uint) error {
	result := a.DB.Delete(&domain.Appetizer{}, id)

	switch {
	case result.Error == nil:
		return nil
	case errors.Is(result.Error, gorm.ErrRecordNotFound):
		return sc.ErrNotFound
	default:
		a.log.WithFields(
			logrus.Fields{"error": result.Error.Error(), "id": id},
		).Error("Error deleting appetizer by id")
		return sc.ErrInternal
	}
}
