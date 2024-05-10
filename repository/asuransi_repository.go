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
	MasterDataPending() []entity.MasterAsuransi
	FindAsuransiByNoMsn(no_msn string) entity.MasterAsuransi
	UpdateAmbilAsuransi(no_msn string, kd_user string)
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
	datas := []entity.MasterAsuransi{}
	ctx := context.Background()
	query := "select no_msn NoMsn, nm_customer11 NamaCustomer from asuransi where sts_asuransi = '' or sts_asuransi is null"
	statement, err := lR.conn.PrepareContext(ctx, query)
	if err != nil {
		fmt.Println(err)
	}
	rows, err := statement.QueryContext(ctx)
	if err != nil {
		fmt.Println("errornya di rows ", err)
		fmt.Println(err)
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

func (lR *asuransiRepository) MasterDataPending() []entity.MasterAsuransi {
	datas := []entity.MasterAsuransi{}
	ctx := context.Background()
	query := "select no_msn NoMsn, nm_customer11 NamaCustomer from asuransi where sts_asuransi = 'P'"
	statement, err := lR.conn.PrepareContext(ctx, query)
	if err != nil {
		fmt.Println(err)
	}
	rows, err := statement.QueryContext(ctx)
	if err != nil {
		fmt.Println("errornya di rows ", err)
		fmt.Println(err)
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
	query := "select no_msn NoMsn, nm_customer11, nm_mtr, tgl_faktur, no_wa, sts_asuransi, alasan_pending, alasan_tdk_berminat,kd_dlr, nm_dlr, kelurahan, kecamatan, kodepos, jns_brg, harga from asuransi where no_msn = ? "
	statement, err := lR.conn.PrepareContext(ctx, query)
	if err != nil {
		fmt.Println(err)
	}
	row := statement.QueryRow(no_msn)
	err = row.Scan(&data.NoMsn, &data.NamaCustomer, &data.NamaMotor, &data.TglFaktur, &data.NoTelepon, &data.Status, &data.AlasanPending, &data.AlasanTdkBerminat, &data.KdDlr, &data.NmDlr, &data.Kelurahan, &data.Kecamatan, &data.Kodepos, &data.JnsBrg, &data.Harga)
	if err != nil {
		fmt.Println("errornya di rows ", err)
		fmt.Println(err)
	}

	return data
}

func (lR *asuransiRepository) UpdateAsuransi(dataUpdate entity.MasterAsuransi) entity.MasterAsuransi {
	KdDlr := ""
	NmDlr := ""
	Kelurahan := ""
	Kecamatan := ""
	Kodepos := ""
	JnsBrg := ""
	if dataUpdate.KdDlr != nil {
		KdDlr = *dataUpdate.KdDlr
	}
	if dataUpdate.NmDlr != nil {
		NmDlr = *dataUpdate.NmDlr
	}
	if dataUpdate.Kelurahan != nil {
		Kelurahan = *dataUpdate.Kelurahan
	}
	if dataUpdate.Kecamatan != nil {
		Kecamatan = *dataUpdate.Kecamatan
	}
	if dataUpdate.Kodepos != nil {
		Kodepos = *dataUpdate.Kodepos
	}
	if dataUpdate.JnsBrg != nil {
		JnsBrg = *dataUpdate.JnsBrg
	}
	ctx := context.Background()
	_, err := lR.conn.ExecContext(ctx, "UPDATE asuransi set sts_asuransi=?, alasan_pending=?, alasan_tdk_berminat=?, kd_dlr=?, nm_dlr=?, kelurahan=?, kecamatan=?, kodepos=?, jns_brg=?, harga=?, kd_user=?, tgl_update=? where no_msn=? ", dataUpdate.Status, dataUpdate.AlasanPending, dataUpdate.AlasanTdkBerminat, KdDlr, NmDlr, Kelurahan, Kecamatan, Kodepos, JnsBrg, dataUpdate.Harga, dataUpdate.KdUser, time.Now().Format("2006-01-02"), dataUpdate.NoMsn)
	if err != nil {
		fmt.Println("ini error update ", err)
	}
	return dataUpdate
}

func (lR *asuransiRepository) UpdateAmbilAsuransi(no_msn string, kd_user string) {
	ctx := context.Background()
	_, err := lR.conn.ExecContext(ctx, "UPDATE asuransi set sts_asuransi='P', tgl_update=?, kd_user=? where no_msn=?", time.Now().Format("2006-01-02"), kd_user, no_msn)
	if err != nil {
		fmt.Println("ini error update ", err)
	}
}
