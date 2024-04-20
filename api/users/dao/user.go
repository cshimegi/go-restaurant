package dao

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	sc "alan/restaurant/shared/consts"
	"alan/restaurant/users/shared/domain"
)

// UserStore is implementation of user store
type UserStore struct {
	DB  *gorm.DB
	log *logrus.Entry
}

func (u *UserStore) Create(c *gin.Context, user domain.User) (*domain.User, error) {
	result := u.DB.Create(&user)

	switch {
	case result.Error == nil:
		return u.GetById(c, user.ID)
	default:
		u.log.WithField("error", result.Error.Error()).Error("Error creating user")
		return nil, sc.ErrInternal
	}
}

// ListAll lists all users
func (u *UserStore) ListAll(c *gin.Context) ([]domain.User, error) {
	var users []domain.User
	result := u.DB.Find(&users)

	switch {
	case result.Error == nil:
		return users, nil
	default:
		u.log.WithField("error", result.Error.Error()).Error("Error listing users from database")
		return nil, sc.ErrInternal
	}
}

func (u *UserStore) GetById(c *gin.Context, id uint) (*domain.User, error) {
	var user domain.User
	result := u.DB.First(&user, id)

	switch {
	case result.Error == nil:
		return &user, nil
	case errors.Is(result.Error, gorm.ErrRecordNotFound):
		return nil, sc.ErrNotFound
	default:
		u.log.WithFields(
			logrus.Fields{"error": result.Error.Error(), "id": id},
		).Error("Error getting user by id")
		return nil, sc.ErrInternal
	}
}

func (u *UserStore) PatchById(c *gin.Context, id uint, user domain.User) (*domain.User, error) {
	var updatedUser domain.User
	result := u.DB.Model(&updatedUser).Where("id = ?", id).Updates(&user)

	switch {
	case result.Error == nil:
		return &updatedUser, nil
	case errors.Is(result.Error, gorm.ErrRecordNotFound):
		return nil, sc.ErrNotFound
	default:
		u.log.WithFields(
			logrus.Fields{"error": result.Error.Error(), "id": id},
		).Error("Error patching user by id")
		return nil, sc.ErrInternal
	}
}

func (u *UserStore) DeleteById(c *gin.Context, id uint) error {
	result := u.DB.Delete(&domain.User{}, id)

	switch {
	case result.Error == nil:
		return nil
	case errors.Is(result.Error, gorm.ErrRecordNotFound):
		return sc.ErrNotFound
	default:
		u.log.WithFields(
			logrus.Fields{"error": result.Error.Error(), "id": id},
		).Error("Error deleting user by id")
		return sc.ErrInternal
	}
}
