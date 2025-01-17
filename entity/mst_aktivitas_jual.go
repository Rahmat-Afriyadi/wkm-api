package entity

type MstAktivitasJual struct {
	Id            string `form:"aktif_jual_r" json:"aktif_jual_r" gorm:"column:aktif_jual_r"`
	AktivitasJual string `form:"nm_aktif_jual_r" json:"nm_aktif_jual_r" gorm:"column:nm_aktif_jual_r"`
}

func (MstAktivitasJual) TableName() string {
	return "mst_aktifjual_renewal"
}