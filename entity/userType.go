package entity

import (
	"time"
	"gorm.io/gorm"
)

type UserType struct {
	ID        int       `gorm:"primary key"`
	Name      string    `json:"name" validation:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
