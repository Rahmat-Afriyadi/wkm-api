package entity

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type MasterProduk struct {
	KdProduk           string       `json:"kd_produk" gorm:"primary_key;column:id_produk"`
	NmProduk           string       `json:"nm_produk" gorm:"column:nm_produk"`
	Deskripsi          string       `json:"deskripsi" gorm:"column:deskripsi"`
	NilaiPertanggungan uint64       `json:"nilai_pertanggungan" gorm:"column:nilai_pertanggungan"`
	Premi              uint64       `json:"premi" gorm:"column:premi"`
	Admin              uint64       `json:"admin" gorm:"column:admin"`
	JnsAsuransi        string       `json:"jns_asuransi" gorm:"column:jns_asuransi"`
	VendorId           string       `json:"vendor_id" form:"vendor_id" gorm:"column:vendor_id"`
	Vendor             MasterVendor `json:"vendor" gorm:"->;references:KdVendor;foreignKey:VendorId"`
	CreatedAt          *time.Time   `form:"created_at" json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt          *time.Time   `form:"updated_at" json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (MasterProduk) TableName() string {
	return "produk"
}

func (u *MasterProduk) BeforeCreate(tx *gorm.DB) (err error) {
	lastTransaksi := MasterProduk{}
	tx.Last(&lastTransaksi)
	u.KdProduk = GenerateIdProduk(lastTransaksi)

	return
}

// func (u *MasterProduk) BeforeUpdate(tx *gorm.DB) (err error) {
// 	minDate, _ := time.Parse("2006-01-2", "1970-01-01")
// 	if u.CreatedAt.Before(minDate) {
// 		u.CreatedAt = nil
// 	}
// 	return
// }

func GenerateIdProduk(produk MasterProduk) string {

	splitId := strings.Split(produk.KdProduk, "-")
	i, err := strconv.Atoi(splitId[1])
	if err != nil {
		fmt.Println("ini error parse string to int ", err)
	}
	i += 1
	idProduk := ""
	if i > 99 {
		idProduk = fmt.Sprintf("%s%d", splitId[0]+"-", i)
	} else if i > 9 {
		idProduk = fmt.Sprintf("%s%d", splitId[0]+"-0", i)
	} else {
		idProduk = fmt.Sprintf("%s%d", splitId[0]+"-00", i)
	}
	return idProduk

}
