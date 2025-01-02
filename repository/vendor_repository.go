package repository

import (
	"errors"
	"fmt"
	"wkm/entity"
	"wkm/utils"

	"gorm.io/gorm"
)

type VendorRepository interface {
	MasterData(search string, limit int, pageParams int) []entity.MasterVendor
	MasterDataCount(search string) int64
	DetailVendor(id string) entity.MasterVendor
	Create(data entity.MasterVendor) error
	Update(data entity.MasterVendor) error
}

type vendorRepository struct {
	conn *gorm.DB
}

func NewVendorRepository(conn *gorm.DB) VendorRepository {
	return &vendorRepository{
		conn: conn,
	}
}

func (lR *vendorRepository) Create(data entity.MasterVendor) error {
	result := lR.conn.Create(&data)
	if result.Error != nil {
		fmt.Println("ini error ", result.Error)
		return result.Error
	} else {
		return nil
	}

}

func (lR *vendorRepository) Update(data entity.MasterVendor) error {
	record := entity.MasterVendor{KdVendor: data.KdVendor}
	lR.conn.Find(&record)
	if record.NmVendor == "" {
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

func (lR *vendorRepository) MasterData(search string, limit int, pageParams int) []entity.MasterVendor {
	datas := []entity.MasterVendor{}
	query := lR.conn.Where("nm_vendor like ? or deskripsi like ?", "%"+search+"%", "%"+search+"%")
	query.Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Find(&datas)
	return datas
}

func (lR *vendorRepository) MasterDataCount(search string) int64 {
	var datas []entity.MasterVendor
	query := lR.conn.Where("nm_vendor like ? or deskripsi like ?", "%"+search+"%", "%"+search+"%")
	query.Select("id_vendor").Find(&datas)
	return int64(len(datas))
}

func (lR *vendorRepository) DetailVendor(id string) entity.MasterVendor {
	vendor := entity.MasterVendor{KdVendor: id}
	lR.conn.Find(&vendor)
	return vendor
}
