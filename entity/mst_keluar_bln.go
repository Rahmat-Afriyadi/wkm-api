package entity

type MstKeluarBln struct {
	Id        string `form:"keluar_bln2" json:"keluar_bln2" gorm:"column:keluar_bln2"`
	KeluarBln string `form:"nm_keluar_bln2" json:"nm_keluar_bln2" gorm:"column:nm_keluar_bln2"`
}

func (MstKeluarBln) TableName() string {
	return "mst_keluar_bln"
}