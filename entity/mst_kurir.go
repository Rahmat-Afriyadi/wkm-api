package entity

type MstKurir struct {
	KdKurir string `form:"kode_kurir" json:"kode_kurir" gorm:"primary_key;column:kode_kurir"`
	NmKurir string `form:"nama_kurir" json:"nama_kurir" gorm:"column:nama_kurir"`
}

func (MstKurir) TableName() string {
	return "mst_kurir"
}
