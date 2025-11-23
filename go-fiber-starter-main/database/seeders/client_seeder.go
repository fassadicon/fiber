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

type ClientSeeder struct{}

func (s *ClientSeeder) Seed(db *gorm.DB) error {
	utils.Logger.Info("✅ seed data from ClientSeeder")

	maxSize := 500
	batchSize := 100
	var clientsData []models.Client

	for i := 0; i < maxSize; i++ {
		var middle *string
		if gofakeit.Bool() {
			m := gofakeit.FirstName()
			middle = &m
		}

		var suffix *string
		if gofakeit.Bool() {
			sfx := "Jr"
			suffix = &sfx
		}

		birth := gofakeit.DateRange(time.Now().AddDate(-60, 0, 0), time.Now().AddDate(-1, 0, 0))
		mobile := gofakeit.Phone()

		data := models.Client{
			GUID:          uuid.MustParse(gofakeit.UUID()),
			FirstName:     gofakeit.FirstName(),
			LastName:      gofakeit.LastName(),
			MiddleName:    middle,
			Suffix:        suffix,
			Birthdate:     &birth,
			PCN:           nil,
			Sex:           ptrBool(gofakeit.Bool()),
			MobileNumber:  &mobile,
			CivilStatusID: ptrInt(rand.Intn(3) + 1),
			OccupationID:  ptrInt(rand.Intn(5) + 1),
			MonthlySalary: ptrFloat(float64(rand.Intn(50000))),
		}

		clientsData = append(clientsData, data)
	}

	clientRepo := new(repository.ClientRepository)
	return clientRepo.InsertMany(db, clientsData, batchSize)
}
