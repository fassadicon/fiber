package seeders

import (
	"go-fiber-starter/app/models"
	"go-fiber-starter/app/repository"
	"go-fiber-starter/utils"

	"gorm.io/gorm"
)

type AssistedThroughSeeder struct{}

func (s *AssistedThroughSeeder) Seed(db *gorm.DB) error {
	utils.Logger.Info("✅ seed data from AssistedThroughSeeder")

	items := []models.AssistedThrough{
		{Name: "Cash"},
		{Name: "In-kind"},
		{Name: "Vouchers"},
	}

	if err := repository.DB.CreateInBatches(&items, len(items)).Error; err != nil {
		return err
	}

	return nil
}
