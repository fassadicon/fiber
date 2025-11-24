package repository

import (
	"go-fiber-starter/app/models"

	"gorm.io/gorm"
)

type AssistedThroughRepository struct{}

func (r *AssistedThroughRepository) GetAll(tx *gorm.DB) ([]models.AssistedThrough, error) {
	var items []models.AssistedThrough
	if err := tx.Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (r *AssistedThroughRepository) InsertMany(tx *gorm.DB, items []models.AssistedThrough, batchSize int) error {
	if err := tx.CreateInBatches(&items, batchSize).Error; err != nil {
		return err
	}

	return nil
}

func (r *AssistedThroughRepository) GetListIDs(tx *gorm.DB) []int {
	var items []models.AssistedThrough
	ids := make([]int, 0)
	if err := tx.Select("id").Find(&items).Error; err != nil {
		return ids
	}

	for _, it := range items {
		ids = append(ids, it.ID)
	}

	return ids
}
