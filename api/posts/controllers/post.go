package controllers

import (
	"alan/blog/posts/shared/domain"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// IPostService is interface of post service
type IPostService interface {
	ListAll(c *gin.Context) ([]domain.Post, error)
}

// PostController defines type of post controller
type PostController struct {
	svc IPostService
}

// NewPostController defines api paths to post service
func NewPostController(svc IPostService) PostController {
	return PostController{
		svc: svc,
	}
}

// ListAll lists all posts
func (p *PostController) ListAll(c *gin.Context) {
	posts, err := p.svc.ListAll(c)

	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": posts,
	})
}
