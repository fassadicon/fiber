package models

import (
	"github.com/google/uuid"
)

type Assistance struct {
	ID                 int       `json:"id"`
	GUID               uuid.UUID `json:"uuid" gorm:"column:uuid;type:uuid;default:gen_random_uuid()"`
	TransactionID      int       `json:"transaction_id"`
	AssistanceTypeID   int       `json:"assistance_type_id"`
	AmountNeeded       *float64  `json:"amount_needed"`
	AmountProvided     *float64  `json:"amount_provided"`
	Purpose            *string   `json:"purpose"`
	ReleaseModeID      *int      `json:"release_mode_id"`
	Diagnosis          *string   `json:"diagnosis"`
	SocialWorkerID     *int      `json:"social_worker_id"`
	ApprovingOfficerID *int      `json:"approving_officer_id"`
}

func (m *Assistance) TableName() string {
	return "assistances"
}
