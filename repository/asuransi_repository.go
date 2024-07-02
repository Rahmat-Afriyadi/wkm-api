package repository

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"
	"wkm/entity"
	"wkm/utils"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type AsuransiRepository interface {
	MasterData(search string, dataSource string, sts string, username string, tgl1 string, tgl2 string, ap string, limit int, pageParams int) []entity.MasterAsuransi
	MasterDataCount(search string, dataSource string, sts string, username string, tgl1 string, tgl2 string, ap string) int64
	FindAsuransiByNoMsn(no_msn string) entity.MasterAsuransi
	UpdateAmbilAsuransi(no_msn string, kd_user string)
	UpdateAsuransi(data entity.MasterAsuransi) entity.MasterAsuransi
	UpdateAsuransiBerminat(no_msn string)
	UpdateAsuransiBatalBayar(no_msn string)
	MasterDataGorm()
	MasterAlasanPending() []entity.MasterAlasanPending
	MasterAlasanTdkBerminat() []entity.MasterAlasanTdkBerminat
	RekapByStatus(u string, tgl1 string, tgl2 string) entity.MasterStatusAsuransi
	RekapByStatusAll(u string, tgl1 string, tgl2 string) entity.MasterStatusAsuransi
	MasterDataRekapTele(tgl1 string, tgl2 string) []entity.MasterRekapTele
	RekapByStatusJenisSource(tglStart string, tglEnd string) []map[string]interface{}
	RekapByStatusKdUser(tglStart string, tglEnd string) []map[string]interface{}
	RekapByAlasanPending(tgl1 string, tgl2 string) []map[string]interface{}
	RekapByAlasanPendingKdUser(tgl1 string, tgl2 string) []map[string]interface{}
	RincianByAlasanPendingKdUser(tgl1 string, tgl2 string) []map[string]interface{}
	RincianByAlasanTidakMinatKdUser(tgl1 string, tgl2 string) []map[string]interface{}
	RekapByAlasanTdkBerminat(tgl1 string, tgl2 string) []map[string]interface{}
	RekapByAlasanTdkBerminatKdUser(tgl1 string, tgl2 string) []map[string]interface{}
	DetailApprovalTransaksi(idTrx string) entity.DetailApproval
	ListApprovalTransaksi(username string, tgl1 string, tgl2 string, search string, stsPembelian int, pageParams int, limit int) []entity.ListApproval
	ListApprovalTransaksiCount(username string, tgl1 string, tgl2 string, search string, stsPembelian int) int64
}

type asuransiRepository struct {
	connG *gorm.DB
}

func NewAsuransiRepository(connG *gorm.DB) AsuransiRepository {
	return &asuransiRepository{
		connG: connG,
	}
}

func (lR *asuransiRepository) MasterDataRekapTele(tgl1 string, tgl2 string) []entity.MasterRekapTele {
	datas := []entity.MasterRekapTele{}
	lR.connG.Raw("select a.name nama, b.* from users a  inner join (select kd_user, count(*) as total,  count(case when sts_asuransi = 'P' then 1 end) as pending,  count(case when sts_asuransi = 'T' then 1 end) as tidak_berminat,  count(case when sts_asuransi = 'O' then 1 end) as berminat  from asuransi where tgl_verifikasi >= ? and tgl_verifikasi <= ? group by kd_user) b on a.username = b.kd_user", tgl1, tgl2).Scan(&datas)
	return datas
}

func (lR *asuransiRepository) DetailApprovalTransaksi(idTrx string) entity.DetailApproval {
	detail := entity.DetailApproval{}
	lR.connG.Raw("select ko.province province_name, ko.province_code province, ko.city city_name, ko.city_code city, ko.subdistrict subdistrict_name, ko.subdistrict_code subdistrict,  t.id_transaksi, m.nm_mtr,  t.sts_pembelian, t.app_trans_id, t.id_produk, p.nm_produk, p.rate, p.admin, t.otr, (t.otr * (p.rate / 100) + admin) premi, t.thn_mtr, t.warna, t.no_msn, t.no_rgk, t.no_plat, t.nik, k.nm_konsumen, k.no_hp, k.alamat from transaksi t inner join produk p on t.id_produk = p.id_produk inner join konsumen k on k.nik = t.nik left join asuransi a on a.no_msn = t.no_msn left join mst_mtr m on m.kd_mdl = t.motorprice_kode left join kota ko on k.kec = ko.subdistrict_code  where t.id_transaksi = ?", idTrx).Find(&detail)

	return detail
}

