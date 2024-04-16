package controller

import (
	"wkm/service"

	"github.com/gofiber/fiber/v2"
)

type LeasController interface {
	MasterData(ctx *fiber.Ctx) error
}

type leasController struct {
	leasService service.LeasService
}

func NewLeasController(aS service.LeasService) LeasController {
	return &leasController{
		leasService: aS,
	}
}

func (tr *leasController) MasterData(ctx *fiber.Ctx) error {
	return ctx.JSON(tr.leasService.MasterData())

}
