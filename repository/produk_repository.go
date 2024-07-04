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
	result := lR.conn.Save(&data)
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
	lR.conn.Find(&produk)
	fmt.Println("ini datanya yaa ", produk)
	return produk
}
