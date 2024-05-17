package controller

import (
	"wkm/service"

	"github.com/gofiber/fiber/v2"
)

type ProdukController interface {
	MasterData(ctx *fiber.Ctx) error
}

type produkController struct {
	produkService service.ProdukService
}

func NewProdukController(aS service.ProdukService) ProdukController {
	return &produkController{
		produkService: aS,
	}
}

func (tr *produkController) MasterData(ctx *fiber.Ctx) error {

	return ctx.JSON(tr.produkService.MasterData(ctx.Query("search"), ctx.Query("jenis_asuransi")))

}
