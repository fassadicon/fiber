package dto

import "go-fiber-starter/utils"

type AssistanceRequestDTO struct {
	TransactionID      int      `json:"transaction_id" validate:"required,numeric"`
	AssistanceTypeID   int      `json:"assistance_type_id" validate:"required,numeric"`
	AmountNeeded       *float64 `json:"amount_needed"`
	AmountProvided     *float64 `json:"amount_provided"`
	Purpose            *string  `json:"purpose"`
	ReleaseModeID      *int     `json:"release_mode_id"`
	Diagnosis          *string  `json:"diagnosis"`
	SocialWorkerID     *int     `json:"social_worker_id"`
	ApprovingOfficerID *int     `json:"approving_officer_id"`
}

func (req *AssistanceRequestDTO) Validate() error {
	return utils.ExtractValidationError(req)
}

type AssistanceResponseDTO struct {
	ID                 int      `json:"id"`
	GUID               string   `json:"guid"`
	TransactionID      int      `json:"transaction_id"`
	AssistanceTypeID   int      `json:"assistance_type_id"`
	AmountNeeded       *float64 `json:"amount_needed"`
	AmountProvided     *float64 `json:"amount_provided"`
	Purpose            *string  `json:"purpose"`
	ReleaseModeID      *int     `json:"release_mode_id"`
	Diagnosis          *string  `json:"diagnosis"`
	SocialWorkerID     *int     `json:"social_worker_id"`
	ApprovingOfficerID *int     `json:"approving_officer_id"`
}
