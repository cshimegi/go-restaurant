package domain

import (
	"time"
)

// User defines user model
type User struct {
	ID        uint   `json:"id" gorm:"primaryKey;"`
	FirtName  string `json:"first_name" validate:"required,first_name" gorm:"not null"`
	LastName  string `json:"last_name" validate:"required,last_name" gorm:"not null"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email" validate:"required,email" gorm:"unique;not null"`
	Password  string `json:"password" validate:"required" gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
