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
	"wkm/request"
	"wkm/utils"
	"wkm/utils/query"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type CustomerMtrRepository interface {
	MasterData(search string, sts string, jns string, username string, limit int, pageParams int) []entity.CustomerMtr

	MasterDataCount(search string, sts string, jns string, username string) int64
	ListAmbilData() []entity.Faktur3
	AmbilData(no_msn string, kd_user string) error
	Show(no_msn string)entity.CustomerMtr
	Update(customer entity.CustomerMtr) (entity.CustomerMtr, error)
	UpdateOkeMembership(customer request.CustomerMtr) (entity.CustomerMtr, error)
}

type customerMtrRepository struct {
	conn     *sql.DB
	connGorm *gorm.DB
	wandaGorm *gorm.DB
}

func NewCustomerMtrRepository(conn *sql.DB, connGorm *gorm.DB,wandaGorm *gorm.DB) CustomerMtrRepository {
	return &customerMtrRepository{
		conn:     conn,
		connGorm: connGorm,
		wandaGorm: wandaGorm,
	}
}

func (cR *customerMtrRepository) MasterData(search string, sts string, jns string, username string, limit int, pageParams int) []entity.CustomerMtr {
	datas := []entity.CustomerMtr{}
	query := cR.connGorm.Where("no_msn like ? or nm_customer_wkm like ? or nm_customer_fkt like ? ", "%"+search+"%","%"+search+"%", "%"+search+"%")
	query.Where("kd_user_ts = ?", username)
	query.Where(fmt.Sprintf("%s = ?", jns),sts)

	query.Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Order("tgl_call_tele desc").Find(&datas)
	return datas

}
func (cR *customerMtrRepository) MasterDataCount(search string, sts string, jns string, username string) int64 {
	datas := []entity.CustomerMtr{}
	query := cR.connGorm.Where("no_msn like ? or nm_customer_wkm like ? or nm_customer_fkt like ? ", "%"+search+"%", "%"+search+"%","%"+search+"%")
	query.Where("kd_user_ts = ?", username)
	query.Where(fmt.Sprintf("%s = ?", jns),sts)

	query.Select("no_msn").Find(&datas)
	return int64(len(datas))
}

func (r *customerMtrRepository) ListAmbilData() []entity.Faktur3 {
	data := []entity.Faktur3{}
	r.connGorm.Select("no_msn").Order("RAND()").Where("sts_renewal is null").Find(&data)
	return data
}

func (r *customerMtrRepository) AmbilData(no_msn string, kd_user string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	queryAmbilData := query.NewQueryAmbilData()
	defer cancel()
	var kdUser sql.NullString
	var stsRenewal sql.NullString
	err := r.conn.QueryRowContext(ctx, "select kd_user, sts_renewal from tr_wms_faktur3 where no_msn = ?", no_msn).Scan(&kdUser, &stsRenewal)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("data tidak ditemukan")
		} else {
			return err
		}
	}
	if stsRenewal.String != "" {
		return fmt.Errorf("data tersebut telah di ambil oleh user lain")
	}
	now := time.Now()
	r.conn.QueryContext(ctx, "update tr_wms_faktur3 set kd_user = ?, tgl_verifikasi = ?, sts_renewal='P' where no_msn = ?", kd_user,now.Format("2006-01-02"),no_msn)
	r.conn.QueryContext(ctx, queryAmbilData, no_msn)
	return nil
}


