package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"strings"
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
	AllStatusMasterData(search string, username string,tgl_bayar1 string, tgl_bayar2 string, limit int, pageParams int) []response.AllStatusResponse
	CreateCustomerFFaktur(no_msn string) error
	AllStatusMasterDataCount(search string, username string,tgl_bayar1 string, tgl_bayar2 string) int64
	MasterData(search string, sts string, jns string, username string, limit int, pageParams int) []entity.CustomerMtr
	MasterDataCount(search string, sts string, jns string, username string) int64
	MasterDataBalikan(search string, tgl1 string, tgl2 string, username string, limit int, pageParams int) []response.TelesalesBalikanResponseList
	MasterDataBalikanCount(search string, tgl1 string, tgl2 string, username string) int64
	MasterDataBalikanKonfirmer(search string, tgl1 string, tgl2 string, limit int, pageParams int) []response.TelesalesBalikanResponseList
	MasterDataBalikanKonfirmerCount(search string,tgl1 string, tgl2 string) int64
	SelfCount(kd_user string) int64
	ListAmbilData() []entity.Faktur3
	AmbilDataBalikan(no_msn string, kd_user string) error
	EmpatAmbilData(no_msn string) error
	AmbilDataAllStatus(no_msn string, kd_user string) error
	AmbilDataBalikanKonfirmer(no_msn string, kd_user string) error
	AmbilData(no_msn string, kd_user string) error
	Show(no_msn string) response.TelesalesResponse
	Update(customer entity.CustomerMtr) (entity.CustomerMtr, error)
	UpdateOkeMembership(customer request.CustomerMtr) (entity.CustomerMtr, error)
	RekapTele(username string, startDate time.Time, endDate time.Time) (response.RekapTele, error)
	ListBerminatMembership(username string, startDate time.Time, endDate time.Time, limit int, pageParams int, search string) ([]response.MinatMembership, int, int, error)
	ListDataAsuransiPA(username string, startDate time.Time, endDate time.Time, limit int, pageParams int, search string) ([]response.ListAsuransi, int, int, error)
	ListDataAsuransiMtr(username string, startDate time.Time, endDate time.Time, limit int, pageParams int, search string) ([]response.ListAsuransi, int, int, error)
	RekapLeaderTs(startDate time.Time, endDate time.Time) (response.RekapLeaderTs, error)
	RekapBerminatPerWilayah(startDate time.Time, endDate time.Time) ([]response.RekapBerminatPerWilayah, int, error)
	RekapTransaksi(startDate, endDate time.Time) ([]response.RekapTransaksi, error)
	RekapStatus(startDate, endDate time.Time) ([]response.RekapStatus, error)
	ListPerformanceTs(startDate, endDate time.Time) ([]response.ListPerformanceTs, []response.ListPerformanceTs, error)
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
	err := cR.conn.QueryRowContext(ctx, "select count(*) from customer_mtr where tgl_call_tele=? and kd_user_ts = ?", now.Format("2006-01-02"), kd_user).Scan(&count)
	if err != nil {
		fmt.Println("ini error yaa ", err.Error())
	}
	return count
}

func (cR *customerMtrRepository) MasterDataBalikan(search string, tgl1 string, tgl2 string, username string, limit int, pageParams int) []response.TelesalesBalikanResponseList {
	if pageParams <= 0 {
		pageParams = 1
	}
	offset := (pageParams - 1) * limit
	datas := []response.TelesalesBalikanResponseList{}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := strings.Builder{}
	query.WriteString("select a.no_msn, a.nm_customer11, b.alasan, a.tgl_update_kartu_balikan,mk.nama_kurir from tr_wms_faktur3 a inner join mst_alasan_belum_bayar_kurir b on a.alasan_belum_bayar2 = b.id left join mst_kurir mk on a.kode_kurir = mk.kode_kurir where a.sts_bayar_renewal = 'B' and a.sts_renewal ='O' and a.kd_user = ? and (a.no_msn like ? or a.nm_customer11 like ?) ")
	conditions := []string{}
	now := time.Now()
	var limitDateKartuBalikan time.Time
	var konfirmerMasuk bool
	cR.conn.QueryRowContext(ctx, "select switch from state where type = 'confirmer_masuk'").Scan(&konfirmerMasuk)
	if konfirmerMasuk {
		if now.Weekday() == time.Monday || now.Weekday() == time.Tuesday || now.Weekday() == time.Wednesday {
			limitDateKartuBalikan = now.AddDate(0, 0, -5)
		} else {
			limitDateKartuBalikan = now.AddDate(0, 0, -3)
		}
		conditions = append(conditions, fmt.Sprintf("and a.tgl_update_kartu_balikan >= '%s'", limitDateKartuBalikan.Format("2006-01-02")))
	}
	if tgl1 != "" && tgl2 != "" {
		conditions = append(conditions, fmt.Sprintf("and a.tgl_update_kartu_balikan >='%s' and a.tgl_update_kartu_balikan <= '%s'",tgl1,tgl2))
	}
	conditions = append(conditions, "order by a.tgl_update_kartu_balikan desc  limit ? offset ?")
	query.WriteString(strings.Join(conditions, " "))

	rows, err := cR.conn.QueryContext(ctx, query.String(), username, "%"+search+"%", "%"+search+"%", limit, offset)
	if err != nil {
		fmt.Println("ini error ", err.Error())
		return datas
	}
	defer rows.Close()
	for rows.Next() {
		var data response.TelesalesBalikanResponseList
		if err := rows.Scan(&data.NoMsn, &data.NmCustomerFkt, &data.AlasanBelumBayar2, &data.TglUpdateKartuBalikan, &data.NamaKurir); err != nil {
			log.Fatal(err)
		}
		datas = append(datas, data)
	}

	// query.Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Order("modified desc").Find(&datas)
	return datas
}
func (cR *customerMtrRepository) MasterDataBalikanCount(search string, tgl1 string, tgl2 string, username string) int64 {
	var count int64
	query := cR.connGorm.Where("no_msn like ? or nm_customer11 like ? ", "%"+search+"%","%"+search+"%")
	query.Where("kd_user = ? and sts_bayar_renewal = 'B' and sts_renewal ='O'", username)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	now := time.Now()
	var limitDateKartuBalikan time.Time
	var konfirmerMasuk bool
	cR.conn.QueryRowContext(ctx, "select switch from state where type = 'confirmer_masuk'").Scan(&konfirmerMasuk)
	if konfirmerMasuk {
		if now.Weekday() == time.Monday || now.Weekday() == time.Tuesday || now.Weekday() == time.Wednesday {
			limitDateKartuBalikan = now.AddDate(0, 0, -5)
		} else {
			limitDateKartuBalikan = now.AddDate(0, 0, -3)
		}
		query.Where("tgl_update_kartu_balikan >= ?", limitDateKartuBalikan.Format("2006-01-02"))
	} else {
		if tgl1 != "" && tgl2 != "" {
			query.Where("tgl_update_kartu_balikan >=? and tgl_update_kartu_balikan <= ?", tgl1, tgl2)
		}
	}
	query.Model(&entity.Faktur3{}).Count(&count)
	return count
}

