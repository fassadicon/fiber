package controllers

import (
	"go-fiber-starter/app/dto"
	"go-fiber-starter/app/services"
	"go-fiber-starter/utils"

	"github.com/gofiber/fiber/v2"
)

type ClientController struct{}

func (c *ClientController) svc() *services.ClientService {
	return new(services.ClientService)
}

func (c *ClientController) List(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ CLIENT LIST")
	return c.svc().List(ctx)
}

func (c *ClientController) Detail(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ CLIENT DETAIL")
	guid, err := utils.ValidateGUIDParams(ctx)
	if err != nil {
		return err
	}

	return c.svc().Detail(ctx, guid)
}

func (c *ClientController) Add(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ CLIENT ADD")
	req := new(dto.ClientRequestDTO)
	if err := ctx.BodyParser(req); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	return c.svc().Add(ctx, *req)
}

func (c *ClientController) Update(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ CLIENT UPDATE")
	guid, err := utils.ValidateGUIDParams(ctx)
	if err != nil {
		return err
	}

	req := new(dto.ClientRequestDTO)
	if err := ctx.BodyParser(req); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	return c.svc().Update(ctx, guid, *req)
}

func (c *ClientController) Delete(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ CLIENT DELETE")
	guid, err := utils.ValidateGUIDParams(ctx)
	if err != nil {
		return err
	}

	return c.svc().Delete(ctx, guid)
}
