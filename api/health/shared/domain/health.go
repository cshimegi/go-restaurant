package domain

import "time"

// ApiInfo defines api info model
type ApiInfo struct {
	Version   string    `json:"version" gorm:"not null"`
	ReleaseAt time.Time `json:"release_at" gorm:"->;<-:create;not null"`
}
