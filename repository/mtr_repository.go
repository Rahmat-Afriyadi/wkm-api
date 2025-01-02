package repository

import (
	"errors"
	"fmt"
	"wkm/entity"
	"wkm/utils"

	"gorm.io/gorm"
)

type MstMtrRepository interface {
	CreateMstMtr(data entity.MstMtr) error
	MasterData(search string, limit int, pageParams int) []entity.MstMtr
	MasterDataCount(search string) int64
	DetailMstMtr(id string) entity.MstMtr
	Update(body entity.MstMtr) error
}

type mstMtrRepository struct {
	conn *gorm.DB
}

func NewMstMtrRepository(conn *gorm.DB) MstMtrRepository {
	return &mstMtrRepository{
		conn: conn,
	}
}

func (lR *mstMtrRepository) DetailMstMtr(id string) entity.MstMtr {
	mstMtr := entity.MstMtr{ID: id}
	lR.conn.Find(&mstMtr)
	return mstMtr
}

func (lR *mstMtrRepository) CreateMstMtr(data entity.MstMtr) error {
	result := lR.conn.Create(&data)
	if result.Error != nil {
		fmt.Println("ini error ", result.Error)
		return result.Error
	} else {
		return nil
	}

}

func (lR *mstMtrRepository) Update(data entity.MstMtr) error {
	record := entity.MstMtr{ID: data.ID}
	lR.conn.Find(&record)
	if record.ProductNama == "" {
		return errors.New("data tidak ditemukan")
	}
	data.CreatedAt = nil
	result := lR.conn.Save(&data)
	if result.Error != nil {
		fmt.Println("ini error ", result.Error)
		return result.Error
	} else {
		return nil
	}
}

func (lR *mstMtrRepository) MasterData(search string, limit int, pageParams int) []entity.MstMtr {
	mstMtr := []entity.MstMtr{}
	query := lR.conn.Where("kd_mdl like ? or nm_mtr like ? or no_mtr like ? or merk like ? ", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	query.Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Find(&mstMtr)
	return mstMtr
}

func (lR *mstMtrRepository) MasterDataCount(search string) int64 {
	var mstMtr []entity.MstMtr
	query := lR.conn.Where("kd_mdl like ? or nm_mtr like ? or no_mtr like ? or merk like ? ", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	query.Select("no_mtr").Find(&mstMtr)
	return int64(len(mstMtr))
}
