package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
	"wkm/entity"
	"wkm/request"

	"gorm.io/gorm"
)

type ParamsUpdateJenisBayar struct {
	NoTandaTerima string
	NamaCustomer  string
}

type Tr3Repository interface {
	DataWABlast(request request.DataWaBlastRequest) []entity.DataWaBlast
	SearchNoMsnByWa(request request.SearchNoMsnByWaRequest) []entity.SearchNoMsnByWa
	UpdateJenisBayar(data []ParamsUpdateJenisBayar, payment_type string, username string)
	UpdateTglAkhirTenor()
	WillBayar(data request.SearchWBRequest) (entity.Faktur3, error)
	UpdateInputBayar(data request.InputBayarRequest) (entity.Faktur3, error)
}

type tr3Repository struct {
	conn     *sql.DB
	connGorm *gorm.DB
}

func NewTr3nRepository(conn *sql.DB, connGorm *gorm.DB) Tr3Repository {
	return &tr3Repository{
		conn:     conn,
		connGorm: connGorm,
	}
}

func (tr *tr3Repository) DataWABlast(request request.DataWaBlastRequest) []entity.DataWaBlast {
	// format (table, filterTypeKodeKerja1, filterTypeKodeKerja2)
	// params (awaltenor, akhirtenor, leas1, leas2, kodekerja1, kodekerja2)
	fmt.Println("ini request ", request)
	datas := []entity.DataWaBlast{}

	tables := []string{
		"tr_wms_faktur2",
		"tr_wms_faktur3",
		"tr_wms_faktur4",
	}
	resultChannel := make(chan *sql.Rows, 0)
	ctx := context.Background()
	kodeKerjaInit := strings.Repeat("?,", len(request.KodeKerja))
	for _, table := range tables {
		go func() {
			query := ""
			// if table != "tr_wms_faktur2" {
			query = fmt.Sprint("select * from (select no_msn, nm_customer11, kd_user, no_yg_dihub_renewal, case when trim(no_wa) REGEXP '^[+62]|^[0-9]*$' and no_wa is not null and no_wa not like '021%' then no_wa when trim(sms_no) REGEXP '^[+62]|^[0-9]*$' and sms_no is not null and sms_no not like '021%' then sms_no when trim(no_telp2) REGEXP '^[+62]|^[0-9]*$' and no_telp2 is not null and no_telp2 not like '021%' then no_telp2 when trim(no_telp1) REGEXP '^[+62]|^[0-9]*$' and no_telp1 is not null and no_telp1 not like '021%' then no_telp1 when trim(no_hp2) REGEXP '^[+62]|^[0-9]*$' and no_hp2 is not null and no_hp2 not like '021%' then no_hp2 when trim(no_hp1)REGEXP '^[+62]|^[0-9]*$' and no_hp1 is not null and no_hp1 not like '021%' then no_hp1 end as no_wa, tgl_akhir_tenor from (select no_msn,nm_customer11,kd_user,case when ket_no_telp1 in ('1','1A','1B') then no_telp1 else no_hp1 end as 'no_telp1', case when ket_no_telp2 in ('1','2') then no_telp2 else no_hp1 end as 'no_telp2',case when no_hp2 is not null and no_hp2 != '' then no_hp2 else no_hp1 end as 'no_hp2', case when ket_no_hp1 in ('1','1A','1B') then no_hp1 else no_hp1 end as 'no_hp1', no_yg_dihub_renewal,sms_no,no_wa,tgl_akhir_tenor from ", table, " where tgl_akhir_tenor>=? and tgl_akhir_tenor<=? and (no_leas =? or no_leas2=?) and (kode_kerja ", request.KodeKerjaFilterType, " (", kodeKerjaInit[:len(kodeKerjaInit)-1], ") or kode_kerja2 ", request.KodeKerjaFilterType, " (", kodeKerjaInit[:len(kodeKerjaInit)-1], "))) t) u where no_wa is not null and no_wa != ''")
			// } else {
			// 	query = fmt.Sprint("select * from (select no_msn, nm_customer11, kd_user, no_yg_dihub_renewal, case when trim(no_telp2) REGEXP '^[+62]|^[0-9]*$' and no_telp2 is not null and no_telp2 not like '021%' then no_telp2 when trim(no_telp1) REGEXP '^[+62]|^[0-9]*$' and no_telp1 is not null and no_telp1 not like '021%' then no_telp1 when trim(no_hp2) REGEXP '^[+62]|^[0-9]*$' and no_hp2 is not null and no_hp2 not like '021%' then no_hp2 when trim(no_hp1) REGEXP '^[+62]|^[0-9]*$' and no_hp1 is not null and no_hp1 not like '021%' then no_hp1 end as no_wa,tgl_akhir_tenor from ", table, " where tgl_akhir_tenor>=? and tgl_akhir_tenor<=? and (no_leas =? or no_leas2=?) and (kode_kerja ", request.KodeKerjaFilterType, " (", kodeKerjaInit[:len(kodeKerjaInit)-1], ") or kode_kerja2 ", request.KodeKerjaFilterType, " (", kodeKerjaInit[:len(kodeKerjaInit)-1], "))) a where no_wa is not null")
			// }
			fmt.Println("ini query yaa ", query)
			statement, err := tr.conn.PrepareContext(ctx, query)
			if err != nil {
				fmt.Println(err)
			}
			defer statement.Close()
			params := []interface{}{
				request.AwalTenor, request.AkhirTenor, request.NoLeas, request.NoLeas,
			}
			for i := 0; i < 2; i++ {
				for _, v := range request.KodeKerja {
					params = append(params, v)
				}
			}

			rows, err := statement.QueryContext(ctx, params...)
			if err != nil {
				fmt.Println("errornya di rows ", err)
				fmt.Println(err)
			}
			resultChannel <- rows
		}()
	}
	for i := 0; i < len(tables); i++ {
		rows := <-resultChannel
		defer rows.Close()
		for rows.Next() {
			iniKdUser := ""
			data := entity.DataWaBlast{}
			if err := rows.Scan(&data.NoMsn, &data.NmCustomer, &iniKdUser, &data.NoYgDiHubRenewal, &data.NoWa, &data.TglAkhirTenor); err != nil {
				fmt.Println("Error scanning row:", err)
				continue
			}
			data.KdUser = &iniKdUser
			// fmt.Println("ini user ", iniKdUser, data.KdUser)
			datas = append(datas, data)
		}
	}
	return datas
}

