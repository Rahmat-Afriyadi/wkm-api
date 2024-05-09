package repository

import (
	"context"
	"database/sql"
	"fmt"
	"wkm/entity"
)

type DlrRepository interface {
	MasterData(search string) []entity.MasterDlr
}

type dlrRepository struct {
	conn *sql.DB
}

func NewDlrRepository(conn *sql.DB) DlrRepository {
	return &dlrRepository{
		conn: conn,
	}
}

func (lR *dlrRepository) MasterData(search string) []entity.MasterDlr {
	fmt.Println("ini query ", search)
	datas := []entity.MasterDlr{}
	ctx := context.Background()
	query := "select kd_dlr, nm_dlr from mst_dealer where kd_dlr like ? or nm_dlr like ? "
	statement, err := lR.conn.PrepareContext(ctx, query)
	if err != nil {
		fmt.Println(err)
	}
	rows, err := statement.QueryContext(ctx, "%"+search+"%", "%"+search+"%")
	if err != nil {
		fmt.Println("errornya di rows ", err)
		fmt.Println(err)
	}
	for rows.Next() {
		var data entity.MasterDlr
		if err := rows.Scan(&data.KdDlr, &data.NmDlr); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		datas = append(datas, data)
	}
	return datas
}
