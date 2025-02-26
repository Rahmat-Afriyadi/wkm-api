package controller

import (
	"strconv"
	"wkm/entity"
	"wkm/request"
	"wkm/service"

	"github.com/gofiber/fiber/v2"
)


type CustomerMtrController interface {
	MasterData(ctx *fiber.Ctx) error
	MasterDataCount(ctx *fiber.Ctx) error
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

func (tr *customerMtrController) MasterData(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	sts := ctx.Query("sts")
	jns := ctx.Query("jns")
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	pageParams, _ := strconv.Atoi(ctx.Query("pageParams"))
	data := tr.customerMtrService.MasterData(search, sts, jns, details.Username, limit, pageParams)
	return ctx.Status(200).JSON(data)
}
func (tr *customerMtrController) MasterDataCount(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	sts := ctx.Query("sts")
	jns := ctx.Query("jns")
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	data := tr.customerMtrService.MasterDataCount(search, sts, jns, details.Username)
	return ctx.Status(200).JSON(data)
}
func (tr *customerMtrController) ListAmbilData(ctx *fiber.Ctx) error {
	data := tr.customerMtrService.ListAmbilData()
	return ctx.Status(200).JSON(fiber.Map{"status":"success","data": data})
}

func (tr *customerMtrController) AmbilData(ctx *fiber.Ctx) error {
	var request entity.Faktur3
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
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
	var request request.CustomerMtr
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request body", "details": err.Error()})
	}
	user := ctx.Locals("user")
	details,_:= user.(entity.User)

	request.KdUserTs = details.Username
	customer, err := tr.customerMtrService.UpdateOkeMembership(request)
	if err != nil {
		return  ctx.Status(400).JSON(fiber.Map{"error": "Invalid request body", "details": err.Error()})
 
	}
	return ctx.Status(200).JSON(fiber.Map{"message": "Berhasil ", "data":customer})
}