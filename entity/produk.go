package entity

type MasterProduk struct {
	KdProduk          string `json:"kd_produk" gorm:"primary_key;column:id_produk"`
	NmProduk          string `json:"nm_produk" gorm:"column:nm_produk"`
	NilaiPertanggunan uint64 `json:"nilai_pertanggungan" gorm:"column:nilai_pertanggungan"`
	Premi             uint64 `json:"premi" gorm:"column:premi"`
	Admin             uint64 `json:"admin" gorm:"column:admin"`
	JnsAsuransi       string `json:"jns_asuransi" gorm:"column:jns_asuransi"`
}

func (MasterProduk) TableName() string {
	return "produk"
}
