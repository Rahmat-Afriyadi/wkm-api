package controller

import (
	"wkm/entity"
	"wkm/service"

	"github.com/gofiber/fiber/v2"
)

type AsuransiPAController interface {
	CreateAsuransiPA(ctx *fiber.Ctx) error
	UpdateAsuransiPA(ctx *fiber.Ctx) error
}

type asuransiPAController struct {
	asuransiPAService service.AsuransiPAService
}

func NewAsuransiPAController(aSP service.AsuransiPAService) AsuransiPAController {
	return &asuransiPAController{
		asuransiPAService: aSP,
	}
}

// Change the receiver to *asuransiPAController
func (aSP *asuransiPAController) CreateAsuransiPA(ctx *fiber.Ctx) error {
	// Decode request JSON into struct
	var asuransiPARequest entity.AsuransiPA
	if err := ctx.BodyParser(&asuransiPARequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}

	// Call service to create AsuransiPA
	err := aSP.asuransiPAService.CreateAsuransiPA(asuransiPARequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to create Asuransi PA",
			"details": err.Error(),
		})
	}

	// Return success response
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Asuransi PA created successfully",
	})
}

func (aSP *asuransiPAController) UpdateAsuransiPA(ctx *fiber.Ctx) error {
	// Decode request JSON into struct
	var asuransiPARequest entity.AsuransiPA
	if err := ctx.BodyParser(&asuransiPARequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}

	id := ctx.Params("id")
    if id == "" {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error":   "Invalid request",
            "details": "id asuransi pa tidak ditemukan",
        })
    }
	// Call service to create AsuransiPA
	err := aSP.asuransiPAService.UpdateAsuransiPA(id, asuransiPARequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to create Asuransi PA",
			"details": err.Error(),
		})
	}

	// Return success response
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Asuransi PA Updated successfully",
	})
}