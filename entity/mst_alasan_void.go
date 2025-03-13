package entity

type MstAlasanVoidKonfirmasi struct {
	Id                   string `form:"id" json:"id" gorm:"column:id"`
	AlasanVoidKonfirmasi string `form:"alasan_void_konfirmasi" json:"alasan_void_konfirmasi" gorm:"column:alasan_void_konfirmasi"`
	Alasan               string `form:"alasan" json:"alasan" gorm:"column:alasan"`
}

func (MstAlasanVoidKonfirmasi) TableName() string {
	return "mst_alasan_void_konfirmasi"
}
