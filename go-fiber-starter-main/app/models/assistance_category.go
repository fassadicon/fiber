package models

type AssistanceCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (m *AssistanceCategory) TableName() string {
	return "assistance_categories"
}
