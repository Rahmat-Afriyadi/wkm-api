package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"wkm/entity"
	"wkm/repository"
	"wkm/request"
	"wkm/response"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type CustomerMtrService interface {
	AllStatusMasterData(search string, username string, tgl_bayar1 string, tgl_bayar2 string, limit int, pageParams int) []response.AllStatusResponse
	AllStatusMasterDataCount(search string, username string, tgl_bayar1 string, tgl_bayar2 string) int64
	MasterData(search string, sts string, jns string, username string, limit int, pageParams int) []entity.CustomerMtr
	MasterDataCount(search string, sts string, jns string, username string) int64
	MasterDataBalikan(search string, tgl1 string, tgl2 string, username string, limit int, pageParams int) []response.TelesalesBalikanResponseList
	MasterDataBalikanCount(search string, tgl1 string, tgl2 string, username string) int64
	MasterDataBalikanKonfirmer(search string, tgl1 string, tgl2 string, limit int, pageParams int) []response.TelesalesBalikanResponseList
	MasterDataBalikanKonfirmerCount(search string, tgl1 string, tgl2 string) int64
	ListAmbilData() []entity.Faktur3
	EmpatAmbilData(no_msn string) error
	AmbilDataBalikan(no_msn string, kd_user string) error
	AmbilDataAllStatus(no_msn string, kd_user string) error
	AmbilDataBalikanKonfirmer(no_msn string, kd_user string) error
	AmbilData(no_msn string, kd_user string) error
	SelfCount(kd_user string) int64
	Show(no_msn string) response.TelesalesResponse
	UpdateOkeMembership(customer request.CustomerMtr) (entity.CustomerMtr, error)
	RekapTele(username string, startDate time.Time, endDate time.Time) (response.RekapTele, error)
	ListBerminatMembership(usrname string, startDate time.Time, endDate time.Time, limit int, pageParams int, search string) ([]response.MinatMembership, int, int, error)
	ListDataAsuransiPA(usrname string, startDate time.Time, endDate time.Time, limit int, pageParams int, search string) ([]response.ListAsuransi, int, int, error)
	ListDataAsuransiMtr(usrname string, startDate time.Time, endDate time.Time, limit int, pageParams int, search string) ([]response.ListAsuransi, int, int, error)
	ExportRekapTele(usrname string, startDate time.Time, endDate time.Time) (string, error)
	RekapLeaderTs(startDate time.Time, endDate time.Time) (response.RekapLeaderTs, error)
	RekapBerminatPerWilayah(startDate time.Time, endDate time.Time) ([]response.RekapBerminatPerWilayah,int, error)
	ExportRekapLeaderTs(startDate time.Time, endDate time.Time) (string, error)
	ListPerformanceTs(startDate time.Time, endDate time.Time) (response.PerformanceTs, error)
	GetRekapStatus(startDate time.Time, endDate time.Time) ([]response.RekapStatus, error)
	ListDataPerKecamatan(startDate time.Time, endDate time.Time,limit int, pageParams int, search string) ([]response.RekapBerminatPerWilayah,int,int,int, error)
}

type customerMtrService struct {
	cR repository.CustomerMtrRepository
}

func NewCustomerMtrService(cR repository.CustomerMtrRepository) CustomerMtrService {
	return &customerMtrService{
		cR: cR,
	}
}

