package entity

import (
	"time"
)

type StockCard struct {
	NoKartu    string    `form:"no_kartu" json:"no_kartu" gorm:"primary_key;column:no_kartu"`
	TglCetak   time.Time `form:"tgl_cetak" json:"tgl_cetak" gorm:"column:tgl_cetak"`
	TglUpdate  time.Time `form:"tgl_update" json:"tgl_update" gorm:"column:tgl_update;autoCreateTime;autoUpdateTime"`
	TglExpired time.Time `form:"tgl_expired" json:"tgl_expired" gorm:"column:tgl_expired"`
	StsKartu   string    `form:"sts_kartu" json:"sts_kartu" gorm:"column:sts_kartu"`
	KdUser     string    `form:"kd_user" json:"kd_user" gorm:"column:kd_user"`
	KdUser4    string    `form:"kd_user4" json:"kd_user4" gorm:"column:kd_user4"`
	NoMsn      string    `form:"no_msn" json:"no_msn" gorm:"column:no_msn"`
	User       User      `json:"user" gorm:"->;references:ID;foreignKey:KdUser"`
	// TglUpdate  *time.Time `json:"tgl_update" gorm:"column:tgl_update;autoCreateTime;autoUpdateTime"`
}

func (StockCard) TableName() string {
	return "stock_card"
}
