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
