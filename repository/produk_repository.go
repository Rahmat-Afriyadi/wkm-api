package repository

import (
	"context"
	"database/sql"
	"fmt"
	"wkm/entity"
)

type ProdukRepository interface {
	MasterData(search string, jenis_asuransi string) []entity.MasterProduk
}

type produkRepository struct {
	conn *sql.DB
}

func NewProdukRepository(conn *sql.DB) ProdukRepository {
	return &produkRepository{
		conn: conn,
	}
}

func (lR *produkRepository) MasterData(search string, jenis_asuransi string) []entity.MasterProduk {
	fmt.Println("ini query ", search)
	datas := []entity.MasterProduk{}
	ctx := context.Background()
	query := "select id_produk, nm_produk from produk where jns_asuransi = ?"
	statement, err := lR.conn.PrepareContext(ctx, query)
	if err != nil {
		fmt.Println(err)
	}
	rows, err := statement.QueryContext(ctx, jenis_asuransi)
	if err != nil {
		fmt.Println("errornya di rows ", err)
		fmt.Println(err)
	}
	for rows.Next() {
		var data entity.MasterProduk
		if err := rows.Scan(&data.KdProduk, &data.NmProduk); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		datas = append(datas, data)
	}
	return datas
}
