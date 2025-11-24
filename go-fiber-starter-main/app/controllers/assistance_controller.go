package controllers

import (
	"errors"
	"go-fiber-starter/app/dto"
	"go-fiber-starter/app/services"
	"go-fiber-starter/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AssistanceController struct{}

func (c *AssistanceController) svc() *services.AssistanceService {
	return new(services.AssistanceService)
}

func (c *AssistanceController) List(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ ASSISTANCE LIST")
	return c.svc().List(ctx)
}

func (c *AssistanceController) Detail(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ ASSISTANCE DETAIL")
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

func (c *AssistanceController) Add(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ ASSISTANCE ADD")
	req := new(dto.AssistanceRequestDTO)
	if err := ctx.BodyParser(req); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	return c.svc().Add(ctx, *req)
}

func (c *AssistanceController) Update(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ ASSISTANCE UPDATE")
	idStr := ctx.Params("id")
	if idStr == "" {
		return utils.JsonErrorValidation(ctx, errors.New("id params is required"))
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	req := new(dto.AssistanceRequestDTO)
	if err := ctx.BodyParser(req); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	return c.svc().Update(ctx, id, *req)
}

func (c *AssistanceController) Delete(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ ASSISTANCE DELETE")
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
