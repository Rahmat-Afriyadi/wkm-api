package repository

import (
	"wkm/entity"

	"gorm.io/gorm"
)

type OtrRepository interface {
	DetailOtrNa(motorprice_kode string, tahun uint16) entity.Otr
	OtrNaList() []entity.Otr
}

type otrRepository struct {
	conn *gorm.DB
}

func NewOtrRepository(conn *gorm.DB) OtrRepository {
	return &otrRepository{
		conn: conn,
	}
}

func (lR *otrRepository) DetailOtrNa(motorprice_kode string, tahun uint16) entity.Otr {
	otr := entity.Otr{}
	lR.conn.Table("otr_na a").Joins("inner join mst_mtr b on a.motorprice_kode = b.kd_mdl").Select("a.tahun, a.motorprice_kode, b.nm_mtr product_nama").Where("a.motorprice_kode = ? and a.tahun = ? ", motorprice_kode, tahun).First(&otr)
	return otr
}

func (lR *otrRepository) OtrNaList() []entity.Otr {
	var otr []entity.Otr
	lR.conn.Table("otr_na").Group("motorprice_kode, tahun").Find(&otr)
	return otr
}
