package entity

import (
	"time"

	"gorm.io/gorm"
)

type CodeSubmission struct {
	ID             int            `json:"id" gorm:"primaryKey"`
	Input          string         `json:"input"`
	ExpectedOutput string         `json:"expected_output"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`

	CodeID int  `json:"code_id"`
	Code   Code `json:"code" gorm:"foreignKey:CodeID"`
}
