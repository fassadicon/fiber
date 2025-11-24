package repository

import (
	"go-fiber-starter/app/models"

	"gorm.io/gorm"
)

type ProgramRepository struct{}

func (r *ProgramRepository) GetAll(tx *gorm.DB) ([]models.Program, error) {
	var items []models.Program
	if err := tx.Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (r *ProgramRepository) InsertMany(tx *gorm.DB, items []models.Program, batchSize int) error {
	if err := tx.CreateInBatches(&items, batchSize).Error; err != nil {
		return err
	}

	return nil
}

func (r *ProgramRepository) GetListIDs(tx *gorm.DB) []int {
	var items []models.Program
	ids := make([]int, 0)
	if err := tx.Select("id").Find(&items).Error; err != nil {
		return ids
	}

	for _, it := range items {
		ids = append(ids, it.ID)
	}

	return ids
}