func (r *customerMtrRepository) Show(no_msn string) entity.CustomerMtr {
	data := entity.CustomerMtr{NoMsn: no_msn}
	r.connGorm.Preload("Memberships", "renewal_ke = (SELECT renewal_ke FROM customer_mtr WHERE customer_mtr.no_msn = membership.no_msn)").Debug().Preload("AsuransiPa").Preload("AsuransiMtr").Find(&data)
	var produkPa entity.MasterProduk
	var produkMtr entity.MasterProduk
	r.wandaGorm.Preload("Vendor").Where("id_produk = ?",data.AsuransiMtr.IDProduk).Find(&produkMtr)
	r.wandaGorm.Preload("Vendor").Where("id_produk = ?",data.AsuransiPa.IDProduk).Find(&produkPa)
	data.AsuransiPa.Produk = produkPa
	data.AsuransiMtr.Produk = produkMtr

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

func (r *customerMtrRepository) UpdateOkeMembership(customer request.CustomerMtr) (entity.CustomerMtr, error) {
	var membership entity.Membership
	var asuransiMtr entity.AsuransiMtr
	var asuransiPa entity.AsuransiPA
	print := 0
	now := time.Now()

	jsonBytes, err := json.Marshal(customer)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return entity.CustomerMtr{}, err
	}

	// Decode JSON bytes into a map
	var customerMtrEntity entity.CustomerMtr
	var existCustomerMtr entity.CustomerMtr
	r.connGorm.Where("no_msn = ? and renewal_ke = ?", customer.NoMsn, customer.RenewalKe).First(&existCustomerMtr)
	var jsonMap map[string]interface{}
	err = json.Unmarshal(jsonBytes, &jsonMap)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return entity.CustomerMtr{}, err
	}
	err = json.Unmarshal(jsonBytes, &customerMtrEntity)
	if err != nil {
		fmt.Println("Error decoding Customer JSON:", err)
		return entity.CustomerMtr{}, err
	}

	// cuman ada tgl_prospect_membership karena butuh update di faktur 3
	// tgl_prospect_asuransi_pa dan tgl_prospect_asuransi_mtr tidak perlu 
	
	if jsonMap["tgl_prospect_membership"] != nil {
		if len(jsonMap["tgl_prospect_membership"].(string)) > 10 {
			jsonMap["tgl_prospect_membership"] = jsonMap["tgl_prospect_membership"].(string)[:10]
		}
	}
	if jsonMap["tgl_janji_bayar"] != nil {
		jsonMap["tgl_janji_bayar"] =jsonMap["tgl_janji_bayar"].(string)[:10]
	}
	if jsonMap["tgl_call_tele"] != nil {
		jsonMap["tgl_call_tele"] =jsonMap["tgl_call_tele"].(string)[:10]
	}
	
	if jsonMap["sts_membership"] == "O" && existCustomerMtr.StsMembership != "O" {
		err = json.Unmarshal(jsonBytes, &membership)
		if err != nil {
			fmt.Println("Error decoding JSON Membership:", err)
			return entity.CustomerMtr{}, err
		}
		if membership.TypeKartu == "E" {
			print = 1
		}
		
		jsonMap["alasan_tdk_membership"] = nil
		jsonMap["alasan_pending_membership"] = nil
		jsonMap["tgl_prospect_membership"] = nil
		if jsonMap["sts_asuransi_pa"] != "O" {
			customerMtrEntity.StsAsuransiPa = "M"
		}
		customerMtrEntity.AlasanTdkMembership = ""
		customerMtrEntity.AlasanPendingMembership = ""
		customerMtrEntity.TglProspectMembership = nil
		r.connGorm.Save(&membership)
	}
	if jsonMap["sts_asuransi_mtr"] == "O" && existCustomerMtr.StsAsuransiMtr != "O" {
		err = json.Unmarshal(jsonBytes, &asuransiMtr)
		if err != nil {
			fmt.Println("Error decoding JSON Membership:", err)
			return entity.CustomerMtr{}, err
		}
		jsonMap["alasan_tdk_asuransi_pa"] = nil
		jsonMap["alasan_pending_asuransi_pa"] = nil
		jsonMap["tgl_prospect_asuransi_pa"] = nil
		customerMtrEntity.AlasanTdkAsuransiMtr = ""
		customerMtrEntity.AlasanPendingAsuransiMtr = ""
		customerMtrEntity.TglProspectAsuransiMtr = nil

		u2, err := uuid.NewV4()
		if err != nil {
			fmt.Println("ini error uuid ", err)
		}
		asuransiMtr.AppTransID = u2.String()
		asuransiMtr.NmCustomer = customer.NmCustomerWkm
		asuransiMtr.IDProduk = customer.IdProdukAsuransIMotor
		asuransiMtr.TglBeli = &now
		asuransiMtr.StsPembelian = "1"
		asuransiMtr.ThnMtr = customer.ThnMtr
		r.connGorm.Save(&asuransiMtr)
	}
	if jsonMap["sts_asuransi_pa"] == "O" &&  existCustomerMtr.StsAsuransiPa != "O"{
		fmt.Println("kesini gk sih ")
		err = json.Unmarshal(jsonBytes, &asuransiPa)
		if err != nil {
			fmt.Println("Error decoding JSON Membership:", err)
			return entity.CustomerMtr{}, err
		}
		jsonMap["alasan_tdk_asuransi_pa"] = nil
		jsonMap["alasan_pending_asuransi_pa"] = nil
		jsonMap["tgl_prospect_asuransi_pa"] = nil
		customerMtrEntity.AlasanTdkAsuransiPa = ""
		customerMtrEntity.AlasanPendingAsuransiPa = ""
		customerMtrEntity.TglProspectAsuransiPa = nil

		u2, err := uuid.NewV4()
		if err != nil {
			fmt.Println("ini error uuid ", err)
		}
		asuransiPa.AppTransID = u2.String()
		asuransiPa.NmCustomer = customer.NmCustomerWkm
		asuransiPa.IDProduk = customer.IdProdukAsuransIPa
		asuransiPa.TglBeli = &now
		asuransiPa.StsPembelian = "1"
		r.connGorm.Save(&asuransiPa)
	}

	stmt, err := r.conn.Prepare("UPDATE tr_wms_faktur3 SET print=?, agama2=?,alamat_bantuan=?,alamat_ktr=?,alamat21=?,alasan_pending_renewal=?,alasan_tdk_renewal=?,alasan_tdk_renewal2=?,email2=?,hobby2=?,jns_klm2=?,kd_card=?,kd_aktivitas_jual_r=?,kec_ktr=?,kec2=?,kel_ktr=?,kel2=?,keluar_bln2=?,kerja_di=?,ket_alamat21=?,ket_no_hp1=?,ket_no_telp1=?,ket_no_telp2=?,sts_kirim=?,kode_didik2=?,kode_kerja2=?,kodepos_ktr=?,kodepos2=?,kota_ktr=?,kota2=?,nama_ktp=?,no_hp2=?,no_telp_ktr2=?,no_telp2=?,no_yg_dihub_renewal=?,rt_ktr=?,rt2=?,rw_ktr=?,rw2=?,sts_kawin=?,sts_renewal=?,tgl_verifikasi=?,tgl_bayar_renewal=?,tgl_prospect=?,tujuan_pakai2=?,sts_jenis_bayar=?, sts_asuransi_pa='O' where no_msn = ? and kd_user =?")
	if err != nil {
		log.Fatal("Error preparing statement:", err)
	}
	defer stmt.Close() // Ensure statement is closed after execution
	res, err := stmt.Exec(print, jsonMap["agama_wkm"],jsonMap["alamat_bantuan_wkm"],jsonMap["alamat_ktr_wkm"],jsonMap["alamat_wkm"],jsonMap["alasan_pending_membership"],jsonMap["alasan_tdk_membership"],jsonMap["alasan_tdk_membership_detail"],jsonMap["email_wkm"],jsonMap["hobby_wkm"],jsonMap["jns_klm_wkm"],jsonMap["jns_membership"],jsonMap["kd_aktivitas_jual_membership"],jsonMap["kec_ktr_wkm"],jsonMap["kec_wkm"],jsonMap["kel_ktr_wkm"],jsonMap["kel_wkm"],jsonMap["keluar_bln_wkm"],jsonMap["kerja_di_wkm"],jsonMap["ket_alamat_wkm"],jsonMap["ket_no_hp_fkt"],jsonMap["ket_no_telp_fkt"],jsonMap["ket_no_telp_wkm"],jsonMap["kirim_ke"],jsonMap["kode_didik_wkm"],jsonMap["kode_kerja_wkm"],jsonMap["kodepos_ktr_wkm"],jsonMap["kodepos_wkm"],jsonMap["kota_ktr_wkm"],jsonMap["kota_wkm"],jsonMap["nm_customer_wkm"],jsonMap["no_hp_wkm"],jsonMap["no_telp_ktr_wkm"],jsonMap["no_telp_wkm"],jsonMap["ket_hub_ts"],jsonMap["rt_ktr_wkm"],jsonMap["rt_wkm"],jsonMap["rw_ktr_wkm"],jsonMap["rw_wkm"],jsonMap["sts_kawin_wkm"],jsonMap["sts_membership"],jsonMap["tgl_call_tele"],jsonMap["tgl_janji_bayar"],jsonMap["tgl_prospect_membership"],jsonMap["tujuan_pakai_wkm"],jsonMap["jns_bayar"], customer.NoMsn, customer.KdUserTs) // Update user with ID 1 to age 28
	if err != nil {
		log.Fatal("Error executing statement:", err)
	}

	// Get the number of affected rows
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatal("Error getting affected rows:", err)
	}

	// if customerMtrEntity.KetNoHpFkt == "1" || customerMtrEntity.KetNoTelpFkt == "1" {
		
	// }
	customerMtrEntity.TglCallTele = &now
	customerMtrEntity.Modified = &now
	customerMtrEntity.JmlCallMembership += 1
	r.connGorm.Save(&customerMtrEntity)
	
	fmt.Printf("Successfully updated %d row(s)\n", rowsAffected)
	return customerMtrEntity, nil
}