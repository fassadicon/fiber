package controllers

import (
	"errors"
	"go-fiber-starter/app/dto"
	"go-fiber-starter/app/services"
	"go-fiber-starter/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AssistanceTypeController struct{}

func (c *AssistanceTypeController) svc() *services.AssistanceTypeService {
	return new(services.AssistanceTypeService)
}

func (c *AssistanceTypeController) List(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ ASSISTANCE TYPE LIST")
	return c.svc().List(ctx)
}

func (c *AssistanceTypeController) Detail(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ ASSISTANCE TYPE DETAIL")
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

func (c *AssistanceTypeController) Add(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ ASSISTANCE TYPE ADD")
	req := new(dto.AssistanceTypeRequestDTO)
	if err := ctx.BodyParser(req); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	return c.svc().Add(ctx, *req)
}

func (c *AssistanceTypeController) Update(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ ASSISTANCE TYPE UPDATE")
	idStr := ctx.Params("id")
	if idStr == "" {
		return utils.JsonErrorValidation(ctx, errors.New("id params is required"))
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	req := new(dto.AssistanceTypeRequestDTO)
	if err := ctx.BodyParser(req); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	return c.svc().Update(ctx, id, *req)
}

func (c *AssistanceTypeController) Delete(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ ASSISTANCE TYPE DELETE")
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