func (lR *asuransiRepository) ListApprovalTransaksi(username string, tgl1 string, tgl2 string, search string, stsPembelian int, pageParams int, limit int) []entity.ListApproval {
	datas := []entity.ListApproval{}
	query := lR.connG.Table("transaksi t").Joins("inner join produk p on t.id_produk = p.id_produk").Joins("inner join konsumen k on k.nik = t.nik").Joins("left join asuransi a on a.no_msn = t.no_msn").Select("t.id_transaksi, k.nm_konsumen, k.no_hp, t.tgl_beli, t.sts_pembelian").Where("t.tgl_beli >= ? and t.tgl_beli <= ? ", tgl1, tgl2)
	if username != "" {
		query.Where("a.kd_user = ?", username)
	}
	if stsPembelian != 0 {
		query.Where("t.sts_pembelian = ?", stsPembelian)
	}
	query.Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Order("t.tgl_beli desc").Find(&datas)
	return datas
}

func (lR *asuransiRepository) ListApprovalTransaksiCount(username string, tgl1 string, tgl2 string, search string, stsPembelian int) int64 {
	datas := []entity.ListApproval{}
	query := lR.connG.Table("transaksi t").Joins("inner join produk p on t.id_produk = p.id_produk").Joins("inner join konsumen k on k.nik = t.nik").Joins("left join asuransi a on a.no_msn = t.no_msn").Select("t.id_transaksi").Where("t.tgl_beli >= ? and t.tgl_beli <= ?", tgl1, tgl2)
	if username != "" {
		query.Where("a.kd_user = ?", username)
	}
	if stsPembelian != 0 {
		query.Where("t.sts_pembelian = ?", stsPembelian)
	}
	query.Find(&datas)
	return int64(len(datas))
}

func (lR *asuransiRepository) MasterData(search string, dataSource string, sts string, username string, tgl1 string, tgl2 string, ap string, limit int, pageParams int) []entity.MasterAsuransi {
	if search == "undefined" {
		search = ""
	}
	if tgl1 == "undefined" {
		tgl1 = ""
	}
	if tgl2 == "undefined" {
		tgl2 = ""
	}
	datas := []entity.MasterAsuransi{}
	filter := entity.MasterAsuransi{JnsSource: dataSource}
	query := lR.connG.Where("no_msn like ? or nm_customer11 like ? or nm_dlr like ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")

	if ap != "undefined" && ap != "" {
		filter.AlasanPending = &ap
	}

	if sts != "all" && sts != "pre" {
		filter.Status = strings.ToUpper(sts)
	}
	if tgl1 != "" && tgl2 != "" {
		fmt.Println("masuk sini gk ", tgl1, tgl2)
		query.Where("tgl_verifikasi >= ? and tgl_verifikasi <= ?", tgl1, tgl2)
	}
	if sts == "pre" {
		query.Where("sts_asuransi = ? or sts_asuransi is null", "")
	}
	if sts != "pre" {
		filter.KdUser = username
	}
	query.Where(&filter).Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Order("tgl_update desc").Find(&datas)
	return datas
}

func (lR *asuransiRepository) MasterDataCount(search string, dataSource string, sts string, username string, tgl1 string, tgl2 string, ap string) int64 {
	if search == "undefined" {
		search = ""
	}
	if tgl1 == "undefined" {
		tgl1 = ""
	}
	if tgl2 == "undefined" {
		tgl2 = ""
	}
	datas := []entity.MasterAsuransi{}
	filter := entity.MasterAsuransi{JnsSource: dataSource}
	query := lR.connG.Where("no_msn like ? or nm_customer11 like ? or nm_dlr like ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")

	if ap != "undefined" && ap != "" {
		filter.AlasanPending = &ap
	}

	if sts != "all" && sts != "pre" {
		filter.Status = strings.ToUpper(sts)
	}
	if tgl1 != "" && tgl2 != "" {
		query.Where("tgl_verifikasi > ? and tgl_verifikasi < ?", tgl1, tgl2)
	}
	if sts == "pre" {
		query.Where("sts_asuransi = ? or sts_asuransi is null", "")
	}
	if sts != "pre" {
		filter.KdUser = username
	}
	query.Where(&filter).Select("no_msn").Find(&datas)
	return int64(len(datas))
}

