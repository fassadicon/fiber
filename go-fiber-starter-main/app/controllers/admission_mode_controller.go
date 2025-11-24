package controllers

import (
	"errors"
	"go-fiber-starter/app/dto"
	"go-fiber-starter/app/services"
	"go-fiber-starter/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AdmissionModeController struct{}

func (c *AdmissionModeController) svc() *services.AdmissionModeService {
	return new(services.AdmissionModeService)
}

func (c *AdmissionModeController) List(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ ADMISSION MODE LIST")
	return c.svc().List(ctx)
}

func (c *AdmissionModeController) Detail(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ ADMISSION MODE DETAIL")
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

func (c *AdmissionModeController) Add(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ ADMISSION MODE ADD")
	req := new(dto.AdmissionModeRequestDTO)
	if err := ctx.BodyParser(req); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	return c.svc().Add(ctx, *req)
}

func (c *AdmissionModeController) Update(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ ADMISSION MODE UPDATE")
	idStr := ctx.Params("id")
	if idStr == "" {
		return utils.JsonErrorValidation(ctx, errors.New("id params is required"))
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	req := new(dto.AdmissionModeRequestDTO)
	if err := ctx.BodyParser(req); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	return c.svc().Update(ctx, id, *req)
}

func (c *AdmissionModeController) Delete(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ ADMISSION MODE DELETE")
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
