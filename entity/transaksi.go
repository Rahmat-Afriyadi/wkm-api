package entity

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Transaksi struct {
	ID             string    `form:"id" json:"id" gorm:"primary_key;column:id_transaksi"`
	IdProduk       string    `form:"id_produk" json:"id_produk" gorm:"column:id_produk"`
	NoMsn          string    `form:"no_msn" json:"no_msn" gorm:"column:no_msn"`
	NoRgk          string    `form:"no_rgk" json:"no_rgk" gorm:"column:no_rgk"`
	Nik            string    `form:"nik" json:"nik" gorm:"column:nik"`
	NoPlat         string    `form:"no_plat" json:"no_plat" gorm:"column:no_plat"`
	AppTransId     string    `form:"app_trans_id" json:"app_trans_id" gorm:"column:app_trans_id"`
	StsPembelian   string    `form:"sts_pembelian" json:"sts_pembelian" gorm:"column:sts_pembelian"`
	Invoice        string    `form:"invoice" json:"invoice" gorm:"column:invoice"`
	PaymentId      string    `form:"payment_id" json:"payment_id" gorm:"column:payment_id"`
	PaymentChannel string    `form:"payment_channel" json:"payment_channel" gorm:"column:payment_channel"`
	MotorPriceKode string    `form:"motorprice_kode" json:"motorprice_kode" gorm:"column:motorprice_kode"`
	Otr            int       `form:"otr" json:"otr" gorm:"column:otr"`
	Amount         int       `form:"amount" json:"amount" gorm:"column:amount"`
	Warna          string    `form:"warna" json:"warna" gorm:"column:warna"`
	ReferralId     string    `form:"referral_id" json:"referral_id" gorm:"column:referral_id"`
	ThnMtr         string    `form:"thn_mtr" json:"thn_mtr" gorm:"column:thn_mtr"`
	TglBeli        string    `gorm:"column:tgl_beli"`
	Created        time.Time `gorm:"default:current_timestamp;column:created"`
	Updated        time.Time `gorm:"default:current_timestamp;column:updated"`
}

func (Transaksi) TableName() string {
	return "transaksi"
}

func (u *Transaksi) BeforeCreate(tx *gorm.DB) (err error) {
	lastTransaksi := Transaksi{}
	tx.Last(&lastTransaksi)
	u.ID = GenerateIdTransaksi(lastTransaksi)
	u.Created = time.Now()

	return
}

func GenerateIdTransaksi(transaksi Transaksi) string {

	i, err := strconv.Atoi(strings.Split(transaksi.ID, "TRN")[1])
	if err != nil {
		fmt.Println("ini error parse string to int ", err)
	}
	i += 1
	idProduk := ""
	if i > 99 {
		idProduk = fmt.Sprintf("%s%d", "TRN", i)
	} else if i > 9 {
		idProduk = fmt.Sprintf("%s%d", "TRN0", i)
	} else {
		idProduk = fmt.Sprintf("%s%d", "TRN00", i)
	}
	return idProduk

}
