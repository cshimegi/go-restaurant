package services

import (
	"github.com/gin-gonic/gin"

	"users/shared/domain"
)

// UserService is interface of user service
type UserService interface {
	ListAll(c *gin.Context) ([]domain.User, error)
}

// UserServiceImpl is implementation of user service
type UserServiceImpl struct{}

// NewUserService is implementaion of user service
func NewUserService() *UserServiceImpl {
	return &UserServiceImpl{}
}

// ListAll lists all users
func (u UserServiceImpl) ListAll(c *gin.Context) ([]domain.User, error) {
	user := domain.User{
		ID:       1,
		FirtName: "Test",
		LastName: "Me",
		Nickname: "",
	}

	return []domain.User{
		user,
	}, nil
}
