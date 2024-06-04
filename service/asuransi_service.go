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
	fmt.Println("ini rekap user ", rekapKdUser)

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

	xlsx.SetCellValue(pendingSheet, "A2", "Alasan")
	xlsx.SetCellValue(pendingSheet, "B2", "Total")
	xlsx.SetColWidth(pendingSheet, "A", "A", 20)
	xlsx.SetColWidth(pendingSheet, "B", "G", 11)

	xlsx.SetCellValue(tdkBerminatSheet, "A2", "Alasan")
	xlsx.SetCellValue(tdkBerminatSheet, "B2", "Total")
	xlsx.SetColWidth(tdkBerminatSheet, "A", "A", 20)
	xlsx.SetColWidth(tdkBerminatSheet, "B", "G", 11)

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

	xlsx.SetCellStyle(pendingSheet, "A1", "G1", headerStyle)
	xlsx.SetCellStyle(tdkBerminatSheet, "A1", "G1", headerStyle)

	xlsx.SetCellStyle(rekapSheet, "A2", "E2", headerStyle)
	xlsx.SetCellStyle(pendingSheet, "A2", "B2", headerStyle)
	xlsx.SetCellStyle(tdkBerminatSheet, "A2", "B2", headerStyle)

	xlsx.SetCellStyle(rekapSheet, "A3", "A4", headerStyle)
	xlsx.SetCellStyle(rekapSheet, "A6", "G6", headerStyle)
	xlsx.SetCellStyle(rekapSheet, "B3", "E4", styleBorder)
	xlsx.MergeCell(rekapSheet, "A1", "G1")
	xlsx.MergeCell(pendingSheet, "A1", "G1")
	xlsx.MergeCell(tdkBerminatSheet, "A1", "G1")

	xlsx.SetCellValue(rekapSheet, "A1", "Report Asuransi Periode "+tgl1+" - "+tgl2)
	xlsx.SetCellValue(pendingSheet, "A1", "Report Asuransi Periode "+tgl1+" - "+tgl2)
	xlsx.SetCellValue(tdkBerminatSheet, "A1", "Report Asuransi Periode "+tgl1+" - "+tgl2)

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

	rowCountKdUser := 7
	for _, each := range rekapKdUser {
		user := s.uR.FindByUsername(each["kd_user"].(string))
		if user.ID == 0 {
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

	rowCountRekapByAlasanPending := 3
	awalRowCountRekapByAlasanPending := 3
	rekapByAlasanPending := s.trR.RekapByAlasanPending(tgl1, tgl2)
	for _, each := range rekapByAlasanPending {
		xlsx.SetCellValue(pendingSheet, fmt.Sprintf("A%d", rowCountRekapByAlasanPending), each["alasan"])
		xlsx.SetCellValue(pendingSheet, fmt.Sprintf("B%d", rowCountRekapByAlasanPending), each["total"])
		rowCountRekapByAlasanPending += 1
	}
	xlsx.SetCellStyle(pendingSheet, fmt.Sprintf("A%d", awalRowCountRekapByAlasanPending), fmt.Sprintf("B%d", rowCountRekapByAlasanPending-1), styleBorder)

	rowCountRekapByAlasanTdkBerminat := 3
	awalRowCountRekapByAlasanTdkBerminat := 3
	rekapByAlasanTdkBerminat := s.trR.RekapByAlasanTdkBerminat(tgl1, tgl2)
	for _, each := range rekapByAlasanTdkBerminat {
		xlsx.SetCellValue(tdkBerminatSheet, fmt.Sprintf("A%d", rowCountRekapByAlasanTdkBerminat), each["alasan"])
		xlsx.SetCellValue(tdkBerminatSheet, fmt.Sprintf("B%d", rowCountRekapByAlasanTdkBerminat), each["total"])
		rowCountRekapByAlasanTdkBerminat += 1
	}
	xlsx.SetCellStyle(tdkBerminatSheet, fmt.Sprintf("A%d", awalRowCountRekapByAlasanTdkBerminat), fmt.Sprintf("B%d", rowCountRekapByAlasanTdkBerminat-1), styleBorder)

	err = xlsx.SaveAs("./file-report-asuransi.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
