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

func (r *AssistedThroughRepository) FindByID(id int) (models.AssistedThrough, error) {
	var m models.AssistedThrough
	if err := DB.First(&m, "id = ?", id).Error; err != nil {
		return m, err
	}

	return m, nil
}

func (r *AssistedThroughRepository) UpdateByID(tx *gorm.DB, id int, storeData models.AssistedThrough) (models.AssistedThrough, error) {
	var m models.AssistedThrough
	if err := tx.Model(&m).Where("id = ?", id).Updates(&storeData).Error; err != nil {
		return m, err
	}

	return m, nil
}

func (r *AssistedThroughRepository) DeleteByID(id int) error {
	var m models.AssistedThrough
	if err := DB.Where("id = ?", id).Delete(&m).Error; err != nil {
		return err
	}

	return nil
}
