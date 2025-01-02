package entity

type MstCard struct {
	KdCard        string `form:"kd_card" json:"kd_card" gorm:"primary_key;column:kd_card"`
	JnsCard       string `form:"jns_card" json:"jns_card" gorm:"column:jns_card"`
	HargaPokok    uint64 `form:"harga_pokok" json:"harga_pokok" gorm:"column:harga_pokok"`
	Asuransi      uint64 `form:"asuransi" json:"asuransi" gorm:"column:asuransi"`
	AsuransiMotor uint64 `form:"asuransi_motor" json:"asuransi_motor" gorm:"column:asuransi_motor"`
}

func (MstCard) TableName() string {
	return "mst_card"
}
