package repository

import (
	"fmt"
	"wkm/entity"
	"time"

	// "wkm/request"

	"gorm.io/gorm"
)

type MstRepository interface {
	MasterAgama() []entity.MstAgama
	MasterTujuPak() []entity.MstTujuPak
	MasterPendidikan() []entity.MstPendidikan
	MasterKeluarBln() []entity.MstKeluarBln
	MasterAktivitasJual() []entity.MstAktivitasJual
	CreateScript(data entity.MstScript, username string) error
	UpdateScript(id string, data entity.MstScript, username string) error
	MasterScript() []entity.MstScript
	ListAllScript() []entity.MstScript
	ViewScript(id string) (entity.MstScript, error)
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

func (mR *mstRepository) ViewScript(id string) (entity.MstScript, error) {
	var data entity.MstScript
	if err := mR.conn.Where("id = ?", id).Find(&data).Error; err != nil {
		return entity.MstScript{}, err
	}
	return data, nil
}
func (mR *mstRepository) MasterScript() []entity.MstScript {
	var data []entity.MstScript
	mR.conn.Where("is_active = ?", 1).Find(&data)
	return data
}

func (mR *mstRepository) ListAllScript() []entity.MstScript {
	var data []entity.MstScript
	mR.conn.Debug().Select("id, title, is_active").Find(&data)
	return data
}

func (mR *mstRepository) CreateScript(data entity.MstScript, username string) error {
	// Buat entitas berdasarkan input data
	now := time.Now()
	script := entity.MstScript{
		Title:    data.Title,
		Script:   data.Script,
		KdUser:   username,
		IsActive: data.IsActive,
		Created: &now,
	}

	// Simpan entitas ke database
	if err := mR.conn.Create(&script).Error; err != nil {
		return fmt.Errorf("failed to create script: %w", err)
	}

	return nil
}

func (mR *mstRepository) UpdateScript(id string, data entity.MstScript, username string) error {
	// Cari script berdasarkan ID
	var script entity.MstScript
	if err := mR.conn.First(&script, "id = ?", id).Error; err != nil {
		return fmt.Errorf("failed to find script with id %s: %w", id, err)
	}
	now := time.Now()
	// Perbarui entitas dengan data baru
	script.Script = data.Script
	script.IsActive = data.IsActive
	script.Modified= &now
	script.ModiBy = username
	

	// Simpan perubahan ke database
	if err := mR.conn.Save(&script).Error; err != nil {
		return fmt.Errorf("failed to update script: %w", err)
	}

	return nil
}
