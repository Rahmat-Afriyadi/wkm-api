package repository

import (
	"wkm/entity"

	"gorm.io/gorm"
)

type MerkRepository interface {
	MasterData(jenis_kendaraan int) []entity.Merk
}

type merkRepository struct {
	conn *gorm.DB
}

func NewMerkRepository(conn *gorm.DB) MerkRepository {
	return &merkRepository{
		conn: conn,
	}
}

func (lR *merkRepository) MasterData(jenis_kendaraan int) []entity.Merk {
	datas := []entity.Merk{}
	if jenis_kendaraan != 0 {
		lR.conn.Where("jenis_kendaraan = ?", jenis_kendaraan).Find(&datas)
	}
	return datas
}
