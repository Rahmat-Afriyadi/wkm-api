package service

import (
	"fmt"
	"wkm/entity"
	"wkm/repository"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type AsuransiService interface {
	MasterData(search string, dataSource string, sts string, username string, tgl1 string, tgl2 string, limit int, pageParams int) []entity.MasterAsuransi
	MasterDataCount(search string, dataSource string, sts string, username string, tgl1 string, tgl2 string) int64
	RekapByStatusKdUser(tgl1 string, tgl2 string) []map[string]interface{}
	FindAsuransiByNoMsn(no_msn string) entity.MasterAsuransi
	UpdateAsuransi(data entity.MasterAsuransi) entity.MasterAsuransi
	UpdateAsuransiBerminat(no_msn string)
	UpdateAsuransiBatalBayar(no_msn string)
	UpdateAmbilAsuransi(no_msn string, kd_user string)
	MasterDataRekapTele() []entity.MasterRekapTele
	RekapByStatus(u string, tgl1 string, tgl2 string) entity.MasterStatusAsuransi
	ExportReport(u string, tgl string)
	MasterAlasanPending() []entity.MasterAlasanPending
	MasterAlasanTdkBerminat() []entity.MasterAlasanTdkBerminat
	DetailApprovalTransaksi(idTrx string) entity.DetailApproval
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

func (s *asuransiService) MasterData(search string, dataSource string, sts string, usename string, tgl1 string, tgl2 string, limit int, pageParams int) []entity.MasterAsuransi {
	return s.trR.MasterData(search, dataSource, sts, usename, tgl1, tgl2, limit, pageParams)
}

func (s *asuransiService) MasterDataCount(search string, dataSource string, sts string, usename string, tgl1 string, tgl2 string) int64 {
	return s.trR.MasterDataCount(search, dataSource, sts, usename, tgl1, tgl2)
}

func (s *asuransiService) DetailApprovalTransaksi(idTrx string) entity.DetailApproval {
	return s.trR.DetailApprovalTransaksi(idTrx)
}

func (s *asuransiService) MasterDataRekapTele() []entity.MasterRekapTele {
	return s.trR.MasterDataRekapTele()
}

func (s *asuransiService) RekapByStatusKdUser(tgl1 string, tgl2 string) []map[string]interface{} {
	datas := s.trR.RekapByStatusKdUser(tgl1, tgl2)
	for _, v := range datas {
		user := s.uR.FindByUsername(v["kd_user"].(string))
		if user.ID == 0 {
			continue
		}
		v["name"] = user.Name
	}
	return datas
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

func (s *asuransiService) RekapByStatus(u string, tgl1 string, tgl2 string) entity.MasterStatusAsuransi {
	return s.trR.RekapByStatus(u, tgl1, tgl2)
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
	rincianPending := s.trR.RincianByAlasanPendingKdUser("2024-05-01", "2024-05-30")
	rincianTdkBerminat := s.trR.RincianByAlasanTidakMinatKdUser("2024-05-01", "2024-05-30")
	masterPending := s.trR.MasterAlasanPending()
	masterAlasanTdkBerminat := s.trR.MasterAlasanTdkBerminat()
	fmt.Println("ini rekap user ", rekapKdUser)

	xlsx := excelize.NewFile()

	rekapSheet := "Rekap"
	pendingSheet := "Pending"
	tdkBerminatSheet := "Tidak Berminat"
	xlsx.SetSheetName(xlsx.GetSheetName(1), rekapSheet)
	xlsx.NewSheet(pendingSheet)
	xlsx.NewSheet(tdkBerminatSheet)
	xlsx.SetCellValue(rekapSheet, "A3", "Source Info")
	xlsx.SetColWidth(rekapSheet, "A", "G", 14)
	xlsx.SetCellValue(rekapSheet, "B3", "Pending")
	xlsx.SetCellValue(rekapSheet, "C3", "Tidak Berminat")
	xlsx.SetCellValue(rekapSheet, "D3", "Berminat")
	xlsx.SetCellValue(rekapSheet, "E3", "Total")
	xlsx.SetCellValue(rekapSheet, "A7", "Id User")
	xlsx.SetCellValue(rekapSheet, "B7", "Nama")
	xlsx.SetCellValue(rekapSheet, "C7", "Pending")
	xlsx.SetCellValue(rekapSheet, "D7", "Tidak Berminat")
	xlsx.SetCellValue(rekapSheet, "E7", "Berminat")
	xlsx.SetCellValue(rekapSheet, "F7", "Total")
	xlsx.SetCellValue(rekapSheet, "G7", "Source Info")

	xlsx.SetCellValue(pendingSheet, "A3", "Alasan")
	xlsx.SetCellValue(pendingSheet, "B3", "Total")
	xlsx.SetColWidth(pendingSheet, "A", "A", 20)
	xlsx.SetColWidth(pendingSheet, "B", "G", 11)

	xlsx.SetCellValue(tdkBerminatSheet, "A3", "Alasan")
	xlsx.SetCellValue(tdkBerminatSheet, "B3", "Total")
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

	xlsx.SetCellStyle(rekapSheet, "A1", "G1", headerStyle)
	xlsx.SetCellStyle(pendingSheet, "A1", "G1", headerStyle)
	xlsx.SetCellStyle(tdkBerminatSheet, "A1", "G1", headerStyle)

	xlsx.SetCellStyle(rekapSheet, "A3", "E3", headerStyle)
	xlsx.SetCellStyle(pendingSheet, "A3", "B3", headerStyle)
	xlsx.SetCellStyle(tdkBerminatSheet, "A3", "B3", headerStyle)

	xlsx.SetCellStyle(rekapSheet, "A4", "A5", headerStyle)
	xlsx.SetCellStyle(rekapSheet, "A7", "G7", headerStyle)
	xlsx.SetCellStyle(rekapSheet, "B4", "E5", styleBorder)
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
		xlsx.SetCellValue(rekapSheet, fmt.Sprintf("A%d", i+4), jenis_source)
		xlsx.SetCellValue(rekapSheet, fmt.Sprintf("B%d", i+4), each["p"])
		xlsx.SetCellValue(rekapSheet, fmt.Sprintf("C%d", i+4), each["t"])
		xlsx.SetCellValue(rekapSheet, fmt.Sprintf("D%d", i+4), each["o"])
		xlsx.SetCellValue(rekapSheet, fmt.Sprintf("E%d", i+4), each["total"])
	}

	rowCountKdUser := 8
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

	rowCountRekapByAlasanPending := 4
	awalRowCountRekapByAlasanPending := 4
	rekapByAlasanPending := s.trR.RekapByAlasanPending(tgl1, tgl2)
	for _, each := range rekapByAlasanPending {
		xlsx.SetCellValue(pendingSheet, fmt.Sprintf("A%d", rowCountRekapByAlasanPending), each["alasan"])
		xlsx.SetCellValue(pendingSheet, fmt.Sprintf("B%d", rowCountRekapByAlasanPending), each["total"])
		rowCountRekapByAlasanPending += 1
	}
	xlsx.SetCellStyle(pendingSheet, fmt.Sprintf("A%d", awalRowCountRekapByAlasanPending), fmt.Sprintf("B%d", rowCountRekapByAlasanPending-1), styleBorder)

	listCol := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	rowCountRekapByAlasanPending += 1
	xlsx.SetCellValue(pendingSheet, fmt.Sprintf("A%d", rowCountRekapByAlasanPending), "Id User")
	xlsx.SetCellValue(pendingSheet, fmt.Sprintf("B%d", rowCountRekapByAlasanPending), "Nama")
	xlsx.SetCellValue(pendingSheet, fmt.Sprintf("C%d", rowCountRekapByAlasanPending), "Tidak Ada Alasan")
	xlsx.SetCellStyle(pendingSheet, fmt.Sprintf("A%d", rowCountRekapByAlasanPending), fmt.Sprintf("%s%d", listCol[len(masterPending)+2], rowCountRekapByAlasanPending), headerStyle)
	for i, v := range masterPending {
		xlsx.SetCellValue(pendingSheet, fmt.Sprintf("%s%d", listCol[i+3], rowCountRekapByAlasanPending), v.Nama)
		xlsx.SetCellValue(pendingSheet, fmt.Sprintf("%s%d", listCol[i+3], rowCountRekapByAlasanPending), v.Nama)
		xlsx.SetColWidth(pendingSheet, listCol[i+3], listCol[i+3], float64(len(v.Nama))+4)
	}
	for _, v := range rincianPending {
		user := s.uR.FindByUsername(v["kd_user"].(string))
		fmt.Println("ini user id nya yaa ", user)
		if user.ID == 0 {
			continue
		}
		xlsx.SetCellValue(pendingSheet, fmt.Sprintf("A%d", rowCountRekapByAlasanPending+1), v["kd_user"])
		xlsx.SetCellValue(pendingSheet, fmt.Sprintf("B%d", rowCountRekapByAlasanPending+1), user.Name)
		xlsx.SetCellValue(pendingSheet, fmt.Sprintf("B%d", rowCountRekapByAlasanPending+1), v["kosong"])
		for j, vj := range masterPending {
			xlsx.SetCellValue(pendingSheet, fmt.Sprintf("%s%d", listCol[j+3], rowCountRekapByAlasanPending+1), v[fmt.Sprintf("%d", vj.Id)])
		}
		rowCountRekapByAlasanPending += 1
	}

	rowCountRekapByAlasanTdkBerminat := 4
	awalRowCountRekapByAlasanTdkBerminat := 4
	rekapByAlasanTdkBerminat := s.trR.RekapByAlasanTdkBerminat(tgl1, tgl2)
	for _, each := range rekapByAlasanTdkBerminat {
		xlsx.SetCellValue(tdkBerminatSheet, fmt.Sprintf("A%d", rowCountRekapByAlasanTdkBerminat), each["alasan"])
		xlsx.SetCellValue(tdkBerminatSheet, fmt.Sprintf("B%d", rowCountRekapByAlasanTdkBerminat), each["total"])
		rowCountRekapByAlasanTdkBerminat += 1
	}
	xlsx.SetCellStyle(tdkBerminatSheet, fmt.Sprintf("A%d", awalRowCountRekapByAlasanTdkBerminat), fmt.Sprintf("B%d", rowCountRekapByAlasanTdkBerminat-1), styleBorder)

	rowCountRekapByAlasanTdkBerminat += 1
	xlsx.SetCellValue(tdkBerminatSheet, fmt.Sprintf("A%d", rowCountRekapByAlasanTdkBerminat), "Id User")
	xlsx.SetCellValue(tdkBerminatSheet, fmt.Sprintf("B%d", rowCountRekapByAlasanTdkBerminat), "Nama")
	xlsx.SetCellStyle(tdkBerminatSheet, fmt.Sprintf("A%d", rowCountRekapByAlasanTdkBerminat), fmt.Sprintf("%s%d", listCol[len(masterAlasanTdkBerminat)+2], rowCountRekapByAlasanTdkBerminat), headerStyle)
	for i, v := range masterAlasanTdkBerminat {
		xlsx.SetCellValue(tdkBerminatSheet, fmt.Sprintf("%s%d", listCol[i+2], rowCountRekapByAlasanTdkBerminat), v.Nama)
		xlsx.SetCellValue(tdkBerminatSheet, fmt.Sprintf("%s%d", listCol[i+2], rowCountRekapByAlasanTdkBerminat), v.Nama)
		xlsx.SetColWidth(tdkBerminatSheet, listCol[i+2], listCol[i+2], float64(len(v.Nama))+4)
	}
	for _, v := range rincianTdkBerminat {
		user := s.uR.FindByUsername(v["kd_user"].(string))
		if user.ID == 0 {
			continue
		}
		xlsx.SetCellValue(tdkBerminatSheet, fmt.Sprintf("A%d", rowCountRekapByAlasanTdkBerminat+1), v["kd_user"])
		xlsx.SetCellValue(tdkBerminatSheet, fmt.Sprintf("B%d", rowCountRekapByAlasanTdkBerminat+1), user.Name)
		for j, vj := range masterAlasanTdkBerminat {
			xlsx.SetCellValue(tdkBerminatSheet, fmt.Sprintf("%s%d", listCol[j+2], rowCountRekapByAlasanTdkBerminat+1), v[fmt.Sprintf("%d", vj.Id)])
		}
		rowCountRekapByAlasanTdkBerminat += 1
	}

	err = xlsx.SaveAs("./file-report-asuransi.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
