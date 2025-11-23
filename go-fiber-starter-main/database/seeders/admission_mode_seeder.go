package seeders

import (
	"go-fiber-starter/app/models"
	"go-fiber-starter/app/repository"
	"go-fiber-starter/utils"

	"gorm.io/gorm"
)

type AdmissionModeSeeder struct{}

func (s *AdmissionModeSeeder) Seed(db *gorm.DB) error {
	utils.Logger.Info("✅ seed data from AdmissionModeSeeder")

	items := []models.AdmissionMode{
		{Name: "Walk-in"},
		{Name: "Referral"},
		{Name: "Outreach"},
	}

	if err := repository.DB.CreateInBatches(&items, len(items)).Error; err != nil {
		return err
	}

	return nil
}
