package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"time"
	"wkm/entity"
	"wkm/request"
	"wkm/response"
	"wkm/utils"
	"wkm/utils/query"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type CustomerMtrRepository interface {
	MasterData(search string, sts string, jns string, username string, limit int, pageParams int) []entity.CustomerMtr

	MasterDataCount(search string, sts string, jns string, username string) int64
	SelfCount(kd_user string) int64
	ListAmbilData() []entity.Faktur3
	AmbilData(no_msn string, kd_user string) error
	Show(no_msn string) entity.CustomerMtr
	Update(customer entity.CustomerMtr) (entity.CustomerMtr, error)
	UpdateOkeMembership(customer request.CustomerMtr) (entity.CustomerMtr, error)
	RekapTele(username string, startDate time.Time, endDate time.Time) (response.RekapTele, error)
	ListBerminatMembership(username string, startDate time.Time, endDate time.Time, limit int, pageParams int, search string) ([]response.MinatMembership, int, int, error)
	ListDataAsuransiPA(username string, startDate time.Time, endDate time.Time, limit int, pageParams int, search string) ([]response.ListAsuransi, int, int, error)
	ListDataAsuransiMtr(username string, startDate time.Time, endDate time.Time, limit int, pageParams int, search string) ([]response.ListAsuransi, int, int, error)
}

type customerMtrRepository struct {
	conn      *sql.DB
	connGorm  *gorm.DB
	wandaGorm *gorm.DB
}

func NewCustomerMtrRepository(conn *sql.DB, connGorm *gorm.DB, wandaGorm *gorm.DB) CustomerMtrRepository {
	return &customerMtrRepository{
		conn:      conn,
		connGorm:  connGorm,
		wandaGorm: wandaGorm,
	}
}

func (cR *customerMtrRepository) SelfCount(kd_user string) int64 {
	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	var count int64
	defer cancel()
	err := cR.conn.QueryRowContext(ctx,"select count(*) from customer_mtr where tgl_call_tele=? and kd_user_ts = ?",now.Format("2006-01-02"), kd_user).Scan(&count)
	if err != nil {
		fmt.Println("ini error yaa ", err.Error())	
	}
	return count
}

