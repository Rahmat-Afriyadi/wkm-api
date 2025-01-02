package controller

import (
	"fmt"
	"path/filepath"
	"strconv"
	"sync"
	"time"
	"wkm/entity"
	"wkm/request"
	"wkm/service"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gofiber/fiber/v2"
)

type TransaksiController interface {
	MasterData(ctx *fiber.Ctx) error
	MasterDataCount(ctx *fiber.Ctx) error
	DetailMstMtr(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	UploadDokumen(ctx *fiber.Ctx) error
	ImportExcell(ctx *fiber.Ctx) error
}

type transaksiController struct {
	transaksiService service.TransaksiService
}

func NewTransaksiController(aS service.TransaksiService) TransaksiController {
	return &transaksiController{
		transaksiService: aS,
	}
}

func ExcelDateToTime(excelDate float64) time.Time {
	// Excel's base date is January 1, 1900
	baseDate := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)

	// Excel incorrectly considers 1900 a leap year, so we need to subtract one day for dates after February 28, 1900
	if excelDate >= 59 { // 59 is 1900-02-28
		excelDate-- // Correct for the leap year bug
	}

	// Convert the Excel date to Go time.Time
	days := int(excelDate)
	return baseDate.AddDate(0, 0, days-1)
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

func (tr *transaksiController) ImportExcell(ctx *fiber.Ctx) error {
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
				// var datas []request.TransaksiRequest
				// if len(rows) < 1 {
				// 	return
				// }
				fmt.Println("ini adalah details ", details.Name)
				for _, v := range rows[4:] {
					num64, _ := strconv.ParseFloat(v[16], 64)
					fmt.Println("ini rows yaa ", ExcelDateToTime(num64))

				}
				// for _, v := range rows[4:] {
				// 	// if len(v) < 9 {
				// 	// 	success = false
				// 	// 	continue
				// 	// }
				// 	datas = append(datas, request.TransaksiRequest{
				// 		// Nik:
				// 		NmKonsumen: v[1],
				// 		// Email:
				// 		// NoHp
				// 		// Alamat
				// 		// TglLahir
				// 		// Admin
				// 		// Amount
				// 		// IdProduk
				// 		// KdMdl
				// 		// NmMtr
				// 		// Warna
				// 		// Tahun
				// 		// Otr
				// 		// Kodepos
				// 		// Kelurahan
				// 		// Kecamatan
				// 		// Kota
				// 		// NoMsn
				// 		// NoRgk
				// 		// NoPlat
				// 	})
				// }
				// fmt.Println("ini dataa  yaa ", datas)
			}(&wg)
		}
	}
	wg.Wait()
	if success {
		return ctx.Status(200).JSON(map[string]string{"message": "Data berhasil di update"})
	}
	return ctx.Status(400).JSON(map[string]string{"message": "Periksa kembali format file anda"})
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
