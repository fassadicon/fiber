package transformer

import (
	"go-fiber-starter/app/dto"
	"go-fiber-starter/app/models"
)

func AssistanceTypeTransformer(m *models.AssistanceType) dto.AssistanceTypeResponseDTO {
	return dto.AssistanceTypeResponseDTO{
		ID:                   m.ID,
		AssistanceCategoryID: m.AssistanceCategoryID,
		Name:                 m.Name,
	}
}

func AssistanceTypeListTransformer(items []models.AssistanceType) (rows []dto.AssistanceTypeResponseDTO) {
	for _, it := range items {
		rows = append(rows, AssistanceTypeTransformer(&it))
	}
	return rows
}
