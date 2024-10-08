package codeSubmission

import (
	"errors"

	"github.com/123-zuleyha/go_backend_project/db/entity"
	"github.com/123-zuleyha/go_backend_project/pkg/utils"
)

type CodeSubmissionService struct {
	repository CodeSubmissionRepository
}

func NewCodeSubmissionService(repository *CodeSubmissionRepository) *CodeSubmissionService {
	return &CodeSubmissionService{repository: *repository}
}

func (s *CodeSubmissionService) GetCodeSubmissions(req *BaseRequest) (*CodeSubmissionResponseDTO, error) {
	codeSubmissions, err := s.repository.GetCodeSubmissions(req)
	if err != nil {
		return nil, err
	}
	codeSubmissionDTOs := []CodeSubmissionDTO{}
	for i := range codeSubmissions {
		codeSubmissionDTO := new(CodeSubmissionDTO)
		err := utils.JSONtoDTO(codeSubmissions[i], codeSubmissionDTO)

		if err != nil {
			return nil, errors.New("failed to convert code submission entity to code submission dto")
		}
		codeSubmissionDTOs = append(codeSubmissionDTOs, *codeSubmissionDTO)
	}

	var resultDTO CodeSubmissionResponseDTO
	resultDTO.Count = len(codeSubmissionDTOs)
	resultDTO.Data = codeSubmissionDTOs

	return &resultDTO, nil
}

func (s *CodeSubmissionService) CreateCodeSubmission(codeSubmissionDTO *CreateCodeSubmissionRequest) (*entity.CodeSubmission, error) {
	codeSubmissionEntity := new(entity.CodeSubmission)
	utils.DTOtoJSON(codeSubmissionDTO, codeSubmissionEntity)

	createdCodeSubmission, err := s.repository.CreateCodeSubmission(codeSubmissionEntity)
	if err != nil {
		return nil, err
	}
	return createdCodeSubmission, nil
}

func (s *CodeSubmissionService) UpdateCodeSubmission(codeSubmissionDTO *UpdateCodeSubmissionRequest) (*entity.CodeSubmission, error) {
	codeSubmissionEntity := new(entity.CodeSubmission)
	if err := utils.DTOtoJSON(codeSubmissionDTO, codeSubmissionEntity); err != nil {
		return nil, err
	}
	err := s.repository.UpdateCodeSubmission(*codeSubmissionEntity)
	if err != nil {
		return nil, err
	}
	return codeSubmissionEntity, nil
}

func (s *CodeSubmissionService) DeleteCodeSubmission(id int) error {
	return s.repository.DeleteCodeSubmission(id)
}

func (s *CodeSubmissionService) GetCodeSubmissionByID(id int) (*CodeSubmissionDTO, error) {
	codeSubmission, err := s.repository.GetCodeSubmissionByID(id)
	if err != nil {
		return nil, err
	}
	codeSubmissionDTO := new(CodeSubmissionDTO)
	err = utils.JSONtoDTO(codeSubmission, codeSubmissionDTO)
	if err != nil {
		return nil, errors.New("failed to convert code submission entity to code submission dto")
	}
	return codeSubmissionDTO, nil
}

func (s *CodeSubmissionService) GetCodeSubmissionsByCodeID(codeID int) (*CodeSubmissionResponseDTO, error) {
	codeSubmissions, err := s.repository.GetCodeSubmissionsByCodeID(codeID)
	if err != nil {
		return nil, err
	}
	codeSubmissionDTOs := []CodeSubmissionDTO{}
	for i := range codeSubmissions {
		codeSubmissionDTO := new(CodeSubmissionDTO)
		err := utils.JSONtoDTO(codeSubmissions[i], codeSubmissionDTO)
		if err != nil {
			return nil, errors.New("failed to convert code submission entity to code submission dto")
		}
		codeSubmissionDTOs = append(codeSubmissionDTOs, *codeSubmissionDTO)
	}

	var resultDTO CodeSubmissionResponseDTO
	resultDTO.Count = len(codeSubmissionDTOs)
	resultDTO.Data = codeSubmissionDTOs

	return &resultDTO, nil
}
