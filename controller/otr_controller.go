package controller

import (
	"fmt"
	"strconv"
	"wkm/entity"
	"wkm/request"
	"wkm/service"

	"github.com/gofiber/fiber/v2"
)

type OtrController interface {
	DetailOtrNa(ctx *fiber.Ctx) error
	OtrNaList(ctx *fiber.Ctx) error
	OtrMstProduk(ctx *fiber.Ctx) error
	OtrMstNa(ctx *fiber.Ctx) error
	CreateOtr(ctx *fiber.Ctx) error
	UpdateOtr(ctx *fiber.Ctx) error
	MasterData(ctx *fiber.Ctx) error
	MasterDataCount(ctx *fiber.Ctx) error
	DetailOtr(ctx *fiber.Ctx) error
	DetailOtrByNoMtr(ctx *fiber.Ctx) error
}

type otrController struct {
	otrService service.OtrService
}

func NewOtrController(aS service.OtrService) OtrController {
	return &otrController{
		otrService: aS,
	}
}

func (tr *otrController) DetailOtr(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	return ctx.JSON(tr.otrService.DetailOtr(id))
}

func (tr *otrController) MasterData(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	pageParams, _ := strconv.Atoi(ctx.Query("pageParams"))
	return ctx.JSON(tr.otrService.MasterData(search, limit, pageParams))
}

func (tr *otrController) MasterDataCount(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	return ctx.JSON(tr.otrService.MasterDataCount(search))
}

func (tr *otrController) OtrNaList(ctx *fiber.Ctx) error {
	return ctx.JSON(tr.otrService.OtrNaList())
}
func (tr *otrController) UpdateOtr(ctx *fiber.Ctx) error {
	var body entity.Otr
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	err = tr.otrService.Update(body)
	if err != nil {
		return ctx.JSON(map[string]interface{}{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil Update"})
}
func (tr *otrController) CreateOtr(ctx *fiber.Ctx) error {
	var body request.CreateOtr
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	tr.otrService.CreateOtr(body)
	return ctx.JSON(map[string]string{"message": "Berhasil create"})
}

func (tr *otrController) OtrMstProduk(ctx *fiber.Ctx) error {
	return ctx.JSON(tr.otrService.OtrMstProduk(ctx.Query("search")))
}
func (tr *otrController) OtrMstNa(ctx *fiber.Ctx) error {
	return ctx.JSON(tr.otrService.OtrMstNa(ctx.Query("search")))
}

func (tr *otrController) DetailOtrNa(ctx *fiber.Ctx) error {
	tahun, _ := strconv.ParseUint(ctx.Query("tahun"), 10, 32)
	return ctx.JSON(tr.otrService.DetailOtrNa(ctx.Query("motorprice_kode"), uint16(tahun)))

}

func (tr *otrController) DetailOtrByNoMtr(ctx *fiber.Ctx) error {
	var otr entity.Otr
	err := ctx.BodyParser(&otr)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	fmt.Println("ni body yaa ", otr.Tahun, otr.NoMtr)
	return ctx.JSON(tr.otrService.DetailOtrByNoMtr(otr.NoMtr, otr.Tahun))

}
