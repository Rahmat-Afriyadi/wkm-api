package entity

type MstTujuPak struct {
	Id      string `form:"id" json:"id" gorm:"column:id"`
	NmTupak string `form:"nm_tupak" json:"nm_tupak" gorm:"column:nm_tujpak"`
}

func (MstTujuPak) TableName() string {
	return "mst_tujuanpakai"
}
