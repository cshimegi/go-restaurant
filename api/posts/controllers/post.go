package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"alan/blog/posts/shared/domain"
)

// IPostService is interface of post service
type IPostService interface {
	ListAll(c *gin.Context) ([]domain.Post, error)
}

// PostController defines type of post controller
type PostController struct {
	svc IPostService
	log *logrus.Entry
}

// NewPostController defines api paths to post service
func NewPostController(svc IPostService, log *logrus.Entry) PostController {
	return PostController{
		svc: svc,
		log: log,
	}
}

// ListAll lists all posts
func (p *PostController) ListAll(c *gin.Context) {
	posts, err := p.svc.ListAll(c)

	if err != nil {
		p.log.WithField("error", err).Error("Error listing posts")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": posts,
	})
}
