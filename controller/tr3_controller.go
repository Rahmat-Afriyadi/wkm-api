package controller

import (
	"fmt"
	"sync"
	"wkm/repository"
	"wkm/request"
	"wkm/service"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gofiber/fiber/v2"
)

type Tr3Controller interface {
	ExportDataWaBlast(ctx *fiber.Ctx) error
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

func (tr *tr3Controller) EditJenisBayar(ctx *fiber.Ctx) error {
	var wg sync.WaitGroup
	form, err := ctx.MultipartForm()
	if err != nil { /* handle error */
		fmt.Println("error form file ", err)
	}
	for _, fileHeaders := range form.File {
		for _, fileHeader := range fileHeaders {
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				file, err := fileHeader.Open()
				if err != nil {
					fmt.Println("ini errornya ", err)
					panic(err)
				}
				xlsx, err := excelize.OpenReader(file)
				if err != nil {
					fmt.Println("ini errornya ", err)
					panic(err)
				}
				rows := xlsx.GetRows(xlsx.GetSheetName(1))
				var datas []repository.ParamsUpdateJenisBayar
				for _, v := range rows[2:] {
					datas = append(datas, repository.ParamsUpdateJenisBayar{
						NoTandaTerima: v[8],
						NamaCustomer:  v[1],
					})
				}
				tr.tr3Service.UpdateJenisBayar(datas)
			}(&wg)
		}
	}
	wg.Wait()
	fmt.Println("Edit sukses")
	return ctx.JSON(map[string]string{"message": "Data berhasil di update"})

}
