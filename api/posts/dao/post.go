package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"alan/blog/posts/shared/domain"
)

// PostStore is implementation of post store
type PostStore struct {
	DB  *gorm.DB
	log *logrus.Entry
}

// ListAll lists all posts
func (p *PostStore) ListAll(c *gin.Context) ([]domain.Post, error) {
	var posts []domain.Post
	result := p.DB.Find(&posts)
	if result.Error != nil {
		p.log.WithField("error", result.Error).Error("Error listing posts from database")
		return nil, result.Error
	}
	return posts, nil
}
