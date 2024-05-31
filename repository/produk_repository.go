package repository

import (
	"wkm/entity"

	"gorm.io/gorm"
)

type ProdukRepository interface {
	MasterData(search string, jenis_asuransi string) []entity.MasterProduk
}

type produkRepository struct {
	conn *gorm.DB
}

func NewProdukRepository(conn *gorm.DB) ProdukRepository {
	return &produkRepository{
		conn: conn,
	}
}

func (lR *produkRepository) MasterData(search string, jenis_asuransi string) []entity.MasterProduk {
	datas := []entity.MasterProduk{}
	lR.conn.Where("jns_asuransi = ?", jenis_asuransi).Find(&datas)
	return datas
}
