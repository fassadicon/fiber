package repository

import (
	"go-fiber-starter/app/models"

	"gorm.io/gorm"
)

type AssistanceRepository struct{}

func (r *AssistanceRepository) InsertMany(tx *gorm.DB, assistances []models.Assistance, batchSize int) error {
	if err := DB.CreateInBatches(&assistances, batchSize).Error; err != nil {
		return err
	}

	return nil
}

func (r *AssistanceRepository) GetListIDs(tx *gorm.DB) []int {
	var items []models.Assistance
	ids := make([]int, 0)
	if err := tx.Select("id").Find(&items).Error; err != nil {
		return ids
	}

	for _, it := range items {
		ids = append(ids, it.ID)
	}

	return ids
}
