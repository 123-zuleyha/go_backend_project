package entity

import (
	"time"

	"gorm.io/gorm"
)

type Answer struct {
	ID        int            `json:"id"`
	Text      string         `json:"text"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	QuestionID int      `json:"question_id"`
	Question   Question `json:"question" gorm:"foreignKey:QuestionID"`

	UserID int  `json:"user_id"`
	User   User `json:"user" gorm:"foreignKey:UserID"`
}
