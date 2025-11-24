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

func (r *ProgramRepository) FindByID(id int) (models.Program, error) {
	var m models.Program
	if err := DB.First(&m, "id = ?", id).Error; err != nil {
		return m, err
	}

	return m, nil
}

func (r *ProgramRepository) UpdateByID(tx *gorm.DB, id int, storeData models.Program) (models.Program, error) {
	var m models.Program
	if err := tx.Model(&m).Where("id = ?", id).Updates(&storeData).Error; err != nil {
		return m, err
	}

	return m, nil
}

func (r *ProgramRepository) DeleteByID(id int) error {
	var m models.Program
	if err := DB.Where("id = ?", id).Delete(&m).Error; err != nil {
		return err
	}

	return nil
}
