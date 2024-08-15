package entity

import (
	"time"
)

type User struct {
	ID        int       `json:"id" gorm:"primary key "`
	FirstName string    `json:"first_name" validation:"required"`
	LastName  string    `json:"last_name" validation:"required"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	UserType  Student   `json:"user_type"`
	Lesson    string    `json:"lesson"`
}
