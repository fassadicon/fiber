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

func (r *BeneficiaryRepository) FindByGUID(guid string) (models.Beneficiary, error) {
	var m models.Beneficiary
	if err := DB.First(&m, "uuid = ?", guid).Error; err != nil {
		return m, err
	}

	return m, nil
}

func (r *BeneficiaryRepository) UpdateByGUID(tx *gorm.DB, guid string, storeData models.Beneficiary) (models.Beneficiary, error) {
	var m models.Beneficiary
	if err := tx.Model(&m).Where("uuid = ?", guid).Updates(&storeData).Error; err != nil {
		return m, err
	}

	return m, nil
}

func (r *BeneficiaryRepository) DeleteByGUID(guid string) error {
	var m models.Beneficiary
	if err := DB.Where("uuid = ?", guid).Delete(&m).Error; err != nil {
		return err
	}

	return nil
}
