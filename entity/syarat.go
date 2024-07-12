package entity

import (
	"fmt"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type Syarat struct {
	IdSyarat string `form:"id_syarat" json:"id_syarat" gorm:"primary_key;column:id_syarat"`
	IdProduk string `form:"id_produk" json:"id_produk" gorm:"column:id_produk"`
	Syarat   string `form:"syarat" json:"syarat" gorm:"column:syarat"`
	// Produk   MasterProduk `json:"produk" gorm:"->;references:KdProduk;foreignKey:IdProduk"`
}

func (Syarat) TableName() string {
	return "syarat"
}

func (u *Syarat) BeforeCreate(tx *gorm.DB) (err error) {
	lastTransaksi := Syarat{}
	tx.Last(&lastTransaksi)
	u.IdSyarat = GenerateIdSyarat(lastTransaksi)

	return
}

// func (u *Syarat) BeforeUpdate(tx *gorm.DB) (err error) {
// 	minDate, _ := time.Parse("2006-01-2", "1970-01-01")
// 	if u.CreatedAt.Before(minDate) {
// 		u.CreatedAt = nil
// 	}
// 	return
// }

func GenerateIdSyarat(syarat Syarat) string {

	splitId := strings.Split(syarat.IdSyarat, "-")
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
