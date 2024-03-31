package domain

import (
	sm "alan/blog/shared/models"
)

// User defines user model
type User struct {
	ID           uint   `json:"id" gorm:"primaryKey;not null"`
	FirstName    string `json:"first_name" validate:"required" gorm:"not null"`
	LastName     string `json:"last_name" validate:"required" gorm:"not null"`
	Nickname     string `json:"nickname"`
	Email        string `json:"email" gorm:"unique;not null"`
	Password     string `json:"password" validate:"required" gorm:"not null"`
	sm.BaseModel `gorm:"embedded"`
}
