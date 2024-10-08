package choice

import (
	"errors"

	"github.com/123-zuleyha/go_backend_project/db/entity"
	"github.com/123-zuleyha/go_backend_project/pkg/utils"
)

type ChoiceService struct {
	repository ChoiceRepository
}

func NewChoiceService(repository *ChoiceRepository) *ChoiceService {
	return &ChoiceService{repository: *repository}
}

func (s *ChoiceService) GetChoices(req *BaseRequest) (*ChoiceResponseDTO, error) {
	choices, err := s.repository.GetChoices(req)
	if err != nil {
		return nil, err
	}
	choiceDTOs := []ChoiceDTO{}
	for i := range choices {
		choiceDTO := new(ChoiceDTO)
		err := utils.JSONtoDTO(choices[i], choiceDTO)

		if err != nil {
			return nil, errors.New("failed to convert choice entity to choice dto")
		}
		choiceDTOs = append(choiceDTOs, *choiceDTO)
	}

	var resultDTO ChoiceResponseDTO
	resultDTO.Count = len(choiceDTOs)
	resultDTO.Data = choiceDTOs

	return &resultDTO, nil
}

func (s *ChoiceService) CreateChoice(choiceDTO *CreateChoiceRequest) (*entity.Choice, error) {
	choiceEntity := new(entity.Choice)
	utils.DTOtoJSON(choiceDTO, choiceEntity)

	createdChoice, err := s.repository.CreateChoice(choiceEntity)
	if err != nil {
		return nil, err
	}
	return createdChoice, nil

}

func (s *ChoiceService) UpdateChoices(choiceDTO *UpdateChoiceRequest) (*entity.Choice, error) {
	choiceEntity := new(entity.Choice)
	utils.DTOtoJSON(choiceDTO, choiceEntity)
	err := s.repository.UpdateChoice(*choiceEntity)
	if err != nil {
		return nil, err
	}
	return choiceEntity, nil
}

func (s *ChoiceService) DeleteChoices(id int) error {
	return s.repository.DeleteChoice(id)
}

func (s *ChoiceService) GetChoicesByQuestionID(req *BaseRequest, questionID int) (*ChoiceResponseDTO, error) {
	choices, err := s.repository.GetChoicesByQuestionID(req, questionID)
	if err != nil {
		return nil, err
	}
	choiceDTOs := []ChoiceDTO{}
	for i := range choices {
		choiceDTO := new(ChoiceDTO)
		err := utils.JSONtoDTO(choices[i], choiceDTO)
		if err != nil {
			return nil, errors.New("failed to convert choice entity to choice dto")
		}
		choiceDTOs = append(choiceDTOs, *choiceDTO)
	}
	var resultDTO ChoiceResponseDTO
	resultDTO.Count = len(choiceDTOs)
	resultDTO.Data = choiceDTOs
	return &resultDTO, nil
}

func (s *ChoiceService) GetChoiceByID(id int) (*ChoiceDTO, error) {
	choice, err := s.repository.GetChoiceByID(id)
	if err != nil {
		return nil, err
	}
	choiceDTO := new(ChoiceDTO)
	err = utils.JSONtoDTO(choice, choiceDTO)
	if err != nil {
		return nil, errors.New("failed to convert choice entity to choice dto")
	}
	return choiceDTO, nil
}

func (r *ChoiceService) GetQuestionsWithChoicesByQuizID(quizID int) ([]Data, error) {
	choices, questions, err := r.repository.GetQuestionsWithChoicesByQuizID(quizID)
	if err != nil {
		return nil, err
	}
	data := []Data{}

	for i := range questions {
		questionDTO := new(QuestionDTO)
		err := utils.JSONtoDTO(questions[i], questionDTO)
		if err != nil {
			return nil, errors.New("failed to convert question entity to question dto")
		}
		data = append(data, Data{Question: *questionDTO})
	}

	for i := range choices {
		for j := range data {
			choiceDTO := new(ChoiceDTO)
			err := utils.JSONtoDTO(choices[i], choiceDTO)
			if err != nil {
				return nil, errors.New("failed to convert choice entity to choice dto")
			}
			if choiceDTO.QuestionID == data[j].Question.ID {
				data[j].Question.Choices = append(data[j].Question.Choices, *choiceDTO)
			}
			choiceDTO = nil
		}
	}

	return data, nil
}
