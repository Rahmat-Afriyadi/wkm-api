package repository

import (
	"errors"
	"time"
	"wkm/entity"
	"wkm/request"
	"wkm/utils"

	"gorm.io/gorm"
)

type ExtendBayarRepository interface {
	MasterData(search string, limit int, pageParams int) []entity.ExtendBayar
	MasterDataCount(search string) int64
	DetailExtendBayar(id string) entity.ExtendBayar
	Create(data request.ExtendBayarRequest) (entity.ExtendBayar, error)
	UpdateFa(data request.ExtendBayarRequest) error
	UpdateLf(data request.ExtendBayarRequest) error
	Delete(id string) error
}

type extendBayarRepository struct {
	conn *gorm.DB
}

func NewExtendBayarRepository(conn *gorm.DB) ExtendBayarRepository {
	return &extendBayarRepository{
		conn: conn,
	}
}

func (lR *extendBayarRepository) Create(data request.ExtendBayarRequest) (entity.ExtendBayar, error) {

	newExtendBayar := entity.ExtendBayar{
		NoMsn:          data.NoMsn,
		KdUserFa:       data.KdUserFa,
		StsApproval:    "P",
		TglPengajuan:   time.Now(),
		TglActualBayar: data.TglActualBayar,
		TglUpdateFa:    time.Now(),
		Deskripsi:      data.Deskripsi,
		RenewalKe:      data.RenewalKe,
	}
	result := lR.conn.Save(&newExtendBayar)
	if result.Error != nil {
		return entity.ExtendBayar{}, result.Error
	}
	return newExtendBayar, nil
}

func (lR *extendBayarRepository) Delete(id string) error {

	result := lR.conn.Where("id", id).Delete(&entity.ExtendBayar{})
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func (lR *extendBayarRepository) UploadDokumen(data entity.ExtendBayar) error {
	return nil

}

func (lR *extendBayarRepository) UpdateFa(data request.ExtendBayarRequest) error {
	extendBayar := entity.ExtendBayar{Id: data.Id}
	lR.conn.Find(&extendBayar)
	if extendBayar.NoMsn == "" {
		return errors.New("data tersebut tidak ditemukan")
	}
	extendBayar.TglActualBayar = data.TglActualBayar
	extendBayar.Deskripsi = data.Deskripsi
	extendBayar.KdUserFa = data.KdUserFa
	result := lR.conn.Save(&extendBayar)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (lR *extendBayarRepository) UpdateLf(data request.ExtendBayarRequest) error {
	extendBayar := entity.ExtendBayar{Id: data.Id}
	lR.conn.Find(&extendBayar)
	if extendBayar.NoMsn == "" {
		return errors.New("data tersebut tidak ditemukan")
	}
	extendBayar.StsApproval = data.StsApproval
	extendBayar.KdUserLf = data.KdUserLf
	lR.conn.Save(&extendBayar)
	return nil
}

func (lR *extendBayarRepository) MasterData(search string, limit int, pageParams int) []entity.ExtendBayar {
	if search == "undefined" {
		search = ""
	}
	datas := []entity.ExtendBayar{}
	query := lR.conn.Where("deskripsi like ?", "%"+search+"%")
	query.Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Preload("Faktur", func(db *gorm.DB) *gorm.DB {
		return db.Select("no_msn, nm_customer11") // Select only id and title from Posts
	}).Find(&datas)
	return datas
}

func (lR *extendBayarRepository) MasterDataCount(search string) int64 {
	var datas []entity.ExtendBayar
	query := lR.conn.Where("deskripsi like ?", "%"+search+"%")
	query.Select("id").Find(&datas)
	return int64(len(datas))
}

func (lR *extendBayarRepository) DetailExtendBayar(id string) entity.ExtendBayar {
	extendBayar := entity.ExtendBayar{Id: id}
	lR.conn.Find(&extendBayar)
	return extendBayar
}
