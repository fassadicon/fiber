package dto

import (
    "time"
    "go-fiber-starter/utils"
)

type ClientRequestDTO struct {
    FirstName     string     `json:"first_name" validate:"required"`
    LastName      string     `json:"last_name" validate:"required"`
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

func (req *ClientRequestDTO) Validate() error {
    return utils.ExtractValidationError(req)
}

type ClientResponseDTO struct {
    ID            int        `json:"id"`
    GUID          string     `json:"guid"`
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
