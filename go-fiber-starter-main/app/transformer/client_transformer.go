package transformer

import (
	"encoding/json"
	"go-fiber-starter/app/dto"
	"go-fiber-starter/app/models"
)

func ClientTransformer(m *models.Client) dto.ClientResponseDTO {
	var resp dto.ClientResponseDTO
	// Use json marshal/unmarshal for simple mapping of similar fields
	jsonResponse, _ := json.Marshal(m)
	json.Unmarshal(jsonResponse, &resp)

	resp.GUID = m.GUID.String()
	return resp
}

func ClientListTransformer(items []models.Client) (rows []dto.ClientResponseDTO) {
	for _, it := range items {
		rows = append(rows, ClientTransformer(&it))
	}
	return rows
}
