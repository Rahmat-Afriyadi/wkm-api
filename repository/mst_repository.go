package repository

import (
	"wkm/entity"

	"gorm.io/gorm"
)

type MstRepository interface {
	MasterAgama() []entity.MstAgama
	MasterTujuPak() []entity.MstTujuPak
	MasterPendidikan() []entity.MstPendidikan
	MasterKeluarBln() []entity.MstKeluarBln
	MasterAktivitasJual() []entity.MstAktivitasJual
}

type mstRepository struct {
	conn *gorm.DB
}

func NewMstRepository(conn *gorm.DB) MstRepository {
	return &mstRepository{
		conn: conn,
	}
}

func (r *mstRepository) MasterAgama() []entity.MstAgama {
	var data []entity.MstAgama 
	r.conn.Find(&data)
	return data
}
func (r *mstRepository) MasterTujuPak() []entity.MstTujuPak {
	var data []entity.MstTujuPak 
	r.conn.Find(&data)
	return data
}
func (r *mstRepository) MasterPendidikan() []entity.MstPendidikan {
	var data []entity.MstPendidikan 
	r.conn.Find(&data)
	return data
}
func (r *mstRepository) MasterKeluarBln() []entity.MstKeluarBln {
	var data []entity.MstKeluarBln 
	r.conn.Find(&data)
	return data
}
func (r *mstRepository) MasterAktivitasJual() []entity.MstAktivitasJual {
	var data []entity.MstAktivitasJual 
	r.conn.Find(&data)
	return data
}