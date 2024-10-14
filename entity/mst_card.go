package entity

type MstCard struct {
	KdCard  string `form:"kd_card" json:"kd_card" gorm:"primary_key;column:kd_card"`
	JnsCard string `form:"jns_card" json:"jns_card" gorm:"column:jns_card"`
}

func (MstCard) TableName() string {
	return "mst_card"
}
