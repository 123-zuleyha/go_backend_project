package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserAnswer struct {
	ID           int                   `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time             `json:"created_at"`
	UpdatedAt    time.Time             `json:"updated_at"`
	DeletedAt    gorm.DeletedAt        `json:"deleted_at"`

	UserID       int                      `json:"user_id"`
	User         User                     `json:"user" gorm:"foreignKey:UserID"`

	QuizID       int                      `json:"quiz_id"`
	Quiz         Quiz                     `json:"quiz" gorm:"foreignKey:QuizID"`

	AnswerID     int                    `json:"answer_id"`
	Answer       Answer                  `json:"answer" gorm:"foreignKey:AnswerID"`
}
