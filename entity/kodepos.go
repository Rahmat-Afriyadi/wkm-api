package entity

type MasterKodepos struct {
	KdPos     string `json:"kd_pos" gorm:"column:kd_pos"`
	KodePos   string `json:"kodepos" gorm:"column:kodepos"`
	Kelurahan string `json:"kelurahan" gorm:"column:kelurahan"`
	Kecamatan string `json:"kecamatan" gorm:"column:kecamatan"`
}

func (MasterKodepos) TableName() string {
	return "kodepos"
}
