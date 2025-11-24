package controllers

import (
	"errors"
	"go-fiber-starter/app/dto"
	"go-fiber-starter/app/services"
	"go-fiber-starter/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AssistanceCategoryController struct{}

func (c *AssistanceCategoryController) svc() *services.AssistanceCategoryService {
	return new(services.AssistanceCategoryService)
}

func (c *AssistanceCategoryController) List(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ ASSISTANCE CATEGORY LIST")
	return c.svc().List(ctx)
}

func (c *AssistanceCategoryController) Detail(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ ASSISTANCE CATEGORY DETAIL")
	idStr := ctx.Params("id")
	if idStr == "" {
		return utils.JsonErrorValidation(ctx, errors.New("id params is required"))
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	return c.svc().Detail(ctx, id)
}

func (c *AssistanceCategoryController) Add(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ ASSISTANCE CATEGORY ADD")
	req := new(dto.AssistanceCategoryRequestDTO)
	if err := ctx.BodyParser(req); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	return c.svc().Add(ctx, *req)
}

func (c *AssistanceCategoryController) Update(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ ASSISTANCE CATEGORY UPDATE")
	idStr := ctx.Params("id")
	if idStr == "" {
		return utils.JsonErrorValidation(ctx, errors.New("id params is required"))
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	req := new(dto.AssistanceCategoryRequestDTO)
	if err := ctx.BodyParser(req); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	return c.svc().Update(ctx, id, *req)
}

func (c *AssistanceCategoryController) Delete(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ ASSISTANCE CATEGORY DELETE")
	idStr := ctx.Params("id")
	if idStr == "" {
		return utils.JsonErrorValidation(ctx, errors.New("id params is required"))
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	return c.svc().Delete(ctx, id)
}
