package repository

import (
	"errors"
	"fmt"
	"wkm/entity"
	"wkm/utils"

	"gorm.io/gorm"
)

type ProdukRepository interface {
	MasterData(search string, jenis_asuransi int, limit int, pageParams int) []entity.MasterProduk
	MasterDataCount(search string, jenis_asuransi int) int64
	DetailProduk(id string) entity.MasterProduk
	Create(data entity.MasterProduk) error
	Update(data entity.MasterProduk) error
	DeleteManfaat(id string) error
	DeleteSyarat(id string) error
	DeletePaket(id string) error
}

type produkRepository struct {
	conn *gorm.DB
}

func NewProdukRepository(conn *gorm.DB) ProdukRepository {
	return &produkRepository{
		conn: conn,
	}
}

func (lR *produkRepository) Create(data entity.MasterProduk) error {
	result := lR.conn.Create(&data)
	if result.Error != nil {
		fmt.Println("ini error ", result.Error)
		return result.Error
	} else {
		return nil
	}

}

func (lR *produkRepository) Update(data entity.MasterProduk) error {
	record := entity.MasterProduk{KdProduk: data.KdProduk}

	lR.conn.Find(&record)
	if record.NmProduk == "" {
		return errors.New("data tidak ditemukan")
	}
	// result := lR.conn.Session(&gorm.Session{FullSaveAssociations: true}).Save(&produk)
	// if result.Error != nil {
	// 	fmt.Println("ini error yaa ", result.Error)
	// }
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

	fmt.Println("ini data semua yaa ", data)
	result := lR.conn.Session(&gorm.Session{FullSaveAssociations: true}).Save(&data)
	if result.Error != nil {
		fmt.Println("ini error ", result.Error)
		return result.Error
	} else {
		return nil
	}
}

func (lR *produkRepository) MasterData(search string, jenis_asuransi int, limit int, pageParams int) []entity.MasterProduk {
	datas := []entity.MasterProduk{}
	query := lR.conn.Where("nm_produk like ? or deskripsi like ?", "%"+search+"%", "%"+search+"%")
	if jenis_asuransi != 0 {
		lR.conn.Where("jns_asuransi = ?", jenis_asuransi)
	}
	query.Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Find(&datas)
	return datas
}

func (lR *produkRepository) MasterDataCount(search string, jenis_asuransi int) int64 {
	var datas []entity.MasterProduk
	query := lR.conn.Where("nm_produk like ? or deskripsi like ?", "%"+search+"%", "%"+search+"%")
	if jenis_asuransi != 0 {
		lR.conn.Where("jns_asuransi = ?", jenis_asuransi)
	}
	query.Select("id_produk").Find(&datas)
	return int64(len(datas))
}

func (lR *produkRepository) DetailProduk(id string) entity.MasterProduk {
	produk := entity.MasterProduk{KdProduk: id}
	lR.conn.Preload("Manfaat").Preload("Syarat").Preload("Paket").Find(&produk)
	// produk.Manfaats[5].Manfaat = "ini coba update yaa SEkali lagi yaa"
	// result := lR.conn.Session(&gorm.Session{FullSaveAssociations: true}).Save(&produk)
	// if result.Error != nil {
	// 	fmt.Println("ini error yaa ", result.Error)
	// }
	return produk
}

func (lR *produkRepository) DeleteManfaat(id string) error {
	result := lR.conn.Where("id_manfaat", id).Delete(&entity.Manfaat{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (lR *produkRepository) DeleteSyarat(id string) error {
	result := lR.conn.Where("id_syarat", id).Delete(&entity.Syarat{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (lR *produkRepository) DeletePaket(id string) error {
	result := lR.conn.Where("id_paket", id).Delete(&entity.Paket{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
