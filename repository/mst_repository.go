package repository

import (
	"fmt"
	"time"

	// "wkm/request"
	"log"
	"wkm/entity"
	"wkm/response"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
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
	MasterAlasanTdkMembership() []entity.MstAlasanTdkMembership
	MasterProdukMembership() []response.Choices
	MasterPromoTransfer() []response.Choices
	MasterHobbies() []response.Choices
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
	mR.conn.Debug().Select("id, title, is_active, created").Find(&data)
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
	script.Title=data.Title
	script.Script= data.Script
	script.IsActive= data.IsActive
	script.Modified= &now
	script.ModiBy = username
	
	// Simpan perubahan ke database
	if err := mR.conn.Save(&script).Error; err != nil {
		return fmt.Errorf("failed to update script: %w", err)
	}

	return nil
}
func (r *mstRepository) MasterAlasanTdkMembership() []entity.MstAlasanTdkMembership {
	var data []entity.MstAlasanTdkMembership 
	r.conn.Find(&data)
	return data
}

func (r *mstRepository) MasterProdukMembership() []response.Choices {
	db, err := r.conn.DB()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	// Ensure the connection is established
	if err := db.Ping(); err != nil {
		log.Fatal("Database not reachable:", err)
	}
	

	// Define the query
	query := "SELECT kd_card AS kode, keterangan AS value, harga_pokok FROM db_wkm.mst_card WHERE status = '1' AND cat_card = 'R'"

	// Execute query
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Error executing query:", err)
	}

	// Process query results
	var cards []response.Choices
	price := 0
	p := message.NewPrinter(language.Indonesian)
	for rows.Next() {
		var card response.Choices
		if err := rows.Scan(&card.Value, &card.Name, &price); err != nil {
			log.Fatal("Error scanning row:", err)
		}
		card.Name += " - " + p.Sprintf("Rp. %d", price)
		cards = append(cards, card)
	}
	return cards
}

func (r *mstRepository) MasterPromoTransfer() []response.Choices {
	db, err := r.conn.DB()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	// Ensure the connection is established
	if err := db.Ping(); err != nil {
		log.Fatal("Database not reachable:", err)
	}

	// Define the query
	query := "select id as 'kode',nama_promo as 'value' from db_wkm.mst_promo where sts='1'"

	// Execute query
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Error executing query:", err)
	}

	// Process query results
	var cards []response.Choices
	for rows.Next() {
		var card response.Choices
		if err := rows.Scan(&card.Value, &card.Name); err != nil {
			log.Fatal("Error scanning row:", err)
		}
		cards = append(cards, card)
	}
	return cards
}

func (r *mstRepository) MasterHobbies() []response.Choices {
	db, err := r.conn.DB()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	// Ensure the connection is established
	if err := db.Ping(); err != nil {
		log.Fatal("Database not reachable:", err)
	}

	// Define the query
	query := "select kode_hobby as 'kode',hobby as 'value' from db_wkm.hobby"

	// Execute query
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Error executing query:", err)
	}

	// Process query results
	var cards []response.Choices
	for rows.Next() {
		var card response.Choices
		if err := rows.Scan(&card.Value, &card.Name); err != nil {
			log.Fatal("Error scanning row:", err)
		}
		cards = append(cards, card)
	}

	return cards
}
