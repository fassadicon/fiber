package seeders

import (
	"go-fiber-starter/app/models"
	"go-fiber-starter/app/repository"
	"go-fiber-starter/utils"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BeneficiarySeeder struct{}

func (s *BeneficiarySeeder) Seed(db *gorm.DB) error {
	utils.Logger.Info("✅ seed data from BeneficiarySeeder")

	maxSize := 300
	batchSize := 100
	var dataItems []models.Beneficiary

	for i := 0; i < maxSize; i++ {
		var middle *string
		if gofakeit.Bool() {
			m := gofakeit.FirstName()
			middle = &m
		}

		birth := gofakeit.DateRange(time.Now().AddDate(-60, 0, 0), time.Now().AddDate(-1, 0, 0))
		mobile := gofakeit.Phone()

		data := models.Beneficiary{
			GUID:          uuid.MustParse(gofakeit.UUID()),
			FirstName:     gofakeit.FirstName(),
			LastName:      gofakeit.LastName(),
			MiddleName:    middle,
			Birthdate:     &birth,
			PCN:           nil,
			Sex:           ptrBool(gofakeit.Bool()),
			MobileNumber:  &mobile,
			CivilStatusID: ptrInt(rand.Intn(3) + 1),
			OccupationID:  ptrInt(rand.Intn(5) + 1),
			MonthlySalary: ptrFloat(float64(rand.Intn(40000))),
		}

		dataItems = append(dataItems, data)
	}

	repo := new(repository.BeneficiaryRepository)
	return repo.InsertMany(db, dataItems, batchSize)
}
