package repository

import (
	"errors"
	"fmt"
	"wkm/entity"
	"wkm/request"
	"wkm/utils"

	"gorm.io/gorm"
)

type TransaksiRepository interface {
	MasterData(search string, jenis_asuransi int, limit int, pageParams int) []entity.MasterProduk
	MasterDataCount(search string, jenis_asuransi int) int64
	DetailTransaksi(id string) entity.MasterProduk
	Create(data request.TransaksiCreateRequest) (entity.Transaksi, error)
	Update(data entity.MasterProduk) error
	UploadLogo(data entity.MasterProduk) error
	DeleteManfaat(id string) error
	DeleteSyarat(id string) error
	DeletePaket(id string) error
}

type transaksiRepository struct {
	conn *gorm.DB
}

func NewTransaksiRepository(conn *gorm.DB) TransaksiRepository {
	return &transaksiRepository{
		conn: conn,
	}
}

func (lR *transaksiRepository) Create(data request.TransaksiCreateRequest) (entity.Transaksi, error) {

	lastManfaat := entity.Transaksi{}
	return lastManfaat, nil
	// lR.conn.Last(&lastManfaat)
	// if lastManfaat.IdManfaat == "" {
	// 	lastManfaat.IdManfaat = "MANFAAT-001"
	// }

}

func (lR *transaksiRepository) UploadLogo(data entity.MasterProduk) error {
	record := entity.MasterProduk{KdProduk: data.KdProduk}

	lR.conn.Find(&record)
	if record.NmProduk == "" {
		return errors.New("data tidak ditemukan")
	}
	record.Logo = data.Logo
	fmt.Println("harus kesini lah yaa ", record)
	result := lR.conn.Save(&record)
	if result.Error != nil {
		fmt.Println("ini error ", result.Error)
		return result.Error
	} else {
		return nil
	}

}

func (lR *transaksiRepository) Update(data entity.MasterProduk) error {
	record := entity.MasterProduk{KdProduk: data.KdProduk}

	lR.conn.Find(&record)
	if record.NmProduk == "" {
		return errors.New("data tidak ditemukan")
	}

	lastManfaat := entity.Manfaat{}
	lR.conn.Last(&lastManfaat)
	if lastManfaat.IdManfaat == "" {
		lastManfaat.IdManfaat = "MANFAAT-001"
	}
	for i, v := range data.Manfaat {
		if v.IdManfaat == "" {
			lastManfaat.IdManfaat = entity.GenerateIdManfaat(lastManfaat)
			data.Manfaat[i].IdManfaat = lastManfaat.IdManfaat
		}
	}

	lastSyarat := entity.Syarat{}
	lR.conn.Last(&lastSyarat)
	if lastSyarat.IdSyarat == "" {
		lastSyarat.IdSyarat = "SYARAT-001"
	}
	for i, v := range data.Syarat {
		if v.IdSyarat == "" {
			lastSyarat.IdSyarat = entity.GenerateIdSyarat(lastSyarat)
			data.Syarat[i].IdSyarat = lastSyarat.IdSyarat
		}
	}

	lastPaket := entity.Paket{}
	lR.conn.Last(&lastPaket)
	if lastPaket.IdPaket == "" {
		lastPaket.IdPaket = "PAKET-001"
	}
	for i, v := range data.Paket {
		if v.IdPaket == "" {
			lastPaket.IdPaket = entity.GenerateIdPaket(lastPaket)
			data.Paket[i].IdPaket = lastPaket.IdPaket
		}
	}
	if data.Logo == "" {
		data.Logo = record.Logo
	}
	result := lR.conn.Session(&gorm.Session{FullSaveAssociations: true}).Save(&data)
	if result.Error != nil {
		fmt.Println("ini error ", result.Error)
		return result.Error
	} else {
		return nil
	}
}

func (lR *transaksiRepository) MasterData(search string, jenis_asuransi int, limit int, pageParams int) []entity.MasterProduk {
	datas := []entity.MasterProduk{}
	query := lR.conn.Where("nm_transaksi like ? or deskripsi like ?", "%"+search+"%", "%"+search+"%")
	if jenis_asuransi != 0 {
		lR.conn.Where("jns_asuransi = ?", jenis_asuransi)
	}
	query.Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Find(&datas)
	return datas
}

func (lR *transaksiRepository) MasterDataCount(search string, jenis_asuransi int) int64 {
	var datas []entity.MasterProduk
	query := lR.conn.Where("nm_transaksi like ? or deskripsi like ?", "%"+search+"%", "%"+search+"%")
	if jenis_asuransi != 0 {
		lR.conn.Where("jns_asuransi = ?", jenis_asuransi)
	}
	query.Select("id_transaksi").Find(&datas)
	return int64(len(datas))
}

func (lR *transaksiRepository) DetailTransaksi(id string) entity.MasterProduk {
	transaksi := entity.MasterProduk{KdProduk: id}
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
