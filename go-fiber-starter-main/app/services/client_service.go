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
	paginate := utils.GetPaginationParams(ctx)
	paginationData, err := s.repo().GetAll(paginate)
	if err != nil {
		return utils.JsonErrorInternal(ctx, err, "E_CLIENT_LIST")
	}

	clients := paginationData.Rows.([]models.Client)
	paginationData.Rows = transformer.ClientListTransformer(clients)
	return utils.JsonPagination(ctx, paginationData)
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

func (s *ClientService) Detail(ctx *fiber.Ctx, guid string) error {
	m, err := s.repo().FindByGUID(guid)
	if err != nil {
		return utils.JsonErrorNotFound(ctx, err)
	}

	return utils.JsonSuccess(ctx, transformer.ClientTransformer(&m))
}

func (s *ClientService) Update(ctx *fiber.Ctx, guid string, req dto.ClientRequestDTO) error {
	if err := req.Validate(); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	_, err := s.repo().FindByGUID(guid)
	if err != nil {
		return utils.JsonErrorNotFound(ctx, err)
	}

	store := models.Client{
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
		return utils.JsonErrorInternal(ctx, err, "E_CLIENT_UPDATE")
	}

	return utils.JsonSuccess(ctx, transformer.ClientTransformer(&updated))
}

func (s *ClientService) Delete(ctx *fiber.Ctx, guid string) error {
	if err := s.repo().DeleteByGUID(guid); err != nil {
		return utils.JsonErrorInternal(ctx, err, "E_CLIENT_DELETE")
	}

	return utils.JsonSuccess(ctx, fiber.Map{"guid": guid})
}