func (cR *customerMtrRepository) MasterData(search string, sts string, jns string, username string, limit int, pageParams int) []entity.CustomerMtr {
	datas := []entity.CustomerMtr{}
	query := cR.connGorm.Select("no_msn, nm_customer_fkt, kd_user_ts, alasan_pending_membership, alasan_pending_asuransi_pa, alasan_pending_asuransi_mtr,tgl_prospect_membership,tgl_prospect_asuransi_pa,tgl_prospect_asuransi_mtr, tgl_call_tele").Where("no_msn like ? or nm_customer_wkm like ? or nm_customer_fkt like ? ", "%"+search+"%","%"+search+"%", "%"+search+"%")
	query.Where("kd_user_ts = ?", username)
	query.Where(fmt.Sprintf("%s = ?", jns), sts)

	query.Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Order("modified desc").Find(&datas)
	return datas
}
func (cR *customerMtrRepository) MasterDataCount(search string, sts string, jns string, username string) int64 {
	datas := []entity.CustomerMtr{}
	query := cR.connGorm.Where("no_msn like ? or nm_customer_wkm like ? or nm_customer_fkt like ? ", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	query.Where("kd_user_ts = ?", username)
	query.Where(fmt.Sprintf("%s = ?", jns), sts)

	query.Select("no_msn").Find(&datas)
	return int64(len(datas))
}
func (r *customerMtrRepository) ListAmbilData() []entity.Faktur3 {
	data := []entity.Faktur3{}
	r.connGorm.Select("no_msn").Order("RAND()").Where("sts_renewal is null").Limit(100).Find(&data)
	return data
}

func (r *customerMtrRepository) AmbilData(no_msn string, kd_user string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	queryAmbilData := query.NewQueryAmbilData()
	queryUpdateAmbilData := query.NewQueryUpdateAmbilData()
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
	if stsRenewal.String != ""  {
		return fmt.Errorf("data tersebut telah di ambil oleh user lain")
	}
	now := time.Now()
	_, err = r.conn.QueryContext(ctx, "update tr_wms_faktur3 set kd_user = ?, tgl_verifikasi = ?, sts_renewal='P' where no_msn = ?", kd_user, now.Format("2006-01-02"), no_msn)
	if err != nil {
		return err
	}

	var count int64
	r.connGorm.Model(&entity.CustomerMtr{}).Where("no_msn = ?", no_msn).Count(&count)
	if count > 0 {
		r.conn.QueryContext(ctx, queryUpdateAmbilData, no_msn)
	} else {
		r.conn.QueryContext(ctx, queryAmbilData, no_msn)
	}
	return nil
}

func (r *customerMtrRepository) Show(no_msn string) entity.CustomerMtr {
	data := entity.CustomerMtr{NoMsn: no_msn}
	r.connGorm.Preload("Memberships", "renewal_ke = (SELECT renewal_ke FROM customer_mtr WHERE customer_mtr.no_msn = membership.no_msn)").Preload("AsuransiPa").Preload("AsuransiMtr").Find(&data)
	var produkPa entity.MasterProduk
	var produkMtr entity.MasterProduk
	r.wandaGorm.Preload("Vendor").Where("id_produk = ?", data.AsuransiMtr.IDProduk).Find(&produkMtr)
	r.wandaGorm.Preload("Vendor").Where("id_produk = ?", data.AsuransiPa.IDProduk).Find(&produkPa)
	data.AsuransiPa.Produk = produkPa
	data.AsuransiMtr.Produk = produkMtr

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r.conn.QueryRowContext(ctx, "select c.no_msn, k.nm_kerja, h.hobby, a.agama, t.nm_tujpak, p.nm_pendidikan, kb.nm_keluar_bln2 from customer_mtr c inner join mst_kerja k on k.kode_kerja2 = c.kode_kerja_fkt inner join hobby h on h.kode_hobby = c.hobby_fkt inner join mst_agama a on a.kd_agama = c.agama_fkt inner join mst_tujuanpakai t on c.tujuan_pakai_fkt = t.id inner join mst_pendidikan p on p.kd_pendidikan = c.kode_didik_fkt inner join mst_keluar_bln kb on kb.keluar_bln2 = c.keluar_bln_fkt where c.no_msn = ?", no_msn).Scan(&no_msn, &data.DescKerjaFkt, &data.DescHobbyFkt, &data.DescAgamaFkt, &data.DescTujuanPakaiFkt, &data.DescDidikFkt, &data.DescKeluarBlnFkt)

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

	if jsonMap["tgl_prospect_membership"] != nil {
		if len(jsonMap["tgl_prospect_membership"].(string)) > 10 {
			jsonMap["tgl_prospect_membership"] = jsonMap["tgl_prospect_membership"].(string)[:10]
		}
	}
	if jsonMap["tgl_janji_bayar"] != nil {
		jsonMap["tgl_janji_bayar"] = jsonMap["tgl_janji_bayar"].(string)[:10]
	}

	jsonMap["tgl_call_tele"] = now.Format("2006-01-02")

	if jsonMap["sts_membership"] == "O" {
		result := r.connGorm.Where("no_msn = ? and renewal_ke = ?", customer.NoMsn, customer.RenewalKe).First(&entity.Membership{});
		if result.Error == gorm.ErrRecordNotFound {
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
	}
	
	if jsonMap["sts_asuransi_mtr"] == "O" {
		result := r.connGorm.Where("no_msn = ? ", customer.NoMsn).First(&entity.AsuransiMtr{});
		if result.Error == gorm.ErrRecordNotFound {
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
		
	}
	if jsonMap["sts_asuransi_pa"] == "O"{
		result := r.connGorm.Where("no_msn = ? ", customer.NoMsn).First(&entity.AsuransiPA{});
		if result.Error == gorm.ErrRecordNotFound {
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
	}

	if jsonMap["nm_customer_fkt"] == jsonMap["nm_customer_wkm"] {
		jsonMap["nm_customer_wkm"] = ""
	}
	stmt, err := r.conn.Prepare("UPDATE tr_wms_faktur3 SET print=?, agama2=?,alamat_bantuan=?,alamat_ktr=?,alamat21=?,alasan_pending_renewal=?,alasan_tdk_renewal=?,alasan_tdk_renewal2=?,email2=?,hobby2=?,jns_klm2=?,kd_card=?,kd_aktivitas_jual_r=?,kec_ktr=?,kec2=?,kel_ktr=?,kel2=?,keluar_bln2=?,kerja_di=?,ket_alamat21=?,ket_no_hp1=?,ket_no_telp1=?,ket_no_telp2=?,sts_kirim=?,kode_didik2=?,kode_kerja2=?,kodepos_ktr=?,kodepos2=?,kota_ktr=?,kota2=?,nama_ktp=?,no_hp2=?,no_telp_ktr2=?,no_telp2=?,no_yg_dihub_renewal=?,rt_ktr=?,rt2=?,rw_ktr=?,rw2=?,sts_kawin=?,sts_renewal=?,tgl_verifikasi=?,tgl_bayar_renewal=?,tgl_prospect=?,tujuan_pakai2=?,sts_jenis_bayar=?, sts_asuransi_pa='O' where no_msn = ? and kd_user =?")
	if err != nil {
		log.Fatal("Error preparing statement:", err)
	}
	defer stmt.Close()                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           // Ensure statement is closed after execution
	res, err := stmt.Exec(print, jsonMap["agama_wkm"], jsonMap["alamat_bantuan_wkm"], jsonMap["alamat_ktr_wkm"], jsonMap["alamat_wkm"], jsonMap["alasan_pending_membership"], jsonMap["alasan_tdk_membership"], jsonMap["alasan_tdk_membership_detail"], jsonMap["email_wkm"], jsonMap["hobby_wkm"], jsonMap["jns_klm_wkm"], jsonMap["jns_membership"], jsonMap["kd_aktivitas_jual_membership"], jsonMap["kec_ktr_wkm"], jsonMap["kec_wkm"], jsonMap["kel_ktr_wkm"], jsonMap["kel_wkm"], jsonMap["keluar_bln_wkm"], jsonMap["kerja_di_wkm"], jsonMap["ket_alamat_wkm"], jsonMap["ket_no_hp_fkt"], jsonMap["ket_no_telp_fkt"], jsonMap["ket_no_telp_wkm"], jsonMap["kirim_ke"], jsonMap["kode_didik_wkm"], jsonMap["kode_kerja_wkm"], jsonMap["kodepos_ktr_wkm"], jsonMap["kodepos_wkm"], jsonMap["kota_ktr_wkm"], jsonMap["kota_wkm"], jsonMap["nm_customer_wkm"], jsonMap["no_hp_wkm"], jsonMap["no_telp_ktr_wkm"], jsonMap["no_telp_wkm"], jsonMap["no_yg_dihub_ts"], jsonMap["rt_ktr_wkm"], jsonMap["rt_wkm"], jsonMap["rw_ktr_wkm"], jsonMap["rw_wkm"], jsonMap["sts_kawin_wkm"], jsonMap["sts_membership"], jsonMap["tgl_call_tele"], jsonMap["tgl_janji_bayar"], jsonMap["tgl_prospect_membership"], jsonMap["tujuan_pakai_wkm"], jsonMap["jns_bayar"], customer.NoMsn, customer.KdUserTs) // Update user with ID 1 to age 28
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

func (r *customerMtrRepository) RekapTele(username string, startDate time.Time, endDate time.Time) (response.RekapTele, error) {
	var rekap response.RekapTele
	// Query untuk jumlah data membership
	query := `SELECT COUNT(*) FROM customer_mtr WHERE kd_user_ts = ? AND tgl_call_tele BETWEEN ? AND ?`
	err := r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.JumlahDataMembership)
	if err != nil {
		return rekap, fmt.Errorf("error fetching jumlah_data_membership: %v", err)
	}

	query = `SELECT COUNT(*) FROM customer_mtr 
          WHERE kd_user_ts = ? 
          AND tgl_call_tele BETWEEN DATE_SUB(?, INTERVAL 1 MONTH) 
          AND DATE_SUB(?, INTERVAL 1 MONTH)`

	err = r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.JumlahDataMembershipBefore)
	if err != nil {
		return rekap, fmt.Errorf("error fetching jumlah_data_membership: %v", err)
	}

	// Query untuk data membership berminat (sts_membership = 'O')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE kd_user_ts = ? AND tgl_call_tele BETWEEN ? AND ? AND sts_membership = 'O'`
	err = r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.DataBerminatMembership)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_berminat_membership: %v", err)
	}

	// Query untuk data membership berminat (sts_membership = 'O')
	query = `SELECT COUNT(DISTINCT c.no_msn) 
				FROM customer_mtr c
				JOIN membership m 
    				ON c.no_msn = m.no_msn
			WHERE c.kd_user_ts = ? 
    			AND c.tgl_call_tele BETWEEN ? AND ? 
    			AND c.sts_membership = 'O'
    			AND m.jns_bayar = 'C'
    			AND m.renewal_ke = (
        			SELECT MAX(m2.renewal_ke)
        			FROM membership m2
        			WHERE m2.no_msn = m.no_msn)`
	err = r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.DataBerminatMembershipCash)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_berminat_membership: %v", err)
	}

	// Query untuk data membership berminat (sts_membership = 'O')
	query = `SELECT COUNT(DISTINCT c.no_msn) 
				FROM customer_mtr c
				JOIN membership m 
    				ON c.no_msn = m.no_msn
			WHERE c.kd_user_ts = ? 
    			AND c.tgl_call_tele BETWEEN ? AND ? 
    			AND c.sts_membership = 'O'
    			AND m.jns_bayar = 'T'
    			AND m.renewal_ke = (
        			SELECT MAX(m2.renewal_ke)
        			FROM membership m2
        			WHERE m2.no_msn = m.no_msn)`
	err = r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.DataBerminatMembershipTransfer)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_berminat_membership: %v", err)
	}

	// Query untuk data membership berminat (sts_membership = 'O')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE kd_user_ts = ? AND tgl_call_tele BETWEEN DATE_SUB(?, INTERVAL 1 MONTH) AND DATE_SUB(?, INTERVAL 1 MONTH) AND sts_membership = 'O'`
	err = r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.DataBerminatMembershipBefore)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_berminat_membership: %v", err)
	}

	// Query untuk data membership sukses (sts_membership = 'O') dan (sts_bayar = 'S' di table membership)
	query = `
	SELECT COALESCE(COUNT(*), 0)
	FROM customer_mtr cm
	JOIN membership m ON cm.no_msn = m.no_msn
	WHERE cm.kd_user_ts = ? 
	AND cm.tgl_call_tele BETWEEN ? AND ? 
	AND cm.sts_membership = 'O' 
	AND m.sts_bayar = 'S'`
	err = r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.DataSuksesMembership)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_sukses_membership: %v", err)
	}

	// Query untuk data membership sukses (sts_membership = 'O') dan (sts_bayar = 'S' di table membership)
	query = `
	SELECT COALESCE(COUNT(*), 0)
	FROM customer_mtr cm
	JOIN membership m ON cm.no_msn = m.no_msn
	WHERE cm.kd_user_ts = ? 
	AND cm.tgl_call_tele BETWEEN DATE_SUB(?, INTERVAL 1 MONTH) AND DATE_SUB(?, INTERVAL 1 MONTH) 
	AND cm.sts_membership = 'O' 
	AND m.sts_bayar = 'S'`
	err = r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.DataSuksesMembershipBefore)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_sukses_membership: %v", err)
	}

	// Query untuk data membership prospect (sts_membership = 'F')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE kd_user_ts = ? AND tgl_call_tele BETWEEN ? AND ? AND sts_membership = 'F'`
	err = r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.DataProspectMembership)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_prospect_membership: %v", err)
	}

	// Query untuk data membership prospect (sts_membership = 'F')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE kd_user_ts = ? AND tgl_call_tele BETWEEN DATE_SUB(?, INTERVAL 1 MONTH) AND DATE_SUB(?, INTERVAL 1 MONTH) AND sts_membership = 'F'`
	err = r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.DataProspectMembershipBefore)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_prospect_membership: %v", err)
	}

	// Query untuk data membership tidak berminat (sts_membership = 'T')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE kd_user_ts = ? AND tgl_call_tele BETWEEN ? AND ? AND sts_membership = 'T'`
	err = r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.DataTidakBerminatMembership)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_tidak_berminat_membership: %v", err)
	}

	// Query untuk data membership tidak berminat (sts_membership = 'T')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE kd_user_ts = ? AND tgl_call_tele BETWEEN DATE_SUB(?, INTERVAL 1 MONTH) AND DATE_SUB(?, INTERVAL 1 MONTH) AND sts_membership = 'T'`
	err = r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.DataTidakBerminatMembershipBefore)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_tidak_berminat_membership: %v", err)
	}

	// Query untuk data membership pending (sts_membership = 'P')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE kd_user_ts = ? AND tgl_call_tele BETWEEN ? AND ? AND sts_membership = 'P'`
	err = r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.DataPendingMembership)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_pending_membership: %v", err)
	}

	// Query untuk data membership pending (sts_membership = 'P')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE kd_user_ts = ? AND tgl_call_tele BETWEEN DATE_SUB(?, INTERVAL 1 MONTH) AND DATE_SUB(?, INTERVAL 1 MONTH) AND sts_membership = 'P'`
	err = r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.DataPendingMembershipBefore)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_pending_membership: %v", err)
	}
	//========================================================
	// Query untuk data asuransi pa berminat (sts_asuransi_pa = 'O')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE kd_user_ts = ? AND tgl_call_tele BETWEEN ? AND ? AND sts_asuransi_pa = 'O'`
	err = r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.DataBerminatPA)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_berminat_membership: %v", err)
	}

	// Query untuk data asuransi pa berminat (sts_asuransi_pa = 'O')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE kd_user_ts = ? AND tgl_call_tele BETWEEN DATE_SUB(?, INTERVAL 1 MONTH) AND DATE_SUB(?, INTERVAL 1 MONTH) AND sts_asuransi_pa = 'O'`
	err = r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.DataBerminatPABefore)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_berminat_membership: %v", err)
	}

	// Query untuk data asuransi_pa sukses (sts_asuransi_pa = 'O') dan (sts_bayar = 'S' di table asuransi_pa)
	query = `
	SELECT COALESCE(COUNT(*), 0)
	FROM customer_mtr cm
	JOIN asuransi_pa a ON cm.no_msn = a.no_msn
	WHERE cm.kd_user_ts = ? 
	AND cm.tgl_call_tele BETWEEN ? AND ? 
	AND cm.sts_asuransi_pa = 'O' 
	AND a.sts_bayar = 'S'`
	err = r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.DataSuksesPA)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_sukses_pa: %v", err)
	}

	// Query untuk data asuransi_pa sukses (sts_asuransi_pa = 'O') dan (sts_bayar = 'S' di table asuransi_pa)
	query = `
	SELECT COALESCE(COUNT(*), 0)
	FROM customer_mtr cm
	JOIN asuransi_pa a ON cm.no_msn = a.no_msn
	WHERE cm.kd_user_ts = ? 
	AND cm.tgl_call_tele BETWEEN DATE_SUB(?, INTERVAL 1 MONTH) AND DATE_SUB(?, INTERVAL 1 MONTH) 
	AND cm.sts_asuransi_pa = 'O' 
	AND a.sts_bayar = 'S'`
	err = r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.DataSuksesPABefore)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_sukses_pa: %v", err)
	}

	// Query untuk data asuransi_pa prospect (sts_asuransi_pa = 'F')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE kd_user_ts = ? AND tgl_call_tele BETWEEN ? AND ? AND sts_asuransi_pa = 'F'`
	err = r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.DataProspectPA)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_prospect_asuransi_pa: %v", err)
	}

	// Query untuk data asuransi_pa prospect (sts_asuransi_pa = 'F')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE kd_user_ts = ? AND tgl_call_tele BETWEEN DATE_SUB(?, INTERVAL 1 MONTH) AND DATE_SUB(?, INTERVAL 1 MONTH) AND sts_asuransi_pa = 'F'`
	err = r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.DataProspectPABefore)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_prospect_asuransi_pa: %v", err)
	}

	// Query untuk data Asuransi PA tidak berminat (sts_asuransi_pa = 'T')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE kd_user_ts = ? AND tgl_call_tele BETWEEN ? AND ? AND sts_asuransi_pa = 'T'`
	err = r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.DataTidakBerminatPA)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_tidak_berminat_pa: %v", err)
	}

	// Query untuk data Asuransi PA tidak berminat (sts_asuransi_pa = 'T')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE kd_user_ts = ? AND tgl_call_tele BETWEEN DATE_SUB(?, INTERVAL 1 MONTH) AND DATE_SUB(?, INTERVAL 1 MONTH) AND sts_asuransi_pa = 'T'`
	err = r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.DataTidakBerminatPABefore)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_tidak_berminat_pa: %v", err)
	}

	// Query untuk data Asuransi PA pending (sts_asuransi_pa = 'P')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE kd_user_ts = ? AND tgl_call_tele BETWEEN ? AND ? AND sts_asuransi_pa = 'P'`
	err = r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.DataPendingPA)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_pending_PA: %v", err)
	}

	// Query untuk data Asuransi PA pending (sts_asuransi_pa = 'P')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE kd_user_ts = ? AND tgl_call_tele BETWEEN DATE_SUB(?, INTERVAL 1 MONTH) AND DATE_SUB(?, INTERVAL 1 MONTH) AND sts_asuransi_pa = 'P'`
	err = r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.DataPendingPABefore)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_pending_PA: %v", err)
	}

	//===================================================

	// Query untuk data asuransi mtr berminat (sts_asuransi_mtr = 'O')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE kd_user_ts = ? AND tgl_call_tele BETWEEN ? AND ? AND sts_asuransi_mtr = 'O'`
	err = r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.DataBerminatMtr)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_berminat_mtr: %v", err)
	}

	// Query untuk data asuransi mtr berminat (sts_asuransi_mtr = 'O')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE kd_user_ts = ? AND tgl_call_tele BETWEEN DATE_SUB(?, INTERVAL 1 MONTH) AND DATE_SUB(?, INTERVAL 1 MONTH) AND sts_asuransi_mtr = 'O'`
	err = r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.DataBerminatMtrBefore)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_berminat_mtr: %v", err)
	}

	// Query untuk data asuransi_pa sukses (sts_asuransi_pa = 'O') dan (sts_bayar = 'S' di table asuransi_pa)
	query = `
	SELECT COALESCE(COUNT(*), 0)
	FROM customer_mtr cm
	JOIN asuransi_mtr a ON cm.no_msn = a.no_msn
	WHERE cm.kd_user_ts = ? 
	AND cm.tgl_call_tele BETWEEN ? AND ? 
	AND cm.sts_asuransi_mtr = 'O' 
	AND a.sts_bayar = 'S'`
	err = r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.DataSuksesMtr)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_sukses_pa: %v", err)
	}

	// Query untuk data asuransi_pa sukses (sts_asuransi_pa = 'O') dan (sts_bayar = 'S' di table asuransi_pa)
	query = `
	SELECT COALESCE(COUNT(*), 0)
	FROM customer_mtr cm
	JOIN asuransi_mtr a ON cm.no_msn = a.no_msn
	WHERE cm.kd_user_ts = ? 
	AND cm.tgl_call_tele BETWEEN DATE_SUB(?, INTERVAL 1 MONTH) AND DATE_SUB(?, INTERVAL 1 MONTH) 
	AND cm.sts_asuransi_mtr = 'O' 
	AND a.sts_bayar = 'S'`
	err = r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.DataSuksesMtrBefore)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_sukses_pa: %v", err)
	}

	// Query untuk data asuransi_mtr prospect (sts_asuransi_mtr = 'F')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE kd_user_ts = ? AND tgl_call_tele BETWEEN ? AND ? AND sts_asuransi_mtr = 'F'`
	err = r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.DataProspectMtr)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_prospect_asuransi_mtr: %v", err)
	}

	// Query untuk data asuransi_mtr prospect (sts_asuransi_mtr = 'F')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE kd_user_ts = ? AND tgl_call_tele BETWEEN DATE_SUB(?, INTERVAL 1 MONTH) AND DATE_SUB(?, INTERVAL 1 MONTH) AND sts_asuransi_mtr = 'F'`
	err = r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.DataProspectMtrBefore)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_prospect_asuransi_mtr: %v", err)
	}

	// Query untuk data Asuransi Mtr tidak berminat (sts_asuransi_mtr = 'T')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE kd_user_ts = ? AND tgl_call_tele BETWEEN ? AND ? AND sts_asuransi_mtr = 'T'`
	err = r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.DataTidakBerminatMtr)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_tidak_berminat_pa: %v", err)
	}

	// Query untuk data Asuransi Mtr tidak berminat (sts_asuransi_mtr = 'T')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE kd_user_ts = ? AND tgl_call_tele BETWEEN DATE_SUB(?, INTERVAL 1 MONTH) AND DATE_SUB(?, INTERVAL 1 MONTH) AND sts_asuransi_mtr = 'T'`
	err = r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.DataTidakBerminatMtrBefore)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_tidak_berminat_pa: %v", err)
	}

	// Query untuk data Asuransi Mtr pending (sts_asuransi_mtr = 'P')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE kd_user_ts = ? AND tgl_call_tele BETWEEN ? AND ? AND sts_asuransi_mtr = 'P'`
	err = r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.DataPendingMtr)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_pending_mtr: %v", err)
	}

	// Query untuk data Asuransi Mtr pending (sts_asuransi_mtr = 'P')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE kd_user_ts = ? AND tgl_call_tele BETWEEN DATE_SUB(?, INTERVAL 1 MONTH) AND DATE_SUB(?, INTERVAL 1 MONTH) AND sts_asuransi_mtr = 'P'`
	err = r.conn.QueryRow(query, username, startDate, endDate).Scan(&rekap.DataPendingMtr)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_pending_mtr: %v", err)
	}
	//================================================================

	rekap.DataBerminatMembershipPerBulan = make(map[int]int)
	for i := 1; i <= 12; i++ {
		rekap.DataBerminatMembershipPerBulan[i] = 0
	}

	// Query untuk menghitung data berminat per bulan dalam tahun ini
	query = `
	SELECT MONTH(tgl_call_tele) AS bulan, COUNT(*) AS jumlah
	FROM customer_mtr
	WHERE kd_user_ts = ? 
	AND YEAR(tgl_call_tele) = YEAR(NOW()) 
	AND sts_membership = 'O'
	GROUP BY MONTH(tgl_call_tele)`

	rows, err := r.conn.Query(query, username)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_berminat_per_bulan: %v", err)
	}
	defer rows.Close()

	// Memasukkan hasil query ke dalam map
	for rows.Next() {
		var bulan, jumlah int
		if err := rows.Scan(&bulan, &jumlah); err != nil {
			return rekap, fmt.Errorf("error scanning data_berminat_per_bulan: %v", err)
		}
		rekap.DataBerminatMembershipPerBulan[bulan] = jumlah
	}

	//Rekap PA
	rekap.DataBerminatPAPerBulan = make(map[int]int)
	for i := 1; i <= 12; i++ {
		rekap.DataBerminatPAPerBulan[i] = 0
	}

	// Query untuk menghitung data berminat per bulan dalam tahun ini
	query = `
	SELECT MONTH(tgl_call_tele) AS bulan, COUNT(*) AS jumlah
	FROM customer_mtr
	WHERE kd_user_ts = ? 
	AND YEAR(tgl_call_tele) = YEAR(NOW()) 
	AND sts_asuransi_pa = 'O'
	GROUP BY MONTH(tgl_call_tele)`

	rows, err = r.conn.Query(query, username)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_berminat_per_bulan: %v", err)
	}
	defer rows.Close()

	// Memasukkan hasil query ke dalam map
	for rows.Next() {
		var bulan, jumlah int
		if err := rows.Scan(&bulan, &jumlah); err != nil {
			return rekap, fmt.Errorf("error scanning data_berminat_per_bulan: %v", err)
		}
		rekap.DataBerminatPAPerBulan[bulan] = jumlah
	}

	//Rekap Mtr
	rekap.DataBerminatMtrPerBulan = make(map[int]int)
	for i := 1; i <= 12; i++ {
		rekap.DataBerminatMtrPerBulan[i] = 0
	}

	// Query untuk menghitung data berminat per bulan dalam tahun ini
	query = `
	SELECT MONTH(tgl_call_tele) AS bulan, COUNT(*) AS jumlah
	FROM customer_mtr
	WHERE kd_user_ts = ? 
	AND YEAR(tgl_call_tele) = YEAR(NOW()) 
	AND sts_asuransi_mtr = 'O'
	GROUP BY MONTH(tgl_call_tele)`

	rows, err = r.conn.Query(query, username)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_berminat_per_bulan: %v", err)
	}
	defer rows.Close()

	// Memasukkan hasil query ke dalam map
	for rows.Next() {
		var bulan, jumlah int
		if err := rows.Scan(&bulan, &jumlah); err != nil {
			return rekap, fmt.Errorf("error scanning data_berminat_per_bulan: %v", err)
		}
		rekap.DataBerminatMtrPerBulan[bulan] = jumlah
	}

	return rekap, nil
}

