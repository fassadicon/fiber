package repository

import (
	"go-fiber-starter/app/models"

	"gorm.io/gorm"
)

type ClientRepository struct{}

func (r *ClientRepository) InsertMany(tx *gorm.DB, clients []models.Client, batchSize int) error {
	if err := DB.CreateInBatches(&clients, batchSize).Error; err != nil {
		return err
	}

	return nil
}

func (r *ClientRepository) GetListIDs(tx *gorm.DB) []int {
	var clients []models.Client
	ids := make([]int, 0)
	if err := tx.Select("id").Find(&clients).Error; err != nil {
		return ids
	}

	for _, c := range clients {
		ids = append(ids, c.ID)
	}

	return ids
}

func (r *ClientRepository) FindByGUID(guid string) (models.Client, error) {
	var m models.Client
	if err := DB.First(&m, "uuid = ?", guid).Error; err != nil {
		return m, err
	}

	return m, nil
}

func (r *ClientRepository) UpdateByGUID(tx *gorm.DB, guid string, storeData models.Client) (models.Client, error) {
	var m models.Client
	if err := tx.Model(&m).Where("uuid = ?", guid).Updates(&storeData).Error; err != nil {
		return m, err
	}

	return m, nil
}

func (r *ClientRepository) DeleteByGUID(guid string) error {
	var m models.Client
	if err := DB.Where("uuid = ?", guid).Delete(&m).Error; err != nil {
		return err
	}

	return nil
}
