package dto

import "go-fiber-starter/utils"

type AssistanceCategoryRequestDTO struct {
	Name string `json:"name" validate:"required"`
}

func (req *AssistanceCategoryRequestDTO) Validate() error {
	return utils.ExtractValidationError(req)
}

type AssistanceCategoryResponseDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