func (r *customerMtrRepository) ListBerminatMembership(username string, startDate time.Time, endDate time.Time, limit int, pageParams int, search string) ([]response.MinatMembership, int, int, error) {
	var results []response.MinatMembership
	var totalRecords int
	var totalFiltered int // Menyimpan jumlah total data setelah filter diterapkan

	offset := (pageParams - 1) * limit

	// Base query SQL dengan JOIN ke mst_card untuk mengambil jns_card
	baseQuery := `
		FROM tr_wms_faktur3 cm
		LEFT JOIN mst_card mc ON cm.kd_card = mc.kd_card
		WHERE cm.kd_user = ? AND cm.sts_renewal = 'O'
		AND cm.tgl_verifikasi BETWEEN ? AND ?`

	// Menambahkan kondisi pencarian jika search tidak kosong
	var params []interface{}
	params = append(params, username, startDate, endDate)

	if search != "" {
		baseQuery += " AND (cm.no_msn LIKE ? OR cm.nm_customer11 LIKE ?)"
		searchPattern := "%" + search + "%"
		params = append(params, searchPattern, searchPattern)
	}

	// Query untuk menghitung total records (jumlah data sebelum filter pencarian)
	countTotalQuery := `SELECT COUNT(*) FROM tr_wms_faktur3 WHERE kd_user = ? AND sts_renewal = 'O' AND tgl_verifikasi BETWEEN ? AND ?`
	err := r.conn.QueryRow(countTotalQuery, username, startDate, endDate).Scan(&totalRecords)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("error executing total count query: %v", err)
	}

	// Query untuk menghitung totalFiltered (jumlah data setelah filter pencarian)
	countFilteredQuery := `SELECT COUNT(*) ` + baseQuery
	err = r.conn.QueryRow(countFilteredQuery, params...).Scan(&totalFiltered)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("error executing filtered count query: %v", err)
	}

	// Menghitung total pages berdasarkan jumlah data setelah filter
	totalPages := int(math.Ceil(float64(totalFiltered) / float64(limit)))

	// Query utama untuk mengambil data
	query := `SELECT 
			cm.no_msn, 
			cm.nm_customer11, 
			cm.sts_jenis_bayar, 
			cm.tgl_bayar_renewal, 
			cm.print, 
			cm.sts_renewal, 
			cm.sts_kartu, 
			cm.sts_bayar_renewal, 
			mc.jns_card AS kd_card ` + baseQuery + " ORDER BY cm.tgl_verifikasi DESC LIMIT ? OFFSET ?"
	params = append(params, limit, offset)

	// Menjalankan query
	rows, err := r.conn.Query(query, params...)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	// Loop untuk membaca hasil query
	for rows.Next() {
		var customer response.MinatMembership
		err := rows.Scan(
			&customer.NoMsn,
			&customer.NmCustomer,
			&customer.StsJnsBayar,
			&customer.TglBayarRenewal,
			&customer.Print,
			&customer.StsRenewal,
			&customer.StsKartu,
			&customer.StsBayarRenewal,
			&customer.KdCard,
		)
		if err != nil {
			return nil, 0, 0, fmt.Errorf("error scanning row: %v", err)
		}
		results = append(results, customer)
	}

	// Jika ada error saat iterasi
	if err := rows.Err(); err != nil {
		return nil, 0, 0, fmt.Errorf("error iterating rows: %v", err)
	}

	return results, totalPages, totalRecords, nil
}

