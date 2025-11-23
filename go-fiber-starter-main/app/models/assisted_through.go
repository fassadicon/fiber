package models

type AssistedThrough struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (m *AssistedThrough) TableName() string {
	return "assisted_throughs"
}
