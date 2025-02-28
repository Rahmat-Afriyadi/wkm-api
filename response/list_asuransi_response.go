package response

import "time"

type ListAsuransi struct {
	NoMsn               string     `form:"no_msn" json:"no_msn" gorm:"primary_key;column:no_msn"`
	NmCustomerWkm          string     `form:"nm_customer_wkm" json:"nm_customer_wkm" gorm:"column:nm_customer_wkm"`
	NmCustomerFkt          string     `form:"nm_customer_fkt" json:"nm_customer_fkt" gorm:"column:nm_customer_fkt"`
	TglBeli        *time.Time `form:"tgl_beli" json:"tgl_beli" gorm:"column:tgl_beli"`
	StsAsuransi        string     `form:"sts_asuransi" json:"sts_asuransi" gorm:"column:sts_asuransi"`
	IdProduk 				string `form:"id_produk" json:"id_produk" gorm:"column:id_produk"`
}