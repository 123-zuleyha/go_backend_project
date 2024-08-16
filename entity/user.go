package entity

import (
	"time"
)

type User struct {
	ID         int       `json:"id" gorm:"primary key "`
	FirstName  string    `json:"first_name" validation:"required"`
	LastName   string    `json:"last_name" validation:"required"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	Username   string    `json:"username"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
	UserType   UserType  `json:"user_type"`
	UserTypeID int       `json:"user_type_id"`
	Lessons    []Lesson  `json:"lessons" gorm:many2many:"user_lessons;"`
}
