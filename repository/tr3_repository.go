package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
	"wkm/config"
	"wkm/entity"
	"wkm/request"
	"wkm/response"

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
	WillBayar(data request.SearchWBRequest) (entity.Faktur3, []int, error)
	UpdateInputBayar(data request.InputBayarRequest) (entity.Faktur3, error)
	DataRenewalRequest(data request.DataRenewalRequest) ([]response.DataRenewalResponse, error)
	ExportDataRenewalBasic(data request.DataRenewalRequest) ([]entity.DataRenewal, error)
	ExportDataRenewalGold(data request.DataRenewalRequest) ([]entity.DataRenewal, error)
	ExportDataRenewalPlatinum(data request.DataRenewalRequest) ([]entity.DataRenewal, error)
	ExportDataRenewalPlatinumPlus(data request.DataRenewalRequest) ([]entity.DataRenewal, error)
	ExportDataAsuransiPlatinumPlus(data request.DataRenewalRequest) ([]entity.DataRenewal, error)
	DataPembayaran(tgl1 string, tgl2 string) []entity.Faktur3
	UpdateInputBayarMembership(data request.InputBayarRequest)  error
	UpdateInputBayarAsuransiPA(data request.InputBayarRequest)  error
	UpdateInputBayarAsuransiMtr(data request.InputBayarRequest)  error
}

type tr3Repository struct {
	conn     *sql.DB
	connGorm *gorm.DB
	wandaGorm *gorm.DB
}

func NewTr3nRepository(conn *sql.DB, connGorm *gorm.DB,wandaGorm *gorm.DB) Tr3Repository {
	return &tr3Repository{
		conn:     conn,
		connGorm: connGorm,
		wandaGorm:wandaGorm,
	}
}

