package domain

import (
	"time"

	userDomain "users/shared/domain"
)

// Post define user's post
type Post struct {
	ID        uint `json:"id" gorm:"primaryKey;"`
	UserID    uint
	User      userDomain.User `gorm:"foreignKey:UserID; references:ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Title     string          `json:"title"`
	Content   string          `json:"content"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
