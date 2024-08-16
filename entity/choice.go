package entity

import (
	"time"

	"gorm.io/gorm"
)

type Choice struct {
	ID             int            `json:"id" gorm:"primaryKey"`
	Text           string         `json:"text"`
	IsCorrect      bool           `json:"is_correct"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`

	QuestionID     int            `json:"question_id"`
	Question       Question       `json:"question" gorm:"foreignKey:QuestionID"`
}
