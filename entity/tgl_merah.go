package entity

import (
	"time"
)

type TglMerah struct {
	ID        uint64     `form:"id" json:"id" gorm:"primary_key;column:id"`
	TglAwal   time.Time  `form:"tgl_awal" json:"tgl_awal" gorm:"column:tgl_awal"`
	TglAkhir  time.Time  `form:"tgl_akhir" json:"tgl_akhir" gorm:"column:tgl_akhir"`
	Deskripsi string     `form:"deskripsi" json:"deskripsi" gorm:"column:deskripsi"`
	KdUser    string     `form:"kd_user" json:"kd_user" gorm:"column:kd_user"`
	User      User       `json:"user" gorm:"->;references:ID;foreignKey:KdUser"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (TglMerah) TableName() string {
	return "tgl_merah"
}
