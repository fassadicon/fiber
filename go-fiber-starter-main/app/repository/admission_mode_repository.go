package repository

import (
	"go-fiber-starter/app/models"

	"gorm.io/gorm"
)

type AdmissionModeRepository struct{}

func (r *AdmissionModeRepository) GetAll(tx *gorm.DB) ([]models.AdmissionMode, error) {
	var items []models.AdmissionMode
	if err := tx.Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (r *AdmissionModeRepository) InsertMany(tx *gorm.DB, items []models.AdmissionMode, batchSize int) error {
	if err := tx.CreateInBatches(&items, batchSize).Error; err != nil {
		return err
	}

	return nil
}

func (r *AdmissionModeRepository) GetListIDs(tx *gorm.DB) []int {
	var items []models.AdmissionMode
	ids := make([]int, 0)
	if err := tx.Select("id").Find(&items).Error; err != nil {
		return ids
	}

	for _, it := range items {
		ids = append(ids, it.ID)
	}

	return ids
}
