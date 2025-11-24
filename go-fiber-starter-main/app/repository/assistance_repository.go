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

func (r *AssistanceRepository) FindByID(id int) (models.Assistance, error) {
	var m models.Assistance
	if err := DB.First(&m, "id = ?", id).Error; err != nil {
		return m, err
	}

	return m, nil
}

func (r *AssistanceRepository) UpdateByID(tx *gorm.DB, id int, storeData models.Assistance) (models.Assistance, error) {
	var m models.Assistance
	if err := tx.Model(&m).Where("id = ?", id).Updates(&storeData).Error; err != nil {
		return m, err
	}

	return m, nil
}

func (r *AssistanceRepository) DeleteByID(id int) error {
	var m models.Assistance
	if err := DB.Where("id = ?", id).Delete(&m).Error; err != nil {
		return err
	}

	return nil
}
