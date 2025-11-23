package models

type AdmissionMode struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (m *AdmissionMode) TableName() string {
	return "admission_modes"
}
