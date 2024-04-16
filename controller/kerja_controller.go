package controller

import (
	"wkm/service"

	"github.com/gofiber/fiber/v2"
)

type KerjaController interface {
	MasterData(ctx *fiber.Ctx) error
}

type kerjaController struct {
	kerjaService service.KerjaService
}

func NewKerjaController(aS service.KerjaService) KerjaController {
	return &kerjaController{
		kerjaService: aS,
	}
}

func (tr *kerjaController) MasterData(ctx *fiber.Ctx) error {
	return ctx.JSON(tr.kerjaService.MasterData())

}
