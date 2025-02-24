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
	MasterData(search string, tgl1 string, tgl2 string, limit int, pageParams int) []entity.ExtendBayar
	MasterDataCount(search string, tgl1 string, tgl2 string) int64
	MasterDataLf(search string, tgl1 string, tgl2 string, sa string, limit int, pageParams int) []entity.ExtendBayar
	MasterDataLfCount(search string, tgl1 string, tgl2 string, sa string) int64
	DetailExtendBayar(id string) entity.ExtendBayar
	Create(data request.ExtendBayarRequest) (entity.ExtendBayar, error)
	UpdateFa(data request.ExtendBayarRequest) error
	UpdateLf(data request.ExtendBayarRequest) error
	Delete(id string, kdUserFa string) error
	UpdateApprovalLf(body request.ExtendBayarApprovalRequest) error
	BulkCreate(datas []entity.ExtendBayar) error
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
	existsExtendBayar := entity.ExtendBayar{}
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
	lR.conn.Where("no_msn = ? and sts_approval = ? ", data.NoMsn, "P").First(&existsExtendBayar)
	if existsExtendBayar.NoMsn == data.NoMsn && existsExtendBayar.StsApproval == "P" {
		return entity.ExtendBayar{}, errors.New("data tersebut sedang diproses")
	}
	result := lR.conn.Save(&newExtendBayar)
	if result.Error != nil {
		return entity.ExtendBayar{}, result.Error
	}
	return newExtendBayar, nil
}

