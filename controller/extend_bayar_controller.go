package controller

import (
	"fmt"
	"strconv"
	"sync"
	"time"
	"wkm/config"
	"wkm/entity"
	"wkm/request"
	"wkm/service"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gofiber/fiber/v2"
)

type ExtendBayarController interface {
	MasterData(ctx *fiber.Ctx) error
	MasterDataCount(ctx *fiber.Ctx) error
	DetailExtendBayar(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	UpdateFa(ctx *fiber.Ctx) error
	UpdateLf(ctx *fiber.Ctx) error
	UploadDokumen(ctx *fiber.Ctx) error
	UpdateApprovalLf(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type extendBayarController struct {
	extendBayarService service.ExtendBayarService
}

func NewExtendBayarController(aS service.ExtendBayarService) ExtendBayarController {
	return &extendBayarController{
		extendBayarService: aS,
	}
}

func (tm *extendBayarController) MasterData(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	tgl1 := ctx.Query("tgl1")
	tgl2 := ctx.Query("tgl2")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	pageParams, _ := strconv.Atoi(ctx.Query("pageParams"))
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	if details.Role.Name == "LEADER_FA" {
		return ctx.JSON(tm.extendBayarService.MasterDataLf(search, tgl1, tgl2, limit, pageParams))
	}
	return ctx.JSON(tm.extendBayarService.MasterData(search, tgl1, tgl2, limit, pageParams))
}

func (tm *extendBayarController) MasterDataCount(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	user := ctx.Locals("user")
	tgl1 := ctx.Query("tgl1")
	tgl2 := ctx.Query("tgl2")
	details, _ := user.(entity.User)
	if details.Role.Name == "LEADER_FA" {
		return ctx.JSON(tm.extendBayarService.MasterDataLfCount(search, tgl1, tgl2))
	}
	return ctx.JSON(tm.extendBayarService.MasterDataCount(search, tgl1, tgl2))
}

func (tm *extendBayarController) DetailExtendBayar(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	return ctx.JSON(tm.extendBayarService.Detail(id))
}

func (tm *extendBayarController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	err := tm.extendBayarService.Delete(id, details.Username)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil Delete"})
}

func (tm *extendBayarController) Create(ctx *fiber.Ctx) error {
	var body request.ExtendBayarRequest
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	user := ctx.Locals("user")
	details, _ := user.(entity.User)

	body.KdUserFa = details.Username
	data, err := tm.extendBayarService.Create(body)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error(), "status": "fail"})
	}
	return ctx.JSON(map[string]interface{}{"message": "Pengajuan Extend Bayar Berhasil", "status": "success", "id": data.Id})
}

func (tm *extendBayarController) UpdateFa(ctx *fiber.Ctx) error {
	var body request.ExtendBayarRequest
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	user := ctx.Locals("user")
	details, _ := user.(entity.User)

	body.KdUserFa = details.Username
	err = tm.extendBayarService.UpdateFa(body)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error(), "status": "fail"})
	}
	return ctx.JSON(map[string]interface{}{"message": "Berhasil update", "status": "success"})
}

func (tm *extendBayarController) UpdateLf(ctx *fiber.Ctx) error {
	var body request.ExtendBayarRequest
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	user := ctx.Locals("user")
	details, _ := user.(entity.User)

	body.KdUserLf = details.Username
	err = tm.extendBayarService.UpdateLf(body)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error(), "status": "fail"})
	}
	return ctx.JSON(map[string]interface{}{"message": "Berhasil update", "status": "success"})
}

func (tm *extendBayarController) UpdateApprovalLf(ctx *fiber.Ctx) error {
	var body request.ExtendBayarApprovalRequest
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser kesini yaa", err)
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	user := ctx.Locals("user")
	details, _ := user.(entity.User)

	if details.Role.Name != "LEADER_FA" {
		return ctx.Status(403).JSON(map[string]string{"message": "Kamu bukan Leader FA", "status": "fail"})
	}
	body.KdUserLf = details.Username
	err = tm.extendBayarService.UpdateApprovalLf(body)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error(), "status": "fail"})
	}
	return ctx.JSON(map[string]interface{}{"message": "Berhasil update", "status": "success"})
}

func (tm *extendBayarController) UploadDokumen(ctx *fiber.Ctx) error {
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
	now := time.Now()

	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	var pesanError string
	connGorm, _ := config.GetConnection()

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
				var datas []entity.ExtendBayar
				if len(rows) < 1 {
					return
				}
				var renewalKe string
				for _, v := range rows[1:] {
					connGorm.Model(&entity.Faktur3{}).Select("sts_cetak3").Where("no_msn", v[0]).Scan(&renewalKe)
					if len(v) < 3 {
						success = false
						continue
					}
					tglActualBayar, err := time.Parse("2006-01-02", v[1])
					if err != nil {
						pesanError = "Ada format tanggal yang tidak sesuai " + v[0]
						success = false
						return
					}
					datas = append(datas, entity.ExtendBayar{
						NoMsn:          v[0],
						TglActualBayar: tglActualBayar,
						Deskripsi:      v[2],
						KdUserFa:       details.Username,
						StsApproval:    "P",
						TglPengajuan:   now,
						TglUpdateFa:    now,
						RenewalKe:      renewalKe,
					})
				}
				err = tm.extendBayarService.CreateFromFile(datas)
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
