package entity

import (
	"time"

	"gorm.io/gorm"
)

type Question struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	Text      string         `json:"text"`
	Type      int            `json:"type"`
	Point     int            `json:"point"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	QuizID    int             `json:"quiz_id"`
	Quiz      Quiz            `json:"quiz" gorm:"foreignKey:QuizID"`
}
