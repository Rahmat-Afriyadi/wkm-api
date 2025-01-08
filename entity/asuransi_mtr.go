package entity

import "time"

type AsuransiMtr struct {
	NoMSN                     string     `json:"no_msn" gorm:"column:no_msn" form:"no_msn"`
	NmCustomer                string    `json:"nm_customer" gorm:"column:nm_customer" form:"nm_customer"`
	StsAsuransiMtr            string    `json:"sts_asuransi_mtr" gorm:"column:sts_asuransi_mtr" form:"sts_asuransi_mtr"`
	IDProduk                  string    `json:"id_produk" gorm:"column:id_produk" form:"id_produk"`
	AppTransID                string    `json:"app_trans_id" gorm:"column:app_trans_id" form:"app_trans_id"`
	TglBeli                   *time.Time `json:"tgl_beli" gorm:"column:tgl_beli" form:"tgl_beli"`
	NoKtpNpwp                 string    `json:"no_ktpnpwp" gorm:"column:no_ktpnpwp" form:"no_ktpnpwp"`
	AlasanPendingAsuransiMtr  string    `json:"alasan_pending_asuransi_mtr" gorm:"column:alasan_pending_asuransi_mtr" form:"alasan_pending_asuransi_mtr"`
	StsPembelian              string    `json:"sts_pembelian" gorm:"column:sts_pembelian" form:"sts_pembelian"`
	OTR                       uint32     `json:"otr" gorm:"column:otr" form:"otr"`
	Amount                    uint32     `json:"amount" gorm:"column:amount" form:"amount"`
	Warna                     string    `json:"warna" gorm:"column:warna" form:"warna"`
	NoRGK                     string    `json:"no_rgk" gorm:"column:no_rgk" form:"no_rgk"`
	NoPolFkt                  string    `json:"no_pol_fkt" gorm:"column:no_pol_fkt" form:"no_pol_fkt"`
	ThnMtr                    uint32    `json:"thn_mtr" gorm:"column:thn_mtr" form:"thn_mtr"`
	NmMtr                     string    `json:"nm_mtr" gorm:"column:nm_mtr" form:"nm_mtr"`
}

func (AsuransiMtr) TableName() string {
	return "asuransi_mtr"
}