func (r *customerMtrRepository) ListDataAsuransiPA(username string, startDate, endDate time.Time, limit int, pageParams int, search string) ([]response.ListAsuransi, int, int, error) {
	var results []response.ListAsuransi
	fmt.Println("tanggal", startDate, endDate)

	// Hitung OFFSET untuk paginasi
	offset := (pageParams - 1) * limit

	// Query untuk menghitung total data tanpa filter pencarian
	countTotalQuery := `
		SELECT COUNT(*) 
		FROM db_wkm.customer_mtr 
		WHERE kd_user_ts = ? 
  		AND tgl_call_tele BETWEEN ? AND ?
  		AND sts_asuransi_pa != 'M'; 
	`
	var totalRecords int
	err := r.conn.QueryRow(countTotalQuery, username, startDate, endDate).Scan(&totalRecords)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("error executing total count query: %v", err)
	}

	// Query untuk menghitung total data setelah filter pencarian diterapkan
	countFilteredQuery := `
		SELECT COUNT(*) 
		FROM db_wkm.customer_mtr cm
		LEFT JOIN db_wkm.asuransi_pa ap ON cm.no_msn = ap.no_msn
		WHERE cm.kd_user_ts = ? 
  		AND cm.tgl_call_tele BETWEEN ? AND ?
  		AND cm.sts_asuransi_pa != 'M'
		AND (? = '' OR cm.no_msn LIKE ? OR cm.nm_customer_fkt LIKE ? OR cm.nm_customer_wkm LIKE ?);
	`
	searchPattern := "%" + search + "%"
	var totalFiltered int
	err = r.conn.QueryRow(countFilteredQuery, username, startDate, endDate, search, searchPattern, searchPattern, searchPattern).Scan(&totalFiltered)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("error executing filtered count query: %v", err)
	}

	// Hitung total halaman berdasarkan jumlah data setelah filter
	totalPages := 0
	if limit > 0 {
		totalPages = (totalFiltered + limit - 1) / limit // Ceiling division
	}

	// Query utama dengan JOIN ke asuransi_pa dan filter tambahan
	query := `
		SELECT 
    		cm.no_msn, 
    		cm.nm_customer_fkt, 
    		cm.nm_customer_wkm, 
    		cm.sts_asuransi_pa AS status_asuransi, 
    		ap.tgl_beli, 
    		ap.id_produk AS produk
		FROM db_wkm.customer_mtr cm
		LEFT JOIN db_wkm.asuransi_pa ap ON cm.no_msn = ap.no_msn
		WHERE cm.kd_user_ts = ? 
  		AND cm.tgl_call_tele BETWEEN ? AND ?
  		AND cm.sts_asuransi_pa != 'M' 
		AND (? = '' OR cm.no_msn LIKE ? OR cm.nm_customer_fkt LIKE ? OR cm.nm_customer_wkm LIKE ?)
		ORDER BY ap.tgl_beli DESC
		LIMIT ? OFFSET ?;
	`

	// Menjalankan query utama
	rows, err := r.conn.Query(query, username, startDate, endDate, search, searchPattern, searchPattern, searchPattern, limit, offset)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	// Loop untuk membaca hasil query
	for rows.Next() {
		var customer response.ListAsuransi
		var idProduk sql.NullString // Menyimpan id_produk sementara

		err := rows.Scan(
			&customer.NoMsn,
			&customer.NmCustomerFkt,
			&customer.NmCustomerWkm,
			&customer.StsAsuransi,
			&customer.TglBeli,
			&idProduk, // Ambil id_produk dulu
		)
		if err != nil {
			return nil, 0, 0, fmt.Errorf("error scanning row: %v", err)
		}

		// **Query kedua menggunakan wandaGorm untuk mendapatkan nm_produk**
		var nmProduk string
		err = r.wandaGorm.Raw(`SELECT nm_produk FROM wanda_asuransi.produk WHERE jns_asuransi = 2 AND id_produk = ?`, idProduk).Scan(&nmProduk).Error

		if err != nil {
			nmProduk = "" // Jika error, kosongkan nama produk
		}

		// Set nm_produk ke struct hasil
		customer.IdProduk = nmProduk

		// Tambahkan ke hasil akhir
		results = append(results, customer)
	}

	// Jika ada error saat iterasi
	if err := rows.Err(); err != nil {
		return nil, 0, 0, fmt.Errorf("error iterating rows: %v", err)
	}

	return results, totalPages, totalRecords, nil
}

