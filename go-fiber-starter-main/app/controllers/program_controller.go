package controllers

import (
	"errors"
	"go-fiber-starter/app/dto"
	"go-fiber-starter/app/services"
	"go-fiber-starter/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProgramController struct{}

func (c *ProgramController) svc() *services.ProgramService {
	return new(services.ProgramService)
}

func (c *ProgramController) List(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ PROGRAM LIST")
	return c.svc().List(ctx)
}

func (c *ProgramController) Detail(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ PROGRAM DETAIL")
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

func (c *ProgramController) Add(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ PROGRAM ADD")
	req := new(dto.ProgramRequestDTO)
	if err := ctx.BodyParser(req); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	return c.svc().Add(ctx, *req)
}

func (c *ProgramController) Update(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ PROGRAM UPDATE")
	idStr := ctx.Params("id")
	if idStr == "" {
		return utils.JsonErrorValidation(ctx, errors.New("id params is required"))
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	req := new(dto.ProgramRequestDTO)
	if err := ctx.BodyParser(req); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	return c.svc().Update(ctx, id, *req)
}

func (c *ProgramController) Delete(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ PROGRAM DELETE")
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
