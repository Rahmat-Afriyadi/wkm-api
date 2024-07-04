package entity

import (
	"time"
)

type Merk struct {
	ID             string     `form:"id" json:"id" gorm:"primary_key;column:id"`
	Merk           string     `form:"merk" json:"merk" gorm:"column:merk;"`
	JenisKendaraan uint8      `form:"jenis_kendaraan" json:"jenis_kendaraan" gorm:"column:jenis_kendaraan;"`
	CreatedAt      *time.Time `form:"created_at" json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      *time.Time `form:"updated_at" json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (Merk) TableName() string {
	return "mst_merk"
}
