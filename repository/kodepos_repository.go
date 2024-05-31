package repository

import (
	"wkm/entity"

	"gorm.io/gorm"
)

type KodeposRepository interface {
	MasterData(search string) []entity.MasterKodepos
}

type kodeposRepository struct {
	conn *gorm.DB
}

func NewKodeposRepository(conn *gorm.DB) KodeposRepository {
	return &kodeposRepository{
		conn: conn,
	}
}

func (lR *kodeposRepository) MasterData(search string) []entity.MasterKodepos {
	datas := []entity.MasterKodepos{}
	lR.conn.Where("kodepos like ? or kelurahan like ? or kecamatan like ?", "%"+search+"%", "%"+search+"%", "%"+search+"%").Limit(15).Find(&datas)
	return datas
}
