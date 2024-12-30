package controller

import (
	// "fmt"
	// "strconv"
	// "wkm/entity"
	"wkm/service"

	"github.com/gofiber/fiber/v2"
)

type MstController interface {
	ListClientUser(ctx *fiber.Ctx) error
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
