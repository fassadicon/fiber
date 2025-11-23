package models

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID            int        `json:"id"`
	GUID          uuid.UUID  `json:"uuid" gorm:"column:uuid;type:uuid;default:gen_random_uuid()"`
	FirstName     string     `json:"first_name"`
	LastName      string     `json:"last_name"`
	MiddleName    *string    `json:"middle_name"`
	Suffix        *string    `json:"suffix"`
	Birthdate     *time.Time `json:"birthdate"`
	PCN           *string    `json:"pcn"`
	Sex           *bool      `json:"sex"`
	MobileNumber  *string    `json:"mobile_number"`
	CivilStatusID *int       `json:"civil_status_id"`
	OccupationID  *int       `json:"occupation_id"`
	MonthlySalary *float64   `json:"monthly_salary"`
}

func (m *Client) TableName() string {
	return "clients"
}
