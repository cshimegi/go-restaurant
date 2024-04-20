package services

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"alan/restaurant/users/dao"
	"alan/restaurant/users/shared/domain"
)

// IUserService is the interface for user API
type IUserService interface {
	Create(c *gin.Context, user domain.User) (*domain.User, error)
	ListAll(c *gin.Context) ([]domain.User, error)
	GetById(c *gin.Context, id uint) (*domain.User, error)
	PatchById(c *gin.Context, id uint, user domain.User) (*domain.User, error)
	DeleteById(c *gin.Context, id uint) error
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

func (u *UserService) Create(c *gin.Context, user domain.User) (*domain.User, error) {
	newUser, err := u.store.Create(c, user)

	switch {
	case err == nil:
		return newUser, nil
	default:
		u.log.WithField("error", err.Error()).Error("failed to create user")
		return nil, err
	}
}

// ListAll lists all users
func (u *UserService) ListAll(c *gin.Context) ([]domain.User, error) {
	users, err := u.store.ListAll(c)

	switch {
	case err == nil:
		return users, nil
	default:
		u.log.WithField("error", err).Error("failed to list all users")
		return nil, err
	}
}

func (u *UserService) GetById(c *gin.Context, id uint) (*domain.User, error) {
	user, err := u.store.GetById(c, id)
	switch {
	case err == nil:
		return user, nil
	default:
		u.log.WithField("error", err).Error("failed to get user by id")
		return nil, err
	}
}

func (u *UserService) PatchById(c *gin.Context, id uint, user domain.User) (*domain.User, error) {
	updatedUser, err := u.store.PatchById(c, id, user)
	switch {
	case err == nil:
		return updatedUser, nil
	default:
		u.log.WithField("error", err).Error("failed to patch user by id")
		return nil, err
	}
}

func (u *UserService) DeleteById(c *gin.Context, id uint) error {
	err := u.store.DeleteById(c, id)
	switch {
	case err == nil:
		return nil
	default:
		u.log.WithField("error", err).Error("failed to delete user by id")
		return err
	}
}