func (cS *customerMtrService) SelfCount(kd_user string) int64 {
	return 	cS.cR.SelfCount(kd_user)
}
func (cS *customerMtrService) MasterDataBalikan(search string,tgl1 string, tgl2 string, username string, limit int, pageParams int) []response.TelesalesBalikanResponseList {
	return cS.cR.MasterDataBalikan(search, tgl1, tgl2, username, limit, pageParams)
}
func (cS *customerMtrService) MasterDataBalikanCount(search string,tgl1 string, tgl2 string, username string) int64 {
	return cS.cR.MasterDataBalikanCount(search, tgl1, tgl2, username)
}
func (cS *customerMtrService) MasterDataBalikanKonfirmer(search string,tgl1 string, tgl2 string, limit int, pageParams int) []response.TelesalesBalikanResponseList {
	return cS.cR.MasterDataBalikanKonfirmer(search, tgl1, tgl2, limit, pageParams)
}
func (cS *customerMtrService) MasterDataBalikanKonfirmerCount(search string,tgl1 string, tgl2 string) int64 {
	return cS.cR.MasterDataBalikanKonfirmerCount(search,tgl1, tgl2,)
}
func (cS *customerMtrService) MasterData(search string, sts string, jns string, username string, limit int, pageParams int) []entity.CustomerMtr {
	return cS.cR.MasterData(search, sts, jns, username, limit, pageParams)
}
func (cS *customerMtrService) MasterDataCount(search string, sts string, jns string, username string) int64 {
	return cS.cR.MasterDataCount(search, sts, jns, username)
}
func (cS *customerMtrService) AllStatusMasterData(search string, username string,tgl_bayar1 string, tgl_bayar2 string, limit int, pageParams int) []response.AllStatusResponse{
	return cS.cR.AllStatusMasterData(search, username, tgl_bayar1, tgl_bayar2, limit, pageParams)
}
func (cS *customerMtrService) AllStatusMasterDataCount(search string, username string,tgl_bayar1 string, tgl_bayar2 string) int64 {
	return cS.cR.AllStatusMasterDataCount(search, username, tgl_bayar1, tgl_bayar2)
}

func (cS *customerMtrService) ListAmbilData() []entity.Faktur3 {
	return cS.cR.ListAmbilData()
}

func (cS *customerMtrService) EmpatAmbilData(no_msn string) error  {
	return cS.cR.EmpatAmbilData(no_msn)
}
func (cS *customerMtrService) AmbilDataBalikan(no_msn string, kd_user string) error {
	return cS.cR.AmbilDataBalikan(no_msn, kd_user)
}
func (cS *customerMtrService) AmbilDataAllStatus(no_msn string, kd_user string) error {
	return cS.cR.AmbilDataAllStatus(no_msn, kd_user)
}
func (cS *customerMtrService) AmbilDataBalikanKonfirmer(no_msn string, kd_user string) error {
	return cS.cR.AmbilDataAllStatus(no_msn, kd_user)
}
func (cS *customerMtrService) AmbilData(no_msn string, kd_user string) error {
	return cS.cR.AmbilData(no_msn, kd_user)
}

func (cS *customerMtrService) Show(no_msn string) response.TelesalesResponse{
	return cS.cR.Show(no_msn)
}
func (cS *customerMtrService) UpdateOkeMembership(customer request.CustomerMtr) (entity.CustomerMtr, error) {
	return cS.cR.UpdateOkeMembership(customer)
}

func (cS *customerMtrService) RekapTele(username string, startDate time.Time, endDate time.Time) (response.RekapTele, error) {
	// Jika startDate atau endDate kosong, set ke tanggal hari ini dengan waktu awal & akhir
	now := time.Now()
	if startDate.IsZero() {
		startDate = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()) // 00:00:00
	}
	if endDate.IsZero() {
		endDate = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, now.Location()) // 23:59:59
	}

	// Memanggil repository untuk mendapatkan data rekap
	rekap, err := cS.cR.RekapTele(username, startDate, endDate)
	if err != nil {
		return response.RekapTele{}, err
	}

	return rekap, nil
}

func (cS *customerMtrService) ListBerminatMembership(username string, startDate time.Time, endDate time.Time, limit int, pageParams int, search string) ([]response.MinatMembership, int, int, error) {
	// Memanggil repository untuk mengambil data
	data, totalPages, totalRecords, err := cS.cR.ListBerminatMembership(username, startDate, endDate, limit, pageParams, search)
	if err != nil {
		return nil, 0, 0, err
	}

	return data, totalPages, totalRecords, nil
}

func (cS *customerMtrService) ListDataAsuransiPA(username string, startDate time.Time, endDate time.Time, limit int, pageParams int, search string) ([]response.ListAsuransi, int, int, error) {
	// Memanggil repository untuk mengambil data
	data, totalPages, totalRecords, err := cS.cR.ListDataAsuransiPA(username, startDate, endDate, limit, pageParams, search)
	if err != nil {
		return nil, 0, 0, err
	}

	return data, totalPages, totalRecords, nil
}