func (lR *asuransiRepository) FindAsuransiByNoMsn(no_msn string) entity.MasterAsuransi {
	data := entity.MasterAsuransi{NoMsn: no_msn}
	transaksi := entity.Transaksi{}
	lR.connG.Joins("left join kota k on asuransi.subdistrict = k.subdistrict_code and asuransi.city = k.city_code and asuransi.province = k.province_code").Select("asuransi.*, k.province province_name, k.city city_name, k.subdistrict subdistrict_name").Find(&data)
	if data.AppTransId != "" {
		lR.connG.Where("app_trans_id = ?", data.AppTransId).First(&transaksi)
		if transaksi.ID != "" {
			data.IdTransaksi = transaksi.ID
		}
	}
	return data
}

func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func (lR *asuransiRepository) UpdateAsuransiBerminat(no_msn string) {

	var konsumen entity.Konsumen
	data := entity.MasterAsuransi{
		NoMsn: no_msn,
	}
	var transaksi entity.Transaksi
	lR.connG.Find(&data)
	if data.AppTransId == "" {
		u2, err := uuid.NewV4()
		if err != nil {
			fmt.Println("ini error uuid ", err)
		}
		data.AppTransId = u2.String()
		statusBelumBayar := "B"
		if *data.StatusBayar == "C" {
			data.StatusBayar = &statusBelumBayar
		}
		result := lR.connG.Save(&data)
		if result.Error != nil {
			fmt.Println("ini error update asuransi ", result.Error)
		}
	}

	if data.Nik != "" {
		konsumen = entity.Konsumen{Nik: data.Nik}
		lR.connG.Find(&konsumen)
		if konsumen.Nama == "" {
			konsumen.Created = time.Now()
		}
		konsumen.Nama = data.NamaCustomer
		konsumen.NoHp = data.NoTelepon
		if data.Province != nil {
			konsumen.Prop = *data.Province
		}
		if data.City != nil {
			konsumen.Kota = *data.City
		}
		if data.Subdistrict != nil {
			konsumen.Kec = *data.Subdistrict
		}
		konsumen.Updated = time.Now()
		result := lR.connG.Save(&konsumen)
		if result.Error != nil {
			fmt.Println("ini error update konsumen ", result.Error)
		}
	}

	if data.AppTransId != "" {
		lR.connG.Where("app_trans_id = ?", data.AppTransId).First(&transaksi)
		if transaksi.ID == "" {
			transaksi := entity.Transaksi{
				IdProduk:     data.JnsBrg,
				Nik:          data.Nik,
				TglBeli:      time.Now().Format("2006-01-02"),
				Updated:      time.Now(),
				AppTransId:   data.AppTransId,
				NoMsn:        data.NoMsn,
				Created:      time.Now(),
				Amount:       int(data.Harga),
				StsPembelian: "1",
			}
			result := lR.connG.Create(&transaksi)
			if result.Error != nil {
				fmt.Println("ini error create transaksi", result.Error)
			}

		}

	}

}

func (lR *asuransiRepository) UpdateAsuransiBatalBayar(no_msn string) {

	data := entity.MasterAsuransi{
		NoMsn: no_msn,
	}
	var transaksi entity.Transaksi
	lR.connG.Find(&data)
	if data.AppTransId == "" {
		u2, err := uuid.NewV4()
		if err != nil {
			fmt.Println("ini error uuid ", err)
		}
		data.AppTransId = u2.String()
		result := lR.connG.Save(&data)
		if result.Error != nil {
			fmt.Println("ini error update asuransi ", result.Error)
		}
	}
	if data.AppTransId != "" {
		lR.connG.Where("app_trans_id = ?", data.AppTransId).First(&transaksi)
		if transaksi.ID != "" {
			result := lR.connG.Delete(&transaksi)
			if result.Error != nil {
				fmt.Println("ini error create transaksi", result.Error)
			}
		}
	}

	batalStatus := "C"
	data.AppTransId = ""
	data.StatusBayar = &batalStatus
	data.Status = "P"
	data.TglBayar = nil
	result := lR.connG.Save(&data)
	if result.Error != nil {
		fmt.Println("ini error pending transaksi", result.Error)
	}

}

func GenerateIdTransaksi(transaksi entity.Transaksi) string {

	i, err := strconv.Atoi(strings.Split(transaksi.ID, "TRN")[1])
	if err != nil {
		fmt.Println("ini error parse string to int ", err)
	}
	i += 1
	idProduk := ""
	if i > 99 {
		idProduk = fmt.Sprintf("%s%d", "TRN", i)
	} else if i > 9 {
		idProduk = fmt.Sprintf("%s%d", "TRN0", i)
	} else {
		idProduk = fmt.Sprintf("%s%d", "TRN00", i)
	}
	return idProduk

}

