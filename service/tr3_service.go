package service

import (
	"fmt"
	"math"
	"strings"
	"time"
	"wkm/entity"
	"wkm/repository"
	"wkm/request"
	"wkm/response"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type Tr3Service interface {
	DataWABlast(t request.DataWaBlastRequest) []entity.DataWaBlast
	SearchNoMsnByWa(t request.SearchNoMsnByWaRequest) []entity.SearchNoMsnByWa
	UpdateJenisBayar(data []repository.ParamsUpdateJenisBayar, payment_type string, username string)
	WillBayar(data request.SearchWBRequest) (entity.Faktur3, error)
	UpdateInputBayar(data request.InputBayarRequest) (entity.Faktur3, error)
	DataRenewal(data request.DataRenewalRequest) ([]response.DataRenewalResponse, error)
	ExportPembayaranRenewal(data request.RangeTanggalRequest) error
	ExportDataRenewal(data request.DataRenewalRequest) (entity.DataRenewal, error)
	ExportDataPlatinumPlus(data request.DataRenewalRequest) (entity.DataRenewal, error)
}

type tr3Service struct {
	trR repository.Tr3Repository
}

func NewTr3Service(tR repository.Tr3Repository) Tr3Service {
	return &tr3Service{
		trR: tR,
	}
}

func GetFormattedAddress(record entity.DataRenewal) string {
	var addressParts []string

	// Check if Alamat is not empty
	if record.Alamat != "" {
		addressParts = append(addressParts, record.Alamat)
	} else if record.Alamat11 != nil && *record.Alamat11 != "" {
		addressParts = append(addressParts, *record.Alamat11)
	}

	// Check RT
	rtText := ""
	if record.Rt != "" {
		rtText = fmt.Sprintf("RT %s", record.Rt)
	} else if record.Rt1 != nil && *record.Rt1 != "" {
		rtText = fmt.Sprintf("RT %s", *record.Rt1)
	}
	if rtText != "" {
		addressParts = append(addressParts, rtText)
	}

	// Check RW
	rwText := ""
	if record.Rw != "" {
		rwText = fmt.Sprintf("RW %s", record.Rw)
	} else if record.Rw1 != nil && *record.Rw1 != "" {
		rwText = fmt.Sprintf("RW %s", *record.Rw1)
	}
	if rwText != "" {
		addressParts = append(addressParts, rwText)
	}

	// Add Kelurahan, Kecamatan, Kota, and KodePos
	if record.Kel != "" {
		addressParts = append(addressParts, record.Kel)
	} else if record.Kel1 != nil && *record.Kel1 != "" {
		addressParts = append(addressParts, *record.Kel1)
	}

	// Add Kecamatan
	if record.Kec != "" {
		addressParts = append(addressParts, record.Kec)
	} else if record.Kec1 != nil && *record.Kec1 != "" {
		addressParts = append(addressParts, *record.Kec1)
	}

	// Add Kota
	if record.Kota != nil && *record.Kota != "" {
		addressParts = append(addressParts, *record.Kota)
	} else if record.Kota1 != nil && *record.Kota1 != "" {
		addressParts = append(addressParts, *record.Kota1)
	}

	// Add Kode Pos
	if record.Kodepos != "" {
		addressParts = append(addressParts, record.Kodepos)
	} else if record.Kodepos1 != nil && *record.Kodepos1 != "" {
		addressParts = append(addressParts, *record.Kodepos1)
	}

	// Join all parts of the address
	formattedAddress := strings.Join(addressParts, " ")
	return formattedAddress
}

func (s *tr3Service) ExportPembayaranRenewal(data request.RangeTanggalRequest) error {
	dataPembayaran := s.trR.DataPembayaran(data.Tgl1, data.Tgl2)
	// Create a new Excel file
	xlsx := excelize.NewFile()

	// Define the sheet name
	sheetName := "REKAP"
	xlsx.NewSheet(sheetName)

	// Remove the default "Sheet1"
	xlsx.DeleteSheet("Sheet1")
	tgl1, _ := time.Parse("2006-01-02", data.Tgl1)
	tgl2, _ := time.Parse("2006-01-02", data.Tgl2)

	headerStyle, _ := xlsx.NewStyle(`{
		"font": {
			"bold": true
		},
		"alignment": {
	        "horizontal": "center",
	        "vertical": "center",
			"wrap_text": true
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
	borderStyle, _ := xlsx.NewStyle(`{
		"alignment": {
	        "horizontal": "center",
	        "vertical": "center"
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

	// Set headers for the sheet, including the new columns
	xlsx.SetCellValue(sheetName, "A1", "Laporan Pembayaran Renewal All "+tgl1.Format("02-Jan-2006")+" sd "+tgl2.Format("02-Jan-2006"))
	headers := []string{"NO", "NAMA MEMBER", "TANGGAL BAYAR", "CARD", "JENIS BAYAR", "KD USER", "CONFIRMER", "KURIR", "NO TANDA TERIMA", "ASURANSI", "ASURANSI MOTOR", "HARGA POKOK", "DPP", "PPN"}
	for i, header := range headers {
		xlsx.SetCellValue(sheetName, fmt.Sprintf("%s2", string('A'+i)), header)
	}

	startRow := 3
	for i, record := range dataPembayaran {
		dpp := float64(record.MstCard.HargaPokok) / 1.11
		xlsx.SetCellValue(sheetName, fmt.Sprintf("A%d", i+3), i+1)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("B%d", i+3), record.NmCustomer)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("C%d", i+3), record.TglBayarRenewalFin.Format("02-Jan-2006"))
		xlsx.SetCellValue(sheetName, fmt.Sprintf("D%d", i+3), record.MstCard.JnsCard)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("E%d", i+3), record.StsJnsBayar)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("F%d", i+3), record.User.FullName)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("G%d", i+3), record.User10.FullName)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("H%d", i+3), record.Kurir.NmKurir)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("I%d", i+3), record.NoTandaTerima)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("J%d", i+3), record.MstCard.Asuransi)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("K%d", i+3), record.MstCard.AsuransiMotor)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("L%d", i+3), record.MstCard.HargaPokok)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("M%d", i+3), math.Round(dpp))
		xlsx.SetCellValue(sheetName, fmt.Sprintf("N%d", i+3), math.Round(dpp*0.11))
		startRow += 1
	}

	xlsx.MergeCell(sheetName, "A1", "N1")
	xlsx.SetCellStyle(sheetName, "A1", "N1", headerStyle)
	xlsx.SetCellStyle(sheetName, "A2", "N2", headerStyle)
	xlsx.SetCellStyle(sheetName, fmt.Sprintf("A%d", 3), fmt.Sprintf("N%d", startRow-1), borderStyle)

	if err := xlsx.SaveAs("./pembayaran-renewal.xlsx"); err != nil {
		return err
	}

	return nil
}

