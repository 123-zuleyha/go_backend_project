package user

import (
	"time"

	"github.com/123-zuleyha/go_backend_project/internal/lesson"
)

type PaginatedUserResponse struct {
	Count int       `json:"count"`
	Data  []UserDTO `json:"data"`
}

type UserDTO struct {
	ID        int          `json:"id"`
	FirstName string       `json:"firstname"`
	LastName  string       `json:"lastname"`
	Username  string       `json:"username"`
	Email     string       `json:"email"`
	UserType  *UserTypeDTO `json:"user_types"`
	CreatedAt time.Time    `json:"createdat"`
	UpdatedAt time.Time    `json:"updatedat"`

	Lessons []lesson.LessonDTO `json:"lessons"`
}

type UserTypeDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateUserRequest struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

type UpdateUserRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}