func (cS *customerMtrService) ListDataAsuransiMtr(username string, startDate time.Time, endDate time.Time, limit int, pageParams int, search string) ([]response.ListAsuransi, int, int, error) {
	// Memanggil repository untuk mengambil data
	data, totalPages, totalRecords, err := cS.cR.ListDataAsuransiMtr(username, startDate, endDate, limit, pageParams, search)
	if err != nil {
		return nil, 0, 0, err
	}

	return data, totalPages, totalRecords, nil
}

func (cS *customerMtrService) ConsumeFonnte(body request.OtpCheck) (map[string]interface{}, error) {
	var client = &http.Client{}
	var data map[string]any
	var param = url.Values{}
	param.Set("target", body.NoHp)
	param.Set("message", fmt.Sprintf("%s%d", "Berikut kode OTP ", body.Otp))
	param.Set("schedule", "0")
	param.Set("delay", "2")
	param.Set("countryCode", "62")
	var payload = bytes.NewBufferString(param.Encode())
	request, err := http.NewRequest("POST", "https://api.fonnte.com/send", payload)
	if err != nil {
		return map[string]any{}, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", "k!ph_r+apphR8kJY@+gS")
	response, err := client.Do(request)
	if err != nil {
		return map[string]any{}, err
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return map[string]any{}, err
	}
	return data, nil
}


func (cS *customerMtrService) ExportRekapTele(username string, startDate time.Time, endDate time.Time) (string, error) {
	data, totalPages, totalRecords, err := cS.cR.ListBerminatMembership(username, startDate, endDate, -1, 1, "")
	if err != nil {
		return "", fmt.Errorf("failed to fetch data: %v", err)
	}

	file := excelize.NewFile()
	sheetName1 := "Membership"
	file.SetSheetName("Sheet1", sheetName1)

	// Header setup for Membership
	headers1 := []string{"No Mesin", "Nama Customer", "Jenis Bayar", "Tanggal Bayar Renewal", "Status", "Jenis Kartu"}
	for colIdx, header := range headers1 {
		cell := fmt.Sprintf("%c1", 'A'+colIdx)
		file.SetCellValue(sheetName1, cell, header)
		file.SetCellStyle(sheetName1, cell, cell, setHeaderStyle(file))
	}

	// Fill data for Membership
	for rowIdx, record := range data {
		stsKartuInt, err := strconv.Atoi(record.StsKartu)
		if err != nil {
			stsKartuInt = 0 // Set default jika error
		}

		stsJnsBayarStr := "Tidak ada data"
		if record.StsJnsBayar == "C" {
			stsJnsBayarStr = "Cash"
		} else if record.StsJnsBayar == "T" {
			stsJnsBayarStr = "Transfer"
		}

		status := determineStatus(int(record.Print), record.StsRenewal, stsKartuInt, record.StsBayarRenewal)

		file.SetCellValue(sheetName1, fmt.Sprintf("A%d", rowIdx+2), record.NoMsn)
		file.SetCellValue(sheetName1, fmt.Sprintf("B%d", rowIdx+2), record.NmCustomer)
		file.SetCellValue(sheetName1, fmt.Sprintf("C%d", rowIdx+2), stsJnsBayarStr)
		file.SetCellValue(sheetName1, fmt.Sprintf("D%d", rowIdx+2), record.TglBayarRenewal.Format("02-January-2006"))
		file.SetCellValue(sheetName1, fmt.Sprintf("E%d", rowIdx+2), status)
		file.SetCellValue(sheetName1, fmt.Sprintf("F%d", rowIdx+2), record.KdCard)
	}

	// Fetch data for Insurance PA
	insuranceDataPA, totalPages, totalRecords, err := cS.cR.ListDataAsuransiPA(username, startDate, endDate, -1, 1, "")
	if err != nil {
		return "", fmt.Errorf("failed to fetch insurance PA data: %v", err)
	}

	sheetName2 := "Asuransi PA"
	file.NewSheet(sheetName2)

	// Header setup for Insurance PA
	headers2 := []string{"No Mesin", "Nama Customer", "Status Asuransi", "Tanggal Beli", "Produk"}
	for colIdx, header := range headers2 {
		cell := fmt.Sprintf("%c1", 'A'+colIdx)
		file.SetCellValue(sheetName2, cell, header)
		file.SetCellStyle(sheetName2, cell, cell, setHeaderStyle(file))
	}

	// Fill data for Insurance PA
	for rowIdx, record := range insuranceDataPA {
		namaCustomer := record.NmCustomerWkm
		if namaCustomer == "" {
			namaCustomer = record.NmCustomerFkt
		}
		tglBeliStr := ""
		if record.TglBeli != nil {
			tglBeliStr = record.TglBeli.Format("02-January-2006")
		}

		idProdukStr := ""
		if record.IdProduk != "" {
			idProdukStr = record.IdProduk
		}

		stsAsuransiStr := "Tidak ada data"
		switch record.StsAsuransi {
		case "P":
			stsAsuransiStr = "Pending"
		case "T":
			stsAsuransiStr = "Tidak Minat"
		case "F":
			stsAsuransiStr = "Prospect"
		case "O":
			stsAsuransiStr = "Oke"
		}

		file.SetCellValue(sheetName2, fmt.Sprintf("A%d", rowIdx+2), record.NoMsn)
		file.SetCellValue(sheetName2, fmt.Sprintf("B%d", rowIdx+2), namaCustomer)
		file.SetCellValue(sheetName2, fmt.Sprintf("C%d", rowIdx+2), stsAsuransiStr)
		file.SetCellValue(sheetName2, fmt.Sprintf("D%d", rowIdx+2), tglBeliStr)
		file.SetCellValue(sheetName2, fmt.Sprintf("E%d", rowIdx+2), idProdukStr)
	}

	// Fetch data for Insurance MTR
	insuranceDataMTR, totalPages, totalRecords, err := cS.cR.ListDataAsuransiMtr(username, startDate, endDate, -1, 1, "")
	if err != nil {
		return "", fmt.Errorf("failed to fetch insurance MTR data: %v", err)
	}

	sheetName3 := "Asuransi MTR"
	file.NewSheet(sheetName3)

	// Header setup for Insurance MTR
	headers3 := []string{"No Mesin", "Nama Customer", "Status Asuransi", "Tanggal Beli", "Produk"}
	for colIdx, header := range headers3 {
		cell := fmt.Sprintf("%c1", 'A'+colIdx)
		file.SetCellValue(sheetName3, cell, header)
		file.SetCellStyle(sheetName3, cell, cell, setHeaderStyle(file))
	}

	// Fill data for Insurance MTR
	for rowIdx, record := range insuranceDataMTR {
		namaCustomer := record.NmCustomerWkm
		if namaCustomer == "" {
			namaCustomer = record.NmCustomerFkt
		}
		tglBeliStr := ""
		if record.TglBeli != nil {
			tglBeliStr = record.TglBeli.Format("02-January-2006")
		}

		idProdukStr := ""
		if record.IdProduk != "" {
			idProdukStr = record.IdProduk
		}

		stsAsuransiStr := "Tidak ada data"
		switch record.StsAsuransi {
		case "P":
			stsAsuransiStr = "Pending"
		case "T":
			stsAsuransiStr = "Tidak Minat"
		case "F":
			stsAsuransiStr = "Prospect"
		case "O":
			stsAsuransiStr = "Oke"
		}

		file.SetCellValue(sheetName3, fmt.Sprintf("A%d", rowIdx+2), record.NoMsn)
		file.SetCellValue(sheetName3, fmt.Sprintf("B%d", rowIdx+2), namaCustomer)
		file.SetCellValue(sheetName3, fmt.Sprintf("C%d", rowIdx+2), stsAsuransiStr)
		file.SetCellValue(sheetName3, fmt.Sprintf("D%d", rowIdx+2), tglBeliStr)
		file.SetCellValue(sheetName3, fmt.Sprintf("E%d", rowIdx+2), idProdukStr)
	}

	// Save file
	fileName := fmt.Sprintf("Export_Rekap_Tele.xlsx")
	if err := file.SaveAs(fileName); err != nil {
		return "", fmt.Errorf("failed to save Excel file: %v", err)
	}
	if totalPages == 0 {
		return fileName, nil
	}
	if totalRecords == 0 {
		return fileName, nil
	}
	return fileName, nil
}

func setHeaderStyle(f *excelize.File) int {
	style, _ := f.NewStyle(`{
		"font": {"bold": true},
		"fill": {"type": "pattern", "color": ["#FFFF00"], "pattern": 1},
		"alignment": {"horizontal": "center"}
	}`)
	return style
}

func determineStatus(print int, stsRenewal string, stsKartu int, stsBayarRenewal string) string {
	if print == 0 && stsRenewal == "O" {
		return "Belum Print TT"
	}
	switch stsKartu {
	case 1:
		return "Print TT"
	case 2:
		return "Bawa Kurir"
	case 3:
		if stsRenewal == "O" && stsBayarRenewal == "S" {
			return "Sukses"
		}
	case 4:
		return "Check Kartu"
	case 6:
		return "Kartu Kembali TS"
	}
	return ""
}

func (cS *customerMtrService) RekapLeaderTs(startDate time.Time, endDate time.Time) (response.RekapLeaderTs, error) {
	// Jika startDate atau endDate kosong, set ke tanggal hari ini dengan waktu awal & akhir
	now := time.Now()
	if startDate.IsZero() {
		startDate = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()) // 00:00:00
	}
	if endDate.IsZero() {
		endDate = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, now.Location()) // 23:59:59
	}

	// Memanggil repository untuk mendapatkan data rekap
	rekap, err := cS.cR.RekapLeaderTs(startDate, endDate)
	if err != nil {
		return response.RekapLeaderTs{}, err
	}

	return rekap, nil
}

