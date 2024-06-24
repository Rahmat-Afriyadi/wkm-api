package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Otr struct {
	ID             string `form:"id" json:"id" gorm:"type:uuid;primary_key;column:id"`
	MotorPriceKode string `form:"motorprice_kode" json:"motorprice_kode" gorm:"column:motorprice_kode;index:idx_model,unique"`
	ProductKode    string `form:"product_kode" json:"product_kode" gorm:"column:product_kode"`
	ProductNama    string `form:"product_nama" json:"product_nama" gorm:"column:product_nama"`
	WmKode         string `form:"wm_kode" json:"wm_kode" gorm:"column:wm_kode"`
	Otr            uint64 `form:"otr" json:"otr" gorm:"column:otr"`
	Tahun          uint16 `form:"tahun" json:"tahun" gorm:"column:tahun;index:idx_model,unique"`
}

func (Otr) TableName() string {
	return "otr"
}

func (b *Otr) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}

type OtrNa struct {
	ID             string `form:"id" json:"id" gorm:"type:uuid;primary_key;column:id"`
	MotorPriceKode string `form:"motorprice_kode" json:"motorprice_kode" gorm:"column:motorprice_kode;index:idx_model,unique"`
	ProductNama    string `form:"nm_mtr" json:"nm_mtr" gorm:"column:nm_mtr"`
	Tahun          uint16 `form:"tahun" json:"tahun" gorm:"column:tahun;index:idx_model,unique"`
}

// func (OtrNa) TableName() string {
// 	return "otr"
// }

func (b *OtrNa) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}
