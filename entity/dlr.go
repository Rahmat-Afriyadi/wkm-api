package entity

type MasterDlr struct {
	KdDlr string `json:"kd_dlr" gorm:"column:kd_dlr"`
	NmDlr string `json:"nm_dlr" gorm:"column:nm_dlr"`
}

func (MasterDlr) TableName() string {
	return "mst_dealer"
}
