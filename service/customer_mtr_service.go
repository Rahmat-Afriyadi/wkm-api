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
	MasterData(search string, sts string, jns string, username string, limit int, pageParams int) []entity.CustomerMtr
	MasterDataCount(search string, sts string, jns string, username string) int64
	MasterDataBalikan(search string, username string, limit int, pageParams int) []response.TelesalesBalikanResponseList
	MasterDataBalikanCount(search string, username string) int64
	ListAmbilData() []entity.Faktur3
	AmbilDataBalikan(no_msn string, kd_user string) error
	AmbilData(no_msn string, kd_user string) error
	SelfCount(kd_user string) int64
	Show(no_msn string) response.TelesalesResponse
	UpdateOkeMembership(customer request.CustomerMtr) (entity.CustomerMtr, error)
	RekapTele(username string, startDate time.Time, endDate time.Time) (response.RekapTele, error)
	ListBerminatMembership(usrname string, startDate time.Time, endDate time.Time, limit int, pageParams int, search string) ([]response.MinatMembership, int, int, error)
	ListDataAsuransiPA(usrname string, startDate time.Time, endDate time.Time, limit int, pageParams int, search string) ([]response.ListAsuransi, int, int, error)
	ListDataAsuransiMtr(usrname string, startDate time.Time, endDate time.Time, limit int, pageParams int, search string) ([]response.ListAsuransi, int, int, error)
	ExportRekapTele(usrname string, startDate time.Time, endDate time.Time) (string, error)
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
func (cS *customerMtrService) MasterDataBalikan(search string, username string, limit int, pageParams int) []response.TelesalesBalikanResponseList {
	return cS.cR.MasterDataBalikan(search, username, limit, pageParams)
}
func (cS *customerMtrService) MasterDataBalikanCount(search string, username string) int64 {
	return cS.cR.MasterDataBalikanCount(search, username)
}
func (cS *customerMtrService) MasterData(search string, sts string, jns string, username string, limit int, pageParams int) []entity.CustomerMtr {
	return cS.cR.MasterData(search, sts, jns, username, limit, pageParams)
}
func (cS *customerMtrService) MasterDataCount(search string, sts string, jns string, username string) int64 {
	return cS.cR.MasterDataCount(search, sts, jns, username)
}

func (cS *customerMtrService) ListAmbilData() []entity.Faktur3 {
	return cS.cR.ListAmbilData()
}

func (cS *customerMtrService) AmbilDataBalikan(no_msn string, kd_user string) error {
	return cS.cR.AmbilDataBalikan(no_msn, kd_user)
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
