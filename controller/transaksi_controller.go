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
	UploadLogo(ctx *fiber.Ctx) error
	DeleteManfaat(ctx *fiber.Ctx) error
	DeleteSyarat(ctx *fiber.Ctx) error
	DeletePaket(ctx *fiber.Ctx) error
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
	jenis_asuransi, _ := strconv.Atoi(ctx.Query("jenis_asuransi"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	pageParams, _ := strconv.Atoi(ctx.Query("pageParams"))
	return ctx.JSON(tr.transaksiService.MasterData(search, jenis_asuransi, limit, pageParams))
}

func (tr *transaksiController) MasterDataCount(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	jenis_asuransi, _ := strconv.Atoi(ctx.Query("jenis_asuransi"))
	return ctx.JSON(tr.transaksiService.MasterDataCount(search, jenis_asuransi))
}

func (tr *transaksiController) DetailMstMtr(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	return ctx.JSON(tr.transaksiService.Detail(id))
}

func (tr *transaksiController) Update(ctx *fiber.Ctx) error {
	var body entity.MasterProduk
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

func (tr *transaksiController) UploadLogo(ctx *fiber.Ctx) error {
	var body entity.MasterProduk
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	fmt.Println("ini body", body)

	file, _ := ctx.FormFile("files")
	if file != nil {
		fileName := fmt.Sprintf("%s_%s", time.Now().Format("20060102150405"), file.Filename)
		filepath := filepath.Join("./uploads", fileName)
		if err := ctx.SaveFile(file, filepath); err != nil {
			fmt.Println("ini error file ", err.Error())
		}
		body.Logo = "/uploads/" + fileName
	}
	fmt.Println("kesini gk yaa dia ", body)
	err = tr.transaksiService.UploadLogo(body)
	if err != nil {
		return ctx.JSON(map[string]interface{}{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil Update"})
}

func (tr *transaksiController) Create(ctx *fiber.Ctx) error {
	var body request.TransaksiCreateRequest
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

func (tr *transaksiController) DeleteManfaat(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := tr.transaksiService.DeleteManfaat(id)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil create"})
}

func (tr *transaksiController) DeleteSyarat(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := tr.transaksiService.DeleteSyarat(id)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil create"})
}
func (tr *transaksiController) DeletePaket(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := tr.transaksiService.DeletePaket(id)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil create"})
}
