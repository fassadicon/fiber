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

func (r *TransactionRepository) FindByID(id int) (models.Transaction, error) {
	var m models.Transaction
	if err := DB.First(&m, "id = ?", id).Error; err != nil {
		return m, err
	}

	return m, nil
}

func (r *TransactionRepository) UpdateByID(tx *gorm.DB, id int, storeData models.Transaction) (models.Transaction, error) {
	var m models.Transaction
	if err := tx.Model(&m).Where("id = ?", id).Updates(&storeData).Error; err != nil {
		return m, err
	}

	return m, nil
}

func (r *TransactionRepository) DeleteByID(id int) error {
	var m models.Transaction
	if err := DB.Where("id = ?", id).Delete(&m).Error; err != nil {
		return err
	}

	return nil
}