func (tr *tr3Repository) SearchNoMsnByWa(request request.SearchNoMsnByWaRequest) []entity.SearchNoMsnByWa {
	datas := []entity.SearchNoMsnByWa{}

	tables := []string{
		"tr_wms_faktur3",
		"tr_wms_faktur4",
	}
	resultChannel := make(chan *sql.Rows)
	ctx := context.Background()
	for _, table := range tables {
		go func() {
			query := fmt.Sprint("select * from (select no_msn, nm_customer11, case when  no_wa like ? then no_wa when sms_no like ? then sms_no when no_telp2 like ? then no_telp2 when no_telp1 like ? then no_telp1 when no_hp2 like ? then no_hp2 when no_hp1 like ? then no_hp1 end as no_wa from ", table, ") t where t.no_wa like ? limit 3")
			statement, err := tr.conn.PrepareContext(ctx, query)
			if err != nil {
				fmt.Println(err)
			}
			defer statement.Close()

			params := []interface{}{}
			paramsCount := 7
			for i := 0; i < paramsCount; i++ {
				params = append(params, "%"+request.NoHp+"%")
			}

			rows, err := statement.QueryContext(ctx, params...)
			if err != nil {
				fmt.Println("errornya di rows ", err)
				fmt.Println(err)
			}
			resultChannel <- rows
		}()
	}
	for i := 0; i < len(tables); i++ {
		rows := <-resultChannel
		defer rows.Close()
		for rows.Next() {
			data := entity.SearchNoMsnByWa{}
			if err := rows.Scan(&data.NoMsn, &data.NmCustomer, &data.NoWa); err != nil {
				fmt.Println("Error scanning row:", err)
				continue
			}
			datas = append(datas, data)
		}
	}
	return datas
}

func Log(content string) {
	fileName := "log.txt"

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content + "\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func LogBayar(content string) {
	fileName := "log/pembayaran/log.txt"

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content + "\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func (tr *tr3Repository) UpdateJenisBayar(data []ParamsUpdateJenisBayar, payment_type string, username string) {
	ctx := context.Background()
	now := time.Now()
	for _, v := range data {
		_, err := tr.conn.ExecContext(ctx, "UPDATE tr_wms_faktur3 set sts_jenis_bayar=? where no_tanda_terima=?", payment_type, v.NoTandaTerima)
		Log(now.Format("2006-01-02") + " " + v.NoTandaTerima + " " + v.NamaCustomer + " " + payment_type + " " + username)
		if err != nil {
			fmt.Println("errornya disin yaa ", err)
			continue
		}

	}
}

