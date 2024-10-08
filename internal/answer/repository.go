package answer

import (
	"github.com/123-zuleyha/go_backend_project/db/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AnswerRepository struct {
	DB *gorm.DB
}

func NewAnswerRepository(db *gorm.DB) *AnswerRepository {
	return &AnswerRepository{DB: db}
}

func (r *AnswerRepository) GetAnswers(req *BaseRequest) ([]entity.Answer, error) {
	var answers []entity.Answer
	query := r.DB
	if req.Limit != 0 {
		query = query.Limit(req.Limit)
	}
	if req.Offset != 0 {
		query = query.Offset(req.Offset)
	}
	err := query.
		Preload("Choice").
		Preload("Question").
		Preload("User").
		Find(&answers).
		Error

	if err != nil {
		return nil, err
	}
	return answers, nil
}

func (r *AnswerRepository) GetAnswersByQuestionID(req *BaseRequest, questionID int) ([]entity.Answer, error) {
	var answers []entity.Answer
	query := r.DB
	if req.Limit != 0 {
		query = query.Limit(req.Limit)
	}
	if req.Offset != 0 {
		query = query.Offset(req.Offset)
	}
	err := query.
		Preload("Choice").
		Preload("Question").
		Preload("User").
		Where("question_id = ?", questionID).
		Find(&answers).
		Error
	if err != nil {
		return nil, err
	}
	return answers, nil
}

func (r *AnswerRepository) GetAnswerByID(id int) (entity.Answer, error) {
	var answer entity.Answer
	err := r.DB.
		Preload("Choice").
		Preload("Question").
		Preload("User").
		Where("id = ?", id).
		First(&answer).
		Error
	return answer, err
}

func (r *AnswerRepository) CreateAnswer(answerEntity *entity.Answer) (*entity.Answer, error) {
	err := r.DB.Create(&answerEntity).Error
	return answerEntity, err
}

func (r *AnswerRepository) UpdateAnswer(answerEntity entity.Answer) error {
	err := r.DB.Omit(clause.Associations).Updates(&answerEntity).Error
	return err
}

func (r *AnswerRepository) DeleteAnswer(id int) error {
	err := r.DB.Delete(&entity.Answer{}, id).Error
	return err
}
