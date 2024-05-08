package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"wkm/entity"
)

type AsuransiRepository interface {
	MasterData() []entity.MasterAsuransi
	FindAsuransiByNoMsn(no_msn string) entity.MasterAsuransi
	UpdateAsuransi(data entity.MasterAsuransi) entity.MasterAsuransi
}

type asuransiRepository struct {
	conn *sql.DB
}

func NewAsuransiRepository(conn *sql.DB) AsuransiRepository {
	return &asuransiRepository{
		conn: conn,
	}
}

func (lR *asuransiRepository) MasterData() []entity.MasterAsuransi {
	var datas []entity.MasterAsuransi
	ctx := context.Background()
	query := "select no_msn NoMsn, nm_customer11 NamaCustomer from asuransi"
	statement, err := lR.conn.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	rows, err := statement.QueryContext(ctx)
	if err != nil {
		fmt.Println("errornya di rows ", err)
		panic(err)
	}

	for rows.Next() {
		var data entity.MasterAsuransi
		if err := rows.Scan(&data.NoMsn, &data.NamaCustomer); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		datas = append(datas, data)
	}
	return datas
}

func (lR *asuransiRepository) FindAsuransiByNoMsn(no_msn string) entity.MasterAsuransi {
	var data entity.MasterAsuransi
	ctx := context.Background()
	query := "select no_msn NoMsn, nm_customer11 NamaCustomer, no_wa NoTelepon, sts_renewal Status, alasan_pending, alasan_tdk_berminat,kd_dlr, nm_dlr, kelurahan, kecamatan, kodepos, jns_brg, harga  from asuransi where no_msn = ?"
	statement, err := lR.conn.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	row := statement.QueryRow(no_msn)
	err = row.Scan(&data.NoMsn, &data.NamaCustomer, &data.NoTelepon, &data.Status, &data.AlasanPending, &data.AlasanTdkBerminat, &data.KdDlr, &data.NmDlr, &data.Kelurahan, &data.Kecamatan, &data.Kelurahan, &data.JnsBrg, &data.Harga)
	if err != nil {
		fmt.Println("errornya di rows ", err)
		panic(err)
	}

	return data
}

func (lR *asuransiRepository) UpdateAsuransi(dataUpdate entity.MasterAsuransi) entity.MasterAsuransi {
	ctx := context.Background()
	_, err := lR.conn.ExecContext(ctx, "UPDATE asuransi set sts_renewal=?, alasan_pending=?, alasan_tdk_berminat=?, kd_dlr=?, nm_dlr=?, kelurahan=?, kecamatan=?, kodepos=?, jns_brg=?, harga=?, kd_user=?, tgl_update=? where no_msn=? ", dataUpdate.Status, dataUpdate.AlasanPending, dataUpdate.AlasanTdkBerminat, *dataUpdate.KdDlr, *dataUpdate.NmDlr, *dataUpdate.Kelurahan, *dataUpdate.Kecamatan, *dataUpdate.Kodepos, *dataUpdate.JnsBrg, dataUpdate.Harga, dataUpdate.KdUser, time.Now().Format("2006-01-02"), dataUpdate.NoMsn)
	if err != nil {
		fmt.Println("ini error update ", err)
	}

	return dataUpdate
}
