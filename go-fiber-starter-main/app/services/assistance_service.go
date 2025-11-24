package services

import (
	"go-fiber-starter/app/dto"
	"go-fiber-starter/app/models"
	"go-fiber-starter/app/repository"
	"go-fiber-starter/app/transformer"
	"go-fiber-starter/utils"

	"github.com/gofiber/fiber/v2"
)

type AssistanceService struct{}

func (s *AssistanceService) repo() *repository.AssistanceRepository {
	return new(repository.AssistanceRepository)
}

func (s *AssistanceService) List(ctx *fiber.Ctx) error {
	var items []models.Assistance
	if err := repository.DB.Find(&items).Error; err != nil {
		return utils.JsonErrorInternal(ctx, err, "E_ASSISTANCE_LIST")
	}

	rows := transformer.AssistanceListTransformer(items)
	return utils.JsonSuccess(ctx, rows)
}

func (s *AssistanceService) Add(ctx *fiber.Ctx, req dto.AssistanceRequestDTO) error {
	if err := req.Validate(); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	m := models.Assistance{
		TransactionID:      req.TransactionID,
		AssistanceTypeID:   req.AssistanceTypeID,
		AmountNeeded:       req.AmountNeeded,
		AmountProvided:     req.AmountProvided,
		Purpose:            req.Purpose,
		ReleaseModeID:      req.ReleaseModeID,
		Diagnosis:          req.Diagnosis,
		SocialWorkerID:     req.SocialWorkerID,
		ApprovingOfficerID: req.ApprovingOfficerID,
	}

	if err := repository.DB.Create(&m).Error; err != nil {
		return utils.JsonErrorInternal(ctx, err, "E_ASSISTANCE_ADD")
	}

	return utils.JsonSuccess(ctx, transformer.AssistanceTransformer(&m))
}

func (s *AssistanceService) Detail(ctx *fiber.Ctx, id int) error {
	m, err := s.repo().FindByID(id)
	if err != nil {
		return utils.JsonErrorNotFound(ctx, err)
	}

	return utils.JsonSuccess(ctx, transformer.AssistanceTransformer(&m))
}

func (s *AssistanceService) Update(ctx *fiber.Ctx, id int, req dto.AssistanceRequestDTO) error {
	if err := req.Validate(); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	_, err := s.repo().FindByID(id)
	if err != nil {
		return utils.JsonErrorNotFound(ctx, err)
	}

	store := models.Assistance{
		TransactionID:      req.TransactionID,
		AssistanceTypeID:   req.AssistanceTypeID,
		AmountNeeded:       req.AmountNeeded,
		AmountProvided:     req.AmountProvided,
		Purpose:            req.Purpose,
		ReleaseModeID:      req.ReleaseModeID,
		Diagnosis:          req.Diagnosis,
		SocialWorkerID:     req.SocialWorkerID,
		ApprovingOfficerID: req.ApprovingOfficerID,
	}

	updated, err := s.repo().UpdateByID(repository.DB, id, store)
	if err != nil {
		return utils.JsonErrorInternal(ctx, err, "E_ASSISTANCE_UPDATE")
	}

	return utils.JsonSuccess(ctx, transformer.AssistanceTransformer(&updated))
}

func (s *AssistanceService) Delete(ctx *fiber.Ctx, id int) error {
	if err := s.repo().DeleteByID(id); err != nil {
		return utils.JsonErrorInternal(ctx, err, "E_ASSISTANCE_DELETE")
	}

	return utils.JsonSuccess(ctx, fiber.Map{"id": id})
}
