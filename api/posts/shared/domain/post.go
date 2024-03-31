package domain

import (
	sm "alan/blog/shared/models"
	ud "alan/blog/users/shared/domain"
)

// Post define user's post
type Post struct {
	ID           uint    `json:"id" gorm:"primaryKey;not null"`
	UserID       uint    `json:"user_id" gorm:"user_id"`
	User         ud.User `json:"user" gorm:"references:ID"`
	Title        string  `json:"title" validate:"required"`
	Content      string  `json:"content" validate:"required"`
	sm.BaseModel `gorm:"embedded"`
}
