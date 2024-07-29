package controller

import (
	"fmt"
	"path/filepath"
	"strconv"
	"time"
	"wkm/entity"
	"wkm/request"
	"wkm/service"

	"github.com/gofiber/fiber/v2"
)

type TransaksiController interface {
	MasterData(ctx *fiber.Ctx) error
	MasterDataCount(ctx *fiber.Ctx) error
	DetailMstMtr(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	UploadDokumen(ctx *fiber.Ctx) error
}

type transaksiController struct {
	transaksiService service.TransaksiService
}

func NewTransaksiController(aS service.TransaksiService) TransaksiController {
	return &transaksiController{
		transaksiService: aS,
	}
}

func (tr *transaksiController) MasterData(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	pageParams, _ := strconv.Atoi(ctx.Query("pageParams"))
	return ctx.JSON(tr.transaksiService.MasterData(search, limit, pageParams))
}

func (tr *transaksiController) MasterDataCount(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	return ctx.JSON(tr.transaksiService.MasterDataCount(search))
}

func (tr *transaksiController) DetailMstMtr(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	return ctx.JSON(tr.transaksiService.Detail(id))
}

func (tr *transaksiController) Update(ctx *fiber.Ctx) error {
	var body request.TransaksiRequest
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	err = tr.transaksiService.Update(body)
	if err != nil {
		return ctx.JSON(map[string]interface{}{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil Update"})
}

func (tr *transaksiController) UploadDokumen(ctx *fiber.Ctx) error {
	var body entity.Transaksi
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	ktp, _ := ctx.FormFile("ktp")
	if ktp != nil {
		fmt.Println("ini file name yaa ", ktp.Filename)
		fileName := fmt.Sprintf("%s_%s", time.Now().Format("20060102150405"), ktp.Filename)
		filepath := filepath.Join("./uploads", fileName)
		if err := ctx.SaveFile(ktp, filepath); err != nil {
			fmt.Println("ini error file ", err.Error())
		}
		body.Ktp = "/uploads/" + fileName
	}

	stnk, _ := ctx.FormFile("stnk")
	if stnk != nil {
		fileName := fmt.Sprintf("%s_%s", time.Now().Format("20060102150405"), stnk.Filename)
		filepath := filepath.Join("./uploads", fileName)
		if err := ctx.SaveFile(stnk, filepath); err != nil {
			fmt.Println("ini error file ", err.Error())
		}
		body.Stnk = "/uploads/" + fileName
	}
	err = tr.transaksiService.UploadDokumen(body)
	if err != nil {
		return ctx.JSON(map[string]interface{}{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil Update"})
}

func (tr *transaksiController) Create(ctx *fiber.Ctx) error {
	var body request.TransaksiRequest
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}

	data, err := tr.transaksiService.Create(body)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil create", "id_transaksi": data.ID})
}
