package dtos

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type (
	CreateExampleResponseDTO struct {
		Example ExampleResponseDTO `json:"example"`
	}
)

func (dto *CreateExampleResponseDTO) Validate() (bool, error) {
	var validate *validator.Validate

	validate = validator.New()
	err := validate.Struct(dto)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (dto *CreateExampleResponseDTO) ToBytes() ([]byte, error) {
	b, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}

	return b, nil
}
