package controller

import (
	"wkm/service"

	"github.com/gofiber/fiber/v2"
)

type KodeposController interface {
	MasterData(ctx *fiber.Ctx) error
}

type kodeposController struct {
	kodeposService service.KodeposService
}

func NewKodeposController(aS service.KodeposService) KodeposController {
	return &kodeposController{
		kodeposService: aS,
	}
}

func (tr *kodeposController) MasterData(ctx *fiber.Ctx) error {
	return ctx.JSON(tr.kodeposService.MasterData(ctx.Query("search")))

}
