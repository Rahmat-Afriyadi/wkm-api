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
	UpdateApprovalLf(ctx *fiber.Ctx) error
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
	tgl1 := ctx.Query("tgl1")
	tgl2 := ctx.Query("tgl2")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	pageParams, _ := strconv.Atoi(ctx.Query("pageParams"))
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	if details.Role.Name == "LEADER_FA" {
		return ctx.JSON(tm.extendBayarService.MasterDataLf(search, tgl1, tgl2, limit, pageParams))
	}
	return ctx.JSON(tm.extendBayarService.MasterData(search, tgl1, tgl2, limit, pageParams))
}

func (tm *extendBayarController) MasterDataCount(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	user := ctx.Locals("user")
	tgl1 := ctx.Query("tgl1")
	tgl2 := ctx.Query("tgl2")
	details, _ := user.(entity.User)
	if details.Role.Name == "LEADER_FA" {
		return ctx.JSON(tm.extendBayarService.MasterDataLfCount(search, tgl1, tgl2))
	}
	return ctx.JSON(tm.extendBayarService.MasterDataCount(search, tgl1, tgl2))
}

func (tm *extendBayarController) DetailExtendBayar(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	return ctx.JSON(tm.extendBayarService.Detail(id))
}

func (tm *extendBayarController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	err := tm.extendBayarService.Delete(id, details.Username)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil Delete"})
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
		return ctx.JSON(map[string]string{"message": err.Error(), "status": "fail"})
	}
	return ctx.JSON(map[string]interface{}{"message": "Pengajuan Extend Bayar Berhasil", "status": "success", "id": data.Id})
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
		return ctx.JSON(map[string]string{"message": err.Error(), "status": "fail"})
	}
	return ctx.JSON(map[string]interface{}{"message": "Berhasil update", "status": "success"})
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
		return ctx.JSON(map[string]string{"message": err.Error(), "status": "fail"})
	}
	return ctx.JSON(map[string]interface{}{"message": "Berhasil update", "status": "success"})
}

func (tm *extendBayarController) UpdateApprovalLf(ctx *fiber.Ctx) error {
	var body request.ExtendBayarApprovalRequest
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser kesini yaa", err)
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	user := ctx.Locals("user")
	details, _ := user.(entity.User)

	if details.Role.Name != "LEADER_FA" {
		return ctx.Status(403).JSON(map[string]string{"message": "Kamu bukan Leader FA", "status": "fail"})
	}
	body.KdUserLf = details.Username
	err = tm.extendBayarService.UpdateApprovalLf(body)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error(), "status": "fail"})
	}
	return ctx.JSON(map[string]interface{}{"message": "Berhasil update", "status": "success"})
}
