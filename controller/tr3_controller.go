package controller

import (
	"fmt"
	"sync"
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
		fmt.Println("ini data ", data, err.Error())
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
