package repository

import (
	"wkm/entity"

	"gorm.io/gorm"
)

type ApprovalRepository interface {
	Update(data entity.DetailApproval)
	MokitaToken() entity.MstToken
	MokitaUpdateToken(token string)
}

type approvalRepository struct {
	conn *gorm.DB
}

func NewApprovalRepository(conn *gorm.DB) ApprovalRepository {
	return &approvalRepository{
		conn: conn,
	}
}

func (lR *approvalRepository) MokitaToken() entity.MstToken {
	token := entity.MstToken{}
	lR.conn.Where("nm_user = ? ", "MOKITA").First(&token)
	return token
}

func (lR *approvalRepository) MokitaUpdateToken(token string) {
	var data entity.MstToken
	lR.conn.Where("nm_user = ? ", "MOKITA").First(&data)
	data.Token = token
	lR.conn.Save(&data)
	// data.Token = token
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
