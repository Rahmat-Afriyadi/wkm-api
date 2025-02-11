package controller

import (
	"fmt"
	"strconv"
	"sync"
	"time"
	"wkm/entity"
	"wkm/request"
	"wkm/service"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gofiber/fiber/v2"
)

type TglMerahController interface {
	MasterData(ctx *fiber.Ctx) error
	MasterDataCount(ctx *fiber.Ctx) error
	DetailTglMerah(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	UploadDokumen(ctx *fiber.Ctx) error
	MinTglBayar(ctx *fiber.Ctx) error
}

type tglMerahController struct {
	tglMerahService service.TglMerahService
}

func NewTglMerahController(aS service.TglMerahService) TglMerahController {
	return &tglMerahController{
		tglMerahService: aS,
	}
}

func (tm *tglMerahController) MasterData(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	tgl1 := ctx.Query("tgl1")
	tgl2 := ctx.Query("tgl2")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	pageParams, _ := strconv.Atoi(ctx.Query("pageParams"))
	return ctx.JSON(tm.tglMerahService.MasterData(search, tgl1, tgl2, limit, pageParams))
}

func (tm *tglMerahController) MasterDataCount(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	tgl1 := ctx.Query("tgl1")
	tgl2 := ctx.Query("tgl2")
	return ctx.JSON(tm.tglMerahService.MasterDataCount(search, tgl1, tgl2))
}

func (tm *tglMerahController) DetailTglMerah(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idVal, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return ctx.JSON(map[string]interface{}{"message": err.Error()})
	}
	return ctx.JSON(tm.tglMerahService.Detail(idVal))
}

func (tm *tglMerahController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idUint, _ := strconv.ParseUint(id, 10, 64)
	err := tm.tglMerahService.Delete(idUint)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil Delete"})
}

func (tm *tglMerahController) Update(ctx *fiber.Ctx) error {
	var body request.TglMerahRequest
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	err = tm.tglMerahService.Update(body)
	if err != nil {
		return ctx.JSON(map[string]interface{}{"message": err.Error(), "status": "fail"})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil Update", "status": "success"})
}

func (tm *tglMerahController) UploadDokumen(ctx *fiber.Ctx) error {
	success := true
	var wg sync.WaitGroup
	form, err := ctx.MultipartForm()
	if err != nil { /* handle error */
		fmt.Println("error form file ", err)
	}
	var data EditJenisBayarRequest
	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	var pesanError string

	for _, fileHeaders := range form.File {
		for _, fileHeader := range fileHeaders {
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				file, err := fileHeader.Open()
				if err != nil {
					fmt.Println("ini errornya ", err)
				}
				xlsx, err := excelize.OpenReader(file)
				if err != nil {
					fmt.Println("ini errornya ", err)
				}
				// rows := xlsx.GetRows(xlsx.GetSheetName(1))
				rows := xlsx.GetRows("Sheet1")
				var datas []entity.TglMerah
				if len(rows) < 1 {
					return
				}
				for _, v := range rows[1:] {
					if len(v) < 3 {
						success = false
						continue
					}
					date1, err := time.Parse("2006-01-02", v[0])
					if err != nil {
						pesanError = "Ada format tanggal yang tidak sesuai " + v[0]
						success = false
						return
					}
					date2, err := time.Parse("2006-01-02", v[1])
					if err != nil {
						pesanError = "Ada format tanggal yang tidak sesuai " + v[1]
						success = false
						return
					}
					datas = append(datas, entity.TglMerah{
						TglAwal:   date1,
						TglAkhir:  date2,
						Deskripsi: v[2],
						KdUser:    details.Username,
					})
				}
				err = tm.tglMerahService.CreateFromFile(datas)
				if err != nil {
					pesanError = err.Error()
					success = false
				}
			}(&wg)
		}
	}
	wg.Wait()
	if success {
		return ctx.Status(200).JSON(map[string]string{"message": "Data berhasil di update"})
	} else {
		return ctx.Status(400).JSON(map[string]string{"message": pesanError})
	}
}

func (tm *tglMerahController) MinTglBayar(ctx *fiber.Ctx) error {
	return ctx.JSON(map[string]interface{}{"min": tm.tglMerahService.MinTglBayar()})
}
func (tm *tglMerahController) Create(ctx *fiber.Ctx) error {
	var body request.TglMerahRequest
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	user := ctx.Locals("user")
	details, _ := user.(entity.User)

	body.KdUser = details.Username
	data, err := tm.tglMerahService.Create(body)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error(), "status": "fail"})
	}
	return ctx.JSON(map[string]interface{}{"message": "Berhasil create", "id_tglMerah": data.ID, "status": "success"})
}
