package services

import (
	"github.com/gin-gonic/gin"

	"posts/shared/domain"
	userDomain "users/shared/domain"
)

// PostService is interface of postServiceImpl
type PostService interface {
	ListAll(c *gin.Context) ([]domain.Post, error)
}

// PostServiceImpl is implementation of post service
type PostServiceImpl struct{}

// NewPostService is used to instantiate postService
func NewPostService() *PostServiceImpl {
	return &PostServiceImpl{}
}

// ListAll lists all posts
func (p PostServiceImpl) ListAll(c *gin.Context) ([]domain.Post, error) {
	post := domain.Post{
		ID:     1,
		UserID: 1,
		User: userDomain.User{
			ID:       1,
			FirtName: "Test",
			LastName: "Me",
			Nickname: "",
		},
	}

	return []domain.Post{
		post,
	}, nil
}
