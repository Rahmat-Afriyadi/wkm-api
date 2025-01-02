package controller

import (
	"wkm/service"

	"github.com/gofiber/fiber/v2"
)

type DlrController interface {
	MasterData(ctx *fiber.Ctx) error
}

type dlrController struct {
	dlrService service.DlrService
}

func NewDlrController(aS service.DlrService) DlrController {
	return &dlrController{
		dlrService: aS,
	}
}

func (tr *dlrController) MasterData(ctx *fiber.Ctx) error {
	return ctx.JSON(tr.dlrService.MasterData(ctx.Query("search")))

}
