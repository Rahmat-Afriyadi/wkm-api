package response

import "time"

type MinatMembership struct {
	NoMsn               string     `form:"no_msn" json:"no_msn" gorm:"primary_key;column:no_msn"`
	NmCustomer          string     `form:"nm_customer11" json:"nm_customer11" gorm:"column:nm_customer11"`
	StsJnsBayar         string     `form:"sts_jenis_bayar" json:"sts_jenis_bayar" gorm:"column:sts_jenis_bayar"`
	TglBayarRenewal         *time.Time `form:"tgl_bayar_renewal" json:"tgl_bayar_renewal" gorm:"column:tgl_bayar_renewal"`
	Print               uint8      `form:"print" json:"print" gorm:"column:print"`
	StsRenewal         string     `form:"sts_renewal" json:"sts_renewal" gorm:"column:sts_renewal"`
	StsKartu            string     `form:"sts_kartu" json:"sts_kartu" gorm:"column:sts_kartu"`
	StsBayarRenewal     string     `form:"sts_bayar_renewal" json:"sts_bayar_renewal" gorm:"column:sts_bayar_renewal"`
	KdCard              string     `form:"kd_card" json:"kd_card" gorm:"column:kd_card"`
}