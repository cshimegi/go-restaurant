package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"users/shared/domain"
)

// ListUsers lists all users
func ListUsers(c *gin.Context) {
	firstName := "Test"
	lastName := "Me"
	nickname := ""
	user := domain.User{
		ID:       1,
		FirtName: &firstName,
		LastName: &lastName,
		Nickname: &nickname,
	}
	c.JSON(http.StatusOK, gin.H{
		"data": []domain.User{
			user,
		},
	})
}
