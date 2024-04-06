package services

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"alan/blog/posts/shared/domain"
)

// IPostStore is interface of post store
type IPostStore interface {
	ListAll(c *gin.Context) ([]domain.Post, error)
}

// PostService is implementation of post service
type PostService struct {
	store IPostStore
	log   *logrus.Entry
}

// NewPostService is used to instantiate postService
func NewPostService(store IPostStore, log *logrus.Entry) *PostService {
	return &PostService{
		store: store,
		log:   log,
	}
}

// ListAll lists all posts
func (p *PostService) ListAll(c *gin.Context) ([]domain.Post, error) {
	posts, err := p.store.ListAll(c)
	if err != nil {
		p.log.WithField("error", err).Error("failed to list all posts")
		return nil, err
	}
	return posts, nil
}
