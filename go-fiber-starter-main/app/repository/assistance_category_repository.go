package repository

import (
	"go-fiber-starter/app/models"

	"gorm.io/gorm"
)

type AssistanceCategoryRepository struct{}

func (r *AssistanceCategoryRepository) GetAll(tx *gorm.DB) ([]models.AssistanceCategory, error) {
	var items []models.AssistanceCategory
	if err := tx.Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (r *AssistanceCategoryRepository) InsertMany(tx *gorm.DB, items []models.AssistanceCategory, batchSize int) error {
	if err := tx.CreateInBatches(&items, batchSize).Error; err != nil {
		return err
	}

	return nil
}

func (r *AssistanceCategoryRepository) GetListIDs(tx *gorm.DB) []int {
	var items []models.AssistanceCategory
	ids := make([]int, 0)
	if err := tx.Select("id").Find(&items).Error; err != nil {
		return ids
	}

	for _, it := range items {
		ids = append(ids, it.ID)
	}

	return ids
}

func (r *AssistanceCategoryRepository) FindByID(id int) (models.AssistanceCategory, error) {
	var m models.AssistanceCategory
	if err := DB.First(&m, "id = ?", id).Error; err != nil {
		return m, err
	}

	return m, nil
}

func (r *AssistanceCategoryRepository) UpdateByID(tx *gorm.DB, id int, storeData models.AssistanceCategory) (models.AssistanceCategory, error) {
	var m models.AssistanceCategory
	if err := tx.Model(&m).Where("id = ?", id).Updates(&storeData).Error; err != nil {
		return m, err
	}

	return m, nil
}

func (r *AssistanceCategoryRepository) DeleteByID(id int) error {
	var m models.AssistanceCategory
	if err := DB.Where("id = ?", id).Delete(&m).Error; err != nil {
		return err
	}

	return nil
}
