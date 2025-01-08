package entity

import "time"

type AsuransiPA struct {
	Id                     string     `json:"id" gorm:"column:Id" form:"id"`
	NoMSN                    string     `json:"no_msn" gorm:"column:no_msn" form:"no_msn"`
	NmCustomer               string    `json:"nm_customer" gorm:"column:nm_customer" form:"nm_customer"`
	StsAsuransiPA            string    `json:"sts_asuransi_pa" gorm:"column:sts_asuransi_pa" form:"sts_asuransi_pa"`
	IDProduk                 string    `json:"id_produk" gorm:"column:id_produk" form:"id_produk"`
	AppTransID               string    `json:"app_trans_id" gorm:"column:app_trans_id" form:"app_trans_id"`
	TglBeli                  *time.Time `json:"tgl_beli" gorm:"column:tgl_beli" form:"tgl_beli"`
	NoKtpNpwp                string    `json:"no_ktpnpwp" gorm:"column:no_ktpnpwp" form:"no_ktpnpwp"`
	AlasanPendingAsuransiPA  string    `json:"alasan_pending_asuransi_pa" gorm:"column:alasan_pending_asuransi_pa" form:"alasan_pending_asuransi_pa"`
	StsPembelian             string    `json:"sts_pembelian" gorm:"column:sts_pembelian" form:"sts_pembelian"`

}

func (AsuransiPA) TableName() string {
	return "asuransi_pa"
}