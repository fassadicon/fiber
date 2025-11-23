package models

type Program struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (m *Program) TableName() string {
	return "programs"
}
