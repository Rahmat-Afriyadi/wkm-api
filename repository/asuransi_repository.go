package repository

import (
	"context"
	"database/sql"
	"fmt"
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
	query := "select no_msn NoMsn, nm_customer11 NamaCustomer, no_wa NoTelepon, sts_renewal Status, alasan  from asuransi where no_msn = ?"
	statement, err := lR.conn.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	row := statement.QueryRow(no_msn)
	err = row.Scan(&data.NoMsn, &data.NamaCustomer, &data.NoTelepon, &data.Status, &data.Alasan)
	if err != nil {
		fmt.Println("errornya di rows ", err)
		panic(err)
	}

	return data
}

func (lR *asuransiRepository) UpdateAsuransi(dataUpdate entity.MasterAsuransi) entity.MasterAsuransi {
	ctx := context.Background()
	lR.conn.ExecContext(ctx, "UPDATE asuransi set sts_renewal=?, alasan=? where no_msn=? ", dataUpdate.Status, dataUpdate.Alasan, dataUpdate.NoMsn)

	return dataUpdate
}
