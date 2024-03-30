package services

import (
	"github.com/gin-gonic/gin"

	"alan/blog/posts/dao"
	"alan/blog/posts/shared/domain"
)

// PostService is implementation of post service
type PostService struct {
	store dao.IPostStore
}

// NewPostService is used to instantiate postService
func NewPostService(store dao.IPostStore) *PostService {
	return &PostService{
		store: store,
	}
}

// ListAll lists all posts
func (p PostService) ListAll(c *gin.Context) ([]domain.Post, error) {
	posts, err := p.store.ListAll(c)
	if err != nil {
		panic(err)
	}
	return posts, nil
}
