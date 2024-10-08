package userAnswer

import (
	"github.com/123-zuleyha/go_backend_project/db/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserAnswerRepository struct {
	DB *gorm.DB
}

func NewUserAnswerRepository(db *gorm.DB) *UserAnswerRepository {
	return &UserAnswerRepository{DB: db}
}

func (r *UserAnswerRepository) GetUserAnswers(req *BaseRequest) ([]entity.UserAnswer, error) {
	var userAnswers []entity.UserAnswer
	query := r.DB
	if req.Limit != 0 {
		query = query.Limit(req.Limit)
	}
	if req.Offset != 0 {
		query = query.Offset(req.Offset)
	}
	err := query.Find(&userAnswers).Error

	if err != nil {
		return nil, err
	}
	return userAnswers, nil
}

func (r *UserAnswerRepository) CreateUserAnswer(userAnswerEntity *entity.UserAnswer) (*entity.UserAnswer, error) {
	err := r.DB.Create(&userAnswerEntity).Error
	return userAnswerEntity, err
}

func (r *UserAnswerRepository) UpdateUserAnswer(userAnswerEntity entity.UserAnswer) error {
	return r.DB.Omit(clause.Associations).Updates(&userAnswerEntity).Error
}

func (r *UserAnswerRepository) DeleteUserAnswer(id int) error {
	return r.DB.Omit(clause.Associations).Delete(&entity.UserAnswer{}, id).Error
}

func (r *UserAnswerRepository) GetUserAnswerByQuestionID(id int) (entity.UserAnswer, error) {
	var userAnswer entity.UserAnswer
	err := r.DB.
		Preload("User").
		Preload("Quiz").
		Preload("Answer").
		Where("question_id = ?", id).
		First(&userAnswer).
		Error
	return userAnswer, err
}

func (r *UserAnswerRepository) GetUserAnswerByID(id int) (entity.UserAnswer, error) {
	var userAnswer entity.UserAnswer
	err := r.DB.
		Preload("User").
		Preload("Quiz").
		Preload("Answer").
		Where("id = ?", id).
		First(&userAnswer).
		Error
	return userAnswer, err
}
