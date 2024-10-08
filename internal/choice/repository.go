package choice

import (
	"github.com/123-zuleyha/go_backend_project/db/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ChoiceRepository struct {
	DB *gorm.DB
}

func NewChoiceRepository(db *gorm.DB) *ChoiceRepository {
	return &ChoiceRepository{DB: db}
}

func (r *ChoiceRepository) GetChoices(req *BaseRequest) ([]entity.Choice, error) {
	var choices []entity.Choice
	query := r.DB
	if req.Limit != 0 {
		query = query.Limit(req.Limit)
	}
	if req.Offset != 0 {
		query = query.Offset(req.Offset)
	}
	err := query.
		Preload("Question").
		Find(&choices).
		Error

	if err != nil {
		return nil, err
	}
	return choices, nil
}

func (r *ChoiceRepository) CreateChoice(choiceEntity *entity.Choice) (*entity.Choice, error) {
	err := r.DB.Omit(clause.Associations).Create(&choiceEntity).Error
	return choiceEntity, err
}

func (r *ChoiceRepository) UpdateChoice(choiceEntity entity.Choice) error {
	err := r.DB.Omit(clause.Associations).Updates(&choiceEntity).Error
	return err
}

func (r *ChoiceRepository) DeleteChoice(id int) error {
	err := r.DB.Delete(&entity.Choice{}, id).Error
	return err
}

func (r *ChoiceRepository) GetChoicesByQuestionID(req *BaseRequest, questionID int) ([]entity.Choice, error) {
	var choices []entity.Choice
	query := r.DB
	if req.Limit != 0 {
		query = query.Limit(req.Limit)
	}
	if req.Offset != 0 {
		query = query.Offset(req.Offset)
	}
	err := query.
		Preload("Question").
		Where("question_id = ?", questionID).
		Find(&choices).
		Error
	if err != nil {
		return nil, err
	}
	return choices, nil
}

func (r *ChoiceRepository) GetChoiceByID(id int) (entity.Choice, error) {
	var choice entity.Choice
	err := r.DB.
		Preload("Question").
		Preload("Question.Quiz").
		Where("id = ?", id).
		First(&choice).
		Error
	return choice, err
}

func (r *ChoiceRepository) GetQuestionsWithChoicesByQuizID(quizID int) ([]entity.Choice, []entity.Question, error) {
	var questions []entity.Question
	err := r.DB.
		Omit(clause.Associations).
		Joins("Quiz").
		Where("quiz_id = ?", quizID).
		Find(&questions).
		Error

	var choices []entity.Choice
	err = r.DB.
		Omit(clause.Associations).
		Joins("Question").
		Joins("Question.Quiz").
		Where("quiz_id = ?", quizID).
		Find(&choices).
		Error

	if err != nil {
		return nil, nil, err
	}
	return choices, questions, nil
}
