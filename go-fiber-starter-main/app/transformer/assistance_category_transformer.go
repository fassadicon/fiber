package transformer

import (
	"go-fiber-starter/app/dto"
	"go-fiber-starter/app/models"
)

func AssistanceCategoryTransformer(m *models.AssistanceCategory) dto.AssistanceCategoryResponseDTO {
	return dto.AssistanceCategoryResponseDTO{
		ID:   m.ID,
		Name: m.Name,
	}
}

func AssistanceCategoryListTransformer(items []models.AssistanceCategory) (rows []dto.AssistanceCategoryResponseDTO) {
	for _, it := range items {
		rows = append(rows, AssistanceCategoryTransformer(&it))
	}
	return rows
}
