package controllers

import (
	"errors"
	"go-fiber-starter/app/dto"
	"go-fiber-starter/app/services"
	"go-fiber-starter/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AssistedThroughController struct{}

func (c *AssistedThroughController) svc() *services.AssistedThroughService {
	return new(services.AssistedThroughService)
}

func (c *AssistedThroughController) List(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ ASSISTED THROUGH LIST")
	return c.svc().List(ctx)
}

func (c *AssistedThroughController) Detail(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ ASSISTED THROUGH DETAIL")
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

func (c *AssistedThroughController) Add(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ ASSISTED THROUGH ADD")
	req := new(dto.AssistedThroughRequestDTO)
	if err := ctx.BodyParser(req); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	return c.svc().Add(ctx, *req)
}

func (c *AssistedThroughController) Update(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ ASSISTED THROUGH UPDATE")
	idStr := ctx.Params("id")
	if idStr == "" {
		return utils.JsonErrorValidation(ctx, errors.New("id params is required"))
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	req := new(dto.AssistedThroughRequestDTO)
	if err := ctx.BodyParser(req); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	return c.svc().Update(ctx, id, *req)
}

func (c *AssistedThroughController) Delete(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ ASSISTED THROUGH DELETE")
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
