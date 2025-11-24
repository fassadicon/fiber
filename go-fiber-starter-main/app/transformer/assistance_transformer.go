package transformer

import (
	"encoding/json"
	"go-fiber-starter/app/dto"
	"go-fiber-starter/app/models"
)

func AssistanceTransformer(m *models.Assistance) dto.AssistanceResponseDTO {
	var resp dto.AssistanceResponseDTO
	jsonResponse, _ := json.Marshal(m)
	json.Unmarshal(jsonResponse, &resp)
	return resp
}

func AssistanceListTransformer(items []models.Assistance) (rows []dto.AssistanceResponseDTO) {
	for _, it := range items {
		rows = append(rows, AssistanceTransformer(&it))
	}
	return rows
}
