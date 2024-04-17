package domain

import (
	sm "alan/restaurant/shared/models"
)

// Appetizer define restaurant's appetizer
type Appetizer struct {
	ID           uint   `json:"id" gorm:"primaryKey;not null"`
	Title        string `json:"title" validate:"required"`
	Price        uint   `json:"price" validate:"required"`
	sm.BaseModel `gorm:"embedded"`
}
