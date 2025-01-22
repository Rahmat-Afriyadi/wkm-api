package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type MstScript struct {
	Id			string		`form:"id" json:"id" gorm:"primary_key;column:id"`
	Title		string		`json:"title" gorm:"column:title" form:"title"`
	Script		string		`json:"script" gorm:"column:script" form:"script"`
	KdUser      string      `json:"kd_user" gorm:"column:kd_user" form:"kd_user"`
	IsActive    uint32		`json:"is_active" gorm:"column:is_active" form:"is_active"`  
	Created     *time.Time  `json:"created" gorm:"column:created" form:"created"`
	Modified    *time.Time `json:"modified" gorm:"column:modified" form:"modified"`
	ModiBy      string      `json:"modi_by" gorm:"column:modi_by" form:"modi_by"`
}

func (MstScript) TableName() string {
	return "mst_scripts"
}

func (b *MstScript) BeforeCreate(tx *gorm.DB) (err error) {
	b.Id = uuid.New().String()
	return
}