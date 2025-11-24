package dto

import "go-fiber-starter/utils"

type AssistedThroughRequestDTO struct {
    Name string `json:"name" validate:"required"`
}

func (req *AssistedThroughRequestDTO) Validate() error {
    return utils.ExtractValidationError(req)
}

type AssistedThroughResponseDTO struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}
