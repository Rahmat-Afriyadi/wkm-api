package controller

import (
	"fmt"
	"strconv"
	"wkm/request"
	"wkm/service"

	"github.com/gofiber/fiber/v2"
)

type StockCardController interface {
	MasterData(ctx *fiber.Ctx) error
	MasterDataCount(ctx *fiber.Ctx) error
	DetailMstMtr(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
}

type stockCardController struct {
	stockCardService service.StockCardService
}

func NewStockCardController(aS service.StockCardService) StockCardController {
	return &stockCardController{
		stockCardService: aS,
	}
}

func (tm *stockCardController) MasterData(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	pageParams, _ := strconv.Atoi(ctx.Query("pageParams"))
	return ctx.JSON(tm.stockCardService.MasterData(search, limit, pageParams))
}

func (tm *stockCardController) MasterDataCount(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	return ctx.JSON(tm.stockCardService.MasterDataCount(search))
}

func (tm *stockCardController) DetailMstMtr(ctx *fiber.Ctx) error {
	noKartu := ctx.Params("noKartu")
	return ctx.JSON(tm.stockCardService.Detail(noKartu))
}

func (tm *stockCardController) Update(ctx *fiber.Ctx) error {
	var body request.StockCardRequest
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	err = tm.stockCardService.Update(body)
	if err != nil {
		return ctx.JSON(map[string]interface{}{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil Update"})
}

func (tm *stockCardController) Create(ctx *fiber.Ctx) error {
	var body request.StockCardRequest
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}

	data, err := tm.stockCardService.Create(body)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(map[string]interface{}{"message": "Berhasil create", "noKartu_stockCard": data.NoKartu})
}
