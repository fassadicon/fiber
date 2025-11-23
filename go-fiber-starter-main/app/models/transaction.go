package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID                int        `json:"id"`
	GUID              uuid.UUID  `json:"uuid" gorm:"column:uuid;type:uuid;default:gen_random_uuid()"`
	ClientID          int        `json:"client_id"`
	BeneficiaryID     *int       `json:"beneficiary_id"`
	Date              *time.Time `json:"date"`
	ProgramID         *int       `json:"program_id"`
	IsReturning       bool       `json:"is_returning"`
	AssistedThroughID *int       `json:"assisted_through_id"`
	AdmissionModeID   *int       `json:"admission_mode_id"`
	ClientAge         *int       `json:"client_age"`
	BeneAge           *int       `json:"bene_age"`
	AmountNeeded      *float64   `json:"amount_needed"`
	AmountProvided    *float64   `json:"amount_provided"`
	ProblemPresented  *string    `json:"problem_presented"`
	Assessment        *string    `json:"assessment"`
}

func (m *Transaction) TableName() string {
	return "transactions"
}
