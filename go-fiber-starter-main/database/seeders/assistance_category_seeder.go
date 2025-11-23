package seeders

import (
	"go-fiber-starter/app/models"
	"go-fiber-starter/app/repository"
	"go-fiber-starter/utils"

	"gorm.io/gorm"
)

type AssistanceCategorySeeder struct{}

func (s *AssistanceCategorySeeder) Seed(db *gorm.DB) error {
	utils.Logger.Info("✅ seed data from AssistanceCategorySeeder")

	items := []models.AssistanceCategory{
		{Name: "Financial"},
		{Name: "Medical"},
		{Name: "Shelter"},
	}

	if err := repository.DB.CreateInBatches(&items, len(items)).Error; err != nil {
		return err
	}

	return nil
}
