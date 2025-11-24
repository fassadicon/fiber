package transformer

import (
	"go-fiber-starter/app/dto"
	"go-fiber-starter/app/models"
)

func ProgramTransformer(m *models.Program) dto.ProgramResponseDTO {
	return dto.ProgramResponseDTO{
		ID:   m.ID,
		Name: m.Name,
	}
}

func ProgramListTransformer(items []models.Program) (rows []dto.ProgramResponseDTO) {
	for _, it := range items {
		rows = append(rows, ProgramTransformer(&it))
	}
	return rows
}
