package seeders

import (
	"go-fiber-starter/app/models"
	"go-fiber-starter/app/repository"
	"go-fiber-starter/utils"
	"math/rand"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AssistanceSeeder struct{}

func (s *AssistanceSeeder) Seed(db *gorm.DB) error {
	utils.Logger.Info("✅ seed data from AssistanceSeeder")

	txRepo := new(repository.TransactionRepository)
	txIDs := txRepo.GetListIDs(db)
	if len(txIDs) == 0 {
		return nil
	}

	atRepo := new(repository.AssistanceRepository)
	types := []int{1}
	if err := db.Table("assistance_types").Select("id").Find(&types).Error; err != nil {
	}

	maxSize := 300
	batchSize := 100
	var items []models.Assistance

	for i := 0; i < maxSize; i++ {
		txID := txIDs[rand.Intn(len(txIDs))]
		data := models.Assistance{
			GUID:             uuid.MustParse(gofakeit.UUID()),
			TransactionID:    txID,
			AssistanceTypeID: types[0],
			AmountNeeded:     ptrFloat(float64(rand.Intn(20000))),
			AmountProvided:   ptrFloat(float64(rand.Intn(15000))),
			Purpose:          ptrString(gofakeit.Sentence(6)),
		}

		items = append(items, data)
	}

	return atRepo.InsertMany(db, items, batchSize)
}
