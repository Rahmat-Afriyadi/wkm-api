package controller

import (
	"fmt"
	"strconv"
	"wkm/entity"
	"wkm/service"

	"github.com/gofiber/fiber/v2"
)

type MstMtrController interface {
	CreateMstMtr(ctx *fiber.Ctx) error
	UpdateMstMtr(ctx *fiber.Ctx) error
	MasterData(ctx *fiber.Ctx) error
	MasterDataCount(ctx *fiber.Ctx) error
	DetailMstMtr(ctx *fiber.Ctx) error
}

type mstMtrController struct {
	mstMtrService service.MstMtrService
}

func NewMstMtrController(aS service.MstMtrService) MstMtrController {
	return &mstMtrController{
		mstMtrService: aS,
	}
}

func (tr *mstMtrController) DetailMstMtr(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	return ctx.JSON(tr.mstMtrService.DetailMstMtr(id))
}

func (tr *mstMtrController) MasterData(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	pageParams, _ := strconv.Atoi(ctx.Query("pageParams"))
	return ctx.JSON(tr.mstMtrService.MasterData(search, limit, pageParams))
}

func (tr *mstMtrController) MasterDataCount(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	return ctx.JSON(tr.mstMtrService.MasterDataCount(search))
}

func (tr *mstMtrController) UpdateMstMtr(ctx *fiber.Ctx) error {
	var body entity.MstMtr
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	err = tr.mstMtrService.Update(body)
	if err != nil {
		return ctx.JSON(map[string]interface{}{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil Update"})
}
func (tr *mstMtrController) CreateMstMtr(ctx *fiber.Ctx) error {
	var body entity.MstMtr
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	err = tr.mstMtrService.CreateMstMtr(body)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil create"})
}
