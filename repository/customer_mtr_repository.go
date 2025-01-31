package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"
	"wkm/entity"
	"wkm/utils/query"

	"gorm.io/gorm"
)

type CustomerMtrRepository interface {
	ListAmbilData() []entity.Faktur3
	AmbilData(no_msn string, kd_user string) error
	Show(no_msn string)entity.CustomerMtr
	Update(customer entity.CustomerMtr) (entity.CustomerMtr, error)
	UpdateOkeMembership(customer entity.CustomerMtr) (entity.CustomerMtr, error)
}

type customerMtrRepository struct {
	conn     *sql.DB
	connGorm *gorm.DB
}

func NewCustomerMtrRepository(conn *sql.DB, connGorm *gorm.DB) CustomerMtrRepository {
	return &customerMtrRepository{
		conn:     conn,
		connGorm: connGorm,
	}
}

func (r *customerMtrRepository) ListAmbilData() []entity.Faktur3 {
	data := []entity.Faktur3{}
	r.connGorm.Select("no_msn").Where("sts_renewal is null").Find(&data)
	return data
}

func (r *customerMtrRepository) AmbilData(no_msn string, kd_user string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	queryAmbilData := query.NewQueryAmbilData()
	defer cancel()
	var kdUser sql.NullString
	err := r.conn.QueryRowContext(ctx, "select kd_user from tr_wms_faktur3 where no_msn = ? and sts_renewal is null", no_msn).Scan(&kdUser)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("data tidak ditemukan")
		} else {
			return err
		}
	}
	if kdUser.String != "" {
		return fmt.Errorf("data tersebut telah di ambil oleh kd_user %s", kdUser.String)
	}
	now := time.Now()
	r.conn.QueryContext(ctx, "update tr_wms_faktur3 set kd_user = ?, tgl_verifikasi = ?, sts_renewal='P' where no_msn = ?", kd_user,now.Format("2006-01-02"),no_msn)
	r.conn.QueryContext(ctx, queryAmbilData, no_msn)
	return nil
}


func (r *customerMtrRepository) Show(no_msn string) entity.CustomerMtr {
	data := entity.CustomerMtr{NoMsn: no_msn}
	r.connGorm.Find(&data)
	return data
}

func (r *customerMtrRepository) Update(customer entity.CustomerMtr) (entity.CustomerMtr, error) {
	data := entity.CustomerMtr{NoMsn: customer.NoMsn}
	r.connGorm.Select("no_msn, nm_customer_fkt").Find(&data)
	if data.NmCustomerFkt == "" {
		return entity.CustomerMtr{}, errors.New("data gk ada maaf")
	}

	r.connGorm.Save(&customer)
	return customer, nil
}

func (r *customerMtrRepository) UpdateOkeMembership(customer entity.CustomerMtr) (entity.CustomerMtr, error) {
	jsonBytes, err := json.Marshal(customer)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return entity.CustomerMtr{}, err
	}

	// Decode JSON bytes into a map
	var jsonMap map[string]interface{}
	err = json.Unmarshal(jsonBytes, &jsonMap)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return entity.CustomerMtr{}, err
	}

	stmt, err := r.conn.Prepare("UPDATE tr_wms_faktur3 SET agama2=?,alamat_bantuan=?,alamat_ktr=?,alamat21=?,alasan_pending_renewal=?,alasan_tdk_renewal=?,alasan_tdk_renewal2=?,email2=?,hobby2=?,jns_klm2=?,kd_card=?,kd_aktivitas_jual_r=?,kec_ktr=?,kec2=?,kel_ktr=?,kel2=?,keluar_bln2=?,kerja_di=?,ket_alamat21=?,ket_no_hp1=?,ket_no_telp1=?,ket_no_telp2=?,sts_kirim=?,kode_didik2=?,kode_kerja2=?,kodepos_ktr=?,kodepos2=?,kota_ktr=?,kota2=?,nama_ktp=?,no_hp2=?,no_telp_ktr2=?,no_telp2=?,no_yg_dihub_renewal=?,rt_ktr=?,rt2=?,rw_ktr=?,rw2=?,sts_kawin=?,sts_renewal=?,tgl_verifikasi=?,tgl_bayar_renewal=?,tgl_prospect=?,tujuan_pakai2=?,sts_jenis_bayar=? where no_msn = ? and kd_user =?")
	if err != nil {
		log.Fatal("Error preparing statement:", err)
	}
	defer stmt.Close() // Ensure statement is closed after execution
	// Execute the prepared statement multiple times with different values
	res, err := stmt.Exec(jsonMap["agama_wkm"],jsonMap["alamat_bantuan_wkm"],jsonMap["alamat_ktr_wkm"],jsonMap["alamat_wkm"],jsonMap["alasan_pending_membership"],jsonMap["alasan_tdk_membership"],jsonMap["alasan_tdk_membership_detail"],jsonMap["email_wkm"],jsonMap["hobby_wkm"],jsonMap["jns_klm_wkm"],jsonMap["jns_memberhip"],jsonMap["kd_aktivitas_jual_membership"],jsonMap["kec_ktr_wkm"],jsonMap["kec_wkm"],jsonMap["kel_ktr_wkm"],jsonMap["kel_wkm"],jsonMap["keluar_bln_wkm"],jsonMap["kerja_di_wkm"],jsonMap["ket_alamat_wkm"],jsonMap["ket_no_hp_fkt"],jsonMap["ket_no_telp_fkt"],jsonMap["ket_no_telp_wkm"],jsonMap["kirim_ke"],jsonMap["kode_didik_wkm"],jsonMap["kode_kerja_wkm"],jsonMap["kodepos_ktr_wkm"],jsonMap["kodepos_wkm"],jsonMap["kota_ktr_wkm"],jsonMap["kota_wkm"],jsonMap["nm_customer_wkm"],jsonMap["no_hp_wkm"],jsonMap["no_telp_ktr_wkm"],jsonMap["no_telp_wkm"],jsonMap["no_yg_dihub_ts"],jsonMap["rt_ktr_wkm"],jsonMap["rt_wkm"],jsonMap["rw_ktr_wkm"],jsonMap["rw_wkm"],jsonMap["sts_kawin_wkm"],jsonMap["sts_membership"],jsonMap["tgl_call_tele"],jsonMap["tgl_janji_bayar"],jsonMap["tgl_prospect_membership"],jsonMap["tujuan_pakai_wkm"],jsonMap["jns_bayar"], customer.NoMsn, customer.KdUserTs) // Update user with ID 1 to age 28
	if err != nil {
		log.Fatal("Error executing statement:", err)
	}

	// Get the number of affected rows
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatal("Error getting affected rows:", err)
	}

	fmt.Printf("Successfully updated %d row(s)\n", rowsAffected)

	r.connGorm.Save(&customer)
	return customer, nil
}