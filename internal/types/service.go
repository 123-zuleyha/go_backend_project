package types

import (
	"errors"

	"github.com/123-zuleyha/go_backend_project/db/entity"
		"github.com/123-zuleyha/go_backend_project/pkg/utils"
)

type UserTypeService struct {
	repository UserTypeRepository
}

func NewUserTypeService(repository *UserTypeRepository) *UserTypeService {
	return &UserTypeService{repository: *repository}
}

func (s *UserTypeService) GetUserTypes(req *BaseRequest) (*UserTypeResponseDTO, error) {
	userTypes, err := s.repository.GetUserTypes(req)
	if err != nil {
		return nil, err
	}
	userTypeDTOs := []UserTypeDTO{}
	for i := range userTypes {
		userTypeDTO := new(UserTypeDTO)
		err := utils.JSONtoDTO(userTypes[i], userTypeDTO)

		if err != nil {
			return nil, errors.New("failed to convert user type entity to user type dto")
		}
		userTypeDTOs = append(userTypeDTOs, *userTypeDTO)
	}

	var resultDTO UserTypeResponseDTO
	resultDTO.Count = len(userTypeDTOs)
	resultDTO.Data = userTypeDTOs

	return &resultDTO, nil
}

func (s *UserTypeService) CreateUserType(userTypeDTO *CreateUserTypeRequest) (*entity.UserType, error) {
	userTypeEntity := new(entity.UserType)
	utils.DTOtoJSON(userTypeDTO, userTypeEntity)

	createdType, err := s.repository.CreateUserType(userTypeEntity)
	if err != nil {
		return nil, err
	}
	return createdType, nil

}