func (lR *asuransiRepository) UpdateAsuransi(dataUpdate entity.MasterAsuransi) entity.MasterAsuransi {
	if dataUpdate.TglBayar != nil {
		if *dataUpdate.TglBayar == "" {
			dataUpdate.TglBayar = nil
		}
	}
	today := time.Now().Format("2006-01-02")
	dataUpdate.TglUpdate = &today
	dataUpdate.TglVerifikasi = time.Now().Format("2006-01-02")
	result := lR.connG.Save(&dataUpdate)
	fmt.Println("ini update error ", result.Error)
	return dataUpdate
}

func (lR *asuransiRepository) UpdateAmbilAsuransi(no_msn string, kd_user string) {
	asuransi := entity.MasterAsuransi{NoMsn: no_msn}
	lR.connG.Find(&asuransi)
	asuransi.KdUser = kd_user
	asuransi.TglVerifikasi = time.Now().Format("2006-01-02")
	asuransi.Status = "P"
	alasanPending := "1"
	asuransi.AlasanPending = &alasanPending
	lR.connG.Save(&asuransi)
}

func (lR *asuransiRepository) MasterDataGorm() {
	data := entity.MasterAsuransiGorm{
		NoMsn: "KF71E1815004",
		Nik:   "ininikaku",
	}
	lR.connG.Save(&data)
}

func (lR *asuransiRepository) RekapByStatus(u string, tgl1 string, tgl2 string) entity.MasterStatusAsuransi {
	var result entity.MasterStatusAsuransi
	query := lR.connG.Select("kd_user, count(*) as total, count(case when sts_asuransi = 'P' then 1 end) as p, count(case when sts_asuransi = 'T' then 1 end) as t, count(case when sts_asuransi = 'O' then 1 end) as o").Where("tgl_verifikasi >= ? and tgl_verifikasi <= ?", tgl1, tgl2)
	if u != "" {
		query.Where("kd_user = ?", u)
	}
	query.Table("asuransi").Group("kd_user").Find(&result)
	return result
}

func (lR *asuransiRepository) RekapByStatusAll(u string, tgl1 string, tgl2 string) entity.MasterStatusAsuransi {
	var result entity.MasterStatusAsuransi
	query := lR.connG.Select("count(case when sts_asuransi = 'P' then 1 end) as p, count(case when sts_asuransi = 'T' then 1 end) as t, count(case when sts_asuransi = 'O' then 1 end) as o").Where("tgl_verifikasi >= ? and tgl_verifikasi <= ?", tgl1, tgl2)
	query.Table("asuransi").Find(&result)
	return result
}

func (lR *asuransiRepository) RekapByStatusJenisSource(tglStart string, tglEnd string) []map[string]interface{} {
	var result []map[string]interface{}
	lR.connG.Select("jenis_source, count(*) as total, count(case when sts_asuransi = 'P' then 1 end) as p, count(case when sts_asuransi = 'T' then 1 end) as t, count(case when sts_asuransi = 'O' then 1 end) as o").Where("tgl_verifikasi >= ?", tglStart).Where("tgl_verifikasi <= ?", tglEnd).Table("asuransi").Group("jenis_source").Find(&result)
	return result
}
func (lR *asuransiRepository) RekapByStatusKdUser(tglStart string, tglEnd string) []map[string]interface{} {
	result := []map[string]interface{}{}
	lR.connG.Select("kd_user, count(*) as total, count(case when sts_asuransi = 'P' then 1 end) as p, count(case when sts_asuransi = 'T' then 1 end) as t, count(case when sts_asuransi = 'O' then 1 end) as o").Where("tgl_verifikasi >= ?", tglStart).Where("tgl_verifikasi <= ?", tglEnd).Where("kd_user is not null and kd_user != ''").Table("asuransi").Group("kd_user").Find(&result)
	return result
}

func (lR *asuransiRepository) MasterAlasanPending() []entity.MasterAlasanPending {
	result := []entity.MasterAlasanPending{}
	lR.connG.Order("id asc").Find(&result)
	return result
}

func (lR *asuransiRepository) MasterAlasanTdkBerminat() []entity.MasterAlasanTdkBerminat {
	result := []entity.MasterAlasanTdkBerminat{}
	lR.connG.Find(&result)
	return result
}

