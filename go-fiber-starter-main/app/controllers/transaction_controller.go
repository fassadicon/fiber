package controllers

import (
	"errors"
	"go-fiber-starter/app/dto"
	"go-fiber-starter/app/services"
	"go-fiber-starter/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TransactionController struct{}

func (c *TransactionController) svc() *services.TransactionService {
	return new(services.TransactionService)
}

func (c *TransactionController) List(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ TRANSACTION LIST")
	return c.svc().List(ctx)
}

func (c *TransactionController) Detail(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ TRANSACTION DETAIL")
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

func (c *TransactionController) Add(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ TRANSACTION ADD")
	req := new(dto.TransactionRequestDTO)
	if err := ctx.BodyParser(req); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	return c.svc().Add(ctx, *req)
}

func (c *TransactionController) Update(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ TRANSACTION UPDATE")
	idStr := ctx.Params("id")
	if idStr == "" {
		return utils.JsonErrorValidation(ctx, errors.New("id params is required"))
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	req := new(dto.TransactionRequestDTO)
	if err := ctx.BodyParser(req); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	return c.svc().Update(ctx, id, *req)
}

func (c *TransactionController) Delete(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ TRANSACTION DELETE")
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
