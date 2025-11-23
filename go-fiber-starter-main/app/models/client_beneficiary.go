package models

type ClientBeneficiary struct {
	ClientID      int `json:"client_id"`
	BeneficiaryID int `json:"beneficiary_id"`
}

func (m *ClientBeneficiary) TableName() string {
	return "client_beneficiary"
}
