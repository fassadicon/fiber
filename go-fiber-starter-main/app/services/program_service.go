package services

import (
	"go-fiber-starter/app/dto"
	"go-fiber-starter/app/models"
	"go-fiber-starter/app/repository"
	"go-fiber-starter/app/transformer"
	"go-fiber-starter/utils"

	"github.com/gofiber/fiber/v2"
)

type ProgramService struct{}

func (s *ProgramService) repo() *repository.ProgramRepository {
	return new(repository.ProgramRepository)
}

func (s *ProgramService) List(ctx *fiber.Ctx) error {
	items, err := s.repo().GetAll(repository.DB)
	if err != nil {
		return utils.JsonErrorInternal(ctx, err, "E_PROGRAM_LIST")
	}

	rows := transformer.ProgramListTransformer(items)
	return utils.JsonSuccess(ctx, rows)
}

func (s *ProgramService) Add(ctx *fiber.Ctx, req dto.ProgramRequestDTO) error {
	if err := req.Validate(); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	m := models.Program{
		Name: req.Name,
	}

	if err := repository.DB.Create(&m).Error; err != nil {
		return utils.JsonErrorInternal(ctx, err, "E_PROGRAM_ADD")
	}

	return utils.JsonSuccess(ctx, transformer.ProgramTransformer(&m))
}

func (s *ProgramService) Detail(ctx *fiber.Ctx, id int) error {
	m, err := s.repo().FindByID(id)
	if err != nil {
		return utils.JsonErrorNotFound(ctx, err)
	}

	return utils.JsonSuccess(ctx, transformer.ProgramTransformer(&m))
}

func (s *ProgramService) Update(ctx *fiber.Ctx, id int, req dto.ProgramRequestDTO) error {
	if err := req.Validate(); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	_, err := s.repo().FindByID(id)
	if err != nil {
		return utils.JsonErrorNotFound(ctx, err)
	}

	store := models.Program{Name: req.Name}
	updated, err := s.repo().UpdateByID(repository.DB, id, store)
	if err != nil {
		return utils.JsonErrorInternal(ctx, err, "E_PROGRAM_UPDATE")
	}

	return utils.JsonSuccess(ctx, transformer.ProgramTransformer(&updated))
}

func (s *ProgramService) Delete(ctx *fiber.Ctx, id int) error {
	if err := s.repo().DeleteByID(id); err != nil {
		return utils.JsonErrorInternal(ctx, err, "E_PROGRAM_DELETE")
	}

	return utils.JsonSuccess(ctx, fiber.Map{"id": id})
}
