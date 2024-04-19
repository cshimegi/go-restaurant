package services

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"alan/restaurant/users/dao"
	"alan/restaurant/users/shared/domain"
)

// IUserService is the interface for user API
type IUserService interface {
	ListAll(c *gin.Context) ([]domain.User, error)
}

// UserService is implementation of user service
type UserService struct {
	store dao.IUserStore
	log   *logrus.Entry
}

// NewUserService is implementation of user service
func NewUserService(store dao.IUserStore, log *logrus.Entry) IUserService {
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
