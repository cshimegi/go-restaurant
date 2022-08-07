package dao

import (
	"users/shared/domain"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// IUserStore is interface of user store
type IUserStore interface {
	ListAll(c *gin.Context) ([]domain.User, error)
}

// UserStore is implementation of user store
type UserStore struct {
	DB *gorm.DB
}

// ListAll lists all users
func (u *UserStore) ListAll(c *gin.Context) ([]domain.User, error) {
	var users []domain.User
	result := u.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
