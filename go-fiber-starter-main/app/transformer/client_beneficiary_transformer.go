package transformer

import (
	"go-fiber-starter/app/dto"
	"go-fiber-starter/app/models"
)

func ClientBeneficiaryTransformer(m *models.ClientBeneficiary) dto.ClientBeneficiaryResponseDTO {
	return dto.ClientBeneficiaryResponseDTO{
		ClientID:      m.ClientID,
		BeneficiaryID: m.BeneficiaryID,
	}
}

func ClientBeneficiaryListTransformer(items []models.ClientBeneficiary) (rows []dto.ClientBeneficiaryResponseDTO) {
	for _, it := range items {
		rows = append(rows, ClientBeneficiaryTransformer(&it))
	}
	return rows
}
