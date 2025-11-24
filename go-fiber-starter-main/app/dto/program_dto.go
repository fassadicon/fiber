package dto

import "go-fiber-starter/utils"

type ProgramRequestDTO struct {
    Name string `json:"name" validate:"required"`
}

func (req *ProgramRequestDTO) Validate() error {
    return utils.ExtractValidationError(req)
}

type ProgramResponseDTO struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}
