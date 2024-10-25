package controller

import (
	"fmt"
	"path/filepath"
	"strconv"
	"time"
	"wkm/entity"
	"wkm/service"

	"github.com/gofiber/fiber/v2"
)

type ProdukController interface {
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

type produkController struct {
	produkService service.ProdukService
}

func NewProdukController(aS service.ProdukService) ProdukController {
	return &produkController{
		produkService: aS,
	}
}

func (tr *produkController) MasterData(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	jenis_asuransi, _ := strconv.Atoi(ctx.Query("jenis_asuransi"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	pageParams, _ := strconv.Atoi(ctx.Query("pageParams"))
	return ctx.JSON(tr.produkService.MasterData(search, jenis_asuransi, limit, pageParams))
}

func (tr *produkController) MasterDataCount(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	jenis_asuransi, _ := strconv.Atoi(ctx.Query("jenis_asuransi"))
	return ctx.JSON(tr.produkService.MasterDataCount(search, jenis_asuransi))
}

func (tr *produkController) DetailMstMtr(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	return ctx.JSON(tr.produkService.Detail(id))
}

func (tr *produkController) Update(ctx *fiber.Ctx) error {
	var body entity.MasterProduk
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	err = tr.produkService.Update(body)
	if err != nil {
		return ctx.JSON(map[string]interface{}{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil Update"})
}

func (tr *produkController) UploadLogo(ctx *fiber.Ctx) error {
	var body entity.MasterProduk
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}

	file, _ := ctx.FormFile("files")
	if file != nil {
		fileName := fmt.Sprintf("%s_%s", time.Now().Format("20060102150405"), file.Filename)
		filepath := filepath.Join("./uploads", fileName)
		if err := ctx.SaveFile(file, filepath); err != nil {
			fmt.Println("ini error file ", err.Error())
		}
		body.Logo = "/uploads/" + fileName
	}
	err = tr.produkService.UploadLogo(body)
	if err != nil {
		return ctx.JSON(map[string]interface{}{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil Update"})
}

func (tr *produkController) Create(ctx *fiber.Ctx) error {
	var body entity.MasterProduk
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}

	data, err := tr.produkService.Create(body)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil create", "id_produk": data.KdProduk})
}

func (tr *produkController) DeleteManfaat(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := tr.produkService.DeleteManfaat(id)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil create"})
}

func (tr *produkController) DeleteSyarat(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := tr.produkService.DeleteSyarat(id)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil create"})
}
func (tr *produkController) DeletePaket(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := tr.produkService.DeletePaket(id)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil create"})
}