func (cR *customerMtrRepository) MasterDataBalikanKonfirmer(search string, tgl1 string, tgl2 string, limit int, pageParams int) []response.TelesalesBalikanResponseList {
	if pageParams <= 0 {
		pageParams = 1
	}
	offset := (pageParams - 1) * limit
	datas := []response.TelesalesBalikanResponseList{}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()


	query := strings.Builder{}
	query.WriteString("select a.no_msn, a.nm_customer11, b.alasan, a.tgl_update_kartu_balikan, mk.nama_kurir from tr_wms_faktur3 a inner join mst_alasan_belum_bayar_kurir b on a.alasan_belum_bayar2 = b.id left join mst_kurir mk on a.kode_kurir = mk.kode_kurir where a.sts_bayar_renewal = 'B' and a.sts_renewal ='O' and (a.no_msn like ? or a.nm_customer11 like ?) ")
	conditions := []string{}
	now := time.Now()
	var limitDateKartuBalikan time.Time
	var konfirmerMasuk bool
	cR.conn.QueryRowContext(ctx, "select switch from state where type = 'confirmer_masuk'").Scan(&konfirmerMasuk)
	if konfirmerMasuk {
		if now.Weekday() == time.Monday || now.Weekday() == time.Tuesday || now.Weekday() == time.Wednesday {
			limitDateKartuBalikan = now.AddDate(0, 0, -5)
		} else {
			limitDateKartuBalikan = now.AddDate(0, 0, -3)
		}
		if tgl1 != "" && tgl2 != "" {
			var lastDate time.Time
			tgl2Date, _ := time.Parse("2006-01-02", tgl2)
			if tgl2Date.After(limitDateKartuBalikan)  {
				lastDate = limitDateKartuBalikan
			}else {
				lastDate = tgl2Date
			}
			conditions = append(conditions, fmt.Sprintf("and a.tgl_update_kartu_balikan >='%s' and a.tgl_update_kartu_balikan <= '%s'",tgl1,lastDate))
		}else {
			conditions = append(conditions, fmt.Sprintf("and a.tgl_update_kartu_balikan <= '%s'", limitDateKartuBalikan.Format("2006-01-02")))
		}
	}else {
		conditions = append(conditions, "and a.no_msn  = 'a'")
	}
	conditions = append(conditions, " order by a.tgl_update_kartu_balikan desc limit ? offset ?")
	query.WriteString(strings.Join(conditions, " "))
	rows, err := cR.conn.QueryContext(ctx,query.String(), "%"+search+"%","%"+search+"%", limit, offset)
	if err != nil {
		fmt.Println("ini error ", err.Error())
		return datas
	}
	defer rows.Close()
	for rows.Next() {
		var data response.TelesalesBalikanResponseList
		if err := rows.Scan(&data.NoMsn, &data.NmCustomerFkt, &data.AlasanBelumBayar2, &data.TglUpdateKartuBalikan, &data.NamaKurir); err != nil {
			log.Fatal(err)
		}
		datas = append(datas, data)
	}


	// query.Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Order("modified desc").Find(&datas)
	return datas
}
func (cR *customerMtrRepository) MasterDataBalikanKonfirmerCount(search string,tgl1 string, tgl2 string) int64 {
	var count int64
	query := cR.connGorm.Where("no_msn like ? or nm_customer11 like ? ", "%"+search+"%","%"+search+"%")
	query.Where("sts_bayar_renewal = 'B' and sts_renewal ='O'")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	now := time.Now()
	var limitDateKartuBalikan time.Time
	var konfirmerMasuk bool
	cR.conn.QueryRowContext(ctx, "select switch from state where type = 'confirmer_masuk'").Scan(&konfirmerMasuk)
	if konfirmerMasuk {
		if now.Weekday() == time.Monday || now.Weekday() == time.Tuesday || now.Weekday() == time.Wednesday {
			limitDateKartuBalikan = now.AddDate(0, 0, -5)
		} else {
			limitDateKartuBalikan = now.AddDate(0, 0, -3)
		}
		if tgl1 != "" && tgl2 != "" {
			var lastDate time.Time
			tgl2Date, _ := time.Parse("2006-01-02", tgl2)
			if tgl2Date.After(limitDateKartuBalikan)  {
				lastDate = limitDateKartuBalikan
			}else {
				lastDate = tgl2Date
			}
			query.Where("tgl_update_kartu_balikan >=? and tgl_update_kartu_balikan <= ?",tgl1,lastDate)
		}else {
			query.Where("tgl_update_kartu_balikan <= ?", limitDateKartuBalikan.Format("2006-01-02"))
		}
	} else {
		query.Where("no_msn='a'")
	}

	query.Model(&entity.Faktur3{}).Count(&count)
	return count
}


func (cR *customerMtrRepository) AllStatusMasterData(search string, username string,tgl_bayar1 string, tgl_bayar2 string, limit int, pageParams int) []response.AllStatusResponse {
	offset := (pageParams - 1) * limit
	datas := []response.AllStatusResponse{}
	count := 0  
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cR.conn.QueryRowContext(ctx, "select count(*) from tr_wms_faktur3 where (no_msn like ? or nm_customer11 like ? ) and kd_user = ? ", "%"+search+"%","%"+search+"%",username).Scan(&count)
	if count > 0 {
		rows, err := cR.conn.QueryContext(ctx,"select no_msn, nm_customer11, sts_renewal, tgl_verifikasi, sts_bayar_renewal, '3' from tr_wms_faktur3 where (no_msn like ? or nm_customer11 like ? ) and kd_user = ? limit ? offset ?", "%"+search+"%","%"+search+"%",username, limit, offset)
		if err != nil {
			fmt.Println("ini error ", err.Error())
			return datas
		}
		defer rows.Close()
		for rows.Next() {
			var data response.AllStatusResponse
			if err := rows.Scan(&data.NoMsn, &data.NmCustomer, &data.StsMembership, &data.TglVerifikasi,&data.StsBayar, &data.FromTable); err != nil {
				log.Fatal(err)
			}
			datas = append(datas, data)
		}
	}else {
		rows, err := cR.conn.QueryContext(ctx,"select no_msn, nm_customer11, sts_renewal, tgl_verifikasi, sts_bayar_renewal, '4' from tr_wms_faktur4 where (no_msn like ? or nm_customer11 like ? ) and kd_user = ? limit ? offset ?", "%"+search+"%","%"+search+"%",username, limit, offset)
		if err != nil {
			fmt.Println("ini error ", err.Error())
			return datas
		}
		defer rows.Close()
		for rows.Next() {
			var data response.AllStatusResponse
			if err := rows.Scan(&data.NoMsn, &data.NmCustomer, &data.StsMembership, &data.TglVerifikasi,&data.StsBayar, &data.FromTable); err != nil {
				log.Fatal(err)
			}
			datas = append(datas, data)
		}
	}
	return datas
}
func (cR *customerMtrRepository) AllStatusMasterDataCount(search string, username string,tgl_bayar1 string, tgl_bayar2 string) int64 {
	count := 0  
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cR.conn.QueryRowContext(ctx, "select count(*) from tr_wms_faktur3 where (no_msn like ? or nm_customer11 like ? ) and kd_user = ? ", "%"+search+"%","%"+search+"%",username).Scan(&count)
	if count<1 {
		cR.conn.QueryRowContext(ctx, "select count(*) from tr_wms_faktur4 where (no_msn like ? or nm_customer11 like ? ) and kd_user = ? ", "%"+search+"%","%"+search+"%",username).Scan(&count)
	}
	return int64(count)
}

