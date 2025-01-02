package entity

import (
	"fmt"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type Manfaat struct {
	IdManfaat string `form:"id_manfaat" json:"id_manfaat" gorm:"primary_key;column:id_manfaat"`
	IdProduk  string `form:"id_produk" json:"id_produk" gorm:"column:id_produk"`
	Manfaat   string `form:"manfaat" json:"manfaat" gorm:"column:manfaat"`
	// Produk    MasterProduk `json:"produk" gorm:"->;references:KdProduk;foreignKey:IdProduk"`
}

func (Manfaat) TableName() string {
	return "manfaat"
}
func (u *Manfaat) BeforeCreate(tx *gorm.DB) (err error) {
	if u.IdManfaat != "" {
		return
	}
	lastTransaksi := Manfaat{}
	tx.Last(&lastTransaksi)
	if lastTransaksi.IdManfaat != "" {
		u.IdManfaat = GenerateIdManfaat(lastTransaksi)
	} else {
		u.IdManfaat = "MANFAAT-001"
	}
	return
}

// func (u *Manfaat) BeforeUpdate(tx *gorm.DB) (err error) {
// 	minDate, _ := time.Parse("2006-01-2", "1970-01-01")
// 	if u.CreatedAt.Before(minDate) {
// 		u.CreatedAt = nil
// 	}
// 	return
// }

func GenerateIdManfaat(manfaat Manfaat) string {

	splitId := strings.Split(manfaat.IdManfaat, "-")
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
