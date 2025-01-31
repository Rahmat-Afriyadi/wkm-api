package controller

import (
	"fmt"
	"wkm/entity"
	"wkm/service"

	"github.com/gofiber/fiber/v2"
)


type CustomerMtrController interface {
	ListAmbilData(ctx *fiber.Ctx) error
	AmbilData(ctx *fiber.Ctx) error
	Show(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	UpdateOkeMembership(ctx *fiber.Ctx) error
}

type customerMtrController struct {
	customerMtrService service.CustomerMtrService
}

func NewCustomerMtrController(aS service.CustomerMtrService) CustomerMtrController {
	return &customerMtrController{
		customerMtrService: aS,
	}
}

func (tr *customerMtrController) ListAmbilData(ctx *fiber.Ctx) error {
	data := tr.customerMtrService.ListAmbilData()
	return ctx.Status(200).JSON(fiber.Map{"status":"success","data": data})
}

func (tr *customerMtrController) AmbilData(ctx *fiber.Ctx) error {
	var request entity.Faktur3
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	fmt.Println("ini user ", details.RoleId)
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request body", "details": err.Error()})
	}
	err := tr.customerMtrService.AmbilData(request.NoMsn, details.Username)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"status":"fail", "message": err.Error()})
	}
	return ctx.Status(200).JSON(fiber.Map{"message": "Berhasil Ambil Data"})
}

func (tr *customerMtrController) Show(ctx *fiber.Ctx) error {
	noMsn := ctx.Params("no_msn")
	data := tr.customerMtrService.Show(noMsn)
	fmt.Println("sho yaa ", data.NoMsn)
	return ctx.Status(200).JSON(fiber.Map{"message": "Berhasil ", "data":data})
}

func (tr *customerMtrController) Update(ctx *fiber.Ctx) error {
	var request entity.CustomerMtr
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request body", "details": err.Error()})
	}
	
	return nil
}

func (tr *customerMtrController) UpdateOkeMembership(ctx *fiber.Ctx) error {
	var request entity.CustomerMtr
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request body", "details": err.Error()})
	}
	customer, err := tr.customerMtrService.UpdateOkeMembership(request)
	if err != nil {
		return  ctx.Status(400).JSON(fiber.Map{"error": "Invalid request body", "details": err.Error()})
 
	}
	return ctx.Status(200).JSON(fiber.Map{"message": "Berhasil ", "data":customer})
}