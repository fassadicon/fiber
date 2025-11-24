package services

import (
	"go-fiber-starter/app/dto"
	"go-fiber-starter/app/models"
	"go-fiber-starter/app/repository"
	"go-fiber-starter/app/transformer"
	"go-fiber-starter/utils"

	"github.com/gofiber/fiber/v2"
)

type BeneficiaryService struct{}

func (s *BeneficiaryService) repo() *repository.BeneficiaryRepository {
	return new(repository.BeneficiaryRepository)
}

func (s *BeneficiaryService) List(ctx *fiber.Ctx) error {
	var items []models.Beneficiary
	if err := repository.DB.Find(&items).Error; err != nil {
		return utils.JsonErrorInternal(ctx, err, "E_BENEFICIARY_LIST")
	}

	rows := transformer.BeneficiaryListTransformer(items)
	return utils.JsonSuccess(ctx, rows)
}

func (s *BeneficiaryService) Add(ctx *fiber.Ctx, req dto.BeneficiaryRequestDTO) error {
	if err := req.Validate(); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	m := models.Beneficiary{
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		MiddleName:    req.MiddleName,
		Suffix:        req.Suffix,
		Birthdate:     req.Birthdate,
		PCN:           req.PCN,
		Sex:           req.Sex,
		MobileNumber:  req.MobileNumber,
		CivilStatusID: req.CivilStatusID,
		OccupationID:  req.OccupationID,
		MonthlySalary: req.MonthlySalary,
	}

	if err := repository.DB.Create(&m).Error; err != nil {
		return utils.JsonErrorInternal(ctx, err, "E_BENEFICIARY_ADD")
	}

	return utils.JsonSuccess(ctx, transformer.BeneficiaryTransformer(&m))
}

func (s *BeneficiaryService) Detail(ctx *fiber.Ctx, guid string) error {
	m, err := s.repo().FindByGUID(guid)
	if err != nil {
		return utils.JsonErrorNotFound(ctx, err)
	}

	return utils.JsonSuccess(ctx, transformer.BeneficiaryTransformer(&m))
}

func (s *BeneficiaryService) Update(ctx *fiber.Ctx, guid string, req dto.BeneficiaryRequestDTO) error {
	if err := req.Validate(); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	_, err := s.repo().FindByGUID(guid)
	if err != nil {
		return utils.JsonErrorNotFound(ctx, err)
	}

	store := models.Beneficiary{
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		MiddleName:    req.MiddleName,
		Suffix:        req.Suffix,
		Birthdate:     req.Birthdate,
		PCN:           req.PCN,
		Sex:           req.Sex,
		MobileNumber:  req.MobileNumber,
		CivilStatusID: req.CivilStatusID,
		OccupationID:  req.OccupationID,
		MonthlySalary: req.MonthlySalary,
	}

	updated, err := s.repo().UpdateByGUID(repository.DB, guid, store)
	if err != nil {
		return utils.JsonErrorInternal(ctx, err, "E_BENEFICIARY_UPDATE")
	}

	return utils.JsonSuccess(ctx, transformer.BeneficiaryTransformer(&updated))
}

func (s *BeneficiaryService) Delete(ctx *fiber.Ctx, guid string) error {
	if err := s.repo().DeleteByGUID(guid); err != nil {
		return utils.JsonErrorInternal(ctx, err, "E_BENEFICIARY_DELETE")
	}

	return utils.JsonSuccess(ctx, fiber.Map{"guid": guid})
}
