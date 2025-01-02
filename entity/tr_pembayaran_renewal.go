package entity

import "time"

type TrPembayaranRenewal struct {
	NoMsn               string    `form:"no_msn" json:"no_msn" gorm:"primary_key;column:no_msn"`
	RenewalKe           string    `form:"renewal_ke" json:"renewal_ke" gorm:"primary_key;column:renewal_ke"`
	NmCustomer          string    `form:"nm_customer" json:"nm_customer" gorm:"column:nm_customer"`
	Kota                string    `form:"kota" json:"kota" gorm:"column:kota"`
	KirimKe             string    `form:"kirim_ke" json:"kirim_ke" gorm:"column:kirim_ke"`
	KdCard              string    `form:"kd_card" json:"kd_card" gorm:"column:kd_card"`
	KdUserTs            string    `form:"kd_user_ts" json:"kd_user_ts" gorm:"column:kd_user_ts"`
	TglJualan           time.Time `form:"tgl_jualan" json:"tgl_jualan" gorm:"type:DATE;default:null;column:tgl_jualan"`
	NoKartu             string    `form:"no_kartu" json:"no_kartu" gorm:"column:no_kartu"`
	NoTandaTerima       string    `form:"no_tanda_terima" json:"no_tanda_terima" gorm:"column:no_tanda_terima"`
	CetakKe             uint8     `form:"cetak_ke" json:"cetak_ke" gorm:"column:cetak_ke"`
	KdUserFa            string    `form:"kd_user_fa" json:"kd_user_fa" gorm:"column:kd_user_fa"`
	TglCetakTandaTerima time.Time `form:"tgl_cetak_tanda_terima" json:"tgl_cetak_tanda_terima" gorm:"type:DATE;default:null;column:tgl_cetak_tanda_terima"`
	KdUserSS            string    `form:"kd_user_ss" json:"kd_user_ss" gorm:"column:kd_user_ss"`
	JnsBayar            string    `form:"jns_bayar" json:"jns_bayar" gorm:"column:jns_bayar"`
	TglBayar            time.Time `form:"tgl_bayar" json:"tgl_bayar" gorm:"type:DATE;default:null;column:tgl_bayar"`
	TglInsert           time.Time `form:"tgl_insert" json:"tgl_insert" gorm:"type:DATE;default:null;column:tgl_insert"`
	TglUpdate           time.Time `form:"tgl_update" json:"tgl_update" gorm:"type:DATE;default:null;column:tgl_update;autoUpdateTime"`
}

func (TrPembayaranRenewal) TableName() string {
	return "tr_pembayaran_renewal"
}
