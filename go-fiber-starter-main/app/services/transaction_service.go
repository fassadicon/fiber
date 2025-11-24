package services

import (
    "go-fiber-starter/app/dto"
    "go-fiber-starter/app/models"
    "go-fiber-starter/app/repository"
    "go-fiber-starter/app/transformer"
    "go-fiber-starter/utils"

    "github.com/gofiber/fiber/v2"
)

type TransactionService struct{}

func (s *TransactionService) repo() *repository.TransactionRepository {
    return new(repository.TransactionRepository)
}

func (s *TransactionService) List(ctx *fiber.Ctx) error {
    var items []models.Transaction
    if err := repository.DB.Find(&items).Error; err != nil {
        return utils.JsonErrorInternal(ctx, err, "E_TRANSACTION_LIST")
    }

    rows := transformer.TransactionListTransformer(items)
    return utils.JsonSuccess(ctx, rows)
}

func (s *TransactionService) Add(ctx *fiber.Ctx, req dto.TransactionRequestDTO) error {
    if err := req.Validate(); err != nil {
        return utils.JsonErrorValidation(ctx, err)
    }

    m := models.Transaction{
        ClientID:          req.ClientID,
        BeneficiaryID:     req.BeneficiaryID,
        Date:              req.Date,
        ProgramID:         req.ProgramID,
        IsReturning:       false,
        AssistedThroughID: req.AssistedThroughID,
        AdmissionModeID:   req.AdmissionModeID,
        ClientAge:         req.ClientAge,
        BeneAge:           req.BeneAge,
        AmountNeeded:      req.AmountNeeded,
        AmountProvided:    req.AmountProvided,
        ProblemPresented:  req.ProblemPresented,
        Assessment:        req.Assessment,
    }

    if req.IsReturning != nil {
        m.IsReturning = *req.IsReturning
    }

    if err := repository.DB.Create(&m).Error; err != nil {
        return utils.JsonErrorInternal(ctx, err, "E_TRANSACTION_ADD")
    }

    return utils.JsonSuccess(ctx, transformer.TransactionTransformer(&m))
}
