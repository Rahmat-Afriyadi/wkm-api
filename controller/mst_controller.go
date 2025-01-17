package controller

import (
	// "fmt"
	// "strconv"
	// "wkm/entity"
	"wkm/response"
	"wkm/service"

	"github.com/gofiber/fiber/v2"
)

type MstController interface {
	ListClientUser(ctx *fiber.Ctx) error
	MasterAgama(ctx *fiber.Ctx) error
	MasterTujuPak(ctx *fiber.Ctx) error
	MasterPendidikan(ctx *fiber.Ctx) error
	MasterKeluarBln(ctx *fiber.Ctx) error
	MasterAktivitasJual(ctx *fiber.Ctx) error
}

type mstController struct {
	mstService service.MstService
}

func NewMstController(mS service.MstService) MstController {
	return &mstController{
		mstService: mS,
	}
}

func (mS *mstController) ListClientUser(ctx *fiber.Ctx) error {
    // Memanggil service untuk mendapatkan daftar semua user
    users, err := mS.mstService.ListClientUser()
    if err != nil {
        // Menangani error jika ada
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    // Mengembalikan hasil daftar user dengan status OK
    return ctx.Status(fiber.StatusOK).JSON(users)
}
func (mS *mstController) MasterAgama(ctx *fiber.Ctx) error {
    data := mS.mstService.MasterAgama()
	var res []response.Choices
	for _, v := range data {
		res = append(res, response.Choices{Name:v.Agama, Value:v.Id})
	}
    return ctx.Status(fiber.StatusOK).JSON(res)
}
func (mS *mstController) MasterTujuPak(ctx *fiber.Ctx) error {
    data := mS.mstService.MasterTujuPak()
	var res []response.Choices
	for _, v := range data {
		res = append(res, response.Choices{Name:v.NmTupak, Value:v.Id})
	}
    return ctx.Status(fiber.StatusOK).JSON(res)
}
func (mS *mstController) MasterPendidikan(ctx *fiber.Ctx) error {
    data := mS.mstService.MasterPendidikan()
	var res []response.Choices
	for _, v := range data {
		res = append(res, response.Choices{Name:v.NmPendidikan, Value:v.Id})
	}
    return ctx.Status(fiber.StatusOK).JSON(res)
}
func (mS *mstController) MasterKeluarBln(ctx *fiber.Ctx) error {
    data := mS.mstService.MasterKeluarBln()
	var res []response.Choices
	for _, v := range data {
		res = append(res, response.Choices{Name:v.KeluarBln, Value:v.Id})
	}
    return ctx.Status(fiber.StatusOK).JSON(res)
}
func (mS *mstController) MasterAktivitasJual(ctx *fiber.Ctx) error {
    data := mS.mstService.MasterAktivitasJual()
	var res []response.Choices
	for _, v := range data {
		res = append(res, response.Choices{Name:v.AktivitasJual, Value:v.Id})
	}
    return ctx.Status(fiber.StatusOK).JSON(res)
}
