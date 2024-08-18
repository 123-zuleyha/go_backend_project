package entity

import (
	"time"

	"gorm.io/gorm"
)

type Quiz struct {
	ID          int            `json:"id" gorm:"primary key"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Duration    int            `json:"duration"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	StartTime   time.Time      `json:"start_time"`
	EndTime     time.Time      `json:"end_time"`

	TeacherID int  `json:"teacher_id"`
	Teacher   User `json:"teacher" gorm:"foreignKey:TeacherID"`

	LessonID int    `json:"lesson_id"`
	Lesson   Lesson `json:"lesson" gorm:"foreignKey:LessonID"`
}
