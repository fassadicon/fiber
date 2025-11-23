package seeders

import (
	"go-fiber-starter/app/models"
	"go-fiber-starter/app/repository"
	"go-fiber-starter/utils"

	"gorm.io/gorm"
)

type ProgramSeeder struct{}

func (s *ProgramSeeder) Seed(db *gorm.DB) error {
	utils.Logger.Info("✅ seed data from ProgramSeeder")

	items := []models.Program{
		{Name: "Emergency Assistance"},
		{Name: "Livelihood Support"},
		{Name: "Medical Aid"},
	}

	if err := repository.DB.CreateInBatches(&items, len(items)).Error; err != nil {
		return err
	}

	return nil
}
