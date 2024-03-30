package services

import (
	"github.com/gin-gonic/gin"

	"alan/blog/users/dao"
	"alan/blog/users/shared/domain"
)

// UserService is implementation of user service
type UserService struct {
	store dao.IUserStore
}

// NewUserService is implementation of user service
func NewUserService(store dao.IUserStore) *UserService {
	return &UserService{
		store: store,
	}
}

// ListAll lists all users
func (u *UserService) ListAll(c *gin.Context) ([]domain.User, error) {
	users, err := u.store.ListAll(c)
	if err != nil {
		panic(err)
	}
	return users, nil
}
