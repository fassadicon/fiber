package dto

import "go-fiber-starter/utils"

type ClientBeneficiaryRequestDTO struct {
    ClientID      int `json:"client_id" validate:"required,numeric"`
    BeneficiaryID int `json:"beneficiary_id" validate:"required,numeric"`
}

func (req *ClientBeneficiaryRequestDTO) Validate() error {
    return utils.ExtractValidationError(req)
}

type ClientBeneficiaryResponseDTO struct {
    ClientID      int `json:"client_id"`
    BeneficiaryID int `json:"beneficiary_id"`
}
