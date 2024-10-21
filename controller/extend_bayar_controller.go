package controller

import (
	"fmt"
	"strconv"
	"wkm/entity"
	"wkm/request"
	"wkm/service"

	"github.com/gofiber/fiber/v2"
)

type ExtendBayarController interface {
	MasterData(ctx *fiber.Ctx) error
	MasterDataCount(ctx *fiber.Ctx) error
	DetailExtendBayar(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	UpdateFa(ctx *fiber.Ctx) error
	UpdateLf(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type extendBayarController struct {
	extendBayarService service.ExtendBayarService
}

func NewExtendBayarController(aS service.ExtendBayarService) ExtendBayarController {
	return &extendBayarController{
		extendBayarService: aS,
	}
}

func (tm *extendBayarController) MasterData(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	pageParams, _ := strconv.Atoi(ctx.Query("pageParams"))
	return ctx.JSON(tm.extendBayarService.MasterData(search, limit, pageParams))
}

func (tm *extendBayarController) MasterDataCount(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	return ctx.JSON(tm.extendBayarService.MasterDataCount(search))
}

func (tm *extendBayarController) DetailExtendBayar(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	return ctx.JSON(tm.extendBayarService.Detail(id))
}

func (tm *extendBayarController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := tm.extendBayarService.Delete(id)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil create"})
}

func (tm *extendBayarController) Create(ctx *fiber.Ctx) error {
	var body request.ExtendBayarRequest
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	user := ctx.Locals("user")
	details, _ := user.(entity.User)

	body.KdUserFa = details.Username
	data, err := tm.extendBayarService.Create(body)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(map[string]interface{}{"message": "Berhasil create", "id_extendBayar": data.Id})
}

func (tm *extendBayarController) UpdateFa(ctx *fiber.Ctx) error {
	var body request.ExtendBayarRequest
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	user := ctx.Locals("user")
	details, _ := user.(entity.User)

	body.KdUserFa = details.Username
	err = tm.extendBayarService.UpdateFa(body)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(map[string]interface{}{"message": "Berhasil update"})
}

func (tm *extendBayarController) UpdateLf(ctx *fiber.Ctx) error {
	var body request.ExtendBayarRequest
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	user := ctx.Locals("user")
	details, _ := user.(entity.User)

	body.KdUserLf = details.Username
	err = tm.extendBayarService.UpdateLf(body)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(map[string]interface{}{"message": "Berhasil update"})
}
