package entity

type MasterProduk struct {
	KdProduk string `json:"kd_produk" gorm:"column:id_produk"`
	NmProduk string `json:"nm_produk" gorm:"column:nm_produk"`
}

func (MasterProduk) TableName() string {
	return "produk"
}
