package repository

import (
	"database/sql"
	"fmt"
	// "time"
	"wkm/entity"
	// "wkm/request"
	// "wkm/response"
	// "encoding/json"
	"github.com/google/uuid"
)

type AsuransiPARepository interface {
	CreateAsuransiPA(data entity.AsuransiPA) error
	UpdateAsuransiPA(id string, data entity.AsuransiPA) error
}

type asuransiPARepository struct {
	conn *sql.DB
}

func NewAsuransiPARepository(conn *sql.DB) AsuransiPARepository {
	return &asuransiPARepository{
		conn: conn,
	}
}

func (ar *asuransiPARepository) CreateAsuransiPA(data entity.AsuransiPA) error {
	data.Id = uuid.New().String()
	query := `
		INSERT INTO asuransi_pa (Id, no_msn, nm_customer, sts_asuransi_pa, id_produk, app_trans_id, tgl_beli, no_ktpnpwp, alasan_pending_asuransi_pa, sts_pembelian)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	// Execute the SQL statement
	_, err := ar.conn.Exec(query, data.Id, data.NoMSN, data.NmCustomer, data.StsAsuransiPA, data.IDProduk, data.AppTransID, data.TglBeli, data.NoKtpNpwp, data.AlasanPendingAsuransiPA, data.StsPembelian)
	if err != nil {
		fmt.Printf("Error creating AsuransiPA: %v", err)
		return err // Return the error to the caller
	}

	return nil
}

func (ar *asuransiPARepository) UpdateAsuransiPA(id string, data entity.AsuransiPA) error {
	query := `
		UPDATE asuransi_pa 
		SET no_msn = ?, nm_customer = ?, sts_asuransi_pa = ?, id_produk = ?, app_trans_id = ?, tgl_beli = ?, no_ktpnpwp = ?, alasan_pending_asuransi_pa = ?, sts_pembelian = ?
		WHERE Id = ?
	`
	// Execute the SQL statement
	_, err := ar.conn.Exec(query, data.NoMSN, data.NmCustomer, data.StsAsuransiPA, data.IDProduk, data.AppTransID, data.TglBeli, data.NoKtpNpwp, data.AlasanPendingAsuransiPA, data.StsPembelian, id)
	if err != nil {
		fmt.Printf("Error updating AsuransiPA with ID %s: %v", id, err)
		return err // Return the error to the caller
	}

	return nil // Return nil if the update was successful
}
