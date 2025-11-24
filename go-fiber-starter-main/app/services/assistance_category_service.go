package services

import (
	"go-fiber-starter/app/dto"
	"go-fiber-starter/app/models"
	"go-fiber-starter/app/repository"
	"go-fiber-starter/app/transformer"
	"go-fiber-starter/utils"

	"github.com/gofiber/fiber/v2"
)

type AssistanceCategoryService struct{}

func (s *AssistanceCategoryService) repo() *repository.AssistanceCategoryRepository {
	return new(repository.AssistanceCategoryRepository)
}

func (s *AssistanceCategoryService) List(ctx *fiber.Ctx) error {
	items, err := s.repo().GetAll(repository.DB)
	if err != nil {
		return utils.JsonErrorInternal(ctx, err, "E_ASSISTANCE_CATEGORY_LIST")
	}

	rows := transformer.AssistanceCategoryListTransformer(items)
	return utils.JsonSuccess(ctx, rows)
}

func (s *AssistanceCategoryService) Add(ctx *fiber.Ctx, req dto.AssistanceCategoryRequestDTO) error {
	if err := req.Validate(); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	m := models.AssistanceCategory{
		Name: req.Name,
	}

	if err := repository.DB.Create(&m).Error; err != nil {
		return utils.JsonErrorInternal(ctx, err, "E_ASSISTANCE_CATEGORY_ADD")
	}

	return utils.JsonSuccess(ctx, transformer.AssistanceCategoryTransformer(&m))
}

func (s *AssistanceCategoryService) Detail(ctx *fiber.Ctx, id int) error {
	m, err := s.repo().FindByID(id)
	if err != nil {
		return utils.JsonErrorNotFound(ctx, err)
	}

	return utils.JsonSuccess(ctx, transformer.AssistanceCategoryTransformer(&m))
}

func (s *AssistanceCategoryService) Update(ctx *fiber.Ctx, id int, req dto.AssistanceCategoryRequestDTO) error {
	if err := req.Validate(); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	_, err := s.repo().FindByID(id)
	if err != nil {
		return utils.JsonErrorNotFound(ctx, err)
	}

	store := models.AssistanceCategory{Name: req.Name}
	updated, err := s.repo().UpdateByID(repository.DB, id, store)
	if err != nil {
		return utils.JsonErrorInternal(ctx, err, "E_ASSISTANCE_CATEGORY_UPDATE")
	}

	return utils.JsonSuccess(ctx, transformer.AssistanceCategoryTransformer(&updated))
}

func (s *AssistanceCategoryService) Delete(ctx *fiber.Ctx, id int) error {
	if err := s.repo().DeleteByID(id); err != nil {
		return utils.JsonErrorInternal(ctx, err, "E_ASSISTANCE_CATEGORY_DELETE")
	}

	return utils.JsonSuccess(ctx, fiber.Map{"id": id})
}