func (tr *tr3Repository) UpdateInputBayar(data request.InputBayarRequest) (entity.Faktur3, error) {
	faktur3 := entity.Faktur3{NoMsn: data.NoMsn}
	tr.connGorm.Find(&faktur3)
	if faktur3.NmCustomer == "" {
		return entity.Faktur3{}, errors.New("data tidak ditemukan")
	}
	kotaTrPembayaranRenewal := ""
	if faktur3.StsKirim == "1" {
		kotaTrPembayaranRenewal = faktur3.Kota
	} else if faktur3.StsKirim == "2" {
		kotaTrPembayaranRenewal = faktur3.KotaKtr
	} else if faktur3.StsKirim == "3" {
		kotaTrPembayaranRenewal = faktur3.KotaSrt1
	}
	trPembayaranRenewal := entity.TrPembayaranRenewal{
		NoMsn:               faktur3.NoMsn,
		RenewalKe:           faktur3.StsCetak3,
		NmCustomer:          faktur3.NmCustomer,
		Kota:                kotaTrPembayaranRenewal,
		CetakKe:             faktur3.Print,
		KirimKe:             faktur3.StsKirim,
		KdCard:              faktur3.KdCard,
		KdUserTs:            faktur3.KdUser,
		NoKartu:             faktur3.NoKartu,
		NoTandaTerima:       faktur3.NoTandaTerima,
		KdUserFa:            data.KdUserFa,
		TglCetakTandaTerima: faktur3.TglCetakTandaTerima,
		KdUserSS:            faktur3.KdUser2,
		JnsBayar:            faktur3.StsJnsBayar,
		TglBayar:            data.TglBayar,
		TglInsert:           time.Now(),
		TglJualan:           faktur3.TglVerifikasi,
	}
	now := time.Now()
	faktur3.TglBayarRenewalFin = &data.TglBayar
	faktur3.TglBayarRenewalFinKeyIn = &now
	faktur3.KdUser2 = data.KdUserFa
	faktur3.StsKartu = "3"
	faktur3.StsBawaKartu = "4"
	faktur3.StsAsuransiPa = "O"
	faktur3.StsBayarAsuransiPa = "S"
	faktur3.StsBayarRenewal = "S"

	if faktur3.StsJnsBayar == "C" {
		stockCard := entity.StockCard{NoKartu: faktur3.NoKartu}
		tr.connGorm.Find(&stockCard)
		faktur3.TglExpired = &stockCard.TglExpired
		stockCard.StsKartu = "3"
		stockCard.NoMsn = faktur3.NoMsn
		stockCard.TglUpdate = time.Now()
		stockCard.KdUser4 = data.KdUserFa
		result := tr.connGorm.Save(&stockCard)
		if result.Error != nil {
			return entity.Faktur3{}, result.Error
		}
	}

	result := tr.connGorm.Save(&trPembayaranRenewal)
	if result.Error != nil {
		return entity.Faktur3{}, result.Error
	}
	result = tr.connGorm.Save(&faktur3)
	if result.Error != nil {
		return entity.Faktur3{}, result.Error
	}

	LogBayar(fmt.Sprint(time.Now().Format("2006-01-02 15:04:05"), " - ", data.NoMsn, " - ", data.KdUserFa))

	return entity.Faktur3{}, nil
}

