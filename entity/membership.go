package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Membership struct {
	Id                      string     `json:"id" gorm:"column:id" form:"id"`
	NoMSN                    string     `json:"no_msn" gorm:"column:no_msn" form:"no_msn"`
	StsMembership            string    `json:"sts_membership" gorm:"column:sts_membership" form:"sts_membership"`
	
	TypeKartu                string    `json:"type_kartu" gorm:"column:type_kartu" form:"type_kartu"`
	JnsMembership            string     `json:"jns_membership" gorm:"column:jns_membership" form:"jns_membership"`
	TglJanjiBayar            *time.Time `json:"tgl_janji_bayar" gorm:"column:tgl_janji_bayar" form:"tgl_janji_bayar"`
	KirimKe                  string     `json:"kirim_ke" gorm:"column:kirim_ke" form:"kirim_ke"`
	JnsBayar                 string    `json:"jns_bayar" gorm:"column:jns_bayar" form:"jns_bayar"`
	KdPromoTransfer          string     `json:"kd_promo_transfer" gorm:"column:kd_promo_transfer" form:"kd_promo_transfer"`
	CetakTtKe                uint32    `json:"cetak_tt_ke" gorm:"column:cetak_tt_ke" form:"cetak_tt_ke"`
	NoTandaTerima            string    `json:"no_tanda_terima" gorm:"column:no_tanda_terima" form:"no_tanda_terima"`
	TglCetakTandaTerima      *time.Time `json:"tgl_cetak_tanda_terima" gorm:"column:tgl_cetak_tanda_terima" form:"tgl_cetak_tanda_terima"`
	KdUserCetakTt            string    `json:"kd_user_cetak_tt" gorm:"column:kd_user_cetak_tt" form:"kd_user_cetak_tt"`
	StsKartu                 string     `json:"sts_kartu" gorm:"column:sts_kartu" form:"sts_kartu"`
	RenewalKe                uint32     `json:"renewal_ke" gorm:"column:renewal_ke" form:"renewal_ke"`
	NoKartu                  string    `json:"no_kartu" gorm:"column:no_kartu" form:"no_kartu"`
	TglExpired               *time.Time `json:"tgl_expired" gorm:"column:tgl_expired" form:"tgl_expired"`
	KodeKurir                string    `json:"kode_kurir" gorm:"column:kode_kurir" form:"kode_kurir"`
	TglAmbilKartu            *time.Time `json:"tgl_ambil_kartu" gorm:"column:tgl_ambil_kartu" form:"tgl_ambil_kartu"`
	KdUserTs              string     `gorm:"column:kd_user_ts;" json:"kd_user_ts" form:"kd_user_ts"`
	KdUserBarcodeBawa        string    `json:"kd_user_barcode_bawa" gorm:"column:kd_user_barcode_bawa" form:"kd_user_barcode_bawa"`
	StsBawaKartu             string     `json:"sts_bawa_kartu" gorm:"column:sts_bawa_kartu" form:"sts_bawa_kartu"`
	StsBayar                 string    `json:"sts_bayar" gorm:"column:sts_bayar" form:"sts_bayar"`
	TglBayar                 *time.Time `json:"tgl_bayar" gorm:"column:tgl_bayar" form:"tgl_bayar"`
	TglInputBayar            *time.Time `json:"tgl_input_bayar" gorm:"column:tgl_input_bayar" form:"tgl_input_bayar"`
	KdUserFa                 string    `json:"kd_user_fa" gorm:"column:kd_user_fa" form:"kd_user_fa"`
	AlasanTdkKurir           string    `json:"alasan_tdk_kurir" gorm:"column:alasan_tdk_kurir" form:"alasan_tdk_kurir"`
	TglKembaliKartu          *time.Time `json:"tgl_kembali_kartu" gorm:"column:tgl_kembali_kartu" form:"tgl_kembali_kartu"`
	KdUserBarcodeKembali     string    `json:"kd_user_barcode_kembali" gorm:"column:kd_user_barcode_kembali" form:"kd_user_barcode_kembali"`
	KdUserCheckKembali       string    `json:"kd_user_check_kembali" gorm:"column:kd_user_check_kembali" form:"kd_user_check_kembali"`
	TglUpdateKartuBalikan    *time.Time `json:"tgl_update_kartu_balikan" gorm:"column:tgl_update_kartu_balikan" form:"tgl_update_kartu_balikan"`
	AlasanTdkKurirDetail     string     `json:"alasan_tdk_kurir_detail" gorm:"column:alasan_tdk_kurir_detail" form:"alasan_tdk_kurir_detail"`
	AlasanTdkTsDetail        string    `json:"alasan_tdk_ts_detail" gorm:"column:alasan_tdk_ts_detail" form:"alasan_tdk_ts_detail"`
	AlasanVoidKonfirmasi     string     `json:"alasan_void_konfirmasi" gorm:"column:alasan_void_konfirmasi" form:"alasan_void_konfirmasi"`
	CreatedAt      *time.Time `form:"created_at" json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      *time.Time `form:"updated_at" json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`


	CustomerMtr               CustomerMtr  `form:"customer_mtr" json:"customer_mtr" gorm:"->;references:NoMsn;foreignKey:NoMSN"`
}

func (Membership) TableName() string {
	return "membership"
}
func (b *Membership) BeforeCreate(tx *gorm.DB) (err error) {
	b.Id = uuid.New().String()
	return
}
