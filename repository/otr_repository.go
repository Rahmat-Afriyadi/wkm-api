package repository

import (
	"fmt"
	"wkm/entity"

	"gorm.io/gorm"
)

type OtrRepository interface {
	DetailOtrNa(motorprice_kode string, tahun uint16) entity.Otr
	OtrNaList() []entity.Otr
	OtrMstProduk(search string) []entity.MstMtr
	OtrMstNa(search string) []entity.OtrNa
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

func (lR *otrRepository) OtrMstProduk(search string) []entity.MstMtr {
	var otr []entity.MstMtr
	lR.conn.Where("nm_mtr like ? or no_mtr like ? or kd_mdl like ?", "%"+search+"%", "%"+search+"%", "%"+search+"%").Limit(15).Find(&otr)
	return otr
}
func (lR *otrRepository) OtrMstNa(search string) []entity.OtrNa {
	var otr []entity.OtrNa
	lR.conn.Raw("select a.*, m.nm_mtr from (select motorprice_kode, tahun from otr_na group by motorprice_kode, tahun) a inner join mst_mtr m  on m.kd_mdl = a.motorprice_kode where a.motorprice_kode like ? or a.tahun like ? or m.nm_mtr like ? limit 15 ", "%"+search+"%", "%"+search+"%", "%"+search+"%").Find(&otr)
	fmt.Println("ini data ", otr)
	return otr
}
