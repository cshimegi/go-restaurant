package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"alan/restaurant/appetizers/shared/domain"
)

// AppetizerStore is implementation of post store
type AppetizerStore struct {
	DB  *gorm.DB
	log *logrus.Entry
}

// ListAll lists all appetizers
func (a *AppetizerStore) ListAll(c *gin.Context) ([]domain.Appetizer, error) {
	var appetizers []domain.Appetizer
	result := a.DB.Find(&appetizers)
	if result.Error != nil {
		a.log.WithField("error", result.Error).Error("Error listing appetizers from database")
		return nil, result.Error
	}
	return appetizers, nil
}
