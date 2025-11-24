package transformer

import (
	"go-fiber-starter/app/dto"
	"go-fiber-starter/app/models"
)

func AdmissionModeTransformer(m *models.AdmissionMode) dto.AdmissionModeResponseDTO {
	return dto.AdmissionModeResponseDTO{
		ID:   m.ID,
		Name: m.Name,
	}
}

func AdmissionModeListTransformer(items []models.AdmissionMode) (rows []dto.AdmissionModeResponseDTO) {
	for _, it := range items {
		rows = append(rows, AdmissionModeTransformer(&it))
	}
	return rows
}