func (lR *extendBayarRepository) Delete(id string, kdUserFa string) error {

	result := lR.conn.Where("id", id).Updates(entity.ExtendBayar{TglUpdateFa: time.Now(), IsDeleted: true, KdUserFa: kdUserFa})
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
	now := time.Now()
	if extendBayar.NoMsn == "" {
		return errors.New("data tersebut tidak ditemukan")
	}
	extendBayar.StsApproval = data.StsApproval
	extendBayar.KdUserLf = data.KdUserLf
	extendBayar.TglUpdateLf = &now
	result := lR.conn.Save(&extendBayar)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (lR *extendBayarRepository) UpdateApprovalLf(data request.ExtendBayarApprovalRequest) error {
	sqlConn, _ := lR.conn.DB()
	faktur3Repository := NewTr3nRepository(sqlConn, lR.conn)
	now := time.Now()
	for _, v := range data.Datas {
		extendBayar := entity.ExtendBayar{Id: v.Id}
		lR.conn.Find(&extendBayar)
		if extendBayar.NoMsn == "" || extendBayar.StsApproval == "O" {
			continue
		}
		if data.StsApproval == "O" {
			_, err := faktur3Repository.UpdateInputBayar(request.InputBayarRequest{TglBayar: extendBayar.TglActualBayar, NoMsn: extendBayar.NoMsn, KdUserFa: extendBayar.KdUserFa})
			if err != nil {
				return err
			}
		}
		extendBayar.StsApproval = data.StsApproval
		extendBayar.KdUserLf = data.KdUserLf
		extendBayar.TglUpdateLf = &now
		lR.conn.Save(&extendBayar)
	}
	return nil
}

func (lR *extendBayarRepository) MasterData(search string, tgl1 string, tgl2 string, limit int, pageParams int) []entity.ExtendBayar {
	if search == "undefined" {
		search = ""
	}
	datas := []entity.ExtendBayar{}
	query := lR.conn.Table("pengajuan_extend_bayar AS a").Joins("JOIN tr_wms_faktur3 as b ON b.no_msn = a.no_msn").Where("a.is_deleted = ?", false).Where("a.deskripsi like ? or b.nm_customer11 like ?", "%"+search+"%", "%"+search+"%")
	if tgl1 != "" && tgl2 != "" {
		query.Where("a.tgl_pengajuan >= ? and a.tgl_pengajuan <= ?", tgl1, tgl2)
	}
	query.Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Preload("Faktur", func(db *gorm.DB) *gorm.DB {
		return db.Select("no_msn, nm_customer11") // Select only id and title from Posts
	}).Find(&datas)
	return datas
}

func (lR *extendBayarRepository) MasterDataCount(search string, tgl1 string, tgl2 string) int64 {
	if search == "undefined" {
		search = ""
	}
	var datas []entity.ExtendBayar
	query := lR.conn.Table("pengajuan_extend_bayar AS a").Joins("JOIN tr_wms_faktur3 as b ON b.no_msn = a.no_msn").Where("a.is_deleted = ?", false).Where("a.deskripsi like ? or b.nm_customer11 like ?", "%"+search+"%", "%"+search+"%")
	if tgl1 != "" && tgl2 != "" {
		query.Where("a.tgl_pengajuan >= ? and a.tgl_pengajuan <= ?", tgl1, tgl2)
	}
	query.Select("id").Find(&datas)
	return int64(len(datas))
}

func (lR *extendBayarRepository) MasterDataLf(search string, tgl1 string, tgl2 string, sa string, limit int, pageParams int) []entity.ExtendBayar {
	if search == "undefined" {
		search = ""
	}
	datas := []entity.ExtendBayar{}
	query := lR.conn.Table("pengajuan_extend_bayar AS a").Joins("JOIN tr_wms_faktur3 as b ON b.no_msn = a.no_msn").Where("a.is_deleted = ?", false).Where("a.deskripsi like ? or a.no_msn like ? or b.nm_customer11 like ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	if sa != "all" {
		query.Where("a.sts_approval = ?", sa)
	}
	if tgl1 != "" && tgl2 != "" {
		query.Where("a.tgl_pengajuan >= ? and a.tgl_pengajuan <= ?", tgl1, tgl2)
	}
	query.Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Preload("Faktur", func(db *gorm.DB) *gorm.DB {
		return db.Select("no_msn, nm_customer11") // Select only id and title from Posts
	}).Find(&datas)
	return datas
}

func (lR *extendBayarRepository) MasterDataLfCount(search string, tgl1 string, tgl2 string, sa string) int64 {
	if search == "undefined" {
		search = ""
	}

	var datas []entity.ExtendBayar
	query := lR.conn.Table("pengajuan_extend_bayar AS a").Joins("JOIN tr_wms_faktur3 as b ON b.no_msn = a.no_msn").Where("a.is_deleted = ?", false).Where("a.deskripsi like ? or  a.no_msn like ? or b.nm_customer11 like ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	if sa != "all" {
		query.Where("a.sts_approval = ?", sa)
	}
	if tgl1 != "" && tgl2 != "" {
		query.Where("a.tgl_pengajuan >= ? and a.tgl_pengajuan <= ?", tgl1, tgl2)
	}
	query.Select("id").Find(&datas)
	return int64(len(datas))
}

func (lR *extendBayarRepository) DetailExtendBayar(id string) entity.ExtendBayar {
	extendBayar := entity.ExtendBayar{Id: id}
	lR.conn.Preload("Faktur").Preload("Faktur.Kartu").Preload("Faktur.Kurir").Preload("Faktur.MstCard").Find(&extendBayar)
	return extendBayar
}

func (lR *extendBayarRepository) BulkCreate(datas []entity.ExtendBayar) error {
	var exist entity.ExtendBayar
	sqlDb, err := lR.conn.DB()
	if err != nil {
		return err
	}
	tr3Repository := NewTr3nRepository(sqlDb, lR.conn)
	for _, value := range datas {
		_, _, err := tr3Repository.WillBayar(request.SearchWBRequest{Kode: value.NoMsn})
		if err != nil {
			return errors.New("Nomor Mesin " + value.NoMsn + " " + err.Error())
		}
		lR.conn.Where("no_msn", value.NoMsn).Where("sts_approval", "P").Where("is_deleted", false).First(&exist)
		if exist.Id != "" {
			return errors.New("Nomor mesin tersebut pending " + value.NoMsn)
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
			if err := tx.Table("pengajuan_extend_bayar").Save(&batch).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
		tx.Commit()
	}

	return nil
}
