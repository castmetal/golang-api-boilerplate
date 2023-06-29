package dtos

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type (
	ListAllExampleDTO struct {
		Offset int `json:"offset" validate:"numeric,gte=0,max=9999"`
		Limit  int `json:"limit" validate:"required,numeric,min=10,max=10000"`
	}
)

func (dto *ListAllExampleDTO) Validate() (bool, error) {
	var validate *validator.Validate

	validate = validator.New()
	err := validate.Struct(dto)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (dto *ListAllExampleDTO) ToBytes() ([]byte, error) {
	b, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}

	return b, nil
}
