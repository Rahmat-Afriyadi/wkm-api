package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AsuransiPA struct {
	Id        string  `form:"id" json:"id" gorm:"primary_key;column:id"`
	NoMSN                    string     `json:"no_msn" gorm:"column:no_msn" form:"no_msn"`
	NmCustomer               string    `json:"nm_customer" gorm:"column:nm_customer" form:"nm_customer"`
	StsAsuransiPA            string    `json:"sts_asuransi_pa" gorm:"column:sts_asuransi_pa" form:"sts_asuransi_pa"`
	StsBayar                 string    `json:"sts_bayar" gorm:"column:sts_bayar" form:"sts_bayar"`
	TglBayar                 *time.Time `json:"tgl_bayar" gorm:"column:tgl_bayar" form:"tgl_bayar"`
	TglInputBayar            *time.Time `json:"tgl_input_bayar" gorm:"column:tgl_input_bayar" form:"tgl_input_bayar"`
	KdUserFa                 string    `json:"kd_user_fa" gorm:"column:kd_user_fa" form:"kd_user_fa"`
	IDProduk                 string    `json:"id_produk" gorm:"column:id_produk" form:"id_produk"`
	Produk  				 MasterProduk `json:"produk" gorm:"->;references:KdProduk;foreignKey:IDProduk" form:"produk"`
	AmountPa                    uint64     `json:"amount_asuransi_pa" gorm:"column:amount" form:"amount_asuransi_pa"`
	AppTransID               string    `json:"app_trans_id" gorm:"column:app_trans_id" form:"app_trans_id"`
	TglBeli                  *time.Time `json:"tgl_beli" gorm:"column:tgl_beli" form:"tgl_beli"`
	NoKtpNpwp                string    `json:"no_ktpnpwp_fkt" gorm:"column:no_ktpnpwp" form:"no_ktpnpwp_fkt"`
	AlasanPendingAsuransiPA  string    `json:"alasan_pending_asuransi_pa" gorm:"column:alasan_pending_asuransi_pa" form:"alasan_pending_asuransi_pa"`
	StsPembelian             string    `json:"sts_pembelian" gorm:"column:sts_pembelian" form:"sts_pembelian"`
	KdUserTs              string     `gorm:"column:kd_user_ts;" json:"kd_user_ts" form:"kd_user_ts"`


}

func (AsuransiPA) TableName() string {
	return "asuransi_pa"
}

func (b *AsuransiPA) BeforeCreate(tx *gorm.DB) (err error) {
	b.Id = uuid.New().String()
	return
}