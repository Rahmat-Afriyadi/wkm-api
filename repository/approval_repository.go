package repository

import (
	"fmt"
	"wkm/entity"

	"gorm.io/gorm"
)

type ApprovalRepository interface {
	Update(data entity.DetailApproval)
}

type approvalRepository struct {
	conn *gorm.DB
}

func NewApprovalRepository(conn *gorm.DB) ApprovalRepository {
	return &approvalRepository{
		conn: conn,
	}
}

func (lR *approvalRepository) Update(data entity.DetailApproval) {
	fmt.Println("iini data update ", data.Status, data.StatusApprove)
	transaksi := entity.Transaksi{ID: data.IdTransaksi}
	lR.conn.Find(&transaksi)
	if transaksi.AppTransId == "" {
		return
	}
	konsumen := entity.Konsumen{Nik: transaksi.Nik}
	lR.conn.Find(&konsumen)
	if konsumen.Nama == "" {
		return
	}
	if data.StatusApprove == "1" && data.Status == "1" {
		transaksi.StsPembelian = "2"
	} else if data.StatusApprove == "0" {
		transaksi.StsPembelian = "4"
	}
	transaksi.Nik = data.Nik
	transaksi.NoRgk = data.NoRgk
	transaksi.NoPlat = data.NoPlat

	lR.conn.Save(&transaksi)

	konsumen.Nik = data.Nik
	konsumen.Nama = data.NamaKonsumen
	konsumen.NoHp = data.NoHp
	konsumen.Alamat = data.Alamat
	if data.Province != nil {
		konsumen.Prop = *data.Province
	}
	if data.City != nil {
		konsumen.Kota = *data.City
	}
	if data.Subdistrict != nil {
		konsumen.Kec = *data.Subdistrict
	}

	lR.conn.Save(&konsumen)
}
