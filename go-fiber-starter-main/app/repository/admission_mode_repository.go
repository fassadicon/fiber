package repository

import (
	"go-fiber-starter/app/models"

	"gorm.io/gorm"
)

type AdmissionModeRepository struct{}

func (r *AdmissionModeRepository) GetAll(tx *gorm.DB) ([]models.AdmissionMode, error) {
	var items []models.AdmissionMode
	if err := tx.Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (r *AdmissionModeRepository) InsertMany(tx *gorm.DB, items []models.AdmissionMode, batchSize int) error {
	if err := tx.CreateInBatches(&items, batchSize).Error; err != nil {
		return err
	}

	return nil
}

func (r *AdmissionModeRepository) GetListIDs(tx *gorm.DB) []int {
	var items []models.AdmissionMode
	ids := make([]int, 0)
	if err := tx.Select("id").Find(&items).Error; err != nil {
		return ids
	}

	for _, it := range items {
		ids = append(ids, it.ID)
	}

	return ids
}

func (r *AdmissionModeRepository) FindByID(id int) (models.AdmissionMode, error) {
	var m models.AdmissionMode
	if err := DB.First(&m, "id = ?", id).Error; err != nil {
		return m, err
	}

	return m, nil
}

func (r *AdmissionModeRepository) UpdateByID(tx *gorm.DB, id int, storeData models.AdmissionMode) (models.AdmissionMode, error) {
	var m models.AdmissionMode
	if err := tx.Model(&m).Where("id = ?", id).Updates(&storeData).Error; err != nil {
		return m, err
	}

	return m, nil
}

func (r *AdmissionModeRepository) DeleteByID(id int) error {
	var m models.AdmissionMode
	if err := DB.Where("id = ?", id).Delete(&m).Error; err != nil {
		return err
	}

	return nil
}