func (r *customerMtrRepository) ListDataAsuransiMtr(username string, startDate, endDate time.Time, limit, pageParams int, search string) ([]response.ListAsuransi, int, int, error) {
	var results []response.ListAsuransi

	// Hitung OFFSET untuk paginasi
	offset := (pageParams - 1) * limit

	// Query untuk menghitung total data sebelum filter pencarian
	countTotalQuery := `
		SELECT COUNT(*) 
		FROM db_wkm.customer_mtr 
		WHERE kd_user_ts = ? 
  		AND tgl_call_tele BETWEEN ? AND ?
  		AND sts_asuransi_mtr != 'M';
	`
	var totalRecords int
	err := r.conn.QueryRow(countTotalQuery, username, startDate, endDate).Scan(&totalRecords)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("error executing total count query: %v", err)
	}

	// Query untuk menghitung total data setelah filter pencarian diterapkan
	countFilteredQuery := `
		SELECT COUNT(*) 
		FROM db_wkm.customer_mtr cm
		LEFT JOIN db_wkm.asuransi_mtr am ON cm.no_msn = am.no_msn
		WHERE cm.kd_user_ts = ? 
  		AND cm.tgl_call_tele BETWEEN ? AND ?
  		AND cm.sts_asuransi_mtr != 'M'
		%s;
	`
	searchCondition := ""
	args := []interface{}{username, startDate, endDate}

	if search != "" {
		searchCondition = "AND (cm.no_msn LIKE ? OR cm.nm_customer_fkt LIKE ? OR cm.nm_customer_wkm LIKE ?)"
		searchValue := "%" + search + "%"
		args = append(args, searchValue, searchValue, searchValue)
	}

	// Format query dengan kondisi pencarian
	countFilteredQuery = fmt.Sprintf(countFilteredQuery, searchCondition)

	// Menghitung total data setelah filter
	var totalFiltered int
	err = r.conn.QueryRow(countFilteredQuery, args...).Scan(&totalFiltered)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("error executing filtered count query: %v", err)
	}

	// Hitung total halaman berdasarkan jumlah data setelah filter
	totalPages := 0
	if limit > 0 {
		totalPages = (totalFiltered + limit - 1) / limit // Ceiling division
	}

	// Query utama dengan JOIN ke asuransi_mtr
	query := `
		SELECT 
    		cm.no_msn, 
    		cm.nm_customer_fkt, 
    		cm.nm_customer_wkm, 
    		cm.sts_asuransi_mtr AS status_asuransi, 
    		am.tgl_beli, 
    		am.id_produk AS produk
		FROM db_wkm.customer_mtr cm
		LEFT JOIN db_wkm.asuransi_mtr am ON cm.no_msn = am.no_msn
		WHERE cm.kd_user_ts = ? 
  		AND cm.tgl_call_tele BETWEEN ? AND ?
  		AND cm.sts_asuransi_mtr != 'M' 
		%s
		ORDER BY am.tgl_beli DESC
		LIMIT ? OFFSET ?;
	`

	// Format query dengan kondisi pencarian
	query = fmt.Sprintf(query, searchCondition)
	args = append(args, limit, offset)

	// Menjalankan query utama
	rows, err := r.conn.Query(query, args...)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	// Loop untuk membaca hasil query
	for rows.Next() {
		var customer response.ListAsuransi
		var idProduk sql.NullString // Menyimpan id_produk sementara

		err := rows.Scan(
			&customer.NoMsn,
			&customer.NmCustomerFkt,
			&customer.NmCustomerWkm,
			&customer.StsAsuransi,
			&customer.TglBeli,
			&idProduk, // Ambil id_produk dulu
		)
		if err != nil {
			return nil, 0, 0, fmt.Errorf("error scanning row: %v", err)
		}

		// **Query kedua menggunakan wandaGorm untuk mendapatkan nm_produk**
		var nmProduk string
		err = r.wandaGorm.Raw(`SELECT nm_produk FROM wanda_asuransi.produk WHERE jns_asuransi = 1 AND id_produk = ?`, idProduk).Scan(&nmProduk).Error

		if err != nil {
			nmProduk = "" // Jika error, kosongkan nama produk
		}

		// Set nm_produk ke struct hasil
		customer.IdProduk = nmProduk

		// Tambahkan ke hasil akhir
		results = append(results, customer)
	}

	// Jika ada error saat iterasi
	if err := rows.Err(); err != nil {
		return nil, 0, 0, fmt.Errorf("error iterating rows: %v", err)
	}

	return results, totalPages, totalRecords, nil
}
