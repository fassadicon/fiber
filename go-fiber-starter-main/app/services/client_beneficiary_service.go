package services

import (
    "go-fiber-starter/app/dto"
    "go-fiber-starter/app/models"
    "go-fiber-starter/app/repository"
    "go-fiber-starter/app/transformer"
    "go-fiber-starter/utils"

    "github.com/gofiber/fiber/v2"
)

type ClientBeneficiaryService struct{}

func (s *ClientBeneficiaryService) repo() *repository.ClientBeneficiaryRepository {
    return new(repository.ClientBeneficiaryRepository)
}

func (s *ClientBeneficiaryService) List(ctx *fiber.Ctx) error {
    items, err := s.repo().GetAll(repository.DB)
    if err != nil {
        return utils.JsonErrorInternal(ctx, err, "E_CLIENT_BENEFICIARY_LIST")
    }

    rows := transformer.ClientBeneficiaryListTransformer(items)
    return utils.JsonSuccess(ctx, rows)
}

func (s *ClientBeneficiaryService) Add(ctx *fiber.Ctx, req dto.ClientBeneficiaryRequestDTO) error {
    if err := req.Validate(); err != nil {
        return utils.JsonErrorValidation(ctx, err)
    }

    m := models.ClientBeneficiary{
        ClientID:      req.ClientID,
        BeneficiaryID: req.BeneficiaryID,
    }

    if err := repository.DB.Create(&m).Error; err != nil {
        return utils.JsonErrorInternal(ctx, err, "E_CLIENT_BENEFICIARY_ADD")
    }

    return utils.JsonSuccess(ctx, transformer.ClientBeneficiaryTransformer(&m))
}
