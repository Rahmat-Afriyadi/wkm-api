package controller

import (
	"wkm/request"
	"wkm/service"

	"github.com/gofiber/fiber/v2"
)

type ECardplusController interface {
	InputBayarEMembership(ctx *fiber.Ctx) error
}

type eCardplusController struct {
	eCardplusService service.ECardplusService
}

func NewECardplusController(aS service.ECardplusService) ECardplusController {
	return &eCardplusController{
		eCardplusService: aS,
	}
}

func (tr *eCardplusController) InputBayarEMembership(ctx *fiber.Ctx) error {
	var body request.InputBayarRequest
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(400).JSON(err)
	}
	
	err := tr.eCardplusService.InputBayarEMembership(body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": err})
	}
	return ctx.JSON(map[string]any{"data": "data", "status": "success"})
}
