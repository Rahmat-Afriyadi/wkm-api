package repository

import (
	"fmt"
	"wkm/entity"
	"wkm/request"
	"wkm/utils"

	"gorm.io/gorm"
)

type ExtendBayarRepository interface {
	MasterData(search string, limit int, pageParams int) []entity.ExtendBayar
	MasterDataCount(search string) int64
	DetailExtendBayar(id uint64) entity.ExtendBayar
	Create(data request.ExtendBayarRequest) (entity.ExtendBayar, error)
	Update(data request.ExtendBayarRequest) error
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

	// newExtendBayar := entity.ExtendBayar{
	// 	TglAwal:   data.TglAwal,
	// 	TglAkhir:  data.TglAkhir,
	// 	KdUser:    data.KdUser,
	// 	Deskripsi: data.Deskripsi,
	// }
	// lR.conn.Save(&newExtendBayar)
	return entity.ExtendBayar{}, nil

}

func (lR *extendBayarRepository) Delete(id uint64) error {

	result := lR.conn.Where("id", id).Delete(&entity.ExtendBayar{})
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func (lR *extendBayarRepository) UploadDokumen(data entity.ExtendBayar) error {
	return nil

}

func (lR *extendBayarRepository) Update(data request.ExtendBayarRequest) error {
	// extendBayar := entity.ExtendBayar{ID: data.Id}
	// lR.conn.Find(&extendBayar)
	// if extendBayar.ID == 0 {
	// 	return errors.New("data tersebut tidak ditemukan")
	// }
	// extendBayar.TglAwal = data.TglAwal
	// extendBayar.TglAkhir = data.TglAkhir
	// extendBayar.Deskripsi = data.Deskripsi
	// extendBayar.KdUser = data.KdUser
	// lR.conn.Save(&extendBayar)
	return nil
}

func (lR *extendBayarRepository) MasterData(search string, limit int, pageParams int) []entity.ExtendBayar {
	datas := []entity.ExtendBayar{}
	query := lR.conn.Where("deskripsi like ?", "%"+search+"%")
	query.Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Find(&datas)
	fmt.Println("ini data yaa ", datas)
	return datas
}

func (lR *extendBayarRepository) MasterDataCount(search string) int64 {
	var datas []entity.ExtendBayar
	query := lR.conn.Where("deskripsi like ?", "%"+search+"%")
	query.Select("id").Find(&datas)
	return int64(len(datas))
}

func (lR *extendBayarRepository) DetailExtendBayar(id uint64) entity.ExtendBayar {
	extendBayar := entity.ExtendBayar{Id: "id"}
	lR.conn.Find(&extendBayar)
	return extendBayar
}
