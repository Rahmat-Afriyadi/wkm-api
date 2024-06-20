package controller

import (
	"fmt"
	"wkm/entity"
	"wkm/service"

	"github.com/gofiber/fiber/v2"
)

type ApprovalController interface {
	Update(ctx *fiber.Ctx) error
	MokitaToken(ctx *fiber.Ctx) error
	MokitaUpdateToken(ctx *fiber.Ctx) error
}

type approvalController struct {
	approvalService service.ApprovalService
}

func NewApprovalController(aS service.ApprovalService) ApprovalController {
	return &approvalController{
		approvalService: aS,
	}
}

func (tr *approvalController) Update(ctx *fiber.Ctx) error {
	var body entity.DetailApproval
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	tr.approvalService.Update(body)
	return ctx.JSON(map[string]interface{}{"message": "Data berhasil dihapus"})

}

func (tr *approvalController) MokitaToken(ctx *fiber.Ctx) error {
	tr.approvalService.MokitaToken()
	return ctx.JSON(tr.approvalService.MokitaToken())

}
func (tr *approvalController) MokitaUpdateToken(ctx *fiber.Ctx) error {
	var body map[string]interface{}
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}

	tr.approvalService.MokitaUpdateToken(body["token"].(string))
	return ctx.JSON(map[string]interface{}{"message": "Berhasil update token"})

}
