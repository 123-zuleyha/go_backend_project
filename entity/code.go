package entity

import (
	"time"

	"gorm.io/gorm"
)

type Code struct {
	ID              int            `json:"id" gorm:"primaryKey"`
	Question        string         `json:"question"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
}
