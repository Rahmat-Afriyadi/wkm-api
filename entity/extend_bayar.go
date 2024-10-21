package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ExtendBayar struct {
	Id        string `form:"id" json:"id" gorm:"primary_key;column:id"`
	NoMsn     string `form:"no_msn" json:"no_msn" gorm:"column:no_msn"`
	KdUserFa  string `form:"kd_user_fa" json:"kd_user_fa" gorm:"column:kd_user_fa"`
	KdUserLf  string `form:"kd_user_lf" json:"kd_user_lf" gorm:"column:kd_user_lf"`
	Deskripsi string `form:"deskripsi" json:"deskripsi" gorm:"column:deskripsi"`
	// P O R
	StsApproval    string     `form:"sts_approval" json:"sts_approval" gorm:"column:sts_approval"`
	TglPengajuan   time.Time  `form:"tgl_pengajuan" json:"tgl_pengajuan" gorm:"type:DATE;default:null;column:tgl_pengajuan"`
	TglActualBayar time.Time  `form:"tgl_actual_bayar" json:"tgl_actual_bayar" gorm:"type:DATE;default:null;column:tgl_actual_bayar"`
	TglUpdateLf    *time.Time `form:"tgl_update_lf" json:"tgl_update_lf" gorm:"type:DATE;default:null;column:tgl_update_lf"`
	TglUpdateFa    time.Time  `form:"tgl_update_fa" json:"tgl_update_fa" gorm:"type:DATE;default:null;column:tgl_update_fa"`
}

func (ExtendBayar) TableName() string {
	return "pengajuan_extend_bayar"
}

func (b *ExtendBayar) BeforeCreate(tx *gorm.DB) (err error) {
	b.Id = uuid.New().String()
	return
}
