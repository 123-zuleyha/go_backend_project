package userAnswer

import "github.com/123-zuleyha/go_backend_project/db/entity"

type BaseRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type BaseResponse struct {
	BaseRequest
	Count int `json:"count"`
}

type CreateUserAnswerRequest struct {
	UserID   int `json:"user_id"`
	QuizID   int `json:"quiz_id"`
	AnswerID int `json:"answer_id"`
}

type UpdateUserAnswerRequest struct {
	AnswerID int `json:"answer_id"`
}

type UserAnswerResponseDTO struct {
	Data []UserAnswerDTO `json:"data"`
	BaseResponse
}

type UserAnswerDTO struct {
	ID     int           `json:"id"`
	User   entity.User   `json:"user"`
	Quiz   entity.Quiz   `json:"quiz"`
	Answer entity.Answer `json:"answer"`
}
