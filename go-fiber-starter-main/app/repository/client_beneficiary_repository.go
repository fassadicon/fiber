package repository

import (
	"go-fiber-starter/app/models"

	"gorm.io/gorm"
)

type ClientBeneficiaryRepository struct{}

func (r *ClientBeneficiaryRepository) GetAll(tx *gorm.DB) ([]models.ClientBeneficiary, error) {
	var items []models.ClientBeneficiary
	if err := tx.Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (r *ClientBeneficiaryRepository) InsertMany(tx *gorm.DB, items []models.ClientBeneficiary, batchSize int) error {
	if err := tx.CreateInBatches(&items, batchSize).Error; err != nil {
		return err
	}

	return nil
}

func (r *ClientBeneficiaryRepository) DeleteByComposite(clientID int, beneficiaryID int) error {
	var item models.ClientBeneficiary
	if err := DB.Where("client_id = ? AND beneficiary_id = ?", clientID, beneficiaryID).Delete(&item).Error; err != nil {
		return err
	}

	return nil
}
