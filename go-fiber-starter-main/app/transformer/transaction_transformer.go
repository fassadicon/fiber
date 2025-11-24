package transformer

import (
	"encoding/json"
	"go-fiber-starter/app/dto"
	"go-fiber-starter/app/models"
)

func TransactionTransformer(m *models.Transaction) dto.TransactionResponseDTO {
	var resp dto.TransactionResponseDTO
	jsonResponse, _ := json.Marshal(m)
	json.Unmarshal(jsonResponse, &resp)

	// GUID is stored as uuid field named GUID in some models; ensure string mapping if present
	// Transaction model uses GUID field named GUID with json tag `uuid`, so we skip here.
	return resp
}

func TransactionListTransformer(items []models.Transaction) (rows []dto.TransactionResponseDTO) {
	for _, it := range items {
		rows = append(rows, TransactionTransformer(&it))
	}
	return rows
}
