package repository

import (
	"go-fiber-starter/app/models"

	"gorm.io/gorm"
)

type AssistanceTypeRepository struct{}

func (r *AssistanceTypeRepository) GetAll(tx *gorm.DB) ([]models.AssistanceType, error) {
	var items []models.AssistanceType
	if err := tx.Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (r *AssistanceTypeRepository) InsertMany(tx *gorm.DB, items []models.AssistanceType, batchSize int) error {
	if err := tx.CreateInBatches(&items, batchSize).Error; err != nil {
		return err
	}

	return nil
}

func (r *AssistanceTypeRepository) GetListIDs(tx *gorm.DB) []int {
	var items []models.AssistanceType
	ids := make([]int, 0)
	if err := tx.Select("id").Find(&items).Error; err != nil {
		return ids
	}

	for _, it := range items {
		ids = append(ids, it.ID)
	}

	return ids
}
