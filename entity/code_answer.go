package entity

import (
	"time"

	"gorm.io/gorm"
)

type CodeAnswer struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	Status    string         `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	UserID int  `json:"user_id"`
	User   User `json:"user" gorm:"foreignKey:UserID"`

	CodeID int  `json:"code_id"`
	Code   Code `json:"code" gorm:"foreignKey:UserID"`
}