func (tr *tr3Repository) ExportDataRenewalBasic(data request.DataRenewalRequest) ([]entity.DataRenewal, error) {
	query := `
	SELECT 
	    twf.kd_dlr, 
	    twf.nm_dlr,
		twf.no_rgk,
	    twf.no_msn, 
	    twf.no_kartu, 
	    twf.nm_customer11, 
	    twf.nama_ktp,
	    mc.jns_card,
	    twf.tgl_mohon,
	    twf.alamat11, 
	    twf.rt1, 
	    twf.rw1, 
	    twf.kel1, 
	    twf.kec1, 
	    twf.kota1, 
	    twf.kodepos1, 
	    twf.alamat21, 
	    twf.rt2, 
	    twf.rw2, 
	    twf.kel2, 
	    twf.kec2, 
	    twf.kota2, 
	    twf.kodepos2, 
	    twf.jns_beli, 
	    twf.tgl_bayar_renewal_fin AS 'tgl_awal', 
	    DATE_ADD(twf.tgl_bayar_renewal_fin, INTERVAL 13 MONTH) AS 'tgl_akhir', 
	    twf.no_tanda_terima
	FROM 
	    db_wkm.tr_wms_faktur3 twf
	JOIN 
	    db_wkm.mst_card mc ON twf.kd_card = mc.kd_card 
	WHERE 
	    YEAR(twf.tgl_bayar_renewal_fin) = ? 
	    AND MONTH(twf.tgl_bayar_renewal_fin) = ?
		AND twf.sts_renewal = 'O' AND twf.sts_bayar_renewal = 'S'
		AND mc.jns_card LIKE '%BASIC%';`

	rows, err := tr.conn.Query(query, data.Year, data.Month)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []entity.DataRenewal

	for rows.Next() {
		var result entity.DataRenewal
		if err := rows.Scan(&result.KdDlr, &result.NmDlr,&result.NoRgk , &result.NoMsn, &result.NoKartu, &result.NmCustomer, &result.NamaKtp,
			&result.JnsCard, &result.TglMohon, &result.Alamat11, &result.Rt1, &result.Rw1,
			&result.Kel1, &result.Kec1, &result.Kota1, &result.Kodepos1, &result.Alamat,
			&result.Rt, &result.Rw, &result.Kel, &result.Kec, &result.Kota, &result.Kodepos,
			&result.JnsBeli, &result.TglAwal, &result.TglAkhir, &result.NoTandaTerima); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (tr *tr3Repository) ExportDataRenewalGold(data request.DataRenewalRequest) ([]entity.DataRenewal, error) {
	query := `
	SELECT 
	    twf.kd_dlr, 
	    twf.nm_dlr, 
		twf.no_rgk,
	    twf.no_msn, 
	    twf.no_kartu, 
	    twf.nm_customer11, 
	    twf.nama_ktp,
	    mc.jns_card,
	    twf.tgl_mohon,
	    twf.alamat11, 
	    twf.rt1, 
	    twf.rw1, 
	    twf.kel1, 
	    twf.kec1, 
	    twf.kota1, 
	    twf.kodepos1, 
	    twf.alamat21, 
	    twf.rt2, 
	    twf.rw2, 
	    twf.kel2, 
	    twf.kec2, 
	    twf.kota2, 
	    twf.kodepos2, 
	    twf.jns_beli, 
	    twf.tgl_bayar_renewal_fin AS 'tgl_awal', 
	    DATE_ADD(twf.tgl_bayar_renewal_fin, INTERVAL 13 MONTH) AS 'tgl_akhir', 
	    twf.no_tanda_terima
	FROM 
	    db_wkm.tr_wms_faktur3 twf
	JOIN 
	    db_wkm.mst_card mc ON twf.kd_card = mc.kd_card 
	WHERE 
	    YEAR(twf.tgl_bayar_renewal_fin) = ? 
	    AND MONTH(twf.tgl_bayar_renewal_fin) = ?
		AND twf.sts_renewal = 'O' AND twf.sts_bayar_renewal = 'S'
		AND mc.jns_card LIKE '%Gold%';`

	rows, err := tr.conn.Query(query, data.Year, data.Month)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []entity.DataRenewal

	for rows.Next() {
		var result entity.DataRenewal
		if err := rows.Scan(&result.KdDlr, &result.NmDlr, &result.NoRgk, &result.NoMsn, &result.NoKartu, &result.NmCustomer, &result.NamaKtp,
			&result.JnsCard, &result.TglMohon, &result.Alamat11, &result.Rt1, &result.Rw1,
			&result.Kel1, &result.Kec1, &result.Kota1, &result.Kodepos1, &result.Alamat,
			&result.Rt, &result.Rw, &result.Kel, &result.Kec, &result.Kota, &result.Kodepos,
			&result.JnsBeli, &result.TglAwal, &result.TglAkhir, &result.NoTandaTerima); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (tr *tr3Repository) ExportDataRenewalPlatinum(data request.DataRenewalRequest) ([]entity.DataRenewal, error) {
	query := `
	SELECT 
	    twf.kd_dlr, 
	    twf.nm_dlr, 
		twf.no_rgk,
	    twf.no_msn, 
	    twf.no_kartu, 
	    twf.nm_customer11, 
	    twf.nama_ktp,
	    mc.jns_card,
	    twf.tgl_mohon,
	    twf.alamat11, 
	    twf.rt1, 
	    twf.rw1, 
	    twf.kel1, 
	    twf.kec1, 
	    twf.kota1, 
	    twf.kodepos1, 
	    twf.alamat21, 
	    twf.rt2, 
	    twf.rw2, 
	    twf.kel2, 
	    twf.kec2, 
	    twf.kota2, 
	    twf.kodepos2, 
	    twf.jns_beli, 
	    twf.tgl_bayar_renewal_fin AS 'tgl_awal', 
	    DATE_ADD(twf.tgl_bayar_renewal_fin, INTERVAL 13 MONTH) AS 'tgl_akhir', 
	    twf.no_tanda_terima
	FROM 
	    db_wkm.tr_wms_faktur3 twf
	JOIN 
	    db_wkm.mst_card mc ON twf.kd_card = mc.kd_card 
	WHERE 
	    YEAR(twf.tgl_bayar_renewal_fin) = ? 
		AND twf.sts_renewal = 'O' AND twf.sts_bayar_renewal = 'S'
	    AND MONTH(twf.tgl_bayar_renewal_fin) = ?
		AND mc.jns_card LIKE '%PLATINUM%' 
	    AND mc.jns_card NOT LIKE '%PLUS%';`

	rows, err := tr.conn.Query(query, data.Year, data.Month)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []entity.DataRenewal

	for rows.Next() {
		var result entity.DataRenewal
		if err := rows.Scan(&result.KdDlr, &result.NmDlr, &result.NoRgk, &result.NoMsn, &result.NoKartu, &result.NmCustomer, &result.NamaKtp,
			&result.JnsCard, &result.TglMohon, &result.Alamat11, &result.Rt1, &result.Rw1,
			&result.Kel1, &result.Kec1, &result.Kota1, &result.Kodepos1, &result.Alamat,
			&result.Rt, &result.Rw, &result.Kel, &result.Kec, &result.Kota, &result.Kodepos,
			&result.JnsBeli, &result.TglAwal, &result.TglAkhir, &result.NoTandaTerima); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (tr *tr3Repository) ExportDataRenewalPlatinumPlus(data request.DataRenewalRequest) ([]entity.DataRenewal, error) {
	query := `
	SELECT 
	    twf.kd_dlr, 
	    twf.nm_dlr, 
		twf.no_rgk,
	    twf.no_msn, 
	    twf.no_kartu,
	    twf.nm_mtr,
	    twf.nm_customer11, 
	    twf.nama_ktp,
	    mc.jns_card,
	    twf.tgl_mohon,
	    twf.alamat11, 
	    twf.rt1, 
	    twf.rw1, 
	    twf.kel1, 
	    twf.kec1, 
	    twf.kota1, 
	    twf.kodepos1, 
	    twf.alamat21, 
	    twf.rt2, 
	    twf.rw2, 
	    twf.kel2, 
	    twf.kec2, 
	    twf.kota2, 
	    twf.kodepos2, 
	    twf.jns_beli, 
	    twf.tgl_bayar_renewal_fin AS 'tgl_awal', 
	    DATE_ADD(twf.tgl_bayar_renewal_fin, INTERVAL 13 MONTH) AS 'tgl_akhir', 
	    twf.no_tanda_terima
	FROM 
	    db_wkm.tr_wms_faktur3 twf
	JOIN 
	    db_wkm.mst_card mc ON twf.kd_card = mc.kd_card 
	WHERE 
	    YEAR(twf.tgl_bayar_renewal_fin) = ? 
	    AND MONTH(twf.tgl_bayar_renewal_fin) = ?
		AND twf.sts_renewal = 'O' AND twf.sts_bayar_renewal = 'S'
		AND mc.jns_card LIKE '%PLUS%';`

	rows, err := tr.conn.Query(query, data.Year, data.Month)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []entity.DataRenewal

	for rows.Next() {
		var result entity.DataRenewal
		if err := rows.Scan(&result.KdDlr, &result.NmDlr, &result.NoRgk, &result.NoMsn, &result.NoKartu, &result.NmMtr, &result.NmCustomer, &result.NamaKtp,
			&result.JnsCard, &result.TglMohon, &result.Alamat11, &result.Rt1, &result.Rw1,
			&result.Kel1, &result.Kec1, &result.Kota1, &result.Kodepos1, &result.Alamat,
			&result.Rt, &result.Rw, &result.Kel, &result.Kec, &result.Kota, &result.Kodepos,
			&result.JnsBeli, &result.TglAwal, &result.TglAkhir, &result.NoTandaTerima); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (tr *tr3Repository) ExportDataAsuransiPlatinumPlus(data request.DataRenewalRequest) ([]entity.DataRenewal, error) {
	query := `
	SELECT 
    twf.kd_dlr, 
    twf.nm_dlr, 
    twf.no_msn, 
    twf.no_kartu,
    twf.no_rgk,
    twf.nm_mtr,
    twf.nm_customer11, 
    twf.nama_ktp,
    mc.jns_card,
    twf.tgl_mohon,
    twf.alamat11, 
    twf.rt1, 
    twf.rw1, 
    twf.kel1, 
    twf.kec1, 
    twf.kota1, 
    twf.kodepos1, 
    twf.alamat21, 
    twf.rt2, 
    twf.rw2, 
    twf.kel2, 
    twf.kec2, 
    twf.kota2, 
    twf.kodepos2, 
    twf.jns_beli, 
    CAST(DATE_FORMAT(twf.tgl_bayar_renewal_fin, '%Y-%m-01') AS DATE) AS 'tgl_awal',
    CAST(DATE_FORMAT(DATE_ADD(twf.tgl_bayar_renewal_fin, INTERVAL 12 MONTH), '%Y-%m-01') AS DATE) AS 'tgl_akhir', 
    twf.no_tanda_terima
FROM 
    db_wkm.tr_wms_faktur3 twf
JOIN 
    db_wkm.mst_card mc ON twf.kd_card = mc.kd_card 
WHERE 
    YEAR(twf.tgl_bayar_renewal_fin) = ? 
    AND MONTH(twf.tgl_bayar_renewal_fin) = ?
    AND twf.sts_renewal = 'O' 
    AND twf.sts_bayar_renewal = 'S'
    AND mc.jns_card LIKE '%PLUS%';`

	rows, err := tr.conn.Query(query, data.Year, data.Month)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []entity.DataRenewal

	for rows.Next() {
		var result entity.DataRenewal
		if err := rows.Scan(&result.KdDlr, &result.NmDlr, &result.NoMsn, &result.NoKartu, &result.NoRgk, &result.NmMtr, &result.NmCustomer, &result.NamaKtp,
			&result.JnsCard, &result.TglMohon, &result.Alamat11, &result.Rt1, &result.Rw1,
			&result.Kel1, &result.Kec1, &result.Kota1, &result.Kodepos1, &result.Alamat,
			&result.Rt, &result.Rw, &result.Kel, &result.Kec, &result.Kota, &result.Kodepos,
			&result.JnsBeli, &result.TglAwal, &result.TglAkhir, &result.NoTandaTerima); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (tr *tr3Repository) DataRenewalRequest(data request.DataRenewalRequest) ([]response.DataRenewalResponse, error) {
	fmt.Println("ini data ", data)
	query := `
    SELECT 
    'PLATINUM PLUS' AS jns_card,
    COUNT(*) AS total_jumlah_data
FROM 
    db_wkm.tr_wms_faktur3 twf
JOIN 
    db_wkm.mst_card mc ON twf.kd_card = mc.kd_card 
WHERE 
	twf.sts_renewal = 'O' 
	AND twf.sts_bayar_renewal = 'S'
    AND YEAR(twf.tgl_bayar_renewal_fin) = ?
    AND MONTH(twf.tgl_bayar_renewal_fin) = ?
    AND mc.jns_card LIKE '%PLUS%'

UNION ALL

SELECT 
    'PLATINUM' AS jns_card,
    SUM(jumlah_data) AS total_jumlah_data
FROM 
    (SELECT 
        COUNT(*) AS jumlah_data
    FROM 
        db_wkm.tr_wms_faktur3 twf
    JOIN 
        db_wkm.mst_card mc ON twf.kd_card = mc.kd_card 
    WHERE 
		twf.sts_renewal = 'O' 
		AND twf.sts_bayar_renewal = 'S'
        AND YEAR(twf.tgl_bayar_renewal_fin) = ?
        AND MONTH(twf.tgl_bayar_renewal_fin) = ?
        AND mc.jns_card LIKE '%PLATINUM%'
        AND mc.jns_card NOT LIKE '%PLUS%'
    GROUP BY 
        mc.jns_card) AS subquery

UNION ALL

SELECT 
    'GOLD' AS jns_card,
    COUNT(*) AS total_jumlah_data
FROM 
    db_wkm.tr_wms_faktur3 twf
JOIN 
    db_wkm.mst_card mc ON twf.kd_card = mc.kd_card 
WHERE 
	twf.sts_renewal = 'O' 
	AND twf.sts_bayar_renewal = 'S'
    AND YEAR(twf.tgl_bayar_renewal_fin) = ? 
    AND MONTH(twf.tgl_bayar_renewal_fin) = ?
    AND mc.jns_card LIKE '%GOLD%'

UNION ALL

SELECT 
    'BASIC' AS jns_card,
    COUNT(*) AS total_jumlah_data
FROM 
    db_wkm.tr_wms_faktur3 twf
JOIN 
    db_wkm.mst_card mc ON twf.kd_card = mc.kd_card 
WHERE 
	twf.sts_renewal = 'O' 
	AND twf.sts_bayar_renewal = 'S'
    AND YEAR(twf.tgl_bayar_renewal_fin) = ?
    AND MONTH(twf.tgl_bayar_renewal_fin) = ?
    AND mc.jns_card LIKE '%BASIC%';
    `
	rows, err := tr.conn.Query(query, data.Year, data.Month, data.Year, data.Month, data.Year, data.Month, data.Year, data.Month)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []response.DataRenewalResponse

	// Mengambil hasil dari query
	for rows.Next() {
		var result response.DataRenewalResponse

		if err := rows.Scan(&result.JnsCard, &result.TotalJumlahData); err != nil {
			return nil, err
		}
		zero := 0
		if result.TotalJumlahData == nil {
			result.TotalJumlahData = &zero
		}
		results = append(results, result)
	}

	// Memeriksa apakah ada kesalahan saat iterasi
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

func (tr *tr3Repository) DataWABlast(request request.DataWaBlastRequest) []entity.DataWaBlast {
	// format (table, filterTypeKodeKerja1, filterTypeKodeKerja2)
	// params (awaltenor, akhirtenor, leas1, leas2, kodekerja1, kodekerja2)
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
	var membership entity.Membership
	faktur3 := entity.Faktur3{NoMsn: data.NoMsn}
	customerMtr := entity.CustomerMtr{NoMsn: data.NoMsn}
	tr.connGorm.Find(&faktur3)
	tr.connGorm.Find(&customerMtr)
	if customerMtr.NmCustomerFkt != "" {
		tr.connGorm.Where("no_msn = ? and renewal_ke = ?", faktur3.NoMsn, faktur3.StsCetak3).Preload("MstCard").First(&membership)
	}
	now := time.Now()
	tglExpired := now.AddDate(1,0,0)


	if faktur3.NmCustomer == "" {
		return entity.Faktur3{}, errors.New("data tidak ditemukan")
	}
	noTT := faktur3.NoTandaTerima
	if membership.TypeKartu == "E" {
		noTT, err :=  entity.GenerateNoTT()
		if err != nil {
			return entity.Faktur3{}, err
		}
		noKartu, err := entity.GenerateEcardNumber(membership.MstCard.JnsCard)
		if err != nil {
			return entity.Faktur3{}, err
		}

		faktur3.NoTandaTerima = noTT
		faktur3.TglCetakTandaTerima = &now
		faktur3.TglExpired = &tglExpired
		faktur3.NoKartu = noKartu

		stockCard := entity.StockCard{NoKartu: noKartu}
		tr.connGorm.Find(&stockCard)
		stockCard.StsKartu = "3"
		stockCard.TglExpired = tglExpired
		stockCard.NoMsn = faktur3.NoMsn
		stockCard.TglUpdate = time.Now()
		stockCard.TglCetak = time.Now()
		stockCard.KdUser4 = data.KdUserFa
		result := tr.connGorm.Save(&stockCard)
		if result.Error != nil {
			return entity.Faktur3{}, result.Error
		}

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
		NoTandaTerima:       noTT,
		KdUserFa:            data.KdUserFa,
		TglCetakTandaTerima: faktur3.TglCetakTandaTerima,
		KdUserSS:            faktur3.KdUser2,
		JnsBayar:            faktur3.StsJnsBayar,
		TglBayar:            data.TglBayar,
		TglInsert:           time.Now(),
		TglJualan:           faktur3.TglVerifikasi,
	}
	faktur3.TglBayarRenewalFin = &data.TglBayar
	faktur3.TglBayarRenewalFinKeyIn = &now
	faktur3.KdUser2 = data.KdUserFa
	faktur3.StsKartu = "3"
	faktur3.StsBawaKartu = "4"
	faktur3.StsAsuransiPa = "O"
	faktur3.StsBayarAsuransiPa = "S"
	faktur3.StsBayarRenewal = "S"

	if faktur3.StsJnsBayar == "C" && membership.TypeKartu != "E" {
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
		gormE,_ := config.GetConnectionECardPlus()
		gormE.Save(&entity.ECardPlusMember{NoKartu: faktur3.NoKartu,NmCustomer: faktur3.NmCustomer,  NoMsn: faktur3.NoMsn, TglExpired: stockCard.TglExpired})
	}

	result := tr.connGorm.Save(&trPembayaranRenewal)
	if result.Error != nil {
		return entity.Faktur3{}, result.Error
	}
	result = tr.connGorm.Save(&faktur3)
	if result.Error != nil {
		return entity.Faktur3{}, result.Error
	}
	err := tr.UpdateInputBayarMembership(data)
	if err != nil {
		return entity.Faktur3{}, err
	}
	
	LogBayar(fmt.Sprint(time.Now().Format("2006-01-02 15:04:05"), " - ", data.NoMsn, " - ", data.KdUserFa))

	return entity.Faktur3{}, nil
}

func (tr *tr3Repository) UpdateInputBayarMembership(data request.InputBayarRequest)  error {
	customerMtrRepo := NewCustomerMtrRepository(tr.conn, tr.connGorm, tr.wandaGorm)
	faktur3 := entity.Faktur3{NoMsn: data.NoMsn}
	customerMtr := entity.CustomerMtr{NoMsn: data.NoMsn}
	tr.connGorm.Find(&faktur3)
	tr.connGorm.Find(&customerMtr)
	if customerMtr.NmCustomerFkt == "" {
		err := customerMtrRepo.CreateCustomerFFaktur(data.NoMsn)
		if err != nil {
			fmt.Println("ini error pindah ", err)
			return err
		}else {
			tr.connGorm.Where("no_msn = ? and renewal_ke = ?", faktur3.NoMsn, faktur3.StsCetak3).First(&customerMtr)
		}
	}
	if customerMtr.NmCustomerFkt != "" {
		var membership entity.Membership
		stockCard := entity.StockCard{NoKartu: faktur3.NoKartu}
		tr.connGorm.Find(&stockCard)
		tr.connGorm.Where("no_msn = ? and renewal_ke = ?", faktur3.NoMsn, faktur3.StsCetak3).First(&membership)
		now := time.Now()
		if membership.Id != "" {
			membership.NoTandaTerima = faktur3.NoTandaTerima
			membership.TglCetakTandaTerima = faktur3.TglCetakTandaTerima
			membership.NoKartu = faktur3.NoKartu
			membership.TglBayar = &data.TglBayar
			membership.TglInputBayar = &now
			membership.KdUserFa = data.KdUserFa
			membership.TglExpired = &stockCard.TglExpired
			membership.KodeKurir = faktur3.KdKurir
			membership.RenewalKe = customerMtr.RenewalKe
			membership.AlasanTdkKurir = faktur3.AlasanBelumBayar2
			membership.KirimKe = faktur3.StsKirim
			membership.StsBayar = "S"
			membership.TglExpired= &stockCard.TglExpired
			tr.connGorm.Save(&membership)
		}else {
			newMembership := entity.Membership{
				NoMSN: faktur3.NoMsn,
				StsMembership: "O",
				TypeKartu: "F",
				KirimKe: faktur3.StsKirim,
				StsKartu: "3",
				TglKonfirmasi: customerMtr.TglKonfirmasi,
				KodeKurir: faktur3.KdKurir,
				KdUserCetakTt: *faktur3.KdUser3,
				KdUserTs: faktur3.KdUser,
				RenewalKe: customerMtr.RenewalKe,
				NoTandaTerima: faktur3.NoTandaTerima,
				TglJanjiBayar: faktur3.TglBayarRenewal,
				JnsMembership: faktur3.KdCard,
				JnsBayar: faktur3.StsJnsBayar,
				TglCetakTandaTerima: faktur3.TglCetakTandaTerima,
				NoKartu: faktur3.NoKartu,
				TglBayar: &data.TglBayar,
				TglInputBayar: &now,
				KdUserFa: data.KdUserFa,
				TglExpired: &stockCard.TglExpired,
				StsBayar: "S",
			}
			tr.connGorm.Save(&newMembership)

		}
	}


	return nil
}
func (tr *tr3Repository) UpdateInputBayarAsuransiPA(data request.InputBayarRequest)  error {
	faktur3 := entity.Faktur3{NoMsn: data.NoMsn}
	customerMtr := entity.CustomerMtr{NoMsn: data.NoMsn}
	tr.connGorm.Find(&faktur3)
	tr.connGorm.Find(&customerMtr)
	now := time.Now()
	if faktur3.NmCustomer == "" {
		return errors.New("data tidak ditemukan")
	}
	if customerMtr.NmCustomerFkt != "" {
		var asuransiPa entity.AsuransiPA
		stockCard := entity.StockCard{NoKartu: faktur3.NoKartu}
		tr.connGorm.Find(&stockCard)
		tr.connGorm.Where("no_msn = ? and (sts_bayar != 'S' or sts_bayar is null)", faktur3.NoMsn).First(&asuransiPa)
		if asuransiPa.Id != "" {
			polisId, err := entity.GeneratePolisPAID(tr.connGorm)
			if err != nil {
				fmt.Println("ini error generate polish ", err.Error())
			}
			asuransiPa.NoPolis = polisId
			asuransiPa.TglBayar = &data.TglBayar
			asuransiPa.TglInputBayar = &now
			asuransiPa.KdUserFa = data.KdUserFa
			asuransiPa.StsBayar = "S"
			tr.connGorm.Save(&asuransiPa)

			gormE,_ := config.GetConnectionECardPlus()
			gormE.Save(&entity.ECardPlusMember{NoKartu: asuransiPa.NoPolis,NmCustomer: faktur3.NmCustomer,  NoMsn: faktur3.NoMsn, TglExpired: stockCard.TglExpired})
		} else {
			polisId, err := entity.GeneratePolisPAID(tr.connGorm)
			if err != nil {
				fmt.Println("ini error generate polish ", err.Error())
			}
			newAsurasiPa := entity.AsuransiPA{
				NoMSN: faktur3.NoMsn,
				NoPolis:polisId,
				TglBayar:&data.TglBayar,
				TglInputBayar:&now,
				KdUserFa:data.KdUserFa,
				StsBayar:"S",
			}
			tr.connGorm.Save(&newAsurasiPa)

		}
	}

	return nil
}
func (tr *tr3Repository) UpdateInputBayarAsuransiMtr(data request.InputBayarRequest)  error {
	faktur3 := entity.Faktur3{NoMsn: data.NoMsn}
	customerMtr := entity.CustomerMtr{NoMsn: data.NoMsn}
	tr.connGorm.Find(&faktur3)
	tr.connGorm.Find(&customerMtr)
	if faktur3.NmCustomer == "" {
		return errors.New("data tidak ditemukan")
	}
	now := time.Now()
	if customerMtr.NmCustomerFkt != "" {
		var asuransiMtr entity.AsuransiMtr
		stockCard := entity.StockCard{NoKartu: faktur3.NoKartu}
		tr.connGorm.Find(&stockCard)
		tr.connGorm.Where("no_msn = ? and (sts_bayar != 'S' or sts_bayar is null)", faktur3.NoMsn).First(&asuransiMtr)
		if asuransiMtr.Id != "" {
			polisId, err := entity.GeneratePolisMtrID(tr.connGorm)
			if err != nil {
				fmt.Println("ini error generate polish ", err.Error())
				return err
			}
			asuransiMtr.NoPolis = polisId
			asuransiMtr.TglBayar = &data.TglBayar
			asuransiMtr.TglInputBayar = &now
			asuransiMtr.KdUserFa = data.KdUserFa
			asuransiMtr.StsBayar = "S"
			tr.connGorm.Save(&asuransiMtr)


			gormE,_ := config.GetConnectionECardPlus()
			gormE.Create(&entity.ECardPlusMember{NoKartu: asuransiMtr.NoPolis,NmCustomer: faktur3.NmCustomer, NoMsn: faktur3.NoMsn, TglExpired: stockCard.TglExpired})

		}
	}
	return nil
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
}

func (lR *tr3Repository) DataPembayaran(tgl1 string, tgl2 string) []entity.Faktur3 {
	var datas []entity.Faktur3
	query := lR.connGorm.Table("tr_wms_faktur3 AS a")
	if tgl1 != "" && tgl2 != "" {
		query.Where("a.tgl_bayar_renewal_fin >= ? and a.tgl_bayar_renewal_fin <= ? and sts_renewal = 'O' and sts_bayar_renewal = 'S'", tgl1, tgl2)
	}
	query.Select("a.no_msn, a.nama_ktp, a.kd_card,a.no_tanda_terima,a.no_kartu, a.nm_customer11, a.tgl_bayar_renewal_fin, a.kd_user, a.kd_user10, a.kode_kurir, a.sts_jenis_bayar").Preload("User").Preload("User10").Preload("Kurir").Preload("MstCard", func(db *gorm.DB) *gorm.DB {
		return db.Select("kd_card,jns_card,harga_pokok,asuransi,asuransi_motor") // Pilih kolom tertentu
	}).Find(&datas)
	return datas
}

func (tr *tr3Repository) WillBayar(data request.SearchWBRequest) (entity.Faktur3, []int, error) {
	var faktur entity.Faktur3
	var asuransiPa entity.AsuransiPA
	var asuransiMtr entity.AsuransiMtr
	var membership entity.Membership
	var bayarApa []int
	var afterMembership []int
	result := tr.connGorm.Select("no_msn,sts_cetak3, no_tanda_terima,sts_bayar_renewal, nm_mtr, nm_customer11,no_telp1,no_hp1,no_kartu,sts_jenis_bayar,sts_kartu,alamat_bantuan,sts_kirim,kd_card,kode_kurir,sts_asuransi_pa,sts_bayar_asuransi_pa,alamat21,kota2,kec2,kel2,rt2,rw2,kodepos2,kerja_di,alamat_ktr,rt_ktr,rw_ktr,kel_ktr,kec_ktr,kodepos_ktr,kota1, tgl_verifikasi, alamat_srt12,alamat_srt11,kota_srt1,kec_srt1,kel_srt1,kodepos_srt1").Where("(replace(no_kartu, ' ','') = ? OR no_msn = ? or no_tanda_terima = ?)", data.Kode, data.Kode, data.Kode).Preload("Kurir").Preload("Kartu").Preload("MstCard").Find(&faktur)
	if result.Error != nil {
		return entity.Faktur3{}, bayarApa,result.Error
	}
	tr.connGorm.Where("no_msn = ? and renewal_ke = ? and sts_bayar != 'S'",faktur.NoMsn, faktur.StsCetak3).First(&membership)
	if membership.TypeKartu == "E" {
		faktur.TypeKartu = "E"
	}

	tr.connGorm.Model(&entity.AsuransiPA{}).Where("no_msn = ? and (sts_bayar != 'S' or sts_bayar is null)", faktur.NoMsn).First(&asuransiPa)
	if asuransiPa.Id != "" {
		var produkPa entity.MasterProduk
		tr.wandaGorm.Preload("Vendor").Where("id_produk = ?",asuransiPa.IDProduk).Find(&produkPa)
		faktur.AsuransiPa = entity.DetailAsuransiPA{IdProduk:asuransiPa.IDProduk,NmVendor:produkPa.Vendor.NmVendor,NmProduk:produkPa.NmProduk, AmountPa: asuransiPa.AmountPa}
		afterMembership = append(afterMembership, 2)
	}
	tr.connGorm.Model(&entity.AsuransiMtr{}).Where("no_msn = ? and (sts_bayar != 'S' or sts_bayar is null)", faktur.NoMsn).First(&asuransiMtr)
	if asuransiMtr.Id != "" {
		var produkMtr entity.MasterProduk
		tr.wandaGorm.Preload("Vendor").Where("id_produk = ?",asuransiMtr.IDProduk).Find(&produkMtr)
		faktur.AsuransiMtr = entity.DetailAsuransiMtr{IdProduk:asuransiMtr.IDProduk,NmVendor:produkMtr.Vendor.NmVendor,NmProduk:produkMtr.NmProduk, NmMtr: asuransiMtr.NmMtr, Otr: asuransiMtr.OTR,AmountOtr: asuransiMtr.Amount }
		afterMembership = append(afterMembership, 3)
	}

	if faktur.StsJnsBayar == "T" && faktur.StsBayarRenewal != "S" {
		bayarApa = append(bayarApa, 1)
		bayarApa = append(bayarApa, afterMembership...)
		return faktur, bayarApa,nil
	}

	var errorValidasiMembership error
	stockCard := entity.StockCard{}
	tr.connGorm.Where("no_kartu = ? ", faktur.NoKartu).Find(&stockCard)
	if faktur.NoKartu == "" && faktur.StsJnsBayar != "T" && (membership.Id != "" && faktur.TypeKartu != "E") {
		errorValidasiMembership =  errors.New("Kartu tidak ditemukan")
	}
	if faktur.StsJnsBayar == "C" && stockCard.NoKartu == "" && (membership.Id != "" && faktur.TypeKartu != "E") {
		errorValidasiMembership = errors.New("Kartu tidak ditemukan di stockCard")
	}
	
	if stockCard.StsKartu == "1" && (membership.Id != "" && faktur.TypeKartu != "E") {
		errorValidasiMembership = errors.New("Belum di barcode bawa")
		} else if stockCard.StsKartu == "3" || faktur.StsBayarRenewal == "S" {
			errorValidasiMembership = errors.New("Sudah dibayar")
	} else if stockCard.StsKartu == "4" && (membership.Id != "" && faktur.TypeKartu != "E") {
		errorValidasiMembership =  errors.New("Posisi kartu di clear data")
	}
	
	
	if len(afterMembership) < 1 && errorValidasiMembership != nil{
		return entity.Faktur3{}, afterMembership, errorValidasiMembership
	}
	if errorValidasiMembership == nil {
		bayarApa = append(bayarApa, 1)
		bayarApa = append(bayarApa, afterMembership...)
	}
	
	if len(bayarApa) < 1  {
		return faktur, afterMembership, nil
	}

	return faktur, bayarApa, nil
}
