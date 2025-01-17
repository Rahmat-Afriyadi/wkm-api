package entity

type MstAgama struct {
	Id    string `form:"kd_agama" json:"kd_agama" gorm:"column:kd_agama"`
	Agama string `form:"agama" json:"agama" gorm:"column:agama"`
}

func (MstAgama) TableName() string {
	return "mst_agama"
}