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
	UpdateAsuransiBerminat(ctx *fiber.Ctx) error
	UpdateAsuransiBatalBayar(ctx *fiber.Ctx) error
	UpdateAmbilAsuransi(ctx *fiber.Ctx) error
	MasterDataRekapTele(ctx *fiber.Ctx) error
	RekapByStatus(ctx *fiber.Ctx) error
}

type asuransiController struct {
	asuransiService service.AsuransiService
}

func NewAsuransiController(aS service.AsuransiService) AsuransiController {
	return &asuransiController{
		asuransiService: aS,
	}
}

func (tr *asuransiController) MasterDataRekapTele(ctx *fiber.Ctx) error {
	return ctx.JSON(tr.asuransiService.MasterDataRekapTele())
}
func (tr *asuransiController) MasterData(ctx *fiber.Ctx) error {
	dataSource := ctx.Query("dataSource")
	sts := ctx.Params("status")
	search := ctx.Query("search")
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	return ctx.JSON(tr.asuransiService.MasterData(search, dataSource, sts, details.Username))
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
	details, _ := user.(entity.User)
	asuransi.KdUser = details.Username
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	tr.asuransiService.UpdateAsuransi(asuransi)
	return ctx.JSON("Hallo guys")

}

func (tr *asuransiController) RekapByStatus(ctx *fiber.Ctx) error {
	user := ctx.Locals("user")
	tgl := ctx.Query("tgl")
	details, _ := user.(entity.User)
	result := tr.asuransiService.RekapByStatus(details.Username, tgl)
	return ctx.JSON(result)

}

func (tr *asuransiController) UpdateAsuransiBerminat(ctx *fiber.Ctx) error {
	var body map[string]interface{}
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	fmt.Println("ini body ", body["no_msn"])
	tr.asuransiService.UpdateAsuransiBerminat(body["no_msn"].(string))
	return ctx.JSON("Hallo guys")

}

func (tr *asuransiController) UpdateAsuransiBatalBayar(ctx *fiber.Ctx) error {
	var body map[string]interface{}
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	fmt.Println("ini body ", body["no_msn"])
	tr.asuransiService.UpdateAsuransiBatalBayar(body["no_msn"].(string))
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
