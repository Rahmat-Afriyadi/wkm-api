package repository

import (
	"wkm/entity"
	"wkm/request"
	"wkm/utils"

	"gorm.io/gorm"
)

type StockCardRepository interface {
	MasterData(search string, limit int, pageParams int) []entity.StockCard
	MasterDataCount(search string) int64
	DetailStockCard(no_kartu string) entity.StockCard
	Create(data request.StockCardRequest) (entity.StockCard, error)
	Update(data request.StockCardRequest) error
	UploadDokumen(data entity.StockCard) error
}

type stockCardRepository struct {
	conn *gorm.DB
}

func NewStockCardRepository(conn *gorm.DB) StockCardRepository {
	return &stockCardRepository{
		conn: conn,
	}
}

func (lR *stockCardRepository) Create(data request.StockCardRequest) (entity.StockCard, error) {

	newStockCard := entity.StockCard{}
	return newStockCard, nil

}

func (lR *stockCardRepository) UploadDokumen(data entity.StockCard) error {
	return nil

}

func (lR *stockCardRepository) Update(data request.StockCardRequest) error {

	return nil
}

func (lR *stockCardRepository) MasterData(search string, limit int, pageParams int) []entity.StockCard {
	datas := []entity.StockCard{}
	query := lR.conn.Where("nik like ? or no_msn like ? or id_stockCard like ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	query.Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Preload("MasterProduk").Preload("Konsumen").Find(&datas)
	return datas
}

func (lR *stockCardRepository) MasterDataCount(search string) int64 {
	var datas []entity.StockCard
	query := lR.conn.Where("nik like ? or no_msn like ? or id_stockCard like ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	query.Select("id_stockCard").Find(&datas)
	return int64(len(datas))
}

func (lR *stockCardRepository) DetailStockCard(no_kartu string) entity.StockCard {
	stockCard := entity.StockCard{NoKartu: no_kartu}
	lR.conn.Preload("MasterProduk").Preload("Konsumen").Preload("MstMtr").Find(&stockCard)
	return stockCard
}
