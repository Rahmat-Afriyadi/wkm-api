package controller

import (
	"fmt"
	"strconv"
	"wkm/entity"
	"wkm/service"

	"github.com/gofiber/fiber/v2"
)

type VendorController interface {
	MasterData(ctx *fiber.Ctx) error
	MasterDataCount(ctx *fiber.Ctx) error
	DetailMstMtr(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
}

type vendorController struct {
	vendorService service.VendorService
}

func NewVendorController(aS service.VendorService) VendorController {
	return &vendorController{
		vendorService: aS,
	}
}

func (tr *vendorController) MasterData(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	pageParams, _ := strconv.Atoi(ctx.Query("pageParams"))
	return ctx.JSON(tr.vendorService.MasterData(search, limit, pageParams))
}

func (tr *vendorController) MasterDataCount(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	return ctx.JSON(tr.vendorService.MasterDataCount(search))
}

func (tr *vendorController) DetailMstMtr(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	return ctx.JSON(tr.vendorService.Detail(id))
}

func (tr *vendorController) Update(ctx *fiber.Ctx) error {
	var body entity.MasterVendor
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	err = tr.vendorService.Update(body)
	if err != nil {
		return ctx.JSON(map[string]interface{}{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil Update"})
}
func (tr *vendorController) Create(ctx *fiber.Ctx) error {
	var body entity.MasterVendor
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	err = tr.vendorService.Create(body)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil create"})
}
