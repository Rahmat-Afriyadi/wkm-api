package repository

import (
	"fmt"
	"wkm/entity"
	"wkm/utils"

	"gorm.io/gorm"
)

type ProdukRepository interface {
	MasterData(search string, jenis_asuransi int, limit int, pageParams int) []entity.MasterProduk
	MasterDataCount(search string, jenis_asuransi int) int64
	DetailMstMtr(id string) entity.MasterProduk
}

type produkRepository struct {
	conn *gorm.DB
}

func NewProdukRepository(conn *gorm.DB) ProdukRepository {
	return &produkRepository{
		conn: conn,
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

func (lR *produkRepository) DetailMstMtr(id string) entity.MasterProduk {
	mstMtr := entity.MasterProduk{KdProduk: id}
	lR.conn.Find(&mstMtr)
	fmt.Println("ini datanya yaa ", mstMtr)
	return mstMtr
}
