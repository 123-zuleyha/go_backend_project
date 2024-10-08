package answer

import "github.com/123-zuleyha/go_backend_project/db/entity"

type BaseRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type BaseResponse struct {
	BaseRequest
	Count int `json:"count"`
}

type CreateAnswerRequest struct {
	Text string `json:"text"`

	QuestionID int `json:"question_id"`
	UserID     int `json:"user_id"`
}

type UpdateAnswerRequest struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type AnswerResponseDTO struct {
	Data []AnswerDTO `json:"data"`
	BaseResponse
}

type AnswerDTO struct {
	ID       int             `json:"id"`
	Text     string          `json:"text"`
	Question entity.Question `json:"question"`
	User     entity.User     `json:"user"`
}
