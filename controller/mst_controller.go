package controller

import (
	"fmt"
	// "strconv"
	"wkm/entity"
	// "wkm/request"
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
	CreateScript(ctx *fiber.Ctx) error
	UpdateScript(ctx *fiber.Ctx) error
	MasterScript(ctx *fiber.Ctx) error
	ListAllScript(ctx *fiber.Ctx) error
	ViewScript(ctx *fiber.Ctx) error
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
func (mS *mstController) MasterScript(ctx *fiber.Ctx) error {
    data := mS.mstService.MasterScript()
    return ctx.Status(fiber.StatusOK).JSON(data)
}
func (mS *mstController) ListAllScript(ctx *fiber.Ctx) error {
    data := mS.mstService.ListAllScript()
    return ctx.Status(fiber.StatusOK).JSON(data)
}
func (mS *mstController) CreateScript(ctx *fiber.Ctx) error {
	// Dekode request JSON menjadi struct
	var scriptRequest entity.MstScript
	if err := ctx.BodyParser(&scriptRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}

	// Ambil user details dari context
	user := ctx.Locals("user")
	details, ok := user.(entity.User) // Pastikan type assertion berhasil
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Unauthorized",
			"details": "Invalid user context",
		})
	}


	err := mS.mstService.CreateScript(scriptRequest, details.Username)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to create script",
			"details": err.Error(),
		})
	}

	// Kembalikan response sukses dengan nomor tiket
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":   "Script created successfully",
	})
}

func (mS *mstController) UpdateScript(ctx *fiber.Ctx) error {
	// Dekode request JSON menjadi struct
	var scriptRequest entity.MstScript
	if err := ctx.BodyParser(&scriptRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}

	id := ctx.Params("id")
    if id == "" {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error":   "Invalid request",
            "details": "script tidak ditemukan",
        })
    }

	user := ctx.Locals("user")
	details, ok := user.(entity.User) // Pastikan type assertion berhasil
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Unauthorized",
			"details": "Invalid user context",
		})
	}

	err := mS.mstService.UpdateScript(id, scriptRequest, details.Username)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to create ticket",
			"details": err.Error(),
		})
	}

	// Kembalikan response sukses dengan nomor tiket
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":   "Script created successfully",
	})
}

func (mS *mstController) ViewScript(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
    if id == "" {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error":   "Invalid request",
            "details": "script tidak ditemukan",
        })
    }

	data, err := mS.mstService.ViewScript(id)
	if err != nil {
		// Jika terjadi error (misalnya script tidak ditemukan), kembalikan error response
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf("Script with id %s not found", id),
		})
	}
    res := []entity.MstScript{
		{
			Id: data.Id,
			Title:    data.Title,
			Script:   data.Script,
			IsActive: data.IsActive,
		},
	}

	// Kembalikan response
	return ctx.Status(fiber.StatusOK).JSON(res)
}

