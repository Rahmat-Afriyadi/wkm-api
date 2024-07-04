package entity

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type MasterVendor struct {
	KdVendor  string    `json:"kd_vendor" gorm:"primary_key;column:id_vendor"`
	NmVendor  string    `json:"nm_vendor" gorm:"column:nm_vendor"`
	Deskripsi string    `json:"deskripsi" gorm:"column:deskripsi"`
	Admin     uint64    `json:"admin" gorm:"column:admin"`
	Dealer    float64   `json:"dealer" gorm:"column:dealer"`
	Wkm       float64   `json:"wkm" gorm:"column:wkm"`
	CreatedAt time.Time `form:"created_at" json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `form:"updated_at" json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (MasterVendor) TableName() string {
	return "vendor"
}

func (u *MasterVendor) BeforeCreate(tx *gorm.DB) (err error) {
	lastTransaksi := MasterVendor{}
	tx.Last(&lastTransaksi)
	u.KdVendor = GenerateIdVendor(lastTransaksi)

	return
}

func GenerateIdVendor(vendor MasterVendor) string {

	splitId := strings.Split(vendor.KdVendor, "-")
	i, err := strconv.Atoi(splitId[1])
	if err != nil {
		fmt.Println("ini error parse string to int ", err)
	}
	i += 1
	idVendor := ""
	if i > 99 {
		idVendor = fmt.Sprintf("%s%d", splitId[0]+"-", i)
	} else if i > 9 {
		idVendor = fmt.Sprintf("%s%d", splitId[0]+"-0", i)
	} else {
		idVendor = fmt.Sprintf("%s%d", splitId[0]+"-00", i)
	}
	return idVendor

}