func (tr *tr3Repository) UpdateTglAkhirTenor() {
	ctx := context.Background()
	data := []string{"tr_wms_faktur2", "tr_wms_faktur3", "tr_wms_faktur4"}
	for _, v := range data {
		tr.conn.ExecContext(ctx, "update "+v+" set tgl_akhir_tenor= date_add(tgl_mohon, interval angsuran2 month) where tgl_akhir_tenor is null and angsuran2 not in ('','0','N')")
		_, err := tr.conn.ExecContext(ctx, "update "+v+" set tgl_akhir_tenor= date_add(tgl_mohon, interval angsuran month) where tgl_akhir_tenor is null and angsuran not in ('','0','N')")
		if err != nil {
			fmt.Println("errornya disin yaa ", err)
			continue
		}

	}

	// pisah

	// queries := []struct {
	// 	Query  string
	// 	Params []interface{}
	// }{
	// 	{"update tr_wms_faktur2 set tgl_akhir_tenor= date_add(tgl_mohon, interval angsuran2 month) where tgl_akhir_tenor is null and angsuran2 not in ('','0','N')", []interface{}{}},
	// 	{"update tr_wms_faktur3 set tgl_akhir_tenor= date_add(tgl_mohon, interval angsuran2 month) where tgl_akhir_tenor is null and angsuran2 not in ('','0','N')", []interface{}{}},
	// 	{"update tr_wms_faktur4 set tgl_akhir_tenor= date_add(tgl_mohon, interval angsuran2 month) where tgl_akhir_tenor is null and angsuran2 not in ('','0','N')", []interface{}{}},
	// 	{"update tr_wms_faktur2 set tgl_akhir_tenor= date_add(tgl_mohon, interval angsuran month) where tgl_akhir_tenor is null and angsuran not in ('','0','N')", []interface{}{}},
	// 	{"update tr_wms_faktur3 set tgl_akhir_tenor= date_add(tgl_mohon, interval angsuran month) where tgl_akhir_tenor is null and angsuran not in ('','0','N')", []interface{}{}},
	// 	{"update tr_wms_faktur4 set tgl_akhir_tenor= date_add(tgl_mohon, interval angsuran month) where tgl_akhir_tenor is null and angsuran not in ('','0','N')", []interface{}{}},
	// }
	// // Execute each query
	// tx := tr.connGorm.Begin()
	// for _, q := range queries {
	// 	result := tx.Exec(q.Query, q.Params...)
	// 	if result.Error != nil {
	// 		tx.Rollback()
	// 		fmt.Println("Error:", result.Error)
	// 	}
	// }
	// tx.Commit()
}

func (tr *tr3Repository) WillBayar(data request.SearchWBRequest) (entity.Faktur3, error) {
	var faktur entity.Faktur3
	result := tr.connGorm.Select("no_msn,sts_cetak3, no_tanda_terima,sts_bayar_renewal, nm_mtr, nm_customer11,no_telp1,no_hp1,no_kartu,sts_jenis_bayar,sts_kartu,alamat_bantuan,sts_kirim,kd_card,kode_kurir,sts_asuransi_pa,sts_bayar_asuransi_pa,alamat21,kota2,kec2,kel2,rt2,rw2,kodepos2,kerja_di,alamat_ktr,rt_ktr,rw_ktr,kel_ktr,kec_ktr,kodepos_ktr,kota1, tgl_verifikasi, alamat_srt12,alamat_srt11,kota_srt1,kec_srt1,kel_srt1,kodepos_srt1").Where("(replace(no_kartu, ' ','') = ? OR no_msn = ? or no_tanda_terima = ?)", data.Kode, data.Kode, data.Kode).Preload("Kurir").Preload("Kartu").Preload("MstCard").Find(&faktur)
	fmt.Println("ini faktur guys ", faktur.NoKartu, faktur.NoMsn, faktur.StsJnsBayar)
	if result.Error != nil {
		return entity.Faktur3{}, result.Error
	}
	stockCard := entity.StockCard{}
	tr.connGorm.Where("no_kartu = ? ", faktur.NoKartu).Find(&stockCard)
	if faktur.StsJnsBayar == "T" && faktur.StsBayarRenewal != "S" {
		return faktur, nil
	}
	if faktur.NoKartu == "" && faktur.StsJnsBayar != "T" {
		return entity.Faktur3{}, errors.New("Kartu tidak ditemukan")
	}
	if faktur.StsJnsBayar == "C" && stockCard.NoKartu == "" {
		return entity.Faktur3{}, errors.New("Kartu tidak ditemukan di stockCard")
	}

	fmt.Println("ini stoccard guys ", stockCard.StsKartu)
	if stockCard.StsKartu == "1" {
		return entity.Faktur3{}, errors.New("Belum di barcode bawa")
	} else if stockCard.StsKartu == "3" || faktur.StsBayarRenewal == "S" {
		return entity.Faktur3{}, errors.New("Sudah dibayar")
	} else if stockCard.StsKartu == "4" {
		return entity.Faktur3{}, errors.New("Posisi kartu di clear data")
	}

	return faktur, nil
}
