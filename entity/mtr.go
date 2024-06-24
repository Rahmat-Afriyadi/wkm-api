package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MstMtr struct {
	ID          string `form:"id" json:"id" gorm:"type:uuid;primary_key;column:id"`
	NoMtr       string `form:"no_mtr" json:"no_mtr" gorm:"column:no_mtr;"`
	KdMdl       string `form:"kd_mdl" json:"kd_mdl" gorm:"column:kd_mdl;"`
	ProductNama string `form:"nm_mtr" json:"nm_mtr" gorm:"column:nm_mtr"`
	CubSport    string `form:"cub_sport" json:"cub_sport" gorm:"column:cub_sport"`
	NmLap       string `form:"nm_lap" json:"nm_lap" gorm:"column:nm_lap"`
	LowHigh     string `form:"low_high" json:"low_high" gorm:"column:low_high"`
	MtrSeries   string `form:"mtr_series" json:"mtr_series" gorm:"column:mtr_series;"`
}

func (MstMtr) TableName() string {
	return "mst_mtr"
}

func (b *MstMtr) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}
