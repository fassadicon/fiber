package transformer

import (
	"go-fiber-starter/app/dto"
	"go-fiber-starter/app/models"
)

func AssistedThroughTransformer(m *models.AssistedThrough) dto.AssistedThroughResponseDTO {
	return dto.AssistedThroughResponseDTO{
		ID:   m.ID,
		Name: m.Name,
	}
}

func AssistedThroughListTransformer(items []models.AssistedThrough) (rows []dto.AssistedThroughResponseDTO) {
	for _, it := range items {
		rows = append(rows, AssistedThroughTransformer(&it))
	}
	return rows
}
