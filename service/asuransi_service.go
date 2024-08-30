package service

import (
	"fmt"
	"strconv"
	"time"
	"wkm/entity"
	"wkm/repository"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type AsuransiService interface {
	MasterData(search string, dataSource string, sts string, username string, tgl1 string, tgl2 string, ap string, limit int, pageParams int) []entity.MasterAsuransi
	MasterDataCount(search string, dataSource string, sts string, username string, tgl1 string, tgl2 string, ap string) int64
	RekapByStatusKdUser(tgl1 string, tgl2 string) []map[string]interface{}
	FindAsuransiByNoMsn(no_msn string) entity.MasterAsuransi
	UpdateAsuransi(data entity.MasterAsuransi) entity.MasterAsuransi
	UpdateAsuransiBerminat(no_msn string)
	UpdateAsuransiBatalBayar(no_msn string)
	UpdateAmbilAsuransi(no_msn string, kd_user string)
	MasterDataRekapTele(tgl1 string, tgl2 string) []entity.MasterRekapTele
	RekapByStatus(u string, tgl1 string, tgl2 string) entity.MasterStatusAsuransi
	RekapByStatusAll(u string, tgl1 string, tgl2 string) entity.MasterStatusAsuransi
	ExportReport(u string, tgl string)
	MasterAlasanPending() []entity.MasterAlasanPending
	MasterAlasanTdkBerminat() []entity.MasterAlasanTdkBerminat
	DetailApprovalTransaksi(idTrx string) entity.DetailApproval
	ListApprovalTransaksi(username string, tgl1 string, tgl2 string, search string, stsPembelian int, pageParams int, limit int) []entity.ListApproval
	ListApprovalTransaksiCount(username string, tgl1 string, tgl2 string, search string, stsPembelian int) int64
	AsuransiMstProduk(search string) []entity.MasterProduk
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

func (s *asuransiService) ListApprovalTransaksi(username string, tgl1 string, tgl2 string, search string, stsPembelian int, pageParams int, limit int) []entity.ListApproval {
	return s.trR.ListApprovalTransaksi(username, tgl1, tgl2, search, stsPembelian, pageParams, limit)
}
func (s *asuransiService) ListApprovalTransaksiCount(username string, tgl1 string, tgl2 string, search string, stsPembelian int) int64 {
	return s.trR.ListApprovalTransaksiCount(username, tgl1, tgl2, search, stsPembelian)
}
func (s *asuransiService) MasterData(search string, dataSource string, sts string, usename string, tgl1 string, tgl2 string, ap string, limit int, pageParams int) []entity.MasterAsuransi {
	return s.trR.MasterData(search, dataSource, sts, usename, tgl1, tgl2, ap, limit, pageParams)
}

func (s *asuransiService) MasterDataCount(search string, dataSource string, sts string, usename string, tgl1 string, tgl2 string, ap string) int64 {
	return s.trR.MasterDataCount(search, dataSource, sts, usename, tgl1, tgl2, ap)
}

func (s *asuransiService) DetailApprovalTransaksi(idTrx string) entity.DetailApproval {
	return s.trR.DetailApprovalTransaksi(idTrx)
}

func (s *asuransiService) MasterDataRekapTele(tgl1 string, tgl2 string) []entity.MasterRekapTele {
	return s.trR.MasterDataRekapTele(tgl1, tgl2)
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

func (s *asuransiService) RekapByStatusAll(u string, tgl1 string, tgl2 string) entity.MasterStatusAsuransi {
	return s.trR.RekapByStatusAll(u, tgl1, tgl2)
}

func (s *asuransiService) MasterAlasanPending() []entity.MasterAlasanPending {
	return s.trR.MasterAlasanPending()
}

func (s *asuransiService) MasterAlasanTdkBerminat() []entity.MasterAlasanTdkBerminat {
	return s.trR.MasterAlasanTdkBerminat()
}

func (s *asuransiService) AsuransiMstProduk(search string) []entity.MasterProduk {
	return s.trR.AsuransiMstProduk(search)
}

func filterPendingByKode(items []entity.MasterAlasanPending, kode int) entity.MasterAlasanPending {
	filtered := entity.MasterAlasanPending{}
	for _, item := range items {
		if item.Id == kode {
			filtered = item
			break
		}
	}
	return filtered
}

func filterTdkBerminatByKode(items []entity.MasterAlasanTdkBerminat, kode int) entity.MasterAlasanTdkBerminat {
	filtered := entity.MasterAlasanTdkBerminat{}
	for _, item := range items {
		if item.Id == kode {
			filtered = item
			break
		}
	}
	return filtered
}

func (s *asuransiService) ExportReport1(tgl1 string, tgl2 string) {

	rekapSourceInfo := s.trR.RekapByStatusJenisSource(tgl1, tgl2)
	rekapKdUser := s.trR.RekapByStatusKdUser(tgl1, tgl2)
	rincianPending := s.trR.RincianByAlasanPendingKdUser(tgl1, tgl2)
	rincianTdkBerminat := s.trR.RincianByAlasanTidakMinatKdUser(tgl1, tgl2)
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
	awalRowCountKdUser := 8
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
	xlsx.SetCellStyle(rekapSheet, fmt.Sprintf("A%d", awalRowCountKdUser), fmt.Sprintf("G%d", rowCountKdUser-1), styleBorder)

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
		if user.ID == 0 {
			continue
		}
		xlsx.SetCellValue(pendingSheet, fmt.Sprintf("A%d", rowCountRekapByAlasanPending+1), v["kd_user"])
		xlsx.SetCellValue(pendingSheet, fmt.Sprintf("B%d", rowCountRekapByAlasanPending+1), user.Name)
		xlsx.SetCellValue(pendingSheet, fmt.Sprintf("C%d", rowCountRekapByAlasanPending+1), v["kosong"])
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
	xlsx.SetCellStyle(tdkBerminatSheet, fmt.Sprintf("A%d", rowCountRekapByAlasanTdkBerminat), fmt.Sprintf("%s%d", listCol[len(masterAlasanTdkBerminat)+1], rowCountRekapByAlasanTdkBerminat), headerStyle)
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

func (s *asuransiService) ExportReport(tgl1 string, tgl2 string) {

	startDate, _ := time.Parse("2006-01-02", tgl1)
	endDate, _ := time.Parse("2006-01-02", tgl2)
	listCol := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O"}
	bulan := []string{"", "Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}
	rekapSourceInfo := s.trR.RekapByStatusJenisSource(tgl1, tgl2)
	rekapKdUser := s.trR.RekapByStatusKdUser(tgl1, tgl2)
	rekapBulanAlasanPending := s.trR.RekapBulanAlasanPending(tgl1, tgl2)
	// rekapByAlasanPending := s.trR.RekapByAlasanPending(tgl1, tgl2)
	// rincianPending := s.trR.RincianByAlasanPendingKdUser(tgl1, tgl2)
	// rincianTdkBerminat := s.trR.RincianByAlasanTidakMinatKdUser(tgl1, tgl2)
	masterPending := s.trR.MasterAlasanPending()
	// masterAlasanTdkBerminat := s.trR.MasterAlasanTdkBerminat()

	xlsx := excelize.NewFile()
	startRow := 4
	rekapSheet := "Rekap"
	tdkBerminatSheet := "Tidak Berminat"
	xlsx.SetSheetName(xlsx.GetSheetName(1), rekapSheet)
	xlsx.NewSheet(tdkBerminatSheet)
	xlsx.SetColWidth(rekapSheet, "A", "G", 14)
	xlsx.SetCellValue(rekapSheet, "A3", "Source Info")
	xlsx.SetCellValue(rekapSheet, "B3", "Bulan")
	xlsx.SetCellValue(rekapSheet, "C3", "Pending")
	xlsx.SetCellValue(rekapSheet, "D3", "Tidak Berminat")
	xlsx.SetCellValue(rekapSheet, "E3", "Berminat")
	xlsx.SetCellValue(rekapSheet, "F3", "Total")

	startMergeRekap := startRow
	for index, value := range rekapSourceInfo {
		currentSource := "Excell"
		if value["jenis_source"] == "W" {
			currentSource = "Wanda"
		}
		if index > 0 {
			if value["jenis_source"] != rekapSourceInfo[index-1]["jenis_source"] {
				xlsx.MergeCell(rekapSheet, fmt.Sprintf("A%d", startMergeRekap), fmt.Sprintf("A%d", startRow-1))
				startMergeRekap = startRow
			}
		}
		if index+1 == len(rekapSourceInfo) {
			xlsx.MergeCell(rekapSheet, fmt.Sprintf("A%d", startMergeRekap), fmt.Sprintf("A%d", startRow))
			startMergeRekap = startRow + 1
		}
		xlsx.SetCellValue(rekapSheet, fmt.Sprintf("A%d", startRow), currentSource)
		xlsx.SetCellValue(rekapSheet, fmt.Sprintf("B%d", startRow), bulan[value["bulan"].(int64)])
		xlsx.SetCellValue(rekapSheet, fmt.Sprintf("C%d", startRow), value["p"])
		xlsx.SetCellValue(rekapSheet, fmt.Sprintf("D%d", startRow), value["t"])
		xlsx.SetCellValue(rekapSheet, fmt.Sprintf("E%d", startRow), value["o"])
		xlsx.SetCellValue(rekapSheet, fmt.Sprintf("F%d", startRow), value["total"])
		startRow += 1
	}

	startRow += 1
	xlsx.SetCellValue(rekapSheet, fmt.Sprintf("A%d", startRow), "Id User")
	xlsx.SetCellValue(rekapSheet, fmt.Sprintf("B%d", startRow), "Nama")
	xlsx.SetCellValue(rekapSheet, fmt.Sprintf("C%d", startRow), "Bulan")
	xlsx.SetCellValue(rekapSheet, fmt.Sprintf("D%d", startRow), "Pending")
	xlsx.SetCellValue(rekapSheet, fmt.Sprintf("E%d", startRow), "Tidak Berminat")
	xlsx.SetCellValue(rekapSheet, fmt.Sprintf("F%d", startRow), "Berminat")
	xlsx.SetCellValue(rekapSheet, fmt.Sprintf("G%d", startRow), "Total")
	xlsx.SetCellValue(rekapSheet, fmt.Sprintf("H%d", startRow), "Source Info")
	startRow += 1
	startMergeRekap = startRow
	for index, value := range rekapKdUser {
		user := s.uR.FindByUsername(value["kd_user"].(string))
		if user.ID == 0 {
			continue
		}
		if index > 0 {
			if value["kd_user"] != rekapKdUser[index-1]["kd_user"] {
				xlsx.MergeCell(rekapSheet, fmt.Sprintf("A%d", startMergeRekap), fmt.Sprintf("A%d", startRow-1))
				xlsx.MergeCell(rekapSheet, fmt.Sprintf("B%d", startMergeRekap), fmt.Sprintf("B%d", startRow-1))
				startMergeRekap = startRow
			}
		}
		if index+1 == len(rekapKdUser) {
			xlsx.MergeCell(rekapSheet, fmt.Sprintf("A%d", startMergeRekap), fmt.Sprintf("A%d", startRow))
			xlsx.MergeCell(rekapSheet, fmt.Sprintf("B%d", startMergeRekap), fmt.Sprintf("B%d", startRow))
			startMergeRekap = startRow + 1
		}
		xlsx.SetCellValue(rekapSheet, fmt.Sprintf("A%d", startRow), value["kd_user"])
		xlsx.SetCellValue(rekapSheet, fmt.Sprintf("B%d", startRow), user.Name)
		xlsx.SetCellValue(rekapSheet, fmt.Sprintf("C%d", startRow), bulan[value["bulan"].(int64)])
		xlsx.SetCellValue(rekapSheet, fmt.Sprintf("D%d", startRow), value["p"])
		xlsx.SetCellValue(rekapSheet, fmt.Sprintf("E%d", startRow), value["t"])
		xlsx.SetCellValue(rekapSheet, fmt.Sprintf("F%d", startRow), value["o"])
		xlsx.SetCellValue(rekapSheet, fmt.Sprintf("G%d", startRow), value["total"])
		xlsx.SetCellValue(rekapSheet, fmt.Sprintf("H%d", startRow), user.DataSource)
		startRow += 1
	}

	pendingSheet := "Pending"
	xlsx.NewSheet(pendingSheet)
	startRow = 4
	startMergeRekap = startRow

	xlsx.SetCellValue(pendingSheet, "A3", "Alasan")
	startCol := 1
	for d := startDate; d.Before(endDate) || d.Equal(endDate); d = d.AddDate(0, 1, 0) {
		xlsx.SetCellValue(pendingSheet, fmt.Sprintf("%s3", listCol[startCol]), bulan[d.Month()])
		startCol += 1
	}
	for _, value := range rekapBulanAlasanPending {
		startCol = 1
		kodeAlasanPending, _ := strconv.Atoi(value["alasan_pending"].(string))
		alasanPending := filterPendingByKode(masterPending, kodeAlasanPending)
		xlsx.SetCellValue(pendingSheet, fmt.Sprintf("A%d", startRow), alasanPending.Nama)
		for d := startDate; d.Before(endDate) || d.Equal(endDate); d = d.AddDate(0, 1, 0) {
			xlsx.SetCellValue(pendingSheet, fmt.Sprintf("%s%d", listCol[startCol], startRow), value[bulan[d.Month()]])
			startCol += 1
		}
		startRow += 1

	}
	// xlsx.SetCellValue(rekapSheet, "B7", "Nama")
	// xlsx.SetCellValue(rekapSheet, "C7", "Pending")
	// xlsx.SetCellValue(rekapSheet, "D7", "Tidak Berminat")
	// xlsx.SetCellValue(rekapSheet, "E7", "Berminat")
	// xlsx.SetCellValue(rekapSheet, "F7", "Total")
	// xlsx.SetCellValue(rekapSheet, "G7", "Source Info")

	// xlsx.SetCellValue(pendingSheet, "A3", "Alasan")
	// xlsx.SetCellValue(pendingSheet, "B3", "Total")
	// xlsx.SetColWidth(pendingSheet, "A", "A", 20)
	// xlsx.SetColWidth(pendingSheet, "B", "G", 11)

	// xlsx.SetCellValue(tdkBerminatSheet, "A3", "Alasan")
	// xlsx.SetCellValue(tdkBerminatSheet, "B3", "Total")
	// xlsx.SetColWidth(tdkBerminatSheet, "A", "A", 20)
	// xlsx.SetColWidth(tdkBerminatSheet, "B", "G", 11)

	// headerStyle, err := xlsx.NewStyle(`{
	// 	"font": {
	// 		"bold": true
	// 	},
	// 	"alignment": {
	//         "horizontal": "center",
	//         "vertical": "center"
	//     },
	// 	"fill": {
	// 		"type": "pattern",
	// 		"color": [
	// 			"#FFFF00"
	// 		],
	// 		"pattern": 1
	// 	},
	// 	"border": [
	// 		{
	// 			"type": "left",
	// 			"style": 1,
	// 			"color": "#000000"
	// 		},
	// 		{
	// 			"type": "top",
	// 			"style": 1,
	// 			"color": "#000000"
	// 		},
	// 		{
	// 			"type": "right",
	// 			"style": 1,
	// 			"color": "#000000"
	// 		},
	// 		{
	// 			"type": "bottom",
	// 			"style": 1,
	// 			"color": "#000000"
	// 		}
	// 	]
	// }`)
	// if err != nil {
	// 	fmt.Println("ini error style ", err)
	// }
	// styleBorder, err := xlsx.NewStyle(`{
	// 	"border": [
	//     {
	//         "type": "left",
	//         "style": 1,
	//         "color": "#000000"
	//     },
	//     {
	//         "type": "top",
	//         "style": 1,
	//         "color": "#000000"
	//     },
	//     {
	//         "type": "right",
	//         "style": 1,
	//         "color": "#000000"
	//     },
	//     {
	//         "type": "bottom",
	//         "style": 1,
	//         "color": "#000000"
	// 	}
	// ]
	// }`)
	// if err != nil {
	// 	fmt.Println("ini error style ", err)
	// }

	// xlsx.SetCellStyle(rekapSheet, "A1", "G1", headerStyle)
	// xlsx.SetCellStyle(pendingSheet, "A1", "G1", headerStyle)
	// xlsx.SetCellStyle(tdkBerminatSheet, "A1", "G1", headerStyle)

	// xlsx.SetCellStyle(rekapSheet, "A3", "E3", headerStyle)
	// xlsx.SetCellStyle(pendingSheet, "A3", "B3", headerStyle)
	// xlsx.SetCellStyle(tdkBerminatSheet, "A3", "B3", headerStyle)

	// xlsx.SetCellStyle(rekapSheet, "A4", "A5", headerStyle)
	// xlsx.SetCellStyle(rekapSheet, "A7", "G7", headerStyle)
	// xlsx.SetCellStyle(rekapSheet, "B4", "E5", styleBorder)
	// xlsx.MergeCell(rekapSheet, "A1", "G1")
	// xlsx.MergeCell(pendingSheet, "A1", "G1")
	// xlsx.MergeCell(tdkBerminatSheet, "A1", "G1")

	// xlsx.SetCellValue(rekapSheet, "A1", "Report Asuransi Periode "+tgl1+" - "+tgl2)
	// xlsx.SetCellValue(pendingSheet, "A1", "Report Asuransi Periode "+tgl1+" - "+tgl2)
	// xlsx.SetCellValue(tdkBerminatSheet, "A1", "Report Asuransi Periode "+tgl1+" - "+tgl2)

	// jenis_source := ""
	// for i, each := range rekapSourceInfo {
	// 	if each["jenis_source"] == "W" {
	// 		jenis_source = "Wanda"
	// 	}
	// 	if each["jenis_source"] == "E" {
	// 		jenis_source = "Excel"
	// 	}
	// 	xlsx.SetCellValue(rekapSheet, fmt.Sprintf("A%d", i+4), jenis_source)
	// 	xlsx.SetCellValue(rekapSheet, fmt.Sprintf("B%d", i+4), each["p"])
	// 	xlsx.SetCellValue(rekapSheet, fmt.Sprintf("C%d", i+4), each["t"])
	// 	xlsx.SetCellValue(rekapSheet, fmt.Sprintf("D%d", i+4), each["o"])
	// 	xlsx.SetCellValue(rekapSheet, fmt.Sprintf("E%d", i+4), each["total"])
	// }

	// rowCountKdUser := 8
	// awalRowCountKdUser := 8
	// for _, each := range rekapKdUser {
	// 	user := s.uR.FindByUsername(each["kd_user"].(string))
	// 	if user.ID == 0 {
	// 		continue
	// 	}
	// 	xlsx.SetCellValue(rekapSheet, fmt.Sprintf("A%d", rowCountKdUser), user.Username)
	// 	xlsx.SetCellValue(rekapSheet, fmt.Sprintf("B%d", rowCountKdUser), user.Name)
	// 	xlsx.SetCellValue(rekapSheet, fmt.Sprintf("C%d", rowCountKdUser), each["p"])
	// 	xlsx.SetCellValue(rekapSheet, fmt.Sprintf("D%d", rowCountKdUser), each["t"])
	// 	xlsx.SetCellValue(rekapSheet, fmt.Sprintf("E%d", rowCountKdUser), each["o"])
	// 	xlsx.SetCellValue(rekapSheet, fmt.Sprintf("F%d", rowCountKdUser), each["total"])
	// 	xlsx.SetCellValue(rekapSheet, fmt.Sprintf("G%d", rowCountKdUser), user.DataSource)
	// 	rowCountKdUser += 1
	// }
	// xlsx.SetCellStyle(rekapSheet, fmt.Sprintf("A%d", awalRowCountKdUser), fmt.Sprintf("G%d", rowCountKdUser-1), styleBorder)

	// rowCountRekapByAlasanPending := 4
	// awalRowCountRekapByAlasanPending := 4
	// rekapByAlasanPending := s.trR.RekapByAlasanPending(tgl1, tgl2)
	// for _, each := range rekapByAlasanPending {
	// 	xlsx.SetCellValue(pendingSheet, fmt.Sprintf("A%d", rowCountRekapByAlasanPending), each["alasan"])
	// 	xlsx.SetCellValue(pendingSheet, fmt.Sprintf("B%d", rowCountRekapByAlasanPending), each["total"])
	// 	rowCountRekapByAlasanPending += 1
	// }
	// xlsx.SetCellStyle(pendingSheet, fmt.Sprintf("A%d", awalRowCountRekapByAlasanPending), fmt.Sprintf("B%d", rowCountRekapByAlasanPending-1), styleBorder)

	// listCol := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	// rowCountRekapByAlasanPending += 1
	// xlsx.SetCellValue(pendingSheet, fmt.Sprintf("A%d", rowCountRekapByAlasanPending), "Id User")
	// xlsx.SetCellValue(pendingSheet, fmt.Sprintf("B%d", rowCountRekapByAlasanPending), "Nama")
	// xlsx.SetCellValue(pendingSheet, fmt.Sprintf("C%d", rowCountRekapByAlasanPending), "Tidak Ada Alasan")
	// xlsx.SetCellStyle(pendingSheet, fmt.Sprintf("A%d", rowCountRekapByAlasanPending), fmt.Sprintf("%s%d", listCol[len(masterPending)+2], rowCountRekapByAlasanPending), headerStyle)
	// for i, v := range masterPending {
	// 	xlsx.SetCellValue(pendingSheet, fmt.Sprintf("%s%d", listCol[i+3], rowCountRekapByAlasanPending), v.Nama)
	// 	xlsx.SetCellValue(pendingSheet, fmt.Sprintf("%s%d", listCol[i+3], rowCountRekapByAlasanPending), v.Nama)
	// 	xlsx.SetColWidth(pendingSheet, listCol[i+3], listCol[i+3], float64(len(v.Nama))+4)
	// }
	// for _, v := range rincianPending {
	// 	user := s.uR.FindByUsername(v["kd_user"].(string))
	// 	if user.ID == 0 {
	// 		continue
	// 	}
	// 	xlsx.SetCellValue(pendingSheet, fmt.Sprintf("A%d", rowCountRekapByAlasanPending+1), v["kd_user"])
	// 	xlsx.SetCellValue(pendingSheet, fmt.Sprintf("B%d", rowCountRekapByAlasanPending+1), user.Name)
	// 	xlsx.SetCellValue(pendingSheet, fmt.Sprintf("C%d", rowCountRekapByAlasanPending+1), v["kosong"])
	// 	for j, vj := range masterPending {
	// 		xlsx.SetCellValue(pendingSheet, fmt.Sprintf("%s%d", listCol[j+3], rowCountRekapByAlasanPending+1), v[fmt.Sprintf("%d", vj.Id)])
	// 	}
	// 	rowCountRekapByAlasanPending += 1
	// }

	// rowCountRekapByAlasanTdkBerminat := 4
	// awalRowCountRekapByAlasanTdkBerminat := 4
	// rekapByAlasanTdkBerminat := s.trR.RekapByAlasanTdkBerminat(tgl1, tgl2)
	// for _, each := range rekapByAlasanTdkBerminat {
	// 	xlsx.SetCellValue(tdkBerminatSheet, fmt.Sprintf("A%d", rowCountRekapByAlasanTdkBerminat), each["alasan"])
	// 	xlsx.SetCellValue(tdkBerminatSheet, fmt.Sprintf("B%d", rowCountRekapByAlasanTdkBerminat), each["total"])
	// 	rowCountRekapByAlasanTdkBerminat += 1
	// }
	// xlsx.SetCellStyle(tdkBerminatSheet, fmt.Sprintf("A%d", awalRowCountRekapByAlasanTdkBerminat), fmt.Sprintf("B%d", rowCountRekapByAlasanTdkBerminat-1), styleBorder)

	// rowCountRekapByAlasanTdkBerminat += 1
	// xlsx.SetCellValue(tdkBerminatSheet, fmt.Sprintf("A%d", rowCountRekapByAlasanTdkBerminat), "Id User")
	// xlsx.SetCellValue(tdkBerminatSheet, fmt.Sprintf("B%d", rowCountRekapByAlasanTdkBerminat), "Nama")
	// xlsx.SetCellStyle(tdkBerminatSheet, fmt.Sprintf("A%d", rowCountRekapByAlasanTdkBerminat), fmt.Sprintf("%s%d", listCol[len(masterAlasanTdkBerminat)+1], rowCountRekapByAlasanTdkBerminat), headerStyle)
	// for i, v := range masterAlasanTdkBerminat {
	// 	xlsx.SetCellValue(tdkBerminatSheet, fmt.Sprintf("%s%d", listCol[i+2], rowCountRekapByAlasanTdkBerminat), v.Nama)
	// 	xlsx.SetCellValue(tdkBerminatSheet, fmt.Sprintf("%s%d", listCol[i+2], rowCountRekapByAlasanTdkBerminat), v.Nama)
	// 	xlsx.SetColWidth(tdkBerminatSheet, listCol[i+2], listCol[i+2], float64(len(v.Nama))+4)
	// }
	// for _, v := range rincianTdkBerminat {
	// 	user := s.uR.FindByUsername(v["kd_user"].(string))
	// 	if user.ID == 0 {
	// 		continue
	// 	}
	// 	xlsx.SetCellValue(tdkBerminatSheet, fmt.Sprintf("A%d", rowCountRekapByAlasanTdkBerminat+1), v["kd_user"])
	// 	xlsx.SetCellValue(tdkBerminatSheet, fmt.Sprintf("B%d", rowCountRekapByAlasanTdkBerminat+1), user.Name)
	// 	for j, vj := range masterAlasanTdkBerminat {
	// 		xlsx.SetCellValue(tdkBerminatSheet, fmt.Sprintf("%s%d", listCol[j+2], rowCountRekapByAlasanTdkBerminat+1), v[fmt.Sprintf("%d", vj.Id)])
	// 	}
	// 	rowCountRekapByAlasanTdkBerminat += 1
	// }

	err := xlsx.SaveAs("./file-report-asuransi.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
