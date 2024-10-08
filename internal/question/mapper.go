package question

import "github.com/123-zuleyha/go_backend_project/db/entity"

type BaseRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type BaseResponse struct {
	BaseRequest
	Count int `json:"count"`
}

type CreateQuestionRequest struct {
	Text  string `json:"text"`
	Type  int    `json:"type"`
	Point int    `json:"point"`

	QuizID int `json:"quiz_id"`
}

type UpdateQuestionRequest struct {
	ID    int    `json:"id"`
	Text  string `json:"text"`
	Type  int    `json:"type"`
	Point int    `json:"point"`
}

type QuestionResponseDTO struct {
	Data []QuestionDTO `json:"data"`
	BaseResponse
}

type QuestionDTO struct {
	ID    int         `json:"id"`
	Text  string      `json:"text"`
	Type  int         `json:"type"`
	Point int         `json:"point"`
	Quiz  entity.Quiz `json:"quiz"`
}

type QuestionWithChoicesDTO struct {
	ID          int    `json:"id"`
	Text        string `json:"text"`
	Type        int    `json:"type"`
	Point       int    `json:"point"`
	IsCorrect   bool   `json:"is_correct"`
	QuizID      int    `json:"quiz_id"`
	QuestionID  int    `json:"question_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	TeacherID   int    `json:"teacher_id"`
	LessonID    int    `json:"lesson_id"`
}
