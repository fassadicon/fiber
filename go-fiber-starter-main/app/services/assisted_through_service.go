package services

import (
	"go-fiber-starter/app/dto"
	"go-fiber-starter/app/models"
	"go-fiber-starter/app/repository"
	"go-fiber-starter/app/transformer"
	"go-fiber-starter/utils"

	"github.com/gofiber/fiber/v2"
)

type AssistedThroughService struct{}

func (s *AssistedThroughService) repo() *repository.AssistedThroughRepository {
	return new(repository.AssistedThroughRepository)
}

func (s *AssistedThroughService) List(ctx *fiber.Ctx) error {
	items, err := s.repo().GetAll(repository.DB)
	if err != nil {
		return utils.JsonErrorInternal(ctx, err, "E_ASSISTED_THROUGH_LIST")
	}

	rows := transformer.AssistedThroughListTransformer(items)
	return utils.JsonSuccess(ctx, rows)
}

func (s *AssistedThroughService) Add(ctx *fiber.Ctx, req dto.AssistedThroughRequestDTO) error {
	if err := req.Validate(); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	m := models.AssistedThrough{
		Name: req.Name,
	}

	if err := repository.DB.Create(&m).Error; err != nil {
		return utils.JsonErrorInternal(ctx, err, "E_ASSISTED_THROUGH_ADD")
	}

	return utils.JsonSuccess(ctx, transformer.AssistedThroughTransformer(&m))
}

func (s *AssistedThroughService) Detail(ctx *fiber.Ctx, id int) error {
	m, err := s.repo().FindByID(id)
	if err != nil {
		return utils.JsonErrorNotFound(ctx, err)
	}

	return utils.JsonSuccess(ctx, transformer.AssistedThroughTransformer(&m))
}

func (s *AssistedThroughService) Update(ctx *fiber.Ctx, id int, req dto.AssistedThroughRequestDTO) error {
	if err := req.Validate(); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	_, err := s.repo().FindByID(id)
	if err != nil {
		return utils.JsonErrorNotFound(ctx, err)
	}

	store := models.AssistedThrough{Name: req.Name}
	updated, err := s.repo().UpdateByID(repository.DB, id, store)
	if err != nil {
		return utils.JsonErrorInternal(ctx, err, "E_ASSISTED_THROUGH_UPDATE")
	}

	return utils.JsonSuccess(ctx, transformer.AssistedThroughTransformer(&updated))
}

func (s *AssistedThroughService) Delete(ctx *fiber.Ctx, id int) error {
	if err := s.repo().DeleteByID(id); err != nil {
		return utils.JsonErrorInternal(ctx, err, "E_ASSISTED_THROUGH_DELETE")
	}

	return utils.JsonSuccess(ctx, fiber.Map{"id": id})
}