func (s *tr3Service) ExportDataPlatinumPlus(data request.DataRenewalRequest) (entity.DataRenewal, error) {
	// Create a new Excel file
	xlsx := excelize.NewFile()

	// Define the sheet name
	sheetName := "PLATINUM PLUS"
	xlsx.NewSheet(sheetName)

	// Remove the default "Sheet1"
	xlsx.DeleteSheet("Sheet1")

	// Set headers for the sheet, including the new columns
	headers := []string{"KD DLR", "NAMA DEALER", "NO MEMBERSHIP", "NO. RANGKA", "MERK", "TYPE", "NO MESIN", "NAMA STNK", "NAMA KARTU", "NAMA TERTANGGUNG", "JNS KARTU", "TGL MOHON", "TGL AWAL", "TGL AKHIR", "ALAMAT"}
	for i, header := range headers {
		xlsx.SetCellValue(sheetName, fmt.Sprintf("%s1", string('A'+i)), header)
	}

	headerStyle, err := xlsx.NewStyle(`{
		"font": {
			"bold": true
		}
	}`)
	if err != nil {
		return entity.DataRenewal{}, err
	}

	// Apply the bold style to the header row
	for i := 0; i < len(headers); i++ {
		xlsx.SetCellStyle(sheetName, fmt.Sprintf("%s1", string('A'+i)), fmt.Sprintf("%s1", string('A'+i)), headerStyle)
	}

	// Fetch data for Platinum Plus
	platinumPlusData, err := s.trR.ExportDataAsuransiPlatinumPlus(data)
	if err != nil {
		return entity.DataRenewal{}, err
	}

	// Write data to the sheet
	for i, record := range platinumPlusData {

		var kdDlrValue string
		if record.KdDlr != nil {
			kdDlrValue = *record.KdDlr // Dereference the pointer
		} else {
			kdDlrValue = "" // Handle nil case
		}

		var tglMohonValue string
		if record.TglMohon != nil {
			tglMohonValue = record.TglMohon.Format("2006-01-02") // Format time.Time to string
		} else {
			tglMohonValue = "" // Handle nil case
		}

		namaKartu := record.NamaKtp
		if namaKartu == "" {
			namaKartu = record.NmCustomer
		}
		var NmMtrValue string
		if record.NmMtr != nil {
			NmMtrValue = *record.NmMtr // Dereference the pointer
		} else {
			NmMtrValue = "" // Handle nil case
		}
		namaTertanggung := record.NamaKtp
		if namaTertanggung == "" {
			namaTertanggung = record.NmCustomer
		}
		formattedAddress := GetFormattedAddress(record)

		// Set values for each column, including the new ones
		xlsx.SetCellValue(sheetName, fmt.Sprintf("A%d", i+2), kdDlrValue)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("B%d", i+2), record.NmDlr)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("C%d", i+2), record.NoKartu)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("D%d", i+2), record.NoRgk) // NO. RANGKA (manual input can be added later)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("E%d", i+2), "HONDA")      // MERK
		xlsx.SetCellValue(sheetName, fmt.Sprintf("F%d", i+2), NmMtrValue)   // TYPE
		xlsx.SetCellValue(sheetName, fmt.Sprintf("G%d", i+2), record.NoMsn)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("H%d", i+2), record.NmCustomer)                    // NAMA STNK
		xlsx.SetCellValue(sheetName, fmt.Sprintf("I%d", i+2), namaKartu)                            // NAMA KARTU
		xlsx.SetCellValue(sheetName, fmt.Sprintf("J%d", i+2), namaTertanggung)                      // NAMA TERTANGGUNG
		xlsx.SetCellValue(sheetName, fmt.Sprintf("K%d", i+2), record.JnsCard)                       // JNS KARTU
		xlsx.SetCellValue(sheetName, fmt.Sprintf("L%d", i+2), tglMohonValue)                        // TGL MOHON
		xlsx.SetCellValue(sheetName, fmt.Sprintf("M%d", i+2), record.TglAwal.Format("2006-01-02"))  // TGL AWAL
		xlsx.SetCellValue(sheetName, fmt.Sprintf("N%d", i+2), record.TglAkhir.Format("2006-01-02")) // TGL AKHIR
		xlsx.SetCellValue(sheetName, fmt.Sprintf("O%d", i+2), formattedAddress)                     // ALAMAT
	}

	// Save the Excel file
	if err := xlsx.SaveAs("./Data_Platinum_Plus.xlsx"); err != nil {
		return entity.DataRenewal{}, err
	}

	return entity.DataRenewal{}, nil
}

