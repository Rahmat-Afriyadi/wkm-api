package controller

import (
	"wkm/response"
	"wkm/service"

	"github.com/gofiber/fiber/v2"
)

type KerjaController interface {
	MasterData(ctx *fiber.Ctx) error
	MasterDataChoices(ctx *fiber.Ctx) error
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
func (tr *kerjaController) MasterDataChoices(ctx *fiber.Ctx) error {
	var res []response.Choices
	data := tr.kerjaService.MasterData()
	for _, v := range data {
		res = append(res, response.Choices{Value: v.Kode, Name: v.Nama})
	}
	return ctx.JSON(res)

}
