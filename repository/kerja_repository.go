package repository

import (
	"context"
	"database/sql"
	"fmt"
	"wkm/entity"
)

type KerjaRepository interface {
	MasterData() []entity.MasterKerja
}

type kerjaRepository struct {
	conn *sql.DB
}

func NewKerjanRepository(conn *sql.DB) KerjaRepository {
	return &kerjaRepository{
		conn: conn,
	}
}

func (lR *kerjaRepository) MasterData() []entity.MasterKerja {
	var datas []entity.MasterKerja
	ctx := context.Background()
	query := "select kode_kerja2, nm_kerja from mst_kerja where kode_kerja2 != ''"
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
		var data entity.MasterKerja
		if err := rows.Scan(&data.Kode, &data.Nama); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		datas = append(datas, data)
	}
	return datas
}
