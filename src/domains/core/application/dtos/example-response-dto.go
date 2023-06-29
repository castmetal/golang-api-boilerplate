package dtos

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type (
	ExampleResponseDTO struct {
		ID        string  `json:"id"`
		Name      string  `json:"name"`
		CreatedAt string  `json:"created_at"`
		UpdatedAt string  `json:"updated_at"`
		DeletedAt *string `json:"deleted_at"`
	}
)

func (dto *ExampleResponseDTO) Validate() (bool, error) {
	var validate *validator.Validate

	validate = validator.New()
	err := validate.Struct(dto)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (dto *ExampleResponseDTO) ToBytes() ([]byte, error) {
	b, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}

	return b, nil
}
