package transformer

import (
	"encoding/json"
	"go-fiber-starter/app/dto"
	"go-fiber-starter/app/models"
)

func BeneficiaryTransformer(m *models.Beneficiary) dto.BeneficiaryResponseDTO {
	var resp dto.BeneficiaryResponseDTO
	jsonResponse, _ := json.Marshal(m)
	json.Unmarshal(jsonResponse, &resp)

	resp.GUID = m.GUID.String()
	return resp
}

func BeneficiaryListTransformer(items []models.Beneficiary) (rows []dto.BeneficiaryResponseDTO) {
	for _, it := range items {
		rows = append(rows, BeneficiaryTransformer(&it))
	}
	return rows
}
