package controllers

import (
	"errors"
	"go-fiber-starter/app/dto"
	"go-fiber-starter/app/services"
	"go-fiber-starter/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ClientBeneficiaryController struct{}

func (c *ClientBeneficiaryController) svc() *services.ClientBeneficiaryService {
	return new(services.ClientBeneficiaryService)
}

func (c *ClientBeneficiaryController) List(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ CLIENT BENEFICIARY LIST")
	return c.svc().List(ctx)
}

func (c *ClientBeneficiaryController) Add(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ CLIENT BENEFICIARY ADD")
	req := new(dto.ClientBeneficiaryRequestDTO)
	if err := ctx.BodyParser(req); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	return c.svc().Add(ctx, *req)
}

func (c *ClientBeneficiaryController) Delete(ctx *fiber.Ctx) error {
	utils.Logger.Info("✅ CLIENT BENEFICIARY DELETE")
	clientStr := ctx.Params("client_id")
	beneStr := ctx.Params("beneficiary_id")
	if clientStr == "" || beneStr == "" {
		return utils.JsonErrorValidation(ctx, errors.New("client_id and beneficiary_id params are required"))
	}
	clientID, err := strconv.Atoi(clientStr)
	if err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}
	beneID, err := strconv.Atoi(beneStr)
	if err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}

	return c.svc().Delete(ctx, clientID, beneID)
}
