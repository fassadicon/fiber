package repository

import (
	"go-fiber-starter/app/models"

	"gorm.io/gorm"
)

type TransactionRepository struct{}

func (r *TransactionRepository) InsertMany(tx *gorm.DB, transactions []models.Transaction, batchSize int) error {
	if err := DB.CreateInBatches(&transactions, batchSize).Error; err != nil {
		return err
	}

	return nil
}

func (r *TransactionRepository) GetListIDs(tx *gorm.DB) []int {
	var items []models.Transaction
	ids := make([]int, 0)
	if err := tx.Select("id").Find(&items).Error; err != nil {
		return ids
	}

	for _, it := range items {
		ids = append(ids, it.ID)
	}

	return ids
}
