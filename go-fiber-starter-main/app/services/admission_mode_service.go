package services

import (
    "go-fiber-starter/app/dto"
    "go-fiber-starter/app/models"
    "go-fiber-starter/app/repository"
    "go-fiber-starter/app/transformer"
    "go-fiber-starter/utils"

    "github.com/gofiber/fiber/v2"
)

type AdmissionModeService struct{}

func (s *AdmissionModeService) repo() *repository.AdmissionModeRepository {
    return new(repository.AdmissionModeRepository)
}

func (s *AdmissionModeService) List(ctx *fiber.Ctx) error {
    items, err := s.repo().GetAll(repository.DB)
    if err != nil {
        return utils.JsonErrorInternal(ctx, err, "E_ADMISSION_MODE_LIST")
    }

    rows := transformer.AdmissionModeListTransformer(items)
    return utils.JsonSuccess(ctx, rows)
}

func (s *AdmissionModeService) Add(ctx *fiber.Ctx, req dto.AdmissionModeRequestDTO) error {
    if err := req.Validate(); err != nil {
        return utils.JsonErrorValidation(ctx, err)
    }

    m := models.AdmissionMode{
        Name: req.Name,
    }

    if err := repository.DB.Create(&m).Error; err != nil {
        return utils.JsonErrorInternal(ctx, err, "E_ADMISSION_MODE_ADD")
    }

    return utils.JsonSuccess(ctx, transformer.AdmissionModeTransformer(&m))
}
