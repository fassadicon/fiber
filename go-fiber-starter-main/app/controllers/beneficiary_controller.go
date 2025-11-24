package controllers

import (
	"go-fiber-starter/app/dto"
	"go-fiber-starter/app/services"
	"go-fiber-starter/utils"

	"github.com/gofiber/fiber/v2"
)

type BeneficiaryController struct{}

func (c *BeneficiaryController) svc() *services.BeneficiaryService {
	return new(services.BeneficiaryService)
}

func (c *BeneficiaryController) List(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ BENEFICIARY LIST")
	return c.svc().List(ctx)
}

func (c *BeneficiaryController) Detail(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ BENEFICIARY DETAIL")
	guid, err := utils.ValidateGUIDParams(ctx)
	if err != nil {
		return err
	}

	return c.svc().Detail(ctx, guid)
}

func (c *BeneficiaryController) Add(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ BENEFICIARY ADD")
	req := new(dto.BeneficiaryRequestDTO)
	if err := ctx.BodyParser(req); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	return c.svc().Add(ctx, *req)
}

func (c *BeneficiaryController) Update(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ BENEFICIARY UPDATE")
	guid, err := utils.ValidateGUIDParams(ctx)
	if err != nil {
		return err
	}

	req := new(dto.BeneficiaryRequestDTO)
	if err := ctx.BodyParser(req); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	return c.svc().Update(ctx, guid, *req)
}

func (c *BeneficiaryController) Delete(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ BENEFICIARY DELETE")
	guid, err := utils.ValidateGUIDParams(ctx)
	if err != nil {
		return err
	}

	return c.svc().Delete(ctx, guid)
}
