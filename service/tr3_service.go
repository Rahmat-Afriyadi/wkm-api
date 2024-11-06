package service

import (
	"fmt"
	"strings"
	"wkm/entity"
	"wkm/repository"
	"wkm/request"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type Tr3Service interface {
	DataWABlast(t request.DataWaBlastRequest) []entity.DataWaBlast
	SearchNoMsnByWa(t request.SearchNoMsnByWaRequest) []entity.SearchNoMsnByWa
	UpdateJenisBayar(data []repository.ParamsUpdateJenisBayar, payment_type string, username string)
	WillBayar(data request.SearchWBRequest) (entity.Faktur3, error)
	UpdateInputBayar(data request.InputBayarRequest) (entity.Faktur3, error)
}

type tr3Service struct {
	trR repository.Tr3Repository
}

func NewTr3Service(tR repository.Tr3Repository) Tr3Service {
	return &tr3Service{
		trR: tR,
	}
}

func (s *tr3Service) SearchNoMsnByWa(t request.SearchNoMsnByWaRequest) []entity.SearchNoMsnByWa {
	return s.trR.SearchNoMsnByWa(t)
}
func (s *tr3Service) DataWABlast(t request.DataWaBlastRequest) []entity.DataWaBlast {
	s.trR.UpdateTglAkhirTenor()
	datas := s.trR.DataWABlast(t)

	xlsx := excelize.NewFile()

	sheet1Name := "Sheet One"
	xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)
	xlsx.SetCellValue(sheet1Name, "A1", "Nomor Mesin")
	xlsx.SetCellValue(sheet1Name, "B1", "Nama Customer")
	xlsx.SetCellValue(sheet1Name, "C1", "KD User")
	xlsx.SetCellValue(sheet1Name, "D1", "Nomor yg Dihubungi Renewal")
	xlsx.SetCellValue(sheet1Name, "E1", "Nomor Wa")
	xlsx.SetCellValue(sheet1Name, "F1", "Tanggal Akhir Tenor")

	for i, each := range datas {
		NoYgDiHubRenewal := ""
		if each.NoYgDiHubRenewal != nil {
			NoYgDiHubRenewal = *each.NoYgDiHubRenewal
		}
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", i+2), each.NoMsn)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", i+2), each.NmCustomer)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", i+2), *each.KdUser)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("D%d", i+2), NoYgDiHubRenewal)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("E%d", i+2), each.NoWa)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("F%d", i+2), each.TglAkhirTenor[:10])
	}

	err := xlsx.SaveAs("./file1.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	return datas
}

func (s *tr3Service) UpdateJenisBayar(data []repository.ParamsUpdateJenisBayar, payment_type string, username string) {
	s.trR.UpdateJenisBayar(data, payment_type, username)
}

func (s *tr3Service) UpdateInputBayar(data request.InputBayarRequest) (entity.Faktur3, error) {
	return s.trR.UpdateInputBayar(data)
}

func (s *tr3Service) WillBayar(data request.SearchWBRequest) (entity.Faktur3, error) {
	data.Kode = strings.ReplaceAll(data.Kode, " ", "")
	if len(data.Kode) > 16 {
		data.Kode = data.Kode[:16]
	}

	return s.trR.WillBayar(data)
}
