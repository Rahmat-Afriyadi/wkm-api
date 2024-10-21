package repository

import (
	"errors"
	"fmt"
	"sort"
	"time"
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
	Delete(id uint64) error
	GetMinTglBayar() time.Time
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

	newTglMerah := entity.TglMerah{
		TglAwal:   data.TglAwal,
		TglAkhir:  data.TglAkhir,
		KdUser:    data.KdUser,
		Deskripsi: data.Deskripsi,
	}
	lR.conn.Save(&newTglMerah)
	return newTglMerah, nil

}

func (lR *tglMerahRepository) Delete(id uint64) error {

	result := lR.conn.Where("id", id).Delete(&entity.TglMerah{})
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func (lR *tglMerahRepository) UploadDokumen(data entity.TglMerah) error {
	return nil

}

func (lR *tglMerahRepository) Update(data request.TglMerahRequest) error {
	tglMerah := entity.TglMerah{ID: data.Id}
	lR.conn.Find(&tglMerah)
	if tglMerah.ID == 0 {
		return errors.New("data tersebut tidak ditemukan")
	}
	tglMerah.TglAwal = data.TglAwal
	tglMerah.TglAkhir = data.TglAkhir
	tglMerah.Deskripsi = data.Deskripsi
	tglMerah.KdUser = data.KdUser
	lR.conn.Save(&tglMerah)
	return nil
}

func (lR *tglMerahRepository) MasterData(search string, limit int, pageParams int) []entity.TglMerah {
	datas := []entity.TglMerah{}
	query := lR.conn.Where("deskripsi like ?", "%"+search+"%")
	query.Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Find(&datas)
	return datas
}

func (lR *tglMerahRepository) MasterDataCount(search string) int64 {
	var datas []entity.TglMerah
	query := lR.conn.Where("deskripsi like ?", "%"+search+"%")
	query.Select("id").Find(&datas)
	return int64(len(datas))
}

func (lR *tglMerahRepository) DetailTglMerah(id uint64) entity.TglMerah {
	tglMerah := entity.TglMerah{ID: id}
	lR.conn.Find(&tglMerah)
	return tglMerah
}

func (lR *tglMerahRepository) BulkCreate(datas []entity.TglMerah) error {
	var exist entity.TglMerah
	for _, value := range datas {
		lR.conn.Where("tgl_awal", value.TglAwal.Format("2006-01-02")).Find(&exist)
		if exist.ID != 0 {
			fmt.Println("kesini gk sih ")
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

func (lR *tglMerahRepository) GetMinTglBayar() time.Time {
	listTglMerah := []entity.TglMerah{}
	today := time.Now()
	min := today.AddDate(0, 0, -1)
	if today.Weekday() == 1 {
		min = today.AddDate(0, 0, -3)
	}
	lR.conn.Where("tgl_akhir = ?", min.Format("2006-01-02")).Find(&listTglMerah)
	listMin := []time.Time{}
	if len(listTglMerah) > 0 {
		for _, v := range listTglMerah {
			listMin = append(listMin, v.TglAwal)
		}
		sort.Slice(listMin, func(i, j int) bool {
			return listMin[i].Before(listMin[j])
		})
		if listMin[0].Weekday() == 1 {
			min = listMin[0].AddDate(0, 0, -3)
		} else {
			min = listMin[0].AddDate(0, 0, -1)
		}
	}
	return min

}
