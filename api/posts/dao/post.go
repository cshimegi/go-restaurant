package dao

import (
	"posts/shared/domain"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// IPostStore is interface of post store
type IPostStore interface {
	ListAll(c *gin.Context) ([]domain.Post, error)
}

// PostStore is implementation of post store
type PostStore struct {
	DB *gorm.DB
}

// ListAll lists all posts
func (p *PostStore) ListAll(c *gin.Context) ([]domain.Post, error) {
	var posts []domain.Post
	result := p.DB.Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}
