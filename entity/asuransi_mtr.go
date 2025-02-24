package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AsuransiMtr struct {
	Id        string  `form:"id" json:"id" gorm:"primary_key;column:id"`

	NoMSN                     string     `json:"no_msn" gorm:"column:no_msn" form:"no_msn"`
	NmCustomer                string    `json:"nm_customer" gorm:"column:nm_customer" form:"nm_customer"`
	StsAsuransiMtr            string    `json:"sts_asuransi_mtr" gorm:"column:sts_asuransi_mtr" form:"sts_asuransi_mtr"`
	StsBayar                 string    `json:"sts_bayar" gorm:"column:sts_bayar" form:"sts_bayar"`
	TglBayar                 *time.Time `json:"tgl_bayar" gorm:"column:tgl_bayar" form:"tgl_bayar"`
	TglInputBayar            *time.Time `json:"tgl_input_bayar" gorm:"column:tgl_input_bayar" form:"tgl_input_bayar"`
	KdUserFa                 string    `json:"kd_user_fa" gorm:"column:kd_user_fa" form:"kd_user_fa"`
	IDProduk                  string    `json:"id_produk" gorm:"column:id_produk" form:"id_produk"`
	Produk  				 MasterProduk `json:"produk" gorm:"->;references:KdProduk;foreignKey:IDProduk" form:"produk"`
	AppTransID                string    `json:"app_trans_id" gorm:"column:app_trans_id" form:"app_trans_id"`
	TglBeli                   *time.Time `json:"tgl_beli" gorm:"column:tgl_beli" form:"tgl_beli"`
	NoKtpNpwp                 string    `json:"no_ktpnpwp_fkt" gorm:"column:no_ktpnpwp" form:"no_ktpnpwp_fkt"`
	AlasanPendingAsuransiMtr  string    `json:"alasan_pending_asuransi_mtr" gorm:"column:alasan_pending_asuransi_mtr" form:"alasan_pending_asuransi_mtr"`
	StsPembelian              string    `json:"sts_pembelian" gorm:"column:sts_pembelian" form:"sts_pembelian"`
	OTR                       uint64     `json:"asuransi_mtr_otr" gorm:"column:otr" form:"asuransi_mtr_otr"`
	Amount                    uint64     `json:"asuransi_mtr_amount" gorm:"column:amount" form:"asuransi_mtr_amount"`
	Warna                     string    `json:"warna" gorm:"column:warna" form:"warna"`
	NoRGK                     string    `json:"no_rgk" gorm:"column:no_rgk" form:"no_rgk"`
	NoPolwKM              string     `gorm:"column:no_pol_wkm;" json:"no_pol_wkm" form:"no_pol_wkm"`
	ThnMtr                    uint32    `json:"thn_mtr" gorm:"column:thn_mtr" form:"thn_mtr"`
	NmMtr                     string    `json:"nm_mtr" gorm:"column:nm_mtr" form:"nm_mtr"`
	KdUserTs              string     `gorm:"column:kd_user_ts;" json:"kd_user_ts" form:"kd_user_ts"`

}

func (AsuransiMtr) TableName() string {
	return "asuransi_mtr"
}

func (b *AsuransiMtr) BeforeCreate(tx *gorm.DB) (err error) {
	b.Id = uuid.New().String()
	return
}
