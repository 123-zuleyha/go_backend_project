package entity

import (
	"time"
)

type Lesson struct {
	ID         int
	Name       string
	Duration   int
	Definition string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
	UserTypeID int
	Teacher    user
	TeacherID  int
}
