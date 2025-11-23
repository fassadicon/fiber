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

type TransactionSeeder struct{}

func (s *TransactionSeeder) Seed(db *gorm.DB) error {
	utils.Logger.Info("✅ seed data from TransactionSeeder")

	clientRepo := new(repository.ClientRepository)
	clients := clientRepo.GetListIDs(db)
	if len(clients) == 0 {
		return nil
	}

	beneRepo := new(repository.BeneficiaryRepository)
	beneficiaries := beneRepo.GetListIDs(db)

	// reference lists
	programs := []int{1}
	if err := db.Table("programs").Select("id").Find(&programs).Error; err != nil {
		// ignore, fallback to default
	}

	assisted := []int{1}
	if err := db.Table("assisted_throughs").Select("id").Find(&assisted).Error; err != nil {
	}

	admission := []int{1}
	if err := db.Table("admission_modes").Select("id").Find(&admission).Error; err != nil {
	}

	maxSize := 400
	batchSize := 100
	var txs []models.Transaction

	for i := 0; i < maxSize; i++ {
		clientID := clients[rand.Intn(len(clients))]
		var beneID *int
		if len(beneficiaries) > 0 && gofakeit.Bool() {
			b := beneficiaries[rand.Intn(len(beneficiaries))]
			beneID = &b
		}

		date := gofakeit.DateRange(time.Now().AddDate(-1, 0, 0), time.Now())
		data := models.Transaction{
			GUID:              uuid.MustParse(gofakeit.UUID()),
			ClientID:          clientID,
			BeneficiaryID:     beneID,
			Date:              &date,
			ProgramID:         ptrInt(programs[0]),
			IsReturning:       gofakeit.Bool(),
			AssistedThroughID: ptrInt(assisted[0]),
			AdmissionModeID:   ptrInt(admission[0]),
			ClientAge:         ptrInt(rand.Intn(80)),
			BeneAge:           ptrInt(rand.Intn(80)),
			AmountNeeded:      ptrFloat(float64(rand.Intn(50000))),
			AmountProvided:    ptrFloat(float64(rand.Intn(40000))),
			ProblemPresented:  ptrString(gofakeit.Sentence(6)),
			Assessment:        ptrString(gofakeit.Sentence(10)),
		}

		txs = append(txs, data)
	}

	repo := new(repository.TransactionRepository)
	return repo.InsertMany(db, txs, batchSize)
}
