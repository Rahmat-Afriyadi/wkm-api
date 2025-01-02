package controller

import (
	"fmt"
	"sync"
	"time"
	"wkm/entity"
	"wkm/repository"
	"wkm/request"
	"wkm/service"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gofiber/fiber/v2"
)

type EditJenisBayarRequest struct {
	PaymentType string `form:"payment_type"`
}

type Tr3Controller interface {
	ExportDataWaBlast(ctx *fiber.Ctx) error
	SearchNoMsnByWa(ctx *fiber.Ctx) error
	EditJenisBayar(ctx *fiber.Ctx) error
	UpdateInputBayar(ctx *fiber.Ctx) error
	WillBayar(ctx *fiber.Ctx) error
	DataRenewal(ctx *fiber.Ctx) error
	ExportDataRenewal(ctx *fiber.Ctx) error
	ExportDataPlatinumPlus(ctx *fiber.Ctx) error
	ExportPembayaranRenewal(ctx *fiber.Ctx) error
}

type tr3Controller struct {
	tr3Service service.Tr3Service
}

func NewTr3Controller(aS service.Tr3Service) Tr3Controller {
	return &tr3Controller{
		tr3Service: aS,
	}
}

func (tr *tr3Controller) ExportDataWaBlast(ctx *fiber.Ctx) error {
	var request request.DataWaBlastRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(400).JSON(err)
	}
	_ = tr.tr3Service.DataWABlast(request)
	return ctx.Download("./file1.xlsx")

}

func (tr *tr3Controller) ExportDataRenewal(ctx *fiber.Ctx) error {
	var request request.DataRenewalRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request body", "details": err.Error()})
	}

	// Memanggil service untuk mengekspor data renewal
	if _, err := tr.tr3Service.ExportDataRenewal(request); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Failed to export data", "details": err.Error()})
	}

	return ctx.Download("./Data_Renewal.xlsx")
	// return ctx.Download("./file1.xlsx")

}
func (tr *tr3Controller) ExportDataPlatinumPlus(ctx *fiber.Ctx) error {
	var request request.DataRenewalRequest
	if err := ctx.BodyParser(&request); err != nil {
		fmt.Println("details service request", err.Error())
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request body", "details": err.Error()})
	}
	// Memanggil service untuk mengekspor data platinum plus
	if _, err := tr.tr3Service.ExportDataPlatinumPlus(request); err != nil {
		fmt.Println("details service ", err.Error())
		return ctx.Status(500).JSON(fiber.Map{"error": "Failed to export platinum plus data", "details": err.Error()})
	}

	return ctx.Download("./Data_Platinum_Plus.xlsx")
	// return ctx.Download("./file1.xlsx")

}

func (tr *tr3Controller) DataRenewal(ctx *fiber.Ctx) error {
	var request request.DataRenewalRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request body", "details": err.Error()})
	}

	// Memanggil service untuk melakukan pembaruan data
	data, err := tr.tr3Service.DataRenewal(request)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Data renewal failed", "details": err.Error()})
	}

	// Menampilkan data dan error di console (untuk debugging)
	fmt.Println(data, err)

	// Mengembalikan data sebagai respons JSON
	return ctx.JSON(fiber.Map{"message": "Data renewal successful", "data": data})
}

func (tr *tr3Controller) SearchNoMsnByWa(ctx *fiber.Ctx) error {
	var request request.SearchNoMsnByWaRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(400).JSON(err)
	}
	data := tr.tr3Service.SearchNoMsnByWa(request)
	return ctx.JSON(data)

}

func (tr *tr3Controller) EditJenisBayar(ctx *fiber.Ctx) error {
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
				rows := xlsx.GetRows("Lap. Pembayaran Renewal All")
				var datas []repository.ParamsUpdateJenisBayar
				if len(rows) < 1 {
					return
				}
				for _, v := range rows[2:] {
					if len(v) < 9 {
						success = false
						continue
					}
					if len(v[8]) < 10 {
						continue
					}
					datas = append(datas, repository.ParamsUpdateJenisBayar{
						NoTandaTerima: v[8],
						NamaCustomer:  v[1],
					})

				}
				tr.tr3Service.UpdateJenisBayar(datas, data.PaymentType, details.Username)
			}(&wg)
		}
	}
	wg.Wait()
	if success {
		return ctx.Status(200).JSON(map[string]string{"message": "Data berhasil di update"})
	}
	return ctx.Status(400).JSON(map[string]string{"message": "Periksa kembali format file anda"})

}

func (tr *tr3Controller) UpdateInputBayar(ctx *fiber.Ctx) error {
	var body request.InputBayarRequest
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(400).JSON(err)
	}
	minLimit, err := time.Parse("2006-01-02", "2024-01-01")
	if err != nil {
		return ctx.Status(400).JSON(map[string]string{"message": err.Error(), "status": "fail"})
	}
	if body.TglBayar.Before(minLimit) {
		return ctx.Status(400).JSON(map[string]string{"message": "Mohon mengisi tanggal bayar", "status": "fail"})
	}
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	body.KdUserFa = details.Username
	data, err := tr.tr3Service.UpdateInputBayar(body)
	if err != nil {
		return ctx.Status(400).JSON(map[string]string{"message": err.Error(), "status": "fail"})
	}
	return ctx.JSON(map[string]interface{}{"data": data, "status": "success"})
}

func (tr *tr3Controller) WillBayar(ctx *fiber.Ctx) error {
	var body request.SearchWBRequest
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(400).JSON(err)
	}
	faktur3, err := tr.tr3Service.WillBayar(body)
	if err != nil {
		return ctx.Status(400).JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(faktur3)
}

func (tr *tr3Controller) ExportPembayaranRenewal(ctx *fiber.Ctx) error {
	var request request.RangeTanggalRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(400).JSON(err)
	}
	tr.tr3Service.ExportPembayaranRenewal(request)
	return ctx.Download("./pembayaran-renewal.xlsx")

}