func (lR *asuransiRepository) RekapByAlasanPending(tgl1 string, tgl2 string) []map[string]interface{} {
	result := []map[string]interface{}{}
	fmt.Println("ini tgl ", tgl1, tgl2)
	lR.connG.Raw("select case when ap.name is null then 'Tidak Ada Alasan' else ap.name end as alasan, a.total from (select alasan_pending, count(*) total from asuransi where sts_asuransi = 'P' and jenis_source = 'W' and tgl_verifikasi >= ? and tgl_verifikasi <= ? group by alasan_pending) a left join mst_alasan_pending ap on ap.id = a.alasan_pending", tgl1, tgl2).Find(&result)
	return result
}

func (lR *asuransiRepository) RekapByAlasanPendingKdUser(tgl1 string, tgl2 string) []map[string]interface{} {
	result := []map[string]interface{}{}
	lR.connG.Raw("select a.kd_user, case when ap.name is null then 'Tidak Ada Alasan' else ap.name end as alasan, a.total from (select kd_user, alasan_pending, count(*) total from asuransi where sts_asuransi = 'P' and jenis_source = 'W' and tgl_verifikasi >= ? and tgl_verifikasi <= ? group by alasan_pending, kd_user) a left join mst_alasan_pending ap on ap.id = a.alasan_pending", tgl1, tgl2).Find(&result)
	return result
}

func (lR *asuransiRepository) RincianByAlasanPendingKdUser(tgl1 string, tgl2 string) []map[string]interface{} {
	result := []map[string]interface{}{}
	a := lR.MasterAlasanPending()
	queryKoloms := ", count(case when alasan_pending = '' then 1 end) as kosong"
	for _, v := range a {
		queryKoloms += fmt.Sprintf(", count(case when alasan_pending = %d then 1 end) as '%d' ", v.Id, v.Id)
	}
	query := "select kd_user, count(*) as total" + queryKoloms + "from asuransi where kd_user != '' and kd_user is not null and sts_asuransi = 'P' and tgl_verifikasi >= ? and tgl_verifikasi <=? group by kd_user "
	fmt.Println("ini query yaa ", query)
	lR.connG.Raw(query, tgl1, tgl2).Find(&result)
	return result
}

func (lR *asuransiRepository) RincianByAlasanTidakMinatKdUser(tgl1 string, tgl2 string) []map[string]interface{} {
	result := []map[string]interface{}{}
	a := lR.MasterAlasanTdkBerminat()
	queryKoloms := " "
	for _, v := range a {
		queryKoloms += fmt.Sprintf(", count(case when alasan_tdk_berminat = %d then 1 end) as '%d' ", v.Id, v.Id)
	}
	query := "select kd_user, count(*) as total" + queryKoloms + "from asuransi where kd_user != '' and kd_user is not null and sts_asuransi = 'T' and tgl_verifikasi >= ? and tgl_verifikasi <=? group by kd_user "
	fmt.Println("ini query yaa ", query)
	lR.connG.Raw(query, tgl1, tgl2).Find(&result)
	return result
}

func (lR *asuransiRepository) RekapByAlasanTdkBerminat(tgl1 string, tgl2 string) []map[string]interface{} {
	result := []map[string]interface{}{}
	lR.connG.Raw("select case when ap.name is null then 'Tidak Ada Alasan' else ap.name end as alasan, a.total from (select alasan_tdk_berminat, count(*) total from asuransi where sts_asuransi = 'T' and jenis_source = 'W' and tgl_verifikasi >= ? and tgl_verifikasi <= ? group by alasan_tdk_berminat) a left join mst_alasan_tdk_berminat ap on ap.id = a.alasan_tdk_berminat", tgl1, tgl2).Find(&result)
	return result
}

func (lR *asuransiRepository) RekapByAlasanTdkBerminatKdUser(tgl1 string, tgl2 string) []map[string]interface{} {
	result := []map[string]interface{}{}
	lR.connG.Raw("select a.kd_user, case when ap.name is null then 'Tidak Ada Alasan' else ap.name end as alasan, a.total from (select kd_user, alasan_tdk_berminat, count(*) total from asuransi where sts_asuransi = 'T' and jenis_source = 'W' and tgl_verifikasi >= ? and tgl_verifikasi <= ? group by alasan_tdk_berminat, kd_user) a left join mst_alasan_tdk_berminat ap on ap.id = a.alasan_tdk_berminat", tgl1, tgl2).Find(&result)
	return result
}