func (cS *customerMtrService) RekapBerminatPerWilayah(startDate time.Time, endDate time.Time) ([]response.RekapBerminatPerWilayah, int, error) {
	// Jika startDate atau endDate kosong, set ke tanggal hari ini dengan waktu awal & akhir
	now := time.Now()
	if startDate.IsZero() {
		startDate = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()) // 00:00:00
	}
	if endDate.IsZero() {
		endDate = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, now.Location()) // 23:59:59
	}

	// Memanggil repository untuk mendapatkan data rekap
	rekap, totalData, err := cS.cR.RekapBerminatPerWilayah(startDate, endDate)
	if err != nil {
		return nil, 0, err
	}

	return rekap, totalData, nil
}

func (cS *customerMtrService) ExportRekapLeaderTs(startDate, endDate time.Time) (string, error) {
	now := time.Now()
	if startDate.IsZero() {
		startDate = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	}
	if endDate.IsZero() {
		endDate = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, now.Location())
	}

	rekap, err := cS.cR.RekapTransaksi(startDate, endDate)
	if err != nil {
		return "", err
	}

	rekapStatus, err := cS.cR.RekapStatus(startDate, endDate)
	if err != nil {
		return "", err
	}

	file := excelize.NewFile()
	headerStyle := setHeaderStyle(file)
	borderStyle := setBorderStyle(file)

	sheet1 := "Rekap 1"
	file.SetSheetName("Sheet1", sheet1)

	headers1 := []string{
		"Nama User", "Jumlah Data", "Renewal OK Cash", "Renewal OK Transfer", "Renewal OK Digital",
		"BASIC", "GOLD", "PLATINUM", "PLATINUMP",
	}

	for i, header := range headers1 {
		col := columnNumberToName(i+1) + "1"
		file.SetCellValue(sheet1, col, header)
		file.SetCellStyle(sheet1, col, col, headerStyle)
	}

	for idx, row := range rekap {
		rowIndex := idx + 2
		values := []interface{}{
			row.NamaUser, row.JmlData, row.RenewalOkCash, row.RenewalOkTransfer,
			row.RenewalOkDigital, row.Basic, row.Gold, row.Platinum, row.PlatinumP,
		}

		for i, val := range values {
			cell := fmt.Sprintf("%s%d", columnNumberToName(i+1), rowIndex)
			file.SetCellValue(sheet1, cell, val)
			file.SetCellStyle(sheet1, cell, cell, borderStyle)
		}
	}

	sheet2 := "Rekap Status"
	file.NewSheet(sheet2)

	headers2 := []string{
		"Kd User", "Jumlah Data", "Sudah Terima", "Belum Terima",
		"Renewal OK Cash Update", "Renewal OK Cash", "Renewal OK Transfer",
		"Pikir-Pikir", "Telp Kembali", "Tidak Diangkat", "Belum Registrasi",
		"Prospek", "Basic", "Gold", "Platinum",
	}

	for i := 1; i <= 24; i++ {
		headers2 = append(headers2, fmt.Sprintf("%d", i))
	}

	for i, header := range headers2 {
		col := columnNumberToName(i+1) + "1"
		file.SetCellValue(sheet2, col, header)
		file.SetCellStyle(sheet2, col, col, headerStyle)
	}

	for idx, row := range rekapStatus {
		rowIndex := idx + 2
		values := []interface{}{
			row.KdUser, row.JmlData, row.SudahTerima, row.BelumTerima,
			row.RenewalOkCashUpdate, row.RenewalOkCash, row.RenewalOkTransfer,
			row.PikirRagu, row.TelpKembali, row.TidakDiangkat, row.BelumRegist,
			row.Prospek, row.Basic, row.Gold, row.Platinum,
		}

		for i, val := range values {
			cell := fmt.Sprintf("%s%d", columnNumberToName(i+1), rowIndex)
			file.SetCellValue(sheet2, cell, val)
			file.SetCellStyle(sheet2, cell, cell, borderStyle)
		}

		// Alasan tidak renewal
		for i := 1; i <= 24; i++ {
			alasanKey := fmt.Sprintf("%d", i)
			colName := columnNumberToName(15 + i)
			cell := fmt.Sprintf("%s%d", colName, rowIndex)

			if count, exists := row.AlasanTidakRenewal[alasanKey]; exists {
				file.SetCellValue(sheet2, cell, count)
			} else {
				file.SetCellValue(sheet2, cell, 0)
			}
			file.SetCellStyle(sheet2, cell, cell, borderStyle)
		}
	}

	fileName := "Export_Rekap_Leader_TS.xlsx"
	if err := file.SaveAs(fileName); err != nil {
		return "", fmt.Errorf("failed to save Excel file: %v", err)
	}

	return fileName, nil
}