func (s *tr3Service) ExportDataRenewal(data request.DataRenewalRequest) (entity.DataRenewal, error) {
	// Create a new Excel file
	xlsx := excelize.NewFile()

	// Define the sheet names
	sheets := []string{"BASIC", "GOLD", "PLATINUM", "PLATINUM PLUS"}
	for _, sheet := range sheets {
		xlsx.NewSheet(sheet)
	}

	xlsx.DeleteSheet("Sheet1")

	// Set headers for each sheet
	headers := []string{"KD DLR", "NAMA DEALER", "NO MESIN", "NO KARTU", "NAMA STNK", "NAMA KARTU", "NAMA TERTANGGUNG", "JNS KARTU", "TGL MOHON", "TGL AWAL", "TGL AKHIR", "ALAMAT"}
	for _, sheet := range sheets {
		for i, header := range headers {
			xlsx.SetCellValue(sheet, fmt.Sprintf("%s1", string('A'+i)), header)
		}
	}

	// Fetch data for each JnsCard type
	basicData, err := s.trR.ExportDataRenewalBasic(data)
	if err != nil {
		return entity.DataRenewal{}, err
	}
	if err := writeDataRenewalToSheet(xlsx, "BASIC", basicData, headers); err != nil {
		return entity.DataRenewal{}, err
	}

	goldData, err := s.trR.ExportDataRenewalGold(data)
	if err != nil {
		return entity.DataRenewal{}, err
	}
	if err := writeDataRenewalToSheet(xlsx, "GOLD", goldData, headers); err != nil {
		return entity.DataRenewal{}, err
	}

	platinumData, err := s.trR.ExportDataRenewalPlatinum(data)
	if err != nil {
		return entity.DataRenewal{}, err
	}
	if err := writeDataRenewalToSheet(xlsx, "PLATINUM", platinumData, headers); err != nil {
		return entity.DataRenewal{}, err
	}

	platinumPlusData, err := s.trR.ExportDataRenewalPlatinumPlus(data)
	if err != nil {
		return entity.DataRenewal{}, err
	}
	if err := writeDataRenewalToSheet(xlsx, "PLATINUM PLUS", platinumPlusData, headers); err != nil {
		return entity.DataRenewal{}, err
	}

	// Save the Excel file
	if err := xlsx.SaveAs("./Data_Renewal.xlsx"); err != nil {
		return entity.DataRenewal{}, err
	}

	return entity.DataRenewal{}, nil
}

