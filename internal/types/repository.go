package types

import (
	"github.com/123-zuleyha/go_backend_project/db/entity"
	"gorm.io/gorm"
)

type UserTypeRepository struct {
	DB *gorm.DB
}

func NewUserTypeRepository(db *gorm.DB) *UserTypeRepository {
	return &UserTypeRepository{DB: db}
}

func (r *UserTypeRepository) GetUserTypes(req *BaseRequest) ([]entity.UserType, error) {
	var userTypes []entity.UserType
	query := r.DB
	if req.Limit != 0 {
		query = query.Limit(req.Limit)
	}
	if req.Offset != 0 {
		query = query.Offset(req.Offset)
	}
	err := query.Find(&userTypes).Error
	if err != nil {
		return nil, err
	}
	return userTypes, nil
}

func (r *UserTypeRepository) CreateUserType(userTypeEntity *entity.UserType) (*entity.UserType, error) {
	err := r.DB.Create(userTypeEntity).Error
	return userTypeEntity, err
}
