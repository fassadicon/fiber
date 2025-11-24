package services

import (
    "go-fiber-starter/app/dto"
    "go-fiber-starter/app/models"
    "go-fiber-starter/app/repository"
    "go-fiber-starter/app/transformer"
    "go-fiber-starter/utils"

    "github.com/gofiber/fiber/v2"
)

type ClientService struct{}

func (s *ClientService) repo() *repository.ClientRepository {
    return new(repository.ClientRepository)
}

func (s *ClientService) List(ctx *fiber.Ctx) error {
    var items []models.Client
    if err := repository.DB.Find(&items).Error; err != nil {
        return utils.JsonErrorInternal(ctx, err, "E_CLIENT_LIST")
    }

    rows := transformer.ClientListTransformer(items)
    return utils.JsonSuccess(ctx, rows)
}

func (s *ClientService) Add(ctx *fiber.Ctx, req dto.ClientRequestDTO) error {
    if err := req.Validate(); err != nil {
        return utils.JsonErrorValidation(ctx, err)
    }

    m := models.Client{
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
        return utils.JsonErrorInternal(ctx, err, "E_CLIENT_ADD")
    }

    return utils.JsonSuccess(ctx, transformer.ClientTransformer(&m))
}
