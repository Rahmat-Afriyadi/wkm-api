package entity

import (
	"database/sql"
	"time"
)

type Konsumen struct {
	Nik       string         `form:"nik" json:"nik" gorm:"primary_key;column:nik"`
	Nama      string         `form:"nm_konsumen" json:"nm_konsumen" gorm:"column:nm_konsumen"`
	NoHp      string         `form:"no_hp" json:"no_hp" gorm:"column:no_hp"`
	Email     string         `form:"email" json:"email" gorm:"column:email"`
	Alamat    string         `form:"alamat" json:"alamat" gorm:"column:alamat"`
	Prop      string         `form:"prop" json:"prop" gorm:"column:prop"`
	Kota      string         `form:"kota" json:"kota" gorm:"column:kota"`
	Kec       string         `form:"kec" json:"kec" gorm:"column:kec"`
	Kodepos   string         `form:"kodepos" json:"kodepos" gorm:"column:kodepos"`
	Kota1     string         `form:"kota1" json:"kota1" gorm:"column:kota1"`
	Kecamatan string         `form:"kecamatan" json:"kecamatan" gorm:"column:kecamatan"`
	Kelurahan string         `form:"kelurahan" json:"kelurahan" gorm:"column:kelurahan"`
	TglLahir  sql.NullString `form:"tgl_lahir" json:"tgl_lahir" gorm:"column:tgl_lahir"`
	Created   time.Time      `gorm:"default:current_timestamp;column:created"`
	Updated   time.Time      `gorm:"column:updated;autoUpdateTime"`
}

func (Konsumen) TableName() string {
	return "konsumen"
}
