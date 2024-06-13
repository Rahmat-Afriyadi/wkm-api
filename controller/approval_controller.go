package controller

import (
	"fmt"
	"wkm/entity"
	"wkm/service"

	"github.com/gofiber/fiber/v2"
)

type ApprovalController interface {
	Update(ctx *fiber.Ctx) error
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