// Helper function to write data to the specified sheet
func writeDataRenewalToSheet(xlsx *excelize.File, sheetName string, data []entity.DataRenewal, headers []string) error {
	// Define a style for bold text
	headerStyle, err := xlsx.NewStyle(`{
		"font": {
			"bold": true
		}
	}`)
	if err != nil {
		return err
	}

	// Write data to the specified sheet
	for i, record := range data {

		var kdDlrValue string
		if record.KdDlr != nil {
			kdDlrValue = *record.KdDlr // Dereference the pointer
		} else {
			kdDlrValue = "" // Handle nil case
		}
		var tglMohonValue string
		if record.TglMohon != nil {
			tglMohonValue = record.TglMohon.Format("2006-01-02") // Format time.Time to string
		} else {
			tglMohonValue = "" // Handle nil case
		}
		namaKartu := record.NamaKtp
		if namaKartu == "" {
			namaKartu = record.NmCustomer
		}

		namaTertanggung := record.NamaKtp
		if namaTertanggung == "" {
			namaTertanggung = record.NmCustomer
		}

		formattedAddress := GetFormattedAddress(record)

		xlsx.SetCellValue(sheetName, fmt.Sprintf("A%d", i+2), kdDlrValue)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("B%d", i+2), record.NmDlr)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("C%d", i+2), record.NoMsn)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("D%d", i+2), record.NoKartu)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("E%d", i+2), record.NmCustomer)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("F%d", i+2), namaKartu)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("G%d", i+2), namaTertanggung)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("H%d", i+2), record.JnsCard)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("I%d", i+2), tglMohonValue)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("J%d", i+2), record.TglAwal.Format("2006-01-02"))
		xlsx.SetCellValue(sheetName, fmt.Sprintf("K%d", i+2), record.TglAkhir.Format("2006-01-02"))
		xlsx.SetCellValue(sheetName, fmt.Sprintf("L%d", i+2), formattedAddress)
	}

	validCount := 0
	for _, record := range data {
		if record.JnsCard != "" { // Pastikan NoMsn tidak kosong
			validCount++
		}
	}

	// Add the card type and count below the data
	startRow := len(data) + 3 // 2 rows below the last data row
	xlsx.SetCellValue(sheetName, fmt.Sprintf("H%d", startRow), "Jenis Kartu")
	xlsx.SetCellValue(sheetName, fmt.Sprintf("I%d", startRow), sheetName) // Assuming all records in this sheet have the same JnsCard

	// Set the count of records in the next row
	xlsx.SetCellValue(sheetName, fmt.Sprintf("H%d", startRow+1), "Jumlah Data")
	xlsx.SetCellValue(sheetName, fmt.Sprintf("I%d", startRow+1), validCount)

	// Apply bold style to headers
	for i := 0; i < len(headers); i++ {
		xlsx.SetCellStyle(sheetName, fmt.Sprintf("%s1", string('A'+i)), fmt.Sprintf("%s1", string('A'+i)), headerStyle)
	}

	// Apply bold style to "Jumlah Data" cell
	xlsx.SetCellStyle(sheetName, fmt.Sprintf("H%d", startRow), fmt.Sprintf("I%d", startRow), headerStyle)
	xlsx.SetCellStyle(sheetName, fmt.Sprintf("H%d", startRow+1), fmt.Sprintf("I%d", startRow+1), headerStyle)

	return nil
}

func (s *tr3Service) DataRenewal(t request.DataRenewalRequest) ([]response.DataRenewalResponse, error) {
	return s.trR.DataRenewalRequest(t)
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
