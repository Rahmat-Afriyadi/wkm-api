package entity

type MstPendidikan struct {
	Id           string `form:"kd_pendidikan" json:"kd_pendidikan" gorm:"column:kd_pendidikan"`
	NmPendidikan string `form:"nm_pendidikan" json:"nm_pendidikan" gorm:"column:nm_pendidikan"`
}

func (MstPendidikan) TableName() string {
	return "mst_pendidikan"
}