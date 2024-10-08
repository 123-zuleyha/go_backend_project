package types

type BaseRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type BaseResponse struct {
	BaseRequest
	Count int `json:"count"`
}

type CreateUserTypeRequest struct {
	Name string `json:"name"`
}

type UpdateUserTypeRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserTypeDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserTypeResponseDTO struct {
	Count int           `json:"count"`
	Data  []UserTypeDTO `json:"data"`
}
