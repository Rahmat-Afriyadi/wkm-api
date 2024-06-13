package repository

import (
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

	transaksi.Nik = data.Nik
	transaksi.NoRgk = data.NoRgk
	transaksi.NoPlat = data.NoPlat

	lR.conn.Save(&transaksi)

	konsumen.Nik = data.Nik
	konsumen.Nama = data.NamaKonsumen
	konsumen.NoHp = data.NoHp
	konsumen.Alamat = data.Alamat

	lR.conn.Save(&konsumen)
}
