package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"wkm/entity"
	"wkm/request"
	"wkm/utils"

	"gorm.io/gorm"
)

type TransaksiRepository interface {
	MasterData(search string, limit int, pageParams int) []entity.Transaksi
	MasterDataCount(search string) int64
	DetailTransaksi(id string) entity.Transaksi
	Create(data request.TransaksiRequest) (entity.Transaksi, error)
	CreateImport(data request.TransaksiRequest) (entity.Transaksi, error)
	Update(data request.TransaksiRequest) error
	UploadDokumen(data entity.Transaksi) error
	GenerateAppTransIdDealer(transaksi entity.Transaksi) string
}

type transaksiRepository struct {
	conn *gorm.DB
}

func NewTransaksiRepository(conn *gorm.DB) TransaksiRepository {
	return &transaksiRepository{
		conn: conn,
	}
}
func (lR *transaksiRepository) GenerateAppTransIdDealer(transaksi entity.Transaksi) string {

	splitId := strings.Split(transaksi.AppTransId, "-")
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

func (lR *transaksiRepository) Create(data request.TransaksiRequest) (entity.Transaksi, error) {
	existTransaksi := entity.Transaksi{}
	lR.conn.Where("nik = ? and id_produk = ?", data.Nik, data.IdProduk).Where("sts_pembelian = ? or sts_pembelian = ?", "1", "2").First(&existTransaksi)
	if existTransaksi.ID != "" {
		return entity.Transaksi{}, errors.New("Transaksi telah ada")
	}
	konsumen := entity.Konsumen{Nik: data.Nik}
	lR.conn.Find(&konsumen)
	konsumen.Nama = data.NmKonsumen
	konsumen.NoHp = data.NoHp
	konsumen.Email = data.Email
	konsumen.Alamat = data.Alamat
	konsumen.Kota1 = data.Kota
	konsumen.Kecamatan = data.Kecamatan
	konsumen.Kelurahan = data.Kelurahan
	konsumen.Kodepos = data.Kodepos
	konsumen.TglLahir = sql.NullString{String: data.TglLahir}
	lR.conn.Save(&konsumen)

	lastDealer := entity.Transaksi{}
	lR.conn.Where("payment_channel = ?", "DEALER").Last(&lastDealer)

	newTransaksi := entity.Transaksi{
		IdProduk:       data.IdProduk,
		NoMsn:          data.NoMsn,
		NoRgk:          data.NoRgk,
		Nik:            konsumen.Nik,
		NoPlat:         data.NoPlat,
		StsPembelian:   "1",
		Invoice:        "",
		PaymentId:      "",
		PaymentChannel: "DEALER",
		MotorPriceKode: data.KdMdl,
		Otr:            data.Otr,
		Amount:         data.Amount,
		Warna:          data.Warna,
		ReferralId:     "",
		ThnMtr:         data.Tahun,
		TglBeli:        time.Now().Format("2006-01-02"),
	}
	if lastDealer.ID == "" {
		newTransaksi.AppTransId = "DEALER-001"
	} else {
		newTransaksi.AppTransId = lR.GenerateAppTransIdDealer(lastDealer)
	}
	resultTrx := lR.conn.Create(&newTransaksi)
	if resultTrx.Error != nil {
		return entity.Transaksi{}, resultTrx.Error
	}
	return newTransaksi, nil

}
func (lR *transaksiRepository) CreateImport(data request.TransaksiRequest) (entity.Transaksi, error) {
	existTransaksi := entity.Transaksi{}
	lR.conn.Where("nik = ? and id_produk = ?", data.Nik, data.IdProduk).Where("sts_pembelian = ? or sts_pembelian = ?", "1", "2").First(&existTransaksi)
	if existTransaksi.ID != "" {
		return entity.Transaksi{}, errors.New("transaksi telah ada")
	}
	konsumen := entity.Konsumen{Nik: data.Nik}
	lR.conn.Find(&konsumen)
	konsumen.Nama = data.NmKonsumen
	konsumen.NoHp = data.NoHp
	konsumen.Email = data.Email
	konsumen.Alamat = data.Alamat
	konsumen.Kota1 = data.Kota
	konsumen.Kecamatan = data.Kecamatan
	konsumen.Kelurahan = data.Kelurahan
	konsumen.Kodepos = data.Kodepos
	konsumen.TglLahir = sql.NullString{String: data.TglLahir}
	lR.conn.Save(&konsumen)

	lastDealer := entity.Transaksi{}
	lR.conn.Where("payment_channel = ?", "DEALER").Last(&lastDealer)

	newTransaksi := entity.Transaksi{
		IdProduk:       data.IdProduk,
		NoMsn:          data.NoMsn,
		NoRgk:          data.NoRgk,
		Nik:            konsumen.Nik,
		NoPlat:         data.NoPlat,
		StsPembelian:   "1",
		Invoice:        "",
		PaymentId:      "",
		PaymentChannel: "DEALER",
		MotorPriceKode: data.KdMdl,
		Otr:            data.Otr,
		Amount:         data.Amount,
		Warna:          data.Warna,
		ReferralId:     "",
		ThnMtr:         data.Tahun,
		TglBeli:        time.Now().Format("2006-01-02"),
	}
	if lastDealer.ID == "" {
		newTransaksi.AppTransId = "DEALER-001"
	} else {
		newTransaksi.AppTransId = lR.GenerateAppTransIdDealer(lastDealer)
	}
	resultTrx := lR.conn.Create(&newTransaksi)
	if resultTrx.Error != nil {
		return entity.Transaksi{}, resultTrx.Error
	}
	return newTransaksi, nil

}

func (lR *transaksiRepository) UploadDokumen(data entity.Transaksi) error {
	record := entity.Transaksi{ID: data.ID}
	lR.conn.Find(&record)
	if record.NoMsn == "" {
		return errors.New("data tidak ditemukan")
	}
	record.Ktp = data.Ktp
	record.Stnk = data.Stnk
	result := lR.conn.Save(&record)
	if result.Error != nil {
		fmt.Println("ini error ", result.Error)
		return result.Error
	} else {
		return nil
	}

}

func (lR *transaksiRepository) Update(data request.TransaksiRequest) error {

	konsumen := entity.Konsumen{Nik: data.Nik}
	lR.conn.Find(&konsumen)
	konsumen.Nama = data.NmKonsumen
	konsumen.NoHp = data.NoHp
	konsumen.Email = data.Email
	konsumen.Alamat = data.Alamat
	konsumen.Kota1 = data.Kota
	konsumen.Kecamatan = data.Kecamatan
	konsumen.Kelurahan = data.Kelurahan
	konsumen.Kodepos = data.Kodepos
	konsumen.TglLahir = sql.NullString{String: data.TglLahir}
	lR.conn.Save(&konsumen)

	transaksi := entity.Transaksi{ID: data.IdTransaksi}
	lR.conn.Find(&transaksi)
	transaksi.IdProduk = data.IdProduk
	transaksi.NoMsn = data.NoMsn
	transaksi.NoRgk = data.NoRgk
	transaksi.Nik = konsumen.Nik
	transaksi.NoPlat = data.NoPlat
	transaksi.StsPembelian = "1"
	transaksi.Invoice = ""
	transaksi.PaymentId = ""
	transaksi.PaymentChannel = "DEALER"
	transaksi.MotorPriceKode = data.KdMdl
	transaksi.Otr = data.Otr
	transaksi.Amount = data.Amount
	transaksi.Warna = data.Warna
	transaksi.ReferralId = ""
	transaksi.ThnMtr = data.Tahun
	transaksi.TglBeli = time.Now().Format("2006-01-02")

	result := lR.conn.Save(&transaksi)
	if result.Error != nil {
		fmt.Println("ini error ", result.Error)
		return result.Error
	} else {
		return nil
	}
}

func (lR *transaksiRepository) MasterData(search string, limit int, pageParams int) []entity.Transaksi {
	datas := []entity.Transaksi{}
	query := lR.conn.Where("nik like ? or no_msn like ? or id_transaksi like ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	query.Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Preload("MasterProduk").Preload("Konsumen").Find(&datas)
	return datas
}

func (lR *transaksiRepository) MasterDataCount(search string) int64 {
	var datas []entity.Transaksi
	query := lR.conn.Where("nik like ? or no_msn like ? or id_transaksi like ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	query.Select("id_transaksi").Find(&datas)
	return int64(len(datas))
}

func (lR *transaksiRepository) DetailTransaksi(id string) entity.Transaksi {
	transaksi := entity.Transaksi{ID: id}
	lR.conn.Preload("MasterProduk").Preload("Konsumen").Preload("MstMtr").Find(&transaksi)
	return transaksi
}
