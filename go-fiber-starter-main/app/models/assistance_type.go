package models

type AssistanceType struct {
	ID                   int                `json:"id"`
	AssistanceCategoryID int                `json:"assistance_category_id"`
	Name                 string             `json:"name"`
	AssistanceCategory   AssistanceCategory `json:"-"`
}

func (m *AssistanceType) TableName() string {
	return "assistance_types"
}
