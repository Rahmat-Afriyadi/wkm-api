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
	MasterData(search string, jenis_asuransi int, limit int, pageParams int) []entity.Transaksi
	MasterDataCount(search string, jenis_asuransi int) int64
	DetailTransaksi(id string) entity.Transaksi
	Create(data request.TransaksiCreateRequest) (entity.Transaksi, error)
	Update(data entity.Transaksi) error
	UploadDokumen(data entity.Transaksi) error
	DeleteManfaat(id string) error
	DeleteSyarat(id string) error
	DeletePaket(id string) error
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

func (lR *transaksiRepository) Create(data request.TransaksiCreateRequest) (entity.Transaksi, error) {
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
	// lR.conn.Last(&lastManfaat)
	// if lastManfaat.IdManfaat == "" {
	// 	lastManfaat.IdManfaat = "MANFAAT-001"
	// }

}

func (lR *transaksiRepository) UploadDokumen(data entity.Transaksi) error {
	record := entity.Transaksi{ID: data.ID}

	lR.conn.Find(&record)
	if record.NoMsn == "" {
		return errors.New("data tidak ditemukan")
	}
	record.Ktp = data.Ktp
	result := lR.conn.Save(&record)
	if result.Error != nil {
		fmt.Println("ini error ", result.Error)
		return result.Error
	} else {
		return nil
	}

}

func (lR *transaksiRepository) Update(data entity.Transaksi) error {
	record := entity.Transaksi{ID: data.ID}

	lR.conn.Find(&record)
	if record.NoMsn == "" {
		return errors.New("data tidak ditemukan")
	}
	result := lR.conn.Session(&gorm.Session{FullSaveAssociations: true}).Save(&data)
	if result.Error != nil {
		fmt.Println("ini error ", result.Error)
		return result.Error
	} else {
		return nil
	}
}

func (lR *transaksiRepository) MasterData(search string, jenis_asuransi int, limit int, pageParams int) []entity.Transaksi {
	datas := []entity.Transaksi{}
	query := lR.conn.Where("nm_transaksi like ? or deskripsi like ?", "%"+search+"%", "%"+search+"%")
	if jenis_asuransi != 0 {
		lR.conn.Where("jns_asuransi = ?", jenis_asuransi)
	}
	query.Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Find(&datas)
	return datas
}

func (lR *transaksiRepository) MasterDataCount(search string, jenis_asuransi int) int64 {
	var datas []entity.Transaksi
	query := lR.conn.Where("nm_transaksi like ? or deskripsi like ?", "%"+search+"%", "%"+search+"%")
	if jenis_asuransi != 0 {
		lR.conn.Where("jns_asuransi = ?", jenis_asuransi)
	}
	query.Select("id_transaksi").Find(&datas)
	return int64(len(datas))
}

func (lR *transaksiRepository) DetailTransaksi(id string) entity.Transaksi {
	transaksi := entity.Transaksi{ID: id}
	lR.conn.Preload("Manfaat").Preload("Syarat").Preload("Paket").Find(&transaksi)
	return transaksi
}

func (lR *transaksiRepository) DeleteManfaat(id string) error {
	result := lR.conn.Where("id_manfaat", id).Delete(&entity.Manfaat{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (lR *transaksiRepository) DeleteSyarat(id string) error {
	result := lR.conn.Where("id_syarat", id).Delete(&entity.Syarat{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (lR *transaksiRepository) DeletePaket(id string) error {
	result := lR.conn.Where("id_paket", id).Delete(&entity.Paket{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