func setBorderStyle(f *excelize.File) int {
	style, _ := f.NewStyle(`{
		"border": [
			{"type": "left", "color": "#000000", "style": 1},
			{"type": "right", "color": "#000000", "style": 1},
			{"type": "top", "color": "#000000", "style": 1},
			{"type": "bottom", "color": "#000000", "style": 1}
		]
	}`)
	return style
}

func columnNumberToName(n int) string {
	result := ""
	for n > 0 {
		n--
		result = string(rune('A'+(n%26))) + result
		n /= 26
	}
	return result
}
func (cS *customerMtrService) ListPerformanceTs(startDate time.Time, endDate time.Time) (response.PerformanceTs, error) {
	// Jika startDate atau endDate kosong, set ke tanggal hari ini dengan waktu awal & akhir
	now := time.Now()
	if startDate.IsZero() {
		startDate = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()) // 00:00:00
	}
	if endDate.IsZero() {
		endDate = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, now.Location()) // 23:59:59
	}

	// Pastikan startDate tidak lebih besar dari endDate
	if startDate.After(endDate) {
		return response.PerformanceTs{}, fmt.Errorf("startDate tidak boleh lebih besar dari endDate")
	}

	// Panggil repository untuk mendapatkan top 5 dan low 5
	top5, low5, err := cS.cR.ListPerformanceTs(startDate, endDate)
	if err != nil {
		return response.PerformanceTs{}, err
	}

	// Return hasil response
	return response.PerformanceTs{
		TopUsers:       top5,
		LowPerformance: low5,
	}, nil
}

func (cS *customerMtrService) GetRekapStatus(startDate time.Time, endDate time.Time) ([]response.RekapStatus, error) {
	return cS.cR.RekapStatus(startDate, endDate)
}

func (cS *customerMtrService) ListDataPerKecamatan(startDate time.Time, endDate time.Time, limit int, pageParams int, search string) ([]response.RekapBerminatPerWilayah, int, int,int, error) {
	// Memanggil repository untuk mengambil data
	data, totalPages, totalRecords,totalRowsPerPage, err := cS.cR.RekapBerminatPerKecamatan(startDate, endDate, limit, pageParams, search)
	if err != nil {
		return nil, 0, 0,0, err
	}

	return data, totalPages, totalRecords,totalRowsPerPage, nil
}