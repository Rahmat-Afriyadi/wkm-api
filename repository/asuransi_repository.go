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
	MasterData(search string, dataSource string, sts string, usename string, limit int, pageParams int) []entity.MasterAsuransi
	MasterDataCount(search string, dataSource string, sts string, username string) int64
	FindAsuransiByNoMsn(no_msn string) entity.MasterAsuransi
	UpdateAmbilAsuransi(no_msn string, kd_user string)
	UpdateAsuransi(data entity.MasterAsuransi) entity.MasterAsuransi
	UpdateAsuransiBerminat(no_msn string)
	UpdateAsuransiBatalBayar(no_msn string)
	MasterDataGorm()
	MasterAlasanPending() []entity.MasterAlasanPending
	MasterAlasanTdkBerminat() []entity.MasterAlasanTdkBerminat
	RekapByStatus(u string, tgl string) entity.MasterStatusAsuransi
	MasterDataRekapTele() []entity.MasterRekapTele
	RekapByStatusJenisSource(tglStart string, tglEnd string) []map[string]interface{}
	RekapByStatusKdUser(tglStart string, tglEnd string) []map[string]interface{}
}

type asuransiRepository struct {
	conn  *sql.DB
	connG *gorm.DB
}

func NewAsuransiRepository(connG *gorm.DB) AsuransiRepository {
	return &asuransiRepository{
		connG: connG,
	}
}

func (lR *asuransiRepository) MasterDataRekapTele() []entity.MasterRekapTele {
	datas := []entity.MasterRekapTele{}
	lR.connG.Raw("select a.name nama, b.* from users a  inner join (select kd_user, count(*) as total,  count(case when sts_asuransi = 'P' then 1 end) as pending,  count(case when sts_asuransi = 'T' then 1 end) as tidak_berminat,  count(case when sts_asuransi = 'O' then 1 end) as berminat  from asuransi where tgl_update = ? group by kd_user) b on a.username = b.kd_user", "2024-05-21").Scan(&datas)
	return datas
}

func (lR *asuransiRepository) MasterData(search string, dataSource string, sts string, username string, limit int, pageParams int) []entity.MasterAsuransi {
	if search == "undefined" {
		search = ""
	}
	datas := []entity.MasterAsuransi{}
	filter := entity.MasterAsuransi{JnsSource: dataSource}
	query := lR.connG.Where("no_msn like ? or nm_customer11 like ? or nm_dlr like ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")

	if sts != "all" && sts != "pre" {
		filter.Status = strings.ToUpper(sts)
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

func (lR *asuransiRepository) MasterDataCount(search string, dataSource string, sts string, username string) int64 {
	if search == "undefined" {
		search = ""
	}
	datas := []entity.MasterAsuransi{}
	filter := entity.MasterAsuransi{JnsSource: dataSource}
	query := lR.connG.Where("no_msn like ? or nm_customer11 like ? or nm_dlr like ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")

	if sts != "all" && sts != "pre" {
		filter.Status = strings.ToUpper(sts)
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
	lR.connG.Find(&data)
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
		if data.TglLahir != nil {
			konsumen.TglLahir = data.TglLahir
		}
		if data.Kecamatan != nil {
			konsumen.Kec = *data.Kecamatan
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
	lR.connG.Save(&asuransi)
}

func (lR *asuransiRepository) MasterDataGorm() {
	data := entity.MasterAsuransiGorm{
		NoMsn: "KF71E1815004",
		Nik:   "ininikaku",
	}
	lR.connG.Save(&data)
}

func (lR *asuransiRepository) RekapByStatus(u string, tgl string) entity.MasterStatusAsuransi {
	var result entity.MasterStatusAsuransi
	lR.connG.Select("kd_user, count(*) as total, count(case when sts_asuransi = 'P' then 1 end) as p, count(case when sts_asuransi = 'T' then 1 end) as t, count(case when sts_asuransi = 'O' then 1 end) as o").Where("kd_user = ?", u).Where("tgl_verifikasi = ?", tgl).Table("asuransi").Group("kd_user").Find(&result)
	return result
}

func (lR *asuransiRepository) RekapByStatusJenisSource(tglStart string, tglEnd string) []map[string]interface{} {
	var result []map[string]interface{}
	lR.connG.Select("jenis_source, count(*) as total, count(case when sts_asuransi = 'P' then 1 end) as p, count(case when sts_asuransi = 'T' then 1 end) as t, count(case when sts_asuransi = 'O' then 1 end) as o").Where("tgl_verifikasi >= ?", tglStart).Where("tgl_verifikasi <= ?", tglEnd).Table("asuransi").Group("jenis_source").Find(&result)
	return result
}
func (lR *asuransiRepository) RekapByStatusKdUser(tglStart string, tglEnd string) []map[string]interface{} {
	var result []map[string]interface{}
	lR.connG.Select("kd_user, count(*) as total, count(case when sts_asuransi = 'P' then 1 end) as p, count(case when sts_asuransi = 'T' then 1 end) as t, count(case when sts_asuransi = 'O' then 1 end) as o").Where("tgl_verifikasi >= ?", tglStart).Where("tgl_verifikasi <= ?", tglEnd).Table("asuransi").Group("kd_user").Find(&result)
	return result
}

func (lR *asuransiRepository) MasterAlasanPending() []entity.MasterAlasanPending {
	result := []entity.MasterAlasanPending{}
	lR.connG.Find(&result)
	return result
}

func (lR *asuransiRepository) MasterAlasanTdkBerminat() []entity.MasterAlasanTdkBerminat {
	result := []entity.MasterAlasanTdkBerminat{}
	lR.connG.Find(&result)
	return result
}
