package services

import (
	"go-fiber-starter/app/dto"
	"go-fiber-starter/app/models"
	"go-fiber-starter/app/repository"
	"go-fiber-starter/app/transformer"
	"go-fiber-starter/utils"

	"github.com/gofiber/fiber/v2"
)

type AssistanceTypeService struct{}

func (s *AssistanceTypeService) repo() *repository.AssistanceTypeRepository {
	return new(repository.AssistanceTypeRepository)
}

func (s *AssistanceTypeService) List(ctx *fiber.Ctx) error {
	items, err := s.repo().GetAll(repository.DB)
	if err != nil {
		return utils.JsonErrorInternal(ctx, err, "E_ASSISTANCE_TYPE_LIST")
	}

	rows := transformer.AssistanceTypeListTransformer(items)
	return utils.JsonSuccess(ctx, rows)
}

func (s *AssistanceTypeService) Add(ctx *fiber.Ctx, req dto.AssistanceTypeRequestDTO) error {
	if err := req.Validate(); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	m := models.AssistanceType{
		AssistanceCategoryID: req.AssistanceCategoryID,
		Name:                 req.Name,
	}

	if err := repository.DB.Create(&m).Error; err != nil {
		return utils.JsonErrorInternal(ctx, err, "E_ASSISTANCE_TYPE_ADD")
	}

	return utils.JsonSuccess(ctx, transformer.AssistanceTypeTransformer(&m))
}

func (s *AssistanceTypeService) Detail(ctx *fiber.Ctx, id int) error {
	m, err := s.repo().FindByID(id)
	if err != nil {
		return utils.JsonErrorNotFound(ctx, err)
	}

	return utils.JsonSuccess(ctx, transformer.AssistanceTypeTransformer(&m))
}

func (s *AssistanceTypeService) Update(ctx *fiber.Ctx, id int, req dto.AssistanceTypeRequestDTO) error {
	if err := req.Validate(); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	_, err := s.repo().FindByID(id)
	if err != nil {
		return utils.JsonErrorNotFound(ctx, err)
	}

	store := models.AssistanceType{AssistanceCategoryID: req.AssistanceCategoryID, Name: req.Name}
	updated, err := s.repo().UpdateByID(repository.DB, id, store)
	if err != nil {
		return utils.JsonErrorInternal(ctx, err, "E_ASSISTANCE_TYPE_UPDATE")
	}

	return utils.JsonSuccess(ctx, transformer.AssistanceTypeTransformer(&updated))
}

func (s *AssistanceTypeService) Delete(ctx *fiber.Ctx, id int) error {
	if err := s.repo().DeleteByID(id); err != nil {
		return utils.JsonErrorInternal(ctx, err, "E_ASSISTANCE_TYPE_DELETE")
	}

	return utils.JsonSuccess(ctx, fiber.Map{"id": id})
}
