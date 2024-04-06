package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"alan/blog/users/shared/domain"
)

// UserStore is implementation of user store
type UserStore struct {
	DB  *gorm.DB
	log *logrus.Entry
}

// ListAll lists all users
func (u *UserStore) ListAll(c *gin.Context) ([]domain.User, error) {
	var users []domain.User
	result := u.DB.Find(&users)
	if result.Error != nil {
		u.log.WithField("error", result.Error).Error("Error listing users from database")
		return nil, result.Error
	}
	return users, nil
}
