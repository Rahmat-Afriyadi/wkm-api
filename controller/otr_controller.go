package controller

import (
	"fmt"
	"strconv"
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
}

type otrController struct {
	otrService service.OtrService
}

func NewOtrController(aS service.OtrService) OtrController {
	return &otrController{
		otrService: aS,
	}
}

func (tr *otrController) OtrNaList(ctx *fiber.Ctx) error {
	return ctx.JSON(tr.otrService.OtrNaList())
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
