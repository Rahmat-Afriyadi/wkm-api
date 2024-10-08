package repository

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"
	"wkm/entity"
	"wkm/request"
)

type ParamsUpdateJenisBayar struct {
	NoTandaTerima string
	NamaCustomer  string
}

type Tr3Repository interface {
	DataWABlast(request request.DataWaBlastRequest) []entity.DataWaBlast
	SearchNoMsnByWa(request request.SearchNoMsnByWaRequest) []entity.SearchNoMsnByWa
	UpdateJenisBayar(data []ParamsUpdateJenisBayar, payment_type string, username string)
}

type tr3Repository struct {
	conn *sql.DB
}

func NewTr3nRepository(conn *sql.DB) Tr3Repository {
	return &tr3Repository{
		conn: conn,
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
			query := fmt.Sprint("select * from (select no_msn, nm_customer11, kd_user, no_yg_dihub_renewal, case when trim(no_wa) REGEXP '^[+62]|^[0-9]*$' and no_wa is not null and no_wa not like '021%' then no_wa when trim(sms_no) REGEXP '^[+62]|^[0-9]*$' and sms_no is not null and sms_no not like '021%' then sms_no when trim(no_telp2) REGEXP '^[+62]|^[0-9]*$' and no_telp2 is not null and no_telp2 not like '021%' then no_telp2 when trim(no_telp1) REGEXP '^[+62]|^[0-9]*$' and no_telp1 is not null and no_telp1 not like '021%' then no_telp1 when trim(no_hp2) REGEXP '^[+62]|^[0-9]*$' and no_hp2 is not null and no_hp2 not like '021%' then no_hp2 when trim(no_hp1)REGEXP '^[+62]|^[0-9]*$' and no_hp1 is not null and no_hp1 not like '021%' then no_hp1 end as no_wa, tgl_akhir_tenor from (select no_msn,nm_customer11,kd_user, case when ket_no_telp1 in ('1','1A','1B') then no_telp1 else null end as 'no_telp1', case when ket_no_telp2 in ('1','2') then no_telp2 else null end as 'no_telp2', case when no_hp2 is not null and no_hp2 != '' then no_hp2 else null end as 'no_hp2', case when ket_no_hp1 in ('1','1A','1B') then no_hp1 else null end as 'no_hp1', no_yg_dihub_renewal,sms_no,no_wa,tgl_akhir_tenor from ", table, " where sts_valid='1' and tgl_akhir_tenor>=? and tgl_akhir_tenor<? and (no_leas =? or no_leas2=?) and (kode_kerja ", request.KodeKerjaFilterType, " (", kodeKerjaInit[:len(kodeKerjaInit)-1], ") or kode_kerja2 ", request.KodeKerjaFilterType, " (", kodeKerjaInit[:len(kodeKerjaInit)-1], "))) t) u where no_wa is not null")
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
			data := entity.DataWaBlast{}
			if err := rows.Scan(&data.NoMsn, &data.NmCustomer, &data.KdUser, &data.NoYgDiHubRenewal, &data.NoWa, &data.TglAkhirTenor); err != nil {
				fmt.Println("Error scanning row:", err)
				continue
			}
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
