package entity

type MstAlasanTdkMembership struct {
	Id                  string `form:"id" json:"id" gorm:"column:id"`
	AlasanTdkMembership string `form:"alasan_tdk_renewal" json:"alasan_tdk_renewal" gorm:"column:alasan_tdk_renewal"`
	Alasan              string `form:"alasan" json:"alasan" gorm:"column:alasan"`
}

func (MstAlasanTdkMembership) TableName() string {
	return "mst_alasan_tidak_renewal_tele"
}