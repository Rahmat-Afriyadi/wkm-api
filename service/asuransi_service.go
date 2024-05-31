package service

import (
	"fmt"
	"wkm/entity"
	"wkm/repository"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type AsuransiService interface {
	MasterData(search string, dataSource string, sts string, usename string, limit int, pageParams int) []entity.MasterAsuransi
	MasterDataCount(search string, dataSource string, sts string, usename string) int64
	FindAsuransiByNoMsn(no_msn string) entity.MasterAsuransi
	UpdateAsuransi(data entity.MasterAsuransi) entity.MasterAsuransi
	UpdateAsuransiBerminat(no_msn string)
	UpdateAsuransiBatalBayar(no_msn string)
	UpdateAmbilAsuransi(no_msn string, kd_user string)
	MasterDataRekapTele() []entity.MasterRekapTele
	RekapByStatus(u string, tgl string) entity.MasterStatusAsuransi
	ExportReport(u string, tgl string)
	MasterAlasanPending() []entity.MasterAlasanPending
	MasterAlasanTdkBerminat() []entity.MasterAlasanTdkBerminat
}

type asuransiService struct {
	trR repository.AsuransiRepository
	uR  repository.UserRepository
}

func NewAsuransiService(tR repository.AsuransiRepository, ur repository.UserRepository) AsuransiService {
	return &asuransiService{
		trR: tR,
		uR:  ur,
	}
}

func (s *asuransiService) MasterData(search string, dataSource string, sts string, usename string, limit int, pageParams int) []entity.MasterAsuransi {
	return s.trR.MasterData(search, dataSource, sts, usename, limit, pageParams)
}

func (s *asuransiService) MasterDataCount(search string, dataSource string, sts string, usename string) int64 {
	return s.trR.MasterDataCount(search, dataSource, sts, usename)
}

func (s *asuransiService) MasterDataRekapTele() []entity.MasterRekapTele {
	return s.trR.MasterDataRekapTele()
}

func (s *asuransiService) FindAsuransiByNoMsn(no_msn string) entity.MasterAsuransi {
	return s.trR.FindAsuransiByNoMsn(no_msn)
}

func (s *asuransiService) UpdateAsuransi(data entity.MasterAsuransi) entity.MasterAsuransi {
	return s.trR.UpdateAsuransi(data)
}

func (s *asuransiService) UpdateAsuransiBerminat(no_msn string) {
	s.trR.UpdateAsuransiBerminat(no_msn)
}

func (s *asuransiService) UpdateAsuransiBatalBayar(no_msn string) {
	s.trR.UpdateAsuransiBatalBayar(no_msn)
}

func (s *asuransiService) UpdateAmbilAsuransi(no_msn string, kd_user string) {
	s.trR.UpdateAmbilAsuransi(no_msn, kd_user)
}

func (s *asuransiService) RekapByStatus(u string, tgl string) entity.MasterStatusAsuransi {
	return s.trR.RekapByStatus(u, tgl)
}

func (s *asuransiService) MasterAlasanPending() []entity.MasterAlasanPending {
	return s.trR.MasterAlasanPending()
}

func (s *asuransiService) MasterAlasanTdkBerminat() []entity.MasterAlasanTdkBerminat {
	return s.trR.MasterAlasanTdkBerminat()
}

func (s *asuransiService) ExportReport(tgl1 string, tgl2 string) {

	rekapSourceInfo := s.trR.RekapByStatusJenisSource("2024-05-01", "2024-05-30")
	rekapKdUser := s.trR.RekapByStatusKdUser("2024-05-01", "2024-05-30")

	xlsx := excelize.NewFile()

	rekapSheet := "Rekap"
	pendingSheet := "Pending"
	tdkBerminatSheet := "Tidak Berminat"
	xlsx.SetSheetName(xlsx.GetSheetName(1), rekapSheet)
	xlsx.NewSheet(pendingSheet)
	xlsx.NewSheet(tdkBerminatSheet)
	xlsx.SetCellValue(rekapSheet, "A2", "Source Info")
	xlsx.SetColWidth(rekapSheet, "A", "G", 14)
	xlsx.SetCellValue(rekapSheet, "B2", "Pending")
	xlsx.SetCellValue(rekapSheet, "C2", "Tidak Berminat")
	xlsx.SetCellValue(rekapSheet, "D2", "Berminat")
	xlsx.SetCellValue(rekapSheet, "E2", "Total")
	xlsx.SetCellValue(rekapSheet, "A6", "Id User")
	xlsx.SetCellValue(rekapSheet, "B6", "Nama")
	xlsx.SetCellValue(rekapSheet, "C6", "Pending")
	xlsx.SetCellValue(rekapSheet, "D6", "Tidak Berminat")
	xlsx.SetCellValue(rekapSheet, "E6", "Berminat")
	xlsx.SetCellValue(rekapSheet, "F6", "Total")
	xlsx.SetCellValue(rekapSheet, "G6", "Source Info")

	headerStyle, err := xlsx.NewStyle(`{
		"font": {
			"bold": true
		},
		"alignment": {
            "horizontal": "center",
            "vertical": "center"
        },
		"fill": {
			"type": "pattern",
			"color": [
				"#FFFF00"
			],
			"pattern": 1
		},
		"border": [
			{
				"type": "left",
				"style": 1,
				"color": "#000000"
			},
			{
				"type": "top",
				"style": 1,
				"color": "#000000"
			},
			{
				"type": "right",
				"style": 1,
				"color": "#000000"
			},
			{
				"type": "bottom",
				"style": 1,
				"color": "#000000"
			}
		]
	}`)
	if err != nil {
		fmt.Println("ini error style ", err)
	}
	styleBorder, err := xlsx.NewStyle(`{
		"border": [
        {
            "type": "left",
            "style": 1,
            "color": "#000000"
        },
        {
            "type": "top",
            "style": 1,
            "color": "#000000"
        },
        {
            "type": "right",
            "style": 1,
            "color": "#000000"
        },
        {
            "type": "bottom",
            "style": 1,
            "color": "#000000"
		}
    ]
	}`)
	if err != nil {
		fmt.Println("ini error style ", err)
	}
	xlsx.SetCellStyle(rekapSheet, "A1", "G1", headerStyle)
	xlsx.SetCellStyle(rekapSheet, "A2", "E2", headerStyle)
	xlsx.SetCellStyle(rekapSheet, "A3", "A4", headerStyle)
	xlsx.SetCellStyle(rekapSheet, "A6", "G6", headerStyle)
	xlsx.SetCellStyle(rekapSheet, "B3", "E4", styleBorder)
	xlsx.MergeCell(rekapSheet, "A1", "G1")

	jenis_source := ""
	for i, each := range rekapSourceInfo {
		if each["jenis_source"] == "W" {
			jenis_source = "Wanda"
		}
		if each["jenis_source"] == "E" {
			jenis_source = "Excel"
		}
		xlsx.SetCellValue(rekapSheet, fmt.Sprintf("A%d", i+3), jenis_source)
		xlsx.SetCellValue(rekapSheet, fmt.Sprintf("B%d", i+3), each["p"])
		xlsx.SetCellValue(rekapSheet, fmt.Sprintf("C%d", i+3), each["t"])
		xlsx.SetCellValue(rekapSheet, fmt.Sprintf("D%d", i+3), each["o"])
		xlsx.SetCellValue(rekapSheet, fmt.Sprintf("E%d", i+3), each["total"])
	}

	xlsx.SetCellValue(rekapSheet, "A1", "Report Asuransi Periode "+tgl1+" - "+tgl2)
	rowCountKdUser := 7
	for _, each := range rekapKdUser {
		user := s.uR.FindByUsername(each["kd_user"].(string))
		if user.ID == 0 {
			fmt.Println("ini count ", rowCountKdUser)
			continue
		}
		xlsx.SetCellValue(rekapSheet, fmt.Sprintf("A%d", rowCountKdUser), user.Username)
		xlsx.SetCellValue(rekapSheet, fmt.Sprintf("B%d", rowCountKdUser), user.Name)
		xlsx.SetCellValue(rekapSheet, fmt.Sprintf("C%d", rowCountKdUser), each["p"])
		xlsx.SetCellValue(rekapSheet, fmt.Sprintf("D%d", rowCountKdUser), each["t"])
		xlsx.SetCellValue(rekapSheet, fmt.Sprintf("E%d", rowCountKdUser), each["o"])
		xlsx.SetCellValue(rekapSheet, fmt.Sprintf("F%d", rowCountKdUser), each["total"])
		xlsx.SetCellValue(rekapSheet, fmt.Sprintf("G%d", rowCountKdUser), user.DataSource)
		rowCountKdUser += 1
	}
	xlsx.SetCellStyle(rekapSheet, fmt.Sprintf("A%d", rowCountKdUser-1), fmt.Sprintf("G%d", rowCountKdUser-1), styleBorder)

	err = xlsx.SaveAs("./file-report-asuransi.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
