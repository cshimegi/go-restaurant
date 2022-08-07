package domain

import (
	"time"

	userDomain "users/shared/domain"
)

// Post define user's post
type Post struct {
	ID        uint `json:"id" gorm:"primaryKey;"`
	UserID    uint
	User      userDomain.User `validate:"required,user" gorm:"foreignKey:UserID; references:ID; constraint:OnDelete:CASCADE;"`
	Title     string          `json:"title" validate:"required,title" gorm:"not null"`
	Content   string          `json:"content" validate:"required,content" gorm:"not null"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
}
