package controller

import (
	"strconv"
	"wkm/service"

	"github.com/gofiber/fiber/v2"
)

type MerkController interface {
	MasterData(ctx *fiber.Ctx) error
}

type merkController struct {
	merkService service.MerkService
}

func NewMerkController(aS service.MerkService) MerkController {
	return &merkController{
		merkService: aS,
	}
}

func (tr *merkController) MasterData(ctx *fiber.Ctx) error {
	jenis_kendaraan, _ := strconv.Atoi(ctx.Params("jenisKendaraan"))
	return ctx.JSON(tr.merkService.MasterData(jenis_kendaraan))
}
