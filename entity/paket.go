package entity

import (
	"fmt"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type Paket struct {
	IdPaket  string `form:"id_paket" json:"id_paket" gorm:"primary_key;column:id_paket"`
	IdProduk string `form:"id_produk" json:"id_produk" gorm:"column:id_produk"`
	Paket    string `form:"paket" json:"paket" gorm:"column:nm_paket"`
	// Produk   MasterProduk `json:"produk" gorm:"->;references:KdProduk;foreignKey:IdProduk"`
}

func (Paket) TableName() string {
	return "paket"
}

func (u *Paket) BeforeCreate(tx *gorm.DB) (err error) {
	if u.IdPaket != "" {
		return
	}
	lastTransaksi := Paket{}
	tx.Last(&lastTransaksi)
	if lastTransaksi.IdPaket != "" {
		u.IdPaket = GenerateIdPaket(lastTransaksi)
	} else {
		u.IdPaket = "PAKET-001"
	}
	return
}

// func (u *Paket) BeforeUpdate(tx *gorm.DB) (err error) {
// 	minDate, _ := time.Parse("2006-01-2", "1970-01-01")
// 	if u.CreatedAt.Before(minDate) {
// 		u.CreatedAt = nil
// 	}
// 	return
// }

func GenerateIdPaket(paket Paket) string {

	splitId := strings.Split(paket.IdPaket, "-")
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
