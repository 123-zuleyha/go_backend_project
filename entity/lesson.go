package entity

import (
	"time"
)

type Lesson struct {
	ID         int       `json:"id" gorm:"primary key"`
	Name       string    `json:"name" validation:"required"`
	LessonCode int       `json:"lesson_code" validation:"required"`
	Definition string    `json:"definition"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
	Teacher    User      `json:"teacher" validation:"required"`
	TeacherID  int       `json:"teacher_id" validation:"required"`
}
