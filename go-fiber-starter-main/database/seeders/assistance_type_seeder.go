package seeders

import (
	"go-fiber-starter/app/models"
	"go-fiber-starter/app/repository"
	"go-fiber-starter/utils"

	"gorm.io/gorm"
)

type AssistanceTypeSeeder struct{}

func (s *AssistanceTypeSeeder) Seed(db *gorm.DB) error {
	utils.Logger.Info("✅ seed data from AssistanceTypeSeeder")

	// map categories to types simply — here we create a few types and
	// link them to the first category (id=1) if present
	items := []models.AssistanceType{
		{AssistanceCategoryID: 1, Name: "Cash Grant"},
		{AssistanceCategoryID: 1, Name: "Transport Allowance"},
		{AssistanceCategoryID: 2, Name: "Medical Reimbursement"},
	}

	if err := repository.DB.CreateInBatches(&items, len(items)).Error; err != nil {
		return err
	}

	return nil
}
