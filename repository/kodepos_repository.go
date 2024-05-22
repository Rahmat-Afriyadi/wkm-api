package repository

import (
	"context"
	"database/sql"
	"fmt"
	"wkm/entity"
)

type KodeposRepository interface {
	MasterData(search string) []entity.MasterKodepos
}

type kodeposRepository struct {
	conn *sql.DB
}

func NewKodeposRepository(conn *sql.DB) KodeposRepository {
	return &kodeposRepository{
		conn: conn,
	}
}

func (lR *kodeposRepository) MasterData(search string) []entity.MasterKodepos {
	fmt.Println("ini query ", search)
	datas := []entity.MasterKodepos{}
	ctx := context.Background()
	query := "select kd_pos, kodepos, kelurahan, kecamatan from kodepos where kodepos like ? or kelurahan like ? or kecamatan like ? limit 20"
	statement, err := lR.conn.PrepareContext(ctx, query)
	if err != nil {
		fmt.Println(err)
	}
	rows, err := statement.QueryContext(ctx, "%"+search+"%", "%"+search+"%", "%"+search+"%")
	if err != nil {
		fmt.Println("errornya di rows ", err)
		fmt.Println(err)
	}
	for rows.Next() {
		var data entity.MasterKodepos
		if err := rows.Scan(&data.KdPos, &data.KodePos, &data.Kelurahan, &data.Kecamatan); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		datas = append(datas, data)
	}
	return datas
}
