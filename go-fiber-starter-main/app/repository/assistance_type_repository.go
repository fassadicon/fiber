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

func (r *AssistanceTypeRepository) FindByID(id int) (models.AssistanceType, error) {
	var m models.AssistanceType
	if err := DB.First(&m, "id = ?", id).Error; err != nil {
		return m, err
	}

	return m, nil
}

func (r *AssistanceTypeRepository) UpdateByID(tx *gorm.DB, id int, storeData models.AssistanceType) (models.AssistanceType, error) {
	var m models.AssistanceType
	if err := tx.Model(&m).Where("id = ?", id).Updates(&storeData).Error; err != nil {
		return m, err
	}

	return m, nil
}

func (r *AssistanceTypeRepository) DeleteByID(id int) error {
	var m models.AssistanceType
	if err := DB.Where("id = ?", id).Delete(&m).Error; err != nil {
		return err
	}

	return nil
}