func (cR *customerMtrRepository) MasterData(search string, sts string, jns string, username string, limit int, pageParams int) []entity.CustomerMtr {
	datas := []entity.CustomerMtr{}
	query := cR.connGorm.Select("no_msn, nm_customer_fkt, kd_user_ts, alasan_pending_membership, alasan_pending_asuransi_pa, alasan_pending_asuransi_mtr,tgl_prospect_membership,tgl_prospect_asuransi_pa,tgl_prospect_asuransi_mtr, tgl_call_tele, tgl_faktur").Where("no_msn like ? or nm_customer_wkm like ? or nm_customer_fkt like ? ", "%"+search+"%","%"+search+"%", "%"+search+"%")
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

func (r *customerMtrRepository) ListAmbilDataBalikan() []entity.Faktur3 {
	data := []entity.Faktur3{}
	r.connGorm.Select("no_msn, nm_customer11, tgl_update_kartu_balikan").Where("sts_bayar_renewal = 'B' and tgl_update_kartu_balikan is not null").Limit(100).Find(&data)
	return data
}

func (r *customerMtrRepository) AmbilDataBalikan(no_msn string, kd_user string) error {
	var customerMtr entity.CustomerMtr
	faktur3 := entity.Faktur3{NoMsn: no_msn}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r.conn.QueryRowContext(ctx, "select no_msn, sts_cetak3 from tr_wms_faktur3 where no_msn = ?", no_msn).Scan(&customerMtr.NoMsn, &customerMtr.RenewalKe)
	r.connGorm.Select("no_msn, sts_jenis_bayar, kd_card, tgl_bayar_renewal,sts_kirim, sts_bayar_renewal, kode_kurir").First(&faktur3)
	r.connGorm.Where("no_msn = ? and renewal_ke = ?", no_msn, customerMtr.RenewalKe).Find(&customerMtr)
	if customerMtr.NmCustomerFkt == "" {
		err := r.CreateCustomerFFaktur(no_msn)
		if err != nil {
			fmt.Println("ini error pindah ", err)
			return err
		}else {
			r.connGorm.Where("no_msn = ? and renewal_ke = ?", customerMtr.NoMsn, customerMtr.RenewalKe).First(&customerMtr)
		}
	}
	if customerMtr.StsMembership == "O" {
		membership, err := r.FirstOrCreateMembership(customerMtr)
		if err != nil {
			return err
		}
		r.conn.QueryRowContext(ctx, "select alasan_belum_bayar2, alasan_belum_bayar_detail_kurir from tr_wms_faktur3 where no_msn = ?", no_msn).Scan(&membership.AlasanTdkKurir, &membership.AlasanTdkKurirDetail)
		if membership.KdUserTs != kd_user {
			return fmt.Errorf("data tersebut telah di ambil oleh user lain")
		}
		if membership.StsBayar == "B" {
			return nil
		}else {
			membership.StsBayar = "B"
			r.connGorm.Save(&membership)
			return nil
		}
	}
	return nil
}

func (r *customerMtrRepository) AmbilDataBalikanKonfirmer(no_msn string, kd_user string) error {
	var customerMtr entity.CustomerMtr
	faktur3 := entity.Faktur3{NoMsn: no_msn}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r.conn.QueryRowContext(ctx, "select no_msn, sts_cetak3 from tr_wms_faktur3 where no_msn = ?", no_msn).Scan(&customerMtr.NoMsn, &customerMtr.RenewalKe)
	r.connGorm.Select("no_msn, sts_jenis_bayar, kd_card, tgl_bayar_renewal,sts_kirim, sts_bayar_renewal, kode_kurir").First(&faktur3)
	r.connGorm.Where("no_msn = ? and renewal_ke = ?", no_msn, customerMtr.RenewalKe).Find(&customerMtr)
	if customerMtr.NmCustomerFkt == "" {
		err := r.CreateCustomerFFaktur(no_msn)
		if err != nil {
			fmt.Println("ini error pindah ", err)
			return err
		}else {
			r.connGorm.Where("no_msn = ? and renewal_ke = ?", customerMtr.NoMsn, customerMtr.RenewalKe).First(&customerMtr)
		}
	}
	if customerMtr.StsMembership == "O" {
		membership, err := r.FirstOrCreateMembership(customerMtr)
		if err != nil {
			return err
		}
		r.conn.QueryRowContext(ctx, "select alasan_belum_bayar2, alasan_belum_bayar_detail_kurir from tr_wms_faktur3 where no_msn = ?", no_msn).Scan(&membership.AlasanTdkKurir, &membership.AlasanTdkKurirDetail)
		
		if membership.StsBayar == "B" {
			return nil
		}else {
			membership.KdUserKonfirmer = kd_user
			membership.StsBayar = "B"
			r.connGorm.Save(&membership)
			return nil
		}
	}
	return nil
}

func (r *customerMtrRepository) FirstOrCreateMembership(customerMtr entity.CustomerMtr) (entity.Membership, error) {
	var membership entity.Membership
	faktur3 := entity.Faktur3{NoMsn: customerMtr.NoMsn}
	r.connGorm.Select("no_msn, sts_jenis_bayar, kd_card, tgl_bayar_renewal,sts_kirim, sts_bayar_renewal, kode_kurir").First(&faktur3)
	r.connGorm.Where("no_msn = ? and renewal_ke = ?", customerMtr.NoMsn, customerMtr.RenewalKe).First(&membership)
	if membership.Id == "" {
		jsonBytes, err := json.Marshal(customerMtr)
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return entity.Membership{},err
		}
		err = json.Unmarshal(jsonBytes, &membership)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return entity.Membership{},err
		}
		membership.JnsMembership = faktur3.KdCard
		membership.JnsBayar = faktur3.StsJnsBayar
		membership.KodeKurir = faktur3.KdKurir
		membership.KirimKe = faktur3.StsKirim
		membership.TglJanjiBayar = faktur3.TglBayarRenewal
		membership.KdUserTs = faktur3.KdUser
		membership.TypeKartu = "F"
		membership.StsBayar = faktur3.StsBayarRenewal
		result :=r.connGorm.Save(&membership)
		if result.Error != nil {
			return entity.Membership{},result.Error
		}
	}
	return membership, nil
}
func (r *customerMtrRepository) AmbilDataAllStatus(no_msn string, kd_user string) error {
	var customerMtr entity.CustomerMtr
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r.conn.QueryRowContext(ctx, "select no_msn, sts_cetak3 from tr_wms_faktur3 where no_msn = ?", no_msn).Scan(&customerMtr.NoMsn, &customerMtr.RenewalKe)
	r.connGorm.Where("no_msn = ? and renewal_ke = ?", customerMtr.NoMsn, customerMtr.RenewalKe).Find(&customerMtr)
	if customerMtr.NmCustomerFkt == "" {
		err := r.CreateCustomerFFaktur(no_msn)
		if err != nil {
			fmt.Println("ini error pindah ", err)
			return err
		}else {
			r.connGorm.Where("no_msn = ? and renewal_ke = ?", no_msn, customerMtr.RenewalKe).First(&customerMtr)
		}
	}
	if customerMtr.KdUserTs != kd_user {
		return fmt.Errorf("data tersebut telah di ambil oleh user lain")
	}
	if customerMtr.StsMembership == "O" {
		_, err := r.FirstOrCreateMembership(customerMtr)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *customerMtrRepository) EmpatAmbilData(no_msn string) error {
	count := 0  
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r.conn.QueryContext(ctx, "insert into tr_wms_faktur3 select * from tr_wms_faktur4 where no_msn = ?", no_msn)
	r.conn.QueryRowContext(ctx, "select count(*) from tr_wms_faktur3 where no_msn= ? ", no_msn).Scan(&count)
	if count > 0 {
		r.conn.QueryContext(ctx, "delete from tr_wms_faktur4 where no_msn = ?", no_msn)
	}
	return nil
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
	if stsRenewal.String != "" {
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

func (r *customerMtrRepository) CreateCustomerFFaktur(no_msn string) error {
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

	var count int64
	r.connGorm.Model(&entity.CustomerMtr{}).Where("no_msn = ?", no_msn).Count(&count)
	if count > 0 {
		_, err := r.conn.QueryContext(ctx, queryUpdateAmbilData, no_msn)
		return err
	} else {
		_, err := r.conn.QueryContext(ctx, queryAmbilData, no_msn)
		return err
	}
}

func (r *customerMtrRepository) Show(no_msn string) response.TelesalesResponse {
	var membership entity.Membership
	var asuransiPa entity.AsuransiPA
	var asuransiMtr entity.AsuransiMtr
	var response response.TelesalesResponse
	var produkPa entity.MasterProduk
	var produkMtr entity.MasterProduk
	
	data := entity.CustomerMtr{NoMsn: no_msn}
	r.connGorm.Find(&data)
	
	jsonBytes, _ := json.Marshal(data)
	err := json.Unmarshal(jsonBytes, &response)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r.conn.QueryRowContext(ctx, "select c.no_msn, k.nm_kerja, h.hobby, a.agama, t.nm_tujpak, p.nm_pendidikan, kb.nm_keluar_bln2 from customer_mtr c inner join mst_kerja k on k.kode_kerja2 = c.kode_kerja_fkt inner join hobby h on h.kode_hobby = c.hobby_fkt inner join mst_agama a on a.kd_agama = c.agama_fkt inner join mst_tujuanpakai t on c.tujuan_pakai_fkt = t.id inner join mst_pendidikan p on p.kd_pendidikan = c.kode_didik_fkt inner join mst_keluar_bln kb on kb.keluar_bln2 = c.keluar_bln_fkt where c.no_msn = ?", no_msn).Scan(&no_msn, &response.DescKerjaFkt, &response.DescHobbyFkt, &response.DescAgamaFkt, &response.DescTujuanPakaiFkt, &response.DescDidikFkt, &response.DescKeluarBlnFkt)

	r.conn.QueryRowContext(ctx, "select b.nama_kurir from tr_wms_faktur3 a left join mst_kurir b on a.kode_kurir=b.kode_kurir where a.no_msn = ?", no_msn).Scan(&response.NamaKurir)
	
	r.connGorm.Where("no_msn = ? and renewal_ke = ?", no_msn, data.RenewalKe).Find(&membership)
	if membership.Id != "" {
		r.conn.QueryRowContext(ctx, "select keterangan AS value, harga_pokok + asuransi_motor + asuransi from mst_card where kd_card = ? ", membership.JnsMembership).Scan(&response.JnsMembershipName)
		response.MembershipID =  membership.Id
		response.KdPromoTransfer =  membership.KdPromoTransfer
		response.KirimKe =  membership.KirimKe
		response.JnsBayar =  membership.JnsBayar
		response.JnsMembership =  membership.JnsMembership
		response.TypeKartu =  membership.TypeKartu
		response.StsBayarMembership =  membership.StsBayar
		response.TglJanjiBayar =  membership.TglJanjiBayar.Format("2006-01-02")
		response.AlasanVoidKonfirmasi = membership.AlasanVoidKonfirmasi

		if membership.StsBayar == "B" {
			r.conn.QueryRowContext(ctx, "select b.alasan from tr_wms_faktur3 a inner join mst_alasan_belum_bayar_kurir b on a.alasan_belum_bayar2 = b.id  where a.no_msn = ? ", no_msn).Scan(&response.DescAlasanKurir)
		}
		response.AlasanDetailKurir = membership.AlasanTdkKurirDetail
	}
	r.connGorm.Where("no_msn = ?", no_msn).Find(&asuransiPa)
	if asuransiPa.Id != "" {
		r.wandaGorm.Preload("Vendor").Where("id_produk = ?", asuransiPa.IDProduk).Find(&produkPa)
		response.AsuransiPAID = asuransiPa.Id
		response.IDProdukAsuransiPA = asuransiPa.IDProduk
		response.NamaProdukAsuransiPA = produkPa.NmProduk
		response.NamaVendorPA = produkPa.Vendor.NmVendor
		response.RatePA = produkPa.Rate
		response.AdminPA = produkPa.Admin
		response.AmountAsuransiPA = utils.FormatRupiah(int(asuransiPa.AmountPa))
	}

	r.connGorm.Where("no_msn = ?", no_msn).Find(&asuransiMtr)
	if asuransiMtr.Id != "" {
		r.wandaGorm.Preload("Vendor").Where("id_produk = ?", asuransiMtr.IDProduk).Find(&produkMtr)
		response.AsuransiMTRID = asuransiMtr.Id
		response.IDProdukAsuransiMTR = asuransiMtr.IDProduk
		response.NamaProdukAsuransiMTR = produkMtr.NmProduk
		response.Warna = asuransiMtr.Warna
		response.NamaVendorMTR = produkMtr.Vendor.NmVendor
		response.RateMTR = produkMtr.Rate
		response.AdminMTR = produkMtr.Admin
	}

	if data.TglFaktur != nil {
		response.DescTglFaktur, _ = utils.FormatTanggalIndo(data.TglFaktur.Format("2006-01-02"))
		response.AsuransiMtrTahun = data.TglFaktur.Year()
	}
	if data.TglLahirFkt != nil {
		response.DescTglLahirFkt, _ = utils.FormatTanggalIndo(data.TglLahirFkt.Format("2006-01-02"))
	}
	response.AsuransiNmMtr = data.NmMtr
	response.AsuransiNoMtr = data.NoMtr

	if data.JnsJualFkt == "1" {
		response.JnsJualFktKet = "Tunai"
	} else if data.JnsJualFkt == "2" {
		response.JnsJualFktKet = "Kredit"
	}

	switch data.KetNoInfo {
	case "1":
		response.NoWa = data.NoHpFkt
	case "2":
		response.NoWa = data.NoHpWkm
	case "4":
		response.NoWa = data.NoTelpFkt
	case "5":
		response.NoWa = data.NoTelpWkm
	}

	switch data.NoYgDihubTs {
    case "1":
        response.NoHub = data.NoHpFkt
    case "2":
        response.NoHub = data.NoHpWkm
    case "4":
        response.NoHub = data.NoTelpFkt
    case "5":
        response.NoHub = data.NoTelpWkm
    }

	switch response.StsBayarMembership {
    case "S":
        response.DescStsBayarMembership = "Sudah Bayar"
    case "C":
        response.DescStsBayarMembership = "Cancel"
    case "B":
        response.DescStsBayarMembership = "Belum Bayar"
    case "":
        response.DescStsBayarMembership = ""
    }

	if response.NoKartu2 != "" {
		if len(response.NoKartu2) >= 4 {
			substr := response.NoKartu2[2:4]
			switch substr {
			case "01":
				response.JnsMembershipSebelum = "Basic"
			case "02":
				response.JnsMembershipSebelum = "Gold"
			case "03":
				response.JnsMembershipSebelum = "Platinum"
			case "23":
				response.JnsMembershipSebelum = "Platinum Plus"
			}
		}
	}
	
	return response
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
	isUpdateConfirmer := false
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

	fmt.Println("ini tgl call tele ", jsonMap["tgl_call_tele"])
	

	if jsonMap["sts_membership"] == "O" {
		result := r.connGorm.Where("no_msn = ? and renewal_ke = ?", customer.NoMsn, customer.RenewalKe).First(&membership);
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
			membership.AlasanVoidKonfirmasi = ""
			membership.KdUserKonfirmer = ""
			customerMtrEntity.AlasanTdkMembership = ""
			customerMtrEntity.AlasanPendingMembership = ""
			customerMtrEntity.TglProspectMembership = nil
			r.connGorm.Save(&membership)
		}
		if membership.StsBayar == "B" {
			isUpdateConfirmer = true
			err = json.Unmarshal(jsonBytes, &membership)
			if err != nil {
				fmt.Println("Error decoding JSON Membership:", err)
				return entity.CustomerMtr{}, err
			}
			membership.TglKonfirmasi = &now
			membership.StsKartu = "6"
			membership.StsBayar = ""
			r.connGorm.Save(&membership)
			stmt, err := r.conn.Prepare("UPDATE tr_wms_faktur3 SET sts_kartu='6', sts_bawa_kartu='5', sts_bayar_renewal='',  kd_user11=?, tgl_konfirmasi=?  where no_msn = ?")
			if err != nil {
				log.Fatal("Error preparing statement:", err)
			}
			defer stmt.Close()                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           // Ensure statement is closed after execution
			_, err = stmt.Exec(customer.KdUserKonfirmer, now.Format("2006-01-02"), customer.NoMsn) // Update user with ID 1 to age 28
			if err != nil {
				log.Fatal("Error executing statement:", err)
			}	

		}

	}

	if jsonMap["sts_membership"] == "T"{
		r.connGorm.Where("no_msn = ? and renewal_ke = ?", customer.NoMsn, customer.RenewalKe).First(&membership);
		if membership.Id != "" {
			if membership.StsBayar == "B" {
				isUpdateConfirmer = true
				membership.StsKartu = "8"
				membership.KdUserKonfirmer = customer.KdUserKonfirmer
				membership.StsMembership = "C"
				membership.TglKonfirmasi = &now
				membership.AlasanTdkTsDetail = customer.AlasanTdkTsDetail
				membership.AlasanVoidKonfirmasi = customer.AlasanVoidKonfirmasi
				r.connGorm.Save(&membership)
				stmt, err := r.conn.Prepare("UPDATE tr_wms_faktur3 SET sts_kartu='8', sts_bawa_kartu='', kd_user11=?, tgl_konfirmasi=?, sts_renewal='C',alasan_belum_bayar_detail_telesales=?,alasan_void_konfirmasi=?  where no_msn = ?")
				if err != nil {
					log.Fatal("Error preparing statement:", err)
				}
				defer stmt.Close()                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           // Ensure statement is closed after execution
				_, err = stmt.Exec(customer.KdUserKonfirmer, now.Format("2006-01-02"), customer.AlasanTdkTsDetail, customer.AlasanVoidKonfirmasi, customer.NoMsn) // Update user with ID 1 to age 28
				if err != nil {
					log.Fatal("Error executing statement:", err)
				}
			}	
		}

	}

	if jsonMap["sts_asuransi_mtr"] == "O" {
		result := r.connGorm.Where("no_msn = ? ", customer.NoMsn).First(&entity.AsuransiMtr{})
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
	if jsonMap["sts_asuransi_pa"] == "O" {
		result := r.connGorm.Where("no_msn = ? ", customer.NoMsn).First(&entity.AsuransiPA{})
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
	if isUpdateConfirmer {
		customerMtrEntity.TglKonfirmasi = &now
	}else {
		jsonMap["tgl_call_tele"] = now.Format("2006-01-02")
		customerMtrEntity.TglCallTele = &now
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

func (r *customerMtrRepository) RekapLeaderTs(startDate time.Time, endDate time.Time) (response.RekapLeaderTs, error) {
	var rekap response.RekapLeaderTs
	// Query untuk jumlah data membership
	query := `SELECT COUNT(*) FROM customer_mtr WHERE tgl_call_tele BETWEEN ? AND ?`
	err := r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.JumlahDataMembership)
	if err != nil {
		return rekap, fmt.Errorf("error fetching jumlah_data_membership: %v", err)
	}

	query = `SELECT COUNT(*) FROM customer_mtr 
          WHERE tgl_call_tele BETWEEN DATE_SUB(?, INTERVAL 1 MONTH) 
          AND DATE_SUB(?, INTERVAL 1 MONTH)`

	err = r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.JumlahDataMembershipBefore)
	if err != nil {
		return rekap, fmt.Errorf("error fetching jumlah_data_membership: %v", err)
	}

	// Query untuk data membership berminat (sts_membership = 'O')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE tgl_call_tele BETWEEN ? AND ? AND sts_membership = 'O'`
	err = r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.DataBerminatMembership)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_berminat_membership: %v", err)
	}

	// Query untuk data membership berminat (sts_membership = 'O')
	query = `SELECT COUNT(DISTINCT c.no_msn) 
				FROM customer_mtr c
				JOIN membership m 
    				ON c.no_msn = m.no_msn
			WHERE c.tgl_call_tele BETWEEN ? AND ? 
    			AND c.sts_membership = 'O'
    			AND m.jns_bayar = 'C'
    			AND m.renewal_ke = (
        			SELECT MAX(m2.renewal_ke)
        			FROM membership m2
        			WHERE m2.no_msn = m.no_msn)`
	err = r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.DataBerminatMembershipCash)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_berminat_membership: %v", err)
	}

	// Query untuk data membership berminat (sts_membership = 'O')
	query = `SELECT COUNT(DISTINCT c.no_msn) 
				FROM customer_mtr c
				JOIN membership m 
    				ON c.no_msn = m.no_msn
			WHERE c.tgl_call_tele BETWEEN ? AND ? 
    			AND c.sts_membership = 'O'
    			AND m.jns_bayar = 'T'
    			AND m.renewal_ke = (
        			SELECT MAX(m2.renewal_ke)
        			FROM membership m2
        			WHERE m2.no_msn = m.no_msn)`
	err = r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.DataBerminatMembershipTransfer)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_berminat_membership: %v", err)
	}

	// Query untuk data membership berminat (sts_membership = 'O')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE tgl_call_tele BETWEEN DATE_SUB(?, INTERVAL 1 MONTH) AND DATE_SUB(?, INTERVAL 1 MONTH) AND sts_membership = 'O'`
	err = r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.DataBerminatMembershipBefore)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_berminat_membership: %v", err)
	}

	// Query untuk data membership sukses (sts_membership = 'O') dan (sts_bayar = 'S' di table membership)
	query = `
	SELECT COALESCE(COUNT(*), 0)
	FROM customer_mtr cm
	JOIN membership m ON cm.no_msn = m.no_msn
	WHERE cm.tgl_call_tele BETWEEN ? AND ? 
	AND cm.sts_membership = 'O' 
	AND m.sts_bayar = 'S'`
	err = r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.DataSuksesMembership)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_sukses_membership: %v", err)
	}

	// Query untuk data membership sukses (sts_membership = 'O') dan (sts_bayar = 'S' di table membership)
	query = `
	SELECT COALESCE(COUNT(*), 0)
	FROM customer_mtr cm
	JOIN membership m ON cm.no_msn = m.no_msn
	WHERE cm.tgl_call_tele BETWEEN DATE_SUB(?, INTERVAL 1 MONTH) AND DATE_SUB(?, INTERVAL 1 MONTH) 
	AND cm.sts_membership = 'O' 
	AND m.sts_bayar = 'S'`
	err = r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.DataSuksesMembershipBefore)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_sukses_membership: %v", err)
	}

	// Query untuk data membership prospect (sts_membership = 'F')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE tgl_call_tele BETWEEN ? AND ? AND sts_membership = 'F'`
	err = r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.DataProspectMembership)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_prospect_membership: %v", err)
	}

	// Query untuk data membership prospect (sts_membership = 'F')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE tgl_call_tele BETWEEN DATE_SUB(?, INTERVAL 1 MONTH) AND DATE_SUB(?, INTERVAL 1 MONTH) AND sts_membership = 'F'`
	err = r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.DataProspectMembershipBefore)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_prospect_membership: %v", err)
	}

	// Query untuk data membership tidak berminat (sts_membership = 'T')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE tgl_call_tele BETWEEN ? AND ? AND sts_membership = 'T'`
	err = r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.DataTidakBerminatMembership)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_tidak_berminat_membership: %v", err)
	}

	// Query untuk data membership tidak berminat (sts_membership = 'T')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE tgl_call_tele BETWEEN DATE_SUB(?, INTERVAL 1 MONTH) AND DATE_SUB(?, INTERVAL 1 MONTH) AND sts_membership = 'T'`
	err = r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.DataTidakBerminatMembershipBefore)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_tidak_berminat_membership: %v", err)
	}

	// Query untuk data membership pending (sts_membership = 'P')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE tgl_call_tele BETWEEN ? AND ? AND sts_membership = 'P'`
	err = r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.DataPendingMembership)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_pending_membership: %v", err)
	}

	// Query untuk data membership pending (sts_membership = 'P')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE tgl_call_tele BETWEEN DATE_SUB(?, INTERVAL 1 MONTH) AND DATE_SUB(?, INTERVAL 1 MONTH) AND sts_membership = 'P'`
	err = r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.DataPendingMembershipBefore)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_pending_membership: %v", err)
	}
	//========================================================
	// Query untuk data asuransi pa berminat (sts_asuransi_pa = 'O')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE tgl_call_tele BETWEEN ? AND ? AND sts_asuransi_pa = 'O'`
	err = r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.DataBerminatPA)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_berminat_membership: %v", err)
	}

	// Query untuk data asuransi pa berminat (sts_asuransi_pa = 'O')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE tgl_call_tele BETWEEN DATE_SUB(?, INTERVAL 1 MONTH) AND DATE_SUB(?, INTERVAL 1 MONTH) AND sts_asuransi_pa = 'O'`
	err = r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.DataBerminatPABefore)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_berminat_membership: %v", err)
	}

	// Query untuk data asuransi_pa sukses (sts_asuransi_pa = 'O') dan (sts_bayar = 'S' di table asuransi_pa)
	query = `
	SELECT COALESCE(COUNT(*), 0)
	FROM customer_mtr cm
	JOIN asuransi_pa a ON cm.no_msn = a.no_msn
	WHERE cm.tgl_call_tele BETWEEN ? AND ? 
	AND cm.sts_asuransi_pa = 'O' 
	AND a.sts_bayar = 'S'`
	err = r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.DataSuksesPA)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_sukses_pa: %v", err)
	}

	// Query untuk data asuransi_pa sukses (sts_asuransi_pa = 'O') dan (sts_bayar = 'S' di table asuransi_pa)
	query = `
	SELECT COALESCE(COUNT(*), 0)
	FROM customer_mtr cm
	JOIN asuransi_pa a ON cm.no_msn = a.no_msn
	WHERE cm.tgl_call_tele BETWEEN DATE_SUB(?, INTERVAL 1 MONTH) AND DATE_SUB(?, INTERVAL 1 MONTH) 
	AND cm.sts_asuransi_pa = 'O' 
	AND a.sts_bayar = 'S'`
	err = r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.DataSuksesPABefore)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_sukses_pa: %v", err)
	}

	// Query untuk data asuransi_pa prospect (sts_asuransi_pa = 'F')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE tgl_call_tele BETWEEN ? AND ? AND sts_asuransi_pa = 'F'`
	err = r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.DataProspectPA)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_prospect_asuransi_pa: %v", err)
	}

	// Query untuk data asuransi_pa prospect (sts_asuransi_pa = 'F')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE tgl_call_tele BETWEEN DATE_SUB(?, INTERVAL 1 MONTH) AND DATE_SUB(?, INTERVAL 1 MONTH) AND sts_asuransi_pa = 'F'`
	err = r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.DataProspectPABefore)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_prospect_asuransi_pa: %v", err)
	}

	// Query untuk data Asuransi PA tidak berminat (sts_asuransi_pa = 'T')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE tgl_call_tele BETWEEN ? AND ? AND sts_asuransi_pa = 'T'`
	err = r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.DataTidakBerminatPA)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_tidak_berminat_pa: %v", err)
	}

	// Query untuk data Asuransi PA tidak berminat (sts_asuransi_pa = 'T')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE tgl_call_tele BETWEEN DATE_SUB(?, INTERVAL 1 MONTH) AND DATE_SUB(?, INTERVAL 1 MONTH) AND sts_asuransi_pa = 'T'`
	err = r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.DataTidakBerminatPABefore)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_tidak_berminat_pa: %v", err)
	}

	// Query untuk data Asuransi PA pending (sts_asuransi_pa = 'P')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE tgl_call_tele BETWEEN ? AND ? AND sts_asuransi_pa = 'P'`
	err = r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.DataPendingPA)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_pending_PA: %v", err)
	}

	// Query untuk data Asuransi PA pending (sts_asuransi_pa = 'P')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE tgl_call_tele BETWEEN DATE_SUB(?, INTERVAL 1 MONTH) AND DATE_SUB(?, INTERVAL 1 MONTH) AND sts_asuransi_pa = 'P'`
	err = r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.DataPendingPABefore)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_pending_PA: %v", err)
	}

	//===================================================

	// Query untuk data asuransi mtr berminat (sts_asuransi_mtr = 'O')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE tgl_call_tele BETWEEN ? AND ? AND sts_asuransi_mtr = 'O'`
	err = r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.DataBerminatMtr)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_berminat_mtr: %v", err)
	}

	// Query untuk data asuransi mtr berminat (sts_asuransi_mtr = 'O')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE tgl_call_tele BETWEEN DATE_SUB(?, INTERVAL 1 MONTH) AND DATE_SUB(?, INTERVAL 1 MONTH) AND sts_asuransi_mtr = 'O'`
	err = r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.DataBerminatMtrBefore)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_berminat_mtr: %v", err)
	}

	// Query untuk data asuransi_pa sukses (sts_asuransi_pa = 'O') dan (sts_bayar = 'S' di table asuransi_pa)
	query = `
	SELECT COALESCE(COUNT(*), 0)
	FROM customer_mtr cm
	JOIN asuransi_mtr a ON cm.no_msn = a.no_msn
	WHERE cm.tgl_call_tele BETWEEN ? AND ? 
	AND cm.sts_asuransi_mtr = 'O' 
	AND a.sts_bayar = 'S'`
	err = r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.DataSuksesMtr)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_sukses_pa: %v", err)
	}

	// Query untuk data asuransi_pa sukses (sts_asuransi_pa = 'O') dan (sts_bayar = 'S' di table asuransi_pa)
	query = `
	SELECT COALESCE(COUNT(*), 0)
	FROM customer_mtr cm
	JOIN asuransi_mtr a ON cm.no_msn = a.no_msn
	WHERE cm.tgl_call_tele BETWEEN DATE_SUB(?, INTERVAL 1 MONTH) AND DATE_SUB(?, INTERVAL 1 MONTH) 
	AND cm.sts_asuransi_mtr = 'O' 
	AND a.sts_bayar = 'S'`
	err = r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.DataSuksesMtrBefore)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_sukses_pa: %v", err)
	}

	// Query untuk data asuransi_mtr prospect (sts_asuransi_mtr = 'F')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE tgl_call_tele BETWEEN ? AND ? AND sts_asuransi_mtr = 'F'`
	err = r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.DataProspectMtr)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_prospect_asuransi_mtr: %v", err)
	}

	// Query untuk data asuransi_mtr prospect (sts_asuransi_mtr = 'F')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE tgl_call_tele BETWEEN DATE_SUB(?, INTERVAL 1 MONTH) AND DATE_SUB(?, INTERVAL 1 MONTH) AND sts_asuransi_mtr = 'F'`
	err = r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.DataProspectMtrBefore)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_prospect_asuransi_mtr: %v", err)
	}

	// Query untuk data Asuransi Mtr tidak berminat (sts_asuransi_mtr = 'T')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE tgl_call_tele BETWEEN ? AND ? AND sts_asuransi_mtr = 'T'`
	err = r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.DataTidakBerminatMtr)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_tidak_berminat_pa: %v", err)
	}

	// Query untuk data Asuransi Mtr tidak berminat (sts_asuransi_mtr = 'T')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE tgl_call_tele BETWEEN DATE_SUB(?, INTERVAL 1 MONTH) AND DATE_SUB(?, INTERVAL 1 MONTH) AND sts_asuransi_mtr = 'T'`
	err = r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.DataTidakBerminatMtrBefore)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_tidak_berminat_pa: %v", err)
	}

	// Query untuk data Asuransi Mtr pending (sts_asuransi_mtr = 'P')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE tgl_call_tele BETWEEN ? AND ? AND sts_asuransi_mtr = 'P'`
	err = r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.DataPendingMtr)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_pending_mtr: %v", err)
	}

	// Query untuk data Asuransi Mtr pending (sts_asuransi_mtr = 'P')
	query = `SELECT COUNT(*) FROM customer_mtr WHERE tgl_call_tele BETWEEN DATE_SUB(?, INTERVAL 1 MONTH) AND DATE_SUB(?, INTERVAL 1 MONTH) AND sts_asuransi_mtr = 'P'`
	err = r.conn.QueryRow(query, startDate, endDate).Scan(&rekap.DataPendingMtr)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_pending_mtr: %v", err)
	}
	//=========================BASIC=======================================

	rekap.DataMemberBasicPerBulan = make(map[int]int)
	for i := 1; i <= 12; i++ {
		rekap.DataMemberBasicPerBulan[i] = 0
	}

	// Query untuk menghitung data berminat per bulan dalam tahun ini
	query = `
		SELECT 
    	MONTH(c.tgl_call_tele) AS bulan, 
    	COUNT(*) AS jumlah
			FROM customer_mtr c
		JOIN membership m ON c.no_msn = m.no_msn
		WHERE 
    		YEAR(c.tgl_call_tele) = YEAR(NOW()) 
    		AND c.sts_membership = 'O'
    		AND m.jns_membership = 'R104'
		GROUP BY MONTH(c.tgl_call_tele)`

	rows, err := r.conn.Query(query)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_member_basic_per_bulan: %v", err)
	}
	defer rows.Close()

	// Memasukkan hasil query ke dalam map
	for rows.Next() {
		var bulan, jumlah int
		if err := rows.Scan(&bulan, &jumlah); err != nil {
			return rekap, fmt.Errorf("error scanning data_member_basic_per_bulan: %v", err)
		}
		rekap.DataMemberBasicPerBulan[bulan] = jumlah
	}

	//========================GOLD======================================

	rekap.DataMemberGoldPerBulan = make(map[int]int)
	for i := 1; i <= 12; i++ {
		rekap.DataMemberGoldPerBulan[i] = 0
	}

	// Query untuk menghitung data berminat per bulan dalam tahun ini
	query = `
		SELECT 
    	MONTH(c.tgl_call_tele) AS bulan, 
    	COUNT(*) AS jumlah
			FROM customer_mtr c
		JOIN membership m ON c.no_msn = m.no_msn
		WHERE 
    		YEAR(c.tgl_call_tele) = YEAR(NOW()) 
    		AND c.sts_membership = 'O'
    		AND m.jns_membership = 'R204'
		GROUP BY MONTH(c.tgl_call_tele)`

	rows, err = r.conn.Query(query)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_member_gold_per_bulan: %v", err)
	}
	defer rows.Close()

	// Memasukkan hasil query ke dalam map
	for rows.Next() {
		var bulan, jumlah int
		if err := rows.Scan(&bulan, &jumlah); err != nil {
			return rekap, fmt.Errorf("error scanning data_member_gold_per_bulan: %v", err)
		}
		rekap.DataMemberGoldPerBulan[bulan] = jumlah
	}

	//==========================PLATINUM====================================
	rekap.DataMemberPlatPerBulan = make(map[int]int)
	for i := 1; i <= 12; i++ {
		rekap.DataMemberPlatPerBulan[i] = 0
	}

	// Query untuk menghitung data berminat per bulan dalam tahun ini
	query = `
		SELECT 
    	MONTH(c.tgl_call_tele) AS bulan, 
    	COUNT(*) AS jumlah
			FROM customer_mtr c
		JOIN membership m ON c.no_msn = m.no_msn
		WHERE 
    		YEAR(c.tgl_call_tele) = YEAR(NOW()) 
    		AND c.sts_membership = 'O'
    		AND m.jns_membership IN ('R314','R315')
		GROUP BY MONTH(c.tgl_call_tele)`

	rows, err = r.conn.Query(query)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_member_plat_per_bulan: %v", err)
	}
	defer rows.Close()

	// Memasukkan hasil query ke dalam map
	for rows.Next() {
		var bulan, jumlah int
		if err := rows.Scan(&bulan, &jumlah); err != nil {
			return rekap, fmt.Errorf("error scanning data_member_plat_per_bulan: %v", err)
		}
		rekap.DataMemberPlatPerBulan[bulan] = jumlah
	}

	//=======================PLATINUM PLUS=======================================

	rekap.DataMemberPlatPlusPerBulan = make(map[int]int)
	for i := 1; i <= 12; i++ {
		rekap.DataMemberPlatPlusPerBulan[i] = 0
	}

	// Query untuk menghitung data berminat per bulan dalam tahun ini
	query = `
		SELECT 
    	MONTH(c.tgl_call_tele) AS bulan, 
    	COUNT(*) AS jumlah
			FROM customer_mtr c
		JOIN membership m ON c.no_msn = m.no_msn
		WHERE 
    		YEAR(c.tgl_call_tele) = YEAR(NOW()) 
    		AND c.sts_membership = 'O'
    		AND m.jns_membership IN ('R318','R319')
		GROUP BY MONTH(c.tgl_call_tele)`

	rows, err = r.conn.Query(query)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_member_plat_plus_per_bulan: %v", err)
	}
	defer rows.Close()

	// Memasukkan hasil query ke dalam map
	for rows.Next() {
		var bulan, jumlah int
		if err := rows.Scan(&bulan, &jumlah); err != nil {
			return rekap, fmt.Errorf("error scanning data_member_plat_plus_per_bulan: %v", err)
		}
		rekap.DataMemberPlatPlusPerBulan[bulan] = jumlah
	}

	//=======================Panda 1=======================================

	//Rekap PA
	rekap.DataPanda1PerBulan = make(map[int]int)
	for i := 1; i <= 12; i++ {
		rekap.DataPanda1PerBulan[i] = 0
	}

	// Query untuk menghitung data berminat per bulan dalam tahun ini
	query = `
	SELECT 
    	MONTH(c.tgl_call_tele) AS bulan, 
    	COUNT(*) AS jumlah
			FROM customer_mtr c
		JOIN asuransi_pa a ON c.no_msn = a.no_msn
		WHERE 
    		YEAR(c.tgl_call_tele) = YEAR(NOW()) 
    		AND c.sts_membership = 'O'
    		AND a.id_produk = 'PRODUK-004'
		GROUP BY MONTH(c.tgl_call_tele)`

	rows, err = r.conn.Query(query)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_panda1_per_bulan: %v", err)
	}
	defer rows.Close()

	// Memasukkan hasil query ke dalam map
	for rows.Next() {
		var bulan, jumlah int
		if err := rows.Scan(&bulan, &jumlah); err != nil {
			return rekap, fmt.Errorf("error scanning data_panda1_per_bulan: %v", err)
		}
		rekap.DataPanda1PerBulan[bulan] = jumlah
	}
	//=======================Panda 2=======================================

	//Rekap PA
	rekap.DataPanda2PerBulan = make(map[int]int)
	for i := 1; i <= 12; i++ {
		rekap.DataPanda1PerBulan[i] = 0
	}

	// Query untuk menghitung data berminat per bulan dalam tahun ini
	query = `
	SELECT 
    	MONTH(c.tgl_call_tele) AS bulan, 
    	COUNT(*) AS jumlah
			FROM customer_mtr c
		JOIN asuransi_pa a ON c.no_msn = a.no_msn
		WHERE 
    		YEAR(c.tgl_call_tele) = YEAR(NOW()) 
    		AND c.sts_membership = 'O'
    		AND a.id_produk = 'PRODUK-005'
		GROUP BY MONTH(c.tgl_call_tele)`

	rows, err = r.conn.Query(query)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_panda2_per_bulan: %v", err)
	}
	defer rows.Close()

	// Memasukkan hasil query ke dalam map
	for rows.Next() {
		var bulan, jumlah int
		if err := rows.Scan(&bulan, &jumlah); err != nil {
			return rekap, fmt.Errorf("error scanning data_panda2_per_bulan: %v", err)
		}
		rekap.DataPanda2PerBulan[bulan] = jumlah
	}
	//=======================Panda 3=======================================
	//Rekap PA
	rekap.DataPanda3PerBulan = make(map[int]int)
	for i := 1; i <= 12; i++ {
		rekap.DataPanda3PerBulan[i] = 0
	}

	// Query untuk menghitung data berminat per bulan dalam tahun ini
	query = `
	SELECT 
    	MONTH(c.tgl_call_tele) AS bulan, 
    	COUNT(*) AS jumlah
			FROM customer_mtr c
		JOIN asuransi_pa a ON c.no_msn = a.no_msn
		WHERE 
    		YEAR(c.tgl_call_tele) = YEAR(NOW()) 
    		AND c.sts_membership = 'O'
    		AND a.id_produk = 'PRODUK-006'
		GROUP BY MONTH(c.tgl_call_tele)`

	rows, err = r.conn.Query(query)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_panda3_per_bulan: %v", err)
	}
	defer rows.Close()

	// Memasukkan hasil query ke dalam map
	for rows.Next() {
		var bulan, jumlah int
		if err := rows.Scan(&bulan, &jumlah); err != nil {
			return rekap, fmt.Errorf("error scanning data_panda3_per_bulan: %v", err)
		}
		rekap.DataPanda3PerBulan[bulan] = jumlah
	}
	//=======================TLO=======================================

	//Rekap Mtr
	rekap.DataTloPerBulan = make(map[int]int)
	for i := 1; i <= 12; i++ {
		rekap.DataTloPerBulan[i] = 0
	}

	// Query untuk menghitung data berminat per bulan dalam tahun ini
	query = `
	SELECT 
    	MONTH(c.tgl_call_tele) AS bulan, 
    	COUNT(*) AS jumlah
			FROM customer_mtr c
		JOIN asuransi_mtr a ON c.no_msn = a.no_msn
		WHERE 
    		YEAR(c.tgl_call_tele) = YEAR(NOW()) 
    		AND c.sts_membership = 'O'
    		AND a.id_produk = 'PRODUK-001'
		GROUP BY MONTH(c.tgl_call_tele)`

	rows, err = r.conn.Query(query)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_tlo_per_bulan: %v", err)
	}
	defer rows.Close()

	// Memasukkan hasil query ke dalam map
	for rows.Next() {
		var bulan, jumlah int
		if err := rows.Scan(&bulan, &jumlah); err != nil {
			return rekap, fmt.Errorf("error scanning data_tlo_per_bulan: %v", err)
		}
		rekap.DataTloPerBulan[bulan] = jumlah
	}

	//=======================TLO +=======================================

	//Rekap Mtr
	rekap.DataTloPlusPerBulan = make(map[int]int)
	for i := 1; i <= 12; i++ {
		rekap.DataTloPlusPerBulan[i] = 0
	}

	// Query untuk menghitung data berminat per bulan dalam tahun ini
	query = `
	SELECT 
    	MONTH(c.tgl_call_tele) AS bulan, 
    	COUNT(*) AS jumlah
			FROM customer_mtr c
		JOIN asuransi_mtr a ON c.no_msn = a.no_msn
		WHERE 
    		YEAR(c.tgl_call_tele) = YEAR(NOW()) 
    		AND c.sts_membership = 'O'
    		AND a.id_produk = 'PRODUK-002'
		GROUP BY MONTH(c.tgl_call_tele)`

	rows, err = r.conn.Query(query)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_tlo_plus_per_bulan: %v", err)
	}
	defer rows.Close()

	// Memasukkan hasil query ke dalam map
	for rows.Next() {
		var bulan, jumlah int
		if err := rows.Scan(&bulan, &jumlah); err != nil {
			return rekap, fmt.Errorf("error scanning data_tlo_plus_per_bulan: %v", err)
		}
		rekap.DataTloPlusPerBulan[bulan] = jumlah
	}
	//=======================Komersial=======================================

	//Rekap Mtr
	rekap.DataKomersialPerBulan = make(map[int]int)
	for i := 1; i <= 12; i++ {
		rekap.DataKomersialPerBulan[i] = 0
	}

	// Query untuk menghitung data berminat per bulan dalam tahun ini
	query = `
	SELECT 
    	MONTH(c.tgl_call_tele) AS bulan, 
    	COUNT(*) AS jumlah
			FROM customer_mtr c
		JOIN asuransi_mtr a ON c.no_msn = a.no_msn
		WHERE 
    		YEAR(c.tgl_call_tele) = YEAR(NOW()) 
    		AND c.sts_membership = 'O'
    		AND a.id_produk = 'PRODUK-003'
		GROUP BY MONTH(c.tgl_call_tele)`

	rows, err = r.conn.Query(query)
	if err != nil {
		return rekap, fmt.Errorf("error fetching data_komersial_per_bulan: %v", err)
	}
	defer rows.Close()

	// Memasukkan hasil query ke dalam map
	for rows.Next() {
		var bulan, jumlah int
		if err := rows.Scan(&bulan, &jumlah); err != nil {
			return rekap, fmt.Errorf("error scanning data_komersial_per_bulan: %v", err)
		}
		rekap.DataKomersialPerBulan[bulan] = jumlah
	}

	return rekap, nil
}
func (r *customerMtrRepository) RekapBerminatPerWilayah(startDate time.Time, endDate time.Time) ([]response.RekapBerminatPerWilayah, int, error) {
	query := `
		SELECT 
    CASE 
        WHEN (m.kirim_ke = 1 AND cm.kota_wkm LIKE 'JAK%') OR 
             (m.kirim_ke = 2 AND cm.kota_ktr_wkm LIKE 'JAK%') 
        THEN 'JAKARTA'
        WHEN (m.kirim_ke = 1 AND cm.kota_wkm LIKE '%TANGERANG%') OR 
             (m.kirim_ke = 2 AND cm.kota_ktr_wkm LIKE '%TANGERANG%') 
        THEN 'TANGERANG'
        ELSE 
            CASE 
                WHEN m.kirim_ke = 1 THEN cm.kota_wkm
                WHEN m.kirim_ke = 2 THEN cm.kota_ktr_wkm
            END
    END AS kota,
    CASE 
        WHEN m.kirim_ke = 1 THEN cm.kec_wkm
        WHEN m.kirim_ke = 2 THEN cm.kec_ktr_wkm
    END AS kecamatan,
    COUNT(*) AS jumlah,
    (SELECT COUNT(*) 
     	FROM db_wkm.customer_mtr 
     	WHERE sts_membership = 'O' 
     	AND tgl_call_tele BETWEEN ? AND ?) AS total_data
	FROM db_wkm.customer_mtr cm
	JOIN db_wkm.membership m ON cm.no_msn = m.no_msn
	WHERE cm.sts_membership = 'O' 
    	AND cm.tgl_call_tele BETWEEN ? AND ?
	GROUP BY kota, kecamatan
	`

	rows, err := r.conn.Query(query, startDate, endDate,startDate, endDate)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var rekap []response.RekapBerminatPerWilayah
	var totalData int

	for rows.Next() {
		var data response.RekapBerminatPerWilayah
		err := rows.Scan(&data.Kota, &data.Kecamatan, &data.Jumlah, &totalData)
		if err != nil {
			return nil, 0, err
		}
		rekap = append(rekap, data)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	return rekap, totalData, nil
}

func (r *customerMtrRepository) RekapTransaksi(startDate, endDate time.Time) ([]response.RekapTransaksi, error) {
	query := `
		SELECT 
    u.name AS nama_user,
    COUNT(CASE WHEN f.NO_MSN <> '' THEN f.NO_MSN ELSE NULL END) AS jml_data,
    COUNT(CASE WHEN f.STS_JENIS_BAYAR = 'C' AND f.STS_RENEWAL = 'O' THEN f.STS_JENIS_BAYAR ELSE NULL END) AS renewal_ok_cash,
    COUNT(CASE WHEN f.STS_JENIS_BAYAR = 'T' AND f.STS_RENEWAL = 'O' THEN f.STS_JENIS_BAYAR ELSE NULL END) AS renewal_ok_transfer,
    COUNT(CASE WHEN f.STS_JENIS_BAYAR = 'Q' AND f.STS_RENEWAL = 'O' THEN f.STS_JENIS_BAYAR ELSE NULL END) AS renewal_ok_digital,
    COUNT(CASE WHEN f.kd_card LIKE 'R1%' AND f.STS_RENEWAL = 'O' THEN f.STS_JENIS_BAYAR ELSE NULL END) AS BASIC,
    COUNT(CASE WHEN f.kd_card LIKE 'R2%' AND f.STS_RENEWAL = 'O' THEN f.STS_JENIS_BAYAR ELSE NULL END) AS GOLD,
    COUNT(CASE WHEN f.kd_card IN ('R301','R302','R303','R304','R305','R314','R315') AND f.STS_RENEWAL = 'O' THEN f.STS_JENIS_BAYAR ELSE NULL END) AS PLATINUM,
    COUNT(CASE WHEN f.kd_card IN ('R307','R306','R308','R309','R316','R317','R318','R319') AND f.STS_RENEWAL = 'O' THEN f.STS_JENIS_BAYAR ELSE NULL END) AS PLATINUMP
FROM db_wkm.tr_wms_faktur3 f
LEFT JOIN users.mst_users u ON f.kd_user = u.username
WHERE 
    f.tgl_verifikasi BETWEEN ? AND ?
    AND f.kd_client = '1' 
GROUP BY f.kd_user, u.name

UNION ALL

SELECT 
    'TOTAL' AS nama_user,
    COUNT(CASE WHEN NO_MSN <> '' THEN NO_MSN ELSE NULL END) AS jml_data,
    COUNT(CASE WHEN STS_JENIS_BAYAR = 'C' AND STS_RENEWAL = 'O' THEN STS_JENIS_BAYAR ELSE NULL END) AS renewal_ok_cash,
    COUNT(CASE WHEN STS_JENIS_BAYAR = 'T' AND STS_RENEWAL = 'O' THEN STS_JENIS_BAYAR ELSE NULL END) AS renewal_ok_transfer,
    COUNT(CASE WHEN STS_JENIS_BAYAR = 'Q' AND STS_RENEWAL = 'O' THEN STS_JENIS_BAYAR ELSE NULL END) AS renewal_ok_digital,
    COUNT(CASE WHEN kd_card LIKE 'R1%' AND STS_RENEWAL = 'O' THEN STS_JENIS_BAYAR ELSE NULL END) AS BASIC,
    COUNT(CASE WHEN kd_card LIKE 'R2%' AND STS_RENEWAL = 'O' THEN STS_JENIS_BAYAR ELSE NULL END) AS GOLD,
    COUNT(CASE WHEN kd_card IN ('R301','R302','R303','R304','R305','R314','R315') AND STS_RENEWAL = 'O' THEN STS_JENIS_BAYAR ELSE NULL END) AS PLATINUM,
    COUNT(CASE WHEN kd_card IN ('R307','R306','R308','R309','R316','R317','R318','R319') AND STS_RENEWAL = 'O' THEN STS_JENIS_BAYAR ELSE NULL END) AS PLATINUMP
FROM db_wkm.tr_wms_faktur3
WHERE 
    tgl_verifikasi BETWEEN ? AND ?
    AND kd_client = '1' 

	`

	rows, err := r.conn.Query(query, startDate, endDate, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []response.RekapTransaksi

	for rows.Next() {
		var data response.RekapTransaksi
		err := rows.Scan(
			&data.NamaUser, &data.JmlData, &data.RenewalOkCash, &data.RenewalOkTransfer,
			&data.RenewalOkDigital, &data.Basic, &data.Gold, &data.Platinum, &data.PlatinumP,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, data)
	}

	return result, nil
}
func (r *customerMtrRepository) ListPerformanceTs(startDate, endDate time.Time) ([]response.ListPerformanceTs, []response.ListPerformanceTs, error) {
	query := `
		SELECT 
    su.nama_user, 
    su.jumlah_sukses,
    ROUND((su.jumlah_sukses / COALESCE(t.total, 1)) * 100, 2) AS contribution
FROM (
    SELECT 
        m.kd_user_ts AS nama_user, 
        COUNT(*) AS jumlah_sukses
    FROM membership m
    WHERE m.sts_bayar = 'S' 
        AND DATE(m.created_at) BETWEEN ? AND ?
    GROUP BY m.kd_user_ts
) su
CROSS JOIN (
    SELECT SUM(jumlah_sukses) AS total 
    FROM (
        SELECT COUNT(*) AS jumlah_sukses
        FROM membership 
        WHERE sts_bayar = 'S' 
            AND DATE(created_at) BETWEEN ? AND ?
        GROUP BY kd_user_ts
    ) AS temp
) t
ORDER BY su.jumlah_sukses DESC
	`

	// Eksekusi query
	rows, err := r.conn.Query(query, startDate, endDate,startDate, endDate)
	if err != nil {
		return nil, nil, fmt.Errorf("gagal mengambil data: %w", err)
	}
	defer rows.Close()

	// Menampung data hasil query
	var performanceList []response.ListPerformanceTs

	for rows.Next() {
		var data response.ListPerformanceTs
		if err := rows.Scan(&data.NamaUser, &data.JumlahSukses, &data.Contribution); err != nil {
			return nil, nil, fmt.Errorf("gagal membaca hasil query: %w", err)
		}
		performanceList = append(performanceList, data)
	}

	// Jika tidak ada data
	if len(performanceList) == 0 {
		return nil, nil, nil
	}

	// Ambil Top 5 & Low 5
	var top5, low5 []response.ListPerformanceTs
	if len(performanceList) > 5 {
		top5 = performanceList[:5]                           // Top 5 user dengan jumlah sukses tertinggi
		low5 = performanceList[len(performanceList)-5:] // Low 5 user dengan jumlah sukses terendah
	} else {
		top5 = performanceList
		low5 = performanceList
	}

	return top5, low5, nil
}

func (r *customerMtrRepository) RekapStatus(startDate, endDate time.Time) ([]response.RekapStatus, error) {
	var results []response.RekapStatus

	// Query utama untuk mendapatkan data rekap
	query := `
		SELECT kd_user,
			COUNT(CASE WHEN NO_MSN<>'' THEN NO_MSN ELSE NULL END) AS jml_data,
			COUNT(CASE WHEN TERIMA_KARTU='S' THEN TERIMA_KARTU ELSE NULL END) AS sudah_terima,
			COUNT(CASE WHEN TERIMA_KARTU='B' THEN TERIMA_KARTU ELSE NULL END) AS belum_terima,
			COUNT(CASE WHEN STS_JENIS_BAYAR='C' AND STS_RENEWAL='O' AND sts_cetak IN ('5','6') THEN STS_JENIS_BAYAR ELSE NULL END) AS renewal_ok_cash_update,
			COUNT(CASE WHEN STS_JENIS_BAYAR='C' AND STS_RENEWAL='O' AND sts_cetak NOT IN ('5','6') THEN STS_JENIS_BAYAR ELSE NULL END) AS renewal_ok_cash,
			COUNT(CASE WHEN STS_JENIS_BAYAR='T' AND STS_RENEWAL='O' THEN STS_JENIS_BAYAR ELSE NULL END) AS renewal_ok_transfer,
			COUNT(CASE WHEN ALASAN_PENDING_RENEWAL='1' AND STS_RENEWAL='P' THEN ALASAN_PENDING_RENEWAL ELSE NULL END) AS pikir2_ragu2,
			COUNT(CASE WHEN ALASAN_PENDING_RENEWAL='2' AND STS_RENEWAL='P' THEN ALASAN_PENDING_RENEWAL ELSE NULL END) AS telp_kembali,
			COUNT(CASE WHEN ALASAN_PENDING_RENEWAL IN ('3','') AND STS_RENEWAL='P' THEN ALASAN_PENDING_RENEWAL ELSE NULL END) AS tdk_diangkat,
			COUNT(CASE WHEN ALASAN_PENDING_RENEWAL='4' AND STS_RENEWAL='P' THEN ALASAN_PENDING_RENEWAL ELSE NULL END) AS blm_regist,
			COUNT(CASE WHEN STS_RENEWAL='F' THEN STS_RENEWAL ELSE NULL END) AS prospek,
			COUNT(CASE WHEN kd_card='16' AND sts_renewal='O' AND sts_jenis_bayar='C' THEN kd_card ELSE NULL END) AS basic,
			COUNT(CASE WHEN kd_card='25' AND sts_renewal='O' AND sts_jenis_bayar='C' THEN kd_card ELSE NULL END) AS gold,
			COUNT(CASE WHEN kd_card='38' AND sts_renewal='O' AND sts_jenis_bayar='C' THEN kd_card ELSE NULL END) AS platinum
		FROM tr_wms_faktur3
		WHERE tgl_verifikasi >= ? AND tgl_verifikasi <= ? AND kd_client = '1'
		GROUP BY kd_user;
	`

	rows, err := r.conn.Query(query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item response.RekapStatus
		err := rows.Scan(
			&item.KdUser,
			&item.JmlData,
			&item.SudahTerima,
			&item.BelumTerima,
			&item.RenewalOkCashUpdate,
			&item.RenewalOkCash,
			&item.RenewalOkTransfer,
			&item.PikirRagu,
			&item.TelpKembali,
			&item.TidakDiangkat,
			&item.BelumRegist,
			&item.Prospek,
			&item.Basic,
			&item.Gold,
			&item.Platinum,
		)
		if err != nil {
			return nil, err
		}

		// Menambahkan alasan tidak renewal dari 1 - 24
		alasanTidakRenewal := make(map[string]int)
		for i := 1; i <= 24; i++ {
			var count int
			queryAlasan := `
				SELECT COUNT(*) FROM tr_wms_faktur3 
				WHERE STS_RENEWAL='T' AND ALASAN_TDK_RENEWAL=? 
				AND tgl_verifikasi >= ? AND tgl_verifikasi <= ? AND kd_client = '1'
			`
			err := r.conn.QueryRow(queryAlasan, i, startDate, endDate).Scan(&count)
			if err != nil {
				return nil, err
			}
			alasanTidakRenewal[fmt.Sprintf("%d", i)] = count
		}

		item.AlasanTidakRenewal = alasanTidakRenewal
		results = append(results, item)
	}

	return results, nil
}

