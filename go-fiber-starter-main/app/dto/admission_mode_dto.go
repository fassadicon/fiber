package dto

import "go-fiber-starter/utils"

type AdmissionModeRequestDTO struct {
	Name string `json:"name" validate:"required"`
}

func (req *AdmissionModeRequestDTO) Validate() error {
	return utils.ExtractValidationError(req)
}

type AdmissionModeResponseDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
