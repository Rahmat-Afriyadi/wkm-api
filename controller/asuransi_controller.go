package controller

import (
	"fmt"
	"wkm/entity"
	"wkm/service"

	"github.com/gofiber/fiber/v2"
)

type AsuransiController interface {
	MasterData(ctx *fiber.Ctx) error
	MasterDataPending(ctx *fiber.Ctx) error
	MasterDataOke(ctx *fiber.Ctx) error
	FindAsuransiByNoMsn(ctx *fiber.Ctx) error
	UpdateAsuransi(ctx *fiber.Ctx) error
	UpdateAmbilAsuransi(ctx *fiber.Ctx) error
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
	dataSource := ctx.Query("dataSource")
	return ctx.JSON(tr.asuransiService.MasterData(dataSource))
}

func (tr *asuransiController) MasterDataPending(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	dataSource := ctx.Query("dataSource")
	return ctx.JSON(tr.asuransiService.MasterDataPending(search, dataSource))
}

func (tr *asuransiController) MasterDataOke(ctx *fiber.Ctx) error {
	dataSource := ctx.Query("dataSource")
	return ctx.JSON(tr.asuransiService.MasterDataOke(dataSource))
}

func (tr *asuransiController) FindAsuransiByNoMsn(ctx *fiber.Ctx) error {
	no_msn := ctx.Params("no_msn")
	return ctx.JSON(tr.asuransiService.FindAsuransiByNoMsn(no_msn))

}

func (tr *asuransiController) UpdateAsuransi(ctx *fiber.Ctx) error {
	var asuransi entity.MasterAsuransi
	err := ctx.BodyParser(&asuransi)
	fmt.Println("ini body ", asuransi)
	user := ctx.Locals("user")
	details, _ := user.(entity.UserAsuransi)
	asuransi.KdUser = details.Username
	fmt.Println("ini kd user ", details.Username)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	tr.asuransiService.UpdateAsuransi(asuransi)
	return ctx.JSON("Hallo guys")

}

func (tr *asuransiController) UpdateAmbilAsuransi(ctx *fiber.Ctx) error {
	var asuransi entity.MasterAsuransi
	err := ctx.BodyParser(&asuransi)
	fmt.Println("ini body ", asuransi)
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	asuransi.KdUser = details.Username
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	tr.asuransiService.UpdateAmbilAsuransi(asuransi.NoMsn, asuransi.KdUser)
	return ctx.JSON("Berhasil")

}
