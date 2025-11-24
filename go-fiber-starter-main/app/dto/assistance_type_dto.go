package dto

import "go-fiber-starter/utils"

type AssistanceTypeRequestDTO struct {
	AssistanceCategoryID int    `json:"assistance_category_id" validate:"required,numeric"`
	Name                 string `json:"name" validate:"required"`
}

func (req *AssistanceTypeRequestDTO) Validate() error {
	return utils.ExtractValidationError(req)
}

type AssistanceTypeResponseDTO struct {
	ID                   int    `json:"id"`
	AssistanceCategoryID int    `json:"assistance_category_id"`
	Name                 string `json:"name"`
}
