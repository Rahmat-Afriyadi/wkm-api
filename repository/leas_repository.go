package repository

import (
	"context"
	"database/sql"
	"fmt"
	"wkm/entity"
)

type LeasRepository interface {
	MasterData() []entity.MasterLeas
}

type leasRepository struct {
	conn *sql.DB
}

func NewLeasnRepository(conn *sql.DB) LeasRepository {
	return &leasRepository{
		conn: conn,
	}
}

func (lR *leasRepository) MasterData() []entity.MasterLeas {
	var datas []entity.MasterLeas
	ctx := context.Background()
	query := "select no_leas2, nm_leasing from mst_leasing where no_leas2 != ''"
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
		var data entity.MasterLeas
		if err := rows.Scan(&data.Kode, &data.Nama); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		datas = append(datas, data)
	}
	return datas
}
