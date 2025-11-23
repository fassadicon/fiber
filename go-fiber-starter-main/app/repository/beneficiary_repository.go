package repository

import (
	"go-fiber-starter/app/models"

	"gorm.io/gorm"
)

type BeneficiaryRepository struct{}

func (r *BeneficiaryRepository) InsertMany(tx *gorm.DB, beneficiaries []models.Beneficiary, batchSize int) error {
	if err := DB.CreateInBatches(&beneficiaries, batchSize).Error; err != nil {
		return err
	}

	return nil
}

func (r *BeneficiaryRepository) GetListIDs(tx *gorm.DB) []int {
	var items []models.Beneficiary
	ids := make([]int, 0)
	if err := tx.Select("id").Find(&items).Error; err != nil {
		return ids
	}

	for _, it := range items {
		ids = append(ids, it.ID)
	}

	return ids
}
