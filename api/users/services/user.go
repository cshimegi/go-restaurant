package services

import (
	"github.com/gin-gonic/gin"

	"users/shared/domain"
)

// UserService is struct of user service
type UserService interface {
	ListAll(c *gin.Context) ([]domain.User, error)
}

// userServiceImpl is struct of user service
type userServiceImpl struct{}

// NewUserService is used to instantiate userService
func NewUserService() UserService {
	return &userServiceImpl{}
}

// ListAll lists all users
func (u userServiceImpl) ListAll(c *gin.Context) ([]domain.User, error) {
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
