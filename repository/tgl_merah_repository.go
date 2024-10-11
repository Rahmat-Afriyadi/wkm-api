package repository

import (
	"errors"
	"wkm/entity"
	"wkm/request"
	"wkm/utils"

	"gorm.io/gorm"
)

type TglMerahRepository interface {
	MasterData(search string, limit int, pageParams int) []entity.TglMerah
	MasterDataCount(search string) int64
	DetailTglMerah(id uint64) entity.TglMerah
	Create(data request.TglMerahRequest) (entity.TglMerah, error)
	Update(data request.TglMerahRequest) error
	UploadDokumen(data entity.TglMerah) error
	BulkCreate(data []entity.TglMerah) error
}

type tglMerahRepository struct {
	conn *gorm.DB
}

func NewTglMerahRepository(conn *gorm.DB) TglMerahRepository {
	return &tglMerahRepository{
		conn: conn,
	}
}

func (lR *tglMerahRepository) Create(data request.TglMerahRequest) (entity.TglMerah, error) {

	newTglMerah := entity.TglMerah{}
	return newTglMerah, nil

}

func (lR *tglMerahRepository) UploadDokumen(data entity.TglMerah) error {
	return nil

}

func (lR *tglMerahRepository) Update(data request.TglMerahRequest) error {

	return nil
}

func (lR *tglMerahRepository) MasterData(search string, limit int, pageParams int) []entity.TglMerah {
	datas := []entity.TglMerah{}
	query := lR.conn.Where("nik like ? or no_msn like ? or id_tglMerah like ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	query.Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Preload("MasterProduk").Preload("Konsumen").Find(&datas)
	return datas
}

func (lR *tglMerahRepository) MasterDataCount(search string) int64 {
	var datas []entity.TglMerah
	query := lR.conn.Where("nik like ? or no_msn like ? or id_tglMerah like ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	query.Select("id_tglMerah").Find(&datas)
	return int64(len(datas))
}

func (lR *tglMerahRepository) DetailTglMerah(id uint64) entity.TglMerah {
	tglMerah := entity.TglMerah{ID: id}
	lR.conn.Preload("MasterProduk").Preload("Konsumen").Preload("MstMtr").Find(&tglMerah)
	return tglMerah
}

func (lR *tglMerahRepository) BulkCreate(datas []entity.TglMerah) error {
	var exist entity.TglMerah
	for _, value := range datas {
		lR.conn.Where("tgl_awal", value.TglAwal.Format("2006-01-02")).Find(&exist)
		if exist.ID != 0 {
			return errors.New("tgl tersebut telah diinput " + value.TglAwal.Format("2006-01-02"))
		}
	}
	if len(datas) > 0 {
		tx := lR.conn.Begin()
		batchSize := 500
		for i := 0; i < len(datas); i += batchSize {
			end := i + batchSize
			if end > len(datas) {
				end = len(datas)
			}

			batch := datas[i:end]
			if err := tx.Table("tgl_merah").Create(&batch).Error; err != nil {
				tx.Rollback()
			}
		}
		tx.Commit()
	}

	return nil
}
