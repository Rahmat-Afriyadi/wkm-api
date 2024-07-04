package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Otr struct {
	ID          string    `form:"id" json:"id" gorm:"type:uuid;primary_key;column:id"`
	KdMdl       string    `form:"motorprice_kode" json:"motorprice_kode" gorm:"column:motorprice_kode;unique_index:idx_model,unique"`
	ProductKode string    `form:"product_kode" json:"product_kode" gorm:"column:product_kode;unique_index:idx_model,unique"`
	ProductNama string    `form:"product_nama" json:"product_nama" gorm:"column:product_nama"`
	WrnKode     string    `form:"wrn_kode" json:"wrn_kode" gorm:"column:wrn_kode"`
	Otr         uint64    `form:"otr" json:"otr" gorm:"column:otr"`
	OtrApi      string    `form:"On_The_Road" json:"On_The_Road" gorm:"-"`
	Tahun       uint16    `form:"tahun" json:"tahun" gorm:"column:tahun;unique_index:idx_model,unique"`
	CreatedAt   time.Time `form:"created_at" json:"created_at" gorm:"column:created_at;autoCreateTime"`
	MstMtr      MstMtr    `json:"mst_mtr" gorm:"references:KdMdl;foreignKey:KdMdl"`
	UpdatedAt   time.Time `form:"updated_at" json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (Otr) TableName() string {
	return "otr"
}

func (b *Otr) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}

type MstOtrNa struct {
	ID             string `form:"id" json:"id" gorm:"type:uuid;primary_key;column:id"`
	MotorPriceKode string `form:"motorprice_kode" json:"motorprice_kode" gorm:"column:motorprice_kode;"`
	Tahun          uint16 `form:"tahun" json:"tahun" gorm:"column:tahun;"`
}

func (MstOtrNa) TableName() string {
	return "otr_na"
}

type OtrNa struct {
	ID             string `form:"id" json:"id" gorm:"type:uuid;primary_key;column:id"`
	MotorPriceKode string `form:"motorprice_kode" json:"motorprice_kode" gorm:"column:motorprice_kode;"`
	ProductNama    string `form:"nm_mtr" json:"nm_mtr" gorm:"column:nm_mtr"`
	Tahun          uint16 `form:"tahun" json:"tahun" gorm:"column:tahun;"`
}

// func (OtrNa) TableName() string {
// 	return "otr"
// }

type ResponseOtr struct {
	Status int8  `json:"id"`
	Data   []Otr `json:"data"`
}
