package models

import "time"

type BaseModel struct {
	CreatedAt time.Time `json:"created_at" gorm:"->;<-:create;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}
