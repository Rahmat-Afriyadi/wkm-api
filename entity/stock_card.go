package entity

import (
	"time"
)

type StockCard struct {
	NoKartu    string     `form:"no_kartu" json:"no_kartu" gorm:"primary_key;column:no_kartu"`
	TglCetak   time.Time  `form:"tgl_cetak" json:"tgl_cetak" gorm:"column:tgl_cetak"`
	TglUpdate  time.Time  `form:"tgl_update" json:"tgl_update" gorm:"column:tgl_update"`
	TglExpired time.Time  `form:"tgl_expired" json:"tgl_expired" gorm:"column:tgl_expired"`
	StsKartu   string     `form:"sts_kartu" json:"sts_kartu" gorm:"column:sts_kartu"`
	KdUser     string     `form:"kd_user" json:"kd_user" gorm:"column:kd_user"`
	User       User       `json:"user" gorm:"->;references:ID;foreignKey:KdUser"`
	CreatedAt  *time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  *time.Time `json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (StockCard) TableName() string {
	return "stock_card"
}
