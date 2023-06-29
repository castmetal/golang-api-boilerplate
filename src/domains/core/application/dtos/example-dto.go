package dtos

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type (
	ExampleDTO struct {
		ID   string `json:"id"`
		Name string `json:"name" validate:"required,min=2"`
	}
)

func (dto *ExampleDTO) Validate() (bool, error) {
	var validate *validator.Validate

	validate = validator.New()
	err := validate.Struct(dto)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (dto *ExampleDTO) ToBytes() ([]byte, error) {
	b, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}

	return b, nil
}
