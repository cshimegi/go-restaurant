package services

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"alan/restaurant/users/shared/domain"
)

// IUserStore is interface of user store
type IUserStore interface {
	ListAll(c *gin.Context) ([]domain.User, error)
}

// UserService is implementation of user service
type UserService struct {
	store IUserStore
	log   *logrus.Entry
}

// NewUserService is implementation of user service
func NewUserService(store IUserStore, log *logrus.Entry) *UserService {
	return &UserService{
		store: store,
		log:   log,
	}
}

// ListAll lists all users
func (u *UserService) ListAll(c *gin.Context) ([]domain.User, error) {
	users, err := u.store.ListAll(c)
	if err != nil {
		u.log.WithField("error", err).Error("failed to list all users")
		return nil, err
	}
	return users, nil
}
