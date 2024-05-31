package repository

import (
	"wkm/entity"

	"gorm.io/gorm"
)

type DlrRepository interface {
	MasterData(search string) []entity.MasterDlr
}

type dlrRepository struct {
	conn *gorm.DB
}

func NewDlrRepository(conn *gorm.DB) DlrRepository {
	return &dlrRepository{
		conn: conn,
	}
}

func (lR *dlrRepository) MasterData(search string) []entity.MasterDlr {
	datas := []entity.MasterDlr{}
	lR.conn.Where("kd_dlr like ? or nm_dlr like ? ", "%"+search+"%", "%"+search+"%").Limit(15).Find(&datas)
	return datas
}
