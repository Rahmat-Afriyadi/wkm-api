package controller

import (
	"fmt"
	"wkm/entity"
	"wkm/service"

	"github.com/gofiber/fiber/v2"
)

type AsuransiController interface {
	MasterData(ctx *fiber.Ctx) error
	FindAsuransiByNoMsn(ctx *fiber.Ctx) error
	UpdateAsuransi(ctx *fiber.Ctx) error
}

type asuransiController struct {
	asuransiService service.AsuransiService
}

func NewAsuransiController(aS service.AsuransiService) AsuransiController {
	return &asuransiController{
		asuransiService: aS,
	}
}

func (tr *asuransiController) MasterData(ctx *fiber.Ctx) error {
	return ctx.JSON(tr.asuransiService.MasterData())

}

func (tr *asuransiController) FindAsuransiByNoMsn(ctx *fiber.Ctx) error {
	no_msn := ctx.Params("no_msn")
	return ctx.JSON(tr.asuransiService.FindAsuransiByNoMsn(no_msn))

}

func (tr *asuransiController) UpdateAsuransi(ctx *fiber.Ctx) error {
	var asuransi entity.MasterAsuransi
	err := ctx.BodyParser(&asuransi)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	return ctx.JSON(tr.asuransiService.UpdateAsuransi(asuransi))

}
