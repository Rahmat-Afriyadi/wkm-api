package controller

import (
	"strconv"
	"wkm/service"

	"github.com/gofiber/fiber/v2"
)

type ProdukController interface {
	MasterData(ctx *fiber.Ctx) error
	MasterDataCount(ctx *fiber.Ctx) error
	DetailMstMtr(ctx *fiber.Ctx) error
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
	search := ctx.Query("search")
	jenis_asuransi, _ := strconv.Atoi(ctx.Query("jenis_asuransi"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	pageParams, _ := strconv.Atoi(ctx.Query("pageParams"))
	return ctx.JSON(tr.produkService.MasterData(search, jenis_asuransi, limit, pageParams))
}

func (tr *produkController) MasterDataCount(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	jenis_asuransi, _ := strconv.Atoi(ctx.Query("jenis_asuransi"))
	return ctx.JSON(tr.produkService.MasterDataCount(search, jenis_asuransi))
}

func (tr *produkController) DetailMstMtr(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	return ctx.JSON(tr.produkService.DetailMstMtr(id))
}
