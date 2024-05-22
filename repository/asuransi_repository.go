package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"
	"wkm/entity"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type AsuransiRepository interface {
	MasterData(dataSource string) []entity.MasterAsuransi
	MasterDataPending(search string, dataSource string) []entity.MasterAsuransi
	MasterDataOke(dataSource string) []entity.MasterAsuransi
	FindAsuransiByNoMsn(no_msn string) entity.MasterAsuransiReal
	UpdateAmbilAsuransi(no_msn string, kd_user string)
	UpdateAsuransi(data entity.MasterAsuransiReal) entity.MasterAsuransiReal
	UpdateAsuransiBerminat(no_msn string)
	UpdateAsuransiBatalBayar(no_msn string)
	MasterDataGorm()
}

type asuransiRepository struct {
	conn  *sql.DB
	connG *gorm.DB
}

func NewAsuransiRepository(conn *sql.DB, connG *gorm.DB) AsuransiRepository {
	return &asuransiRepository{
		conn:  conn,
		connG: connG,
	}
}

func (lR *asuransiRepository) MasterData(dataSource string) []entity.MasterAsuransi {
	datas := []entity.MasterAsuransi{}
	ctx := context.Background()
	query := "select no_msn NoMsn, nm_customer11 NamaCustomer, jenis_asuransi from asuransi where (sts_asuransi = '' or sts_asuransi is null) and jenis_source = ?"
	statement, err := lR.conn.PrepareContext(ctx, query)
	if err != nil {
		fmt.Println(err)
	}
	rows, err := statement.QueryContext(ctx, dataSource)
	if err != nil {
		fmt.Println("errornya di rows ", err)
		fmt.Println(err)
	}

	for rows.Next() {
		var data entity.MasterAsuransi
		if err := rows.Scan(&data.NoMsn, &data.NamaCustomer, &data.JnsAsuransi); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		datas = append(datas, data)
	}
	return datas
}

func (lR *asuransiRepository) MasterDataPending(search string, dataSource string) []entity.MasterAsuransi {
	fmt.Println("ini query ", search)
	datas := []entity.MasterAsuransi{}
	ctx := context.Background()
	query := "select no_msn NoMsn, nm_customer11 NamaCustomer, nm_dlr from asuransi where sts_asuransi = 'P' and (no_msn like ? or nm_customer11 like ? or nm_dlr like ?) and jenis_source = ?"
	statement, err := lR.conn.PrepareContext(ctx, query)
	if err != nil {
		fmt.Println(err)
	}
	rows, err := statement.QueryContext(ctx, "%"+search+"%", "%"+search+"%", "%"+search+"%", dataSource)
	if err != nil {
		fmt.Println("errornya di rows ", err)
		fmt.Println(err)
	}

	for rows.Next() {
		var data entity.MasterAsuransi
		if err := rows.Scan(&data.NoMsn, &data.NamaCustomer, &data.NmDlr); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		datas = append(datas, data)
	}
	return datas
}

func (lR *asuransiRepository) MasterDataOke(dataSource string) []entity.MasterAsuransi {
	datas := []entity.MasterAsuransi{}
	ctx := context.Background()
	query := "select no_msn NoMsn, nm_customer11 NamaCustomer, nm_dlr from asuransi where sts_asuransi = 'O' and jenis_source=?"
	statement, err := lR.conn.PrepareContext(ctx, query)
	if err != nil {
		fmt.Println(err)
	}
	rows, err := statement.QueryContext(ctx, dataSource)
	if err != nil {
		fmt.Println("errornya di rows ", err)
		fmt.Println(err)
	}

	for rows.Next() {
		var data entity.MasterAsuransi
		if err := rows.Scan(&data.NoMsn, &data.NamaCustomer, &data.NmDlr); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		datas = append(datas, data)
	}
	return datas
}

func (lR *asuransiRepository) FindAsuransiByNoMsn(no_msn string) entity.MasterAsuransiReal {
	data := entity.MasterAsuransiReal{NoMsn: no_msn}
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
	data := entity.MasterAsuransiReal{
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
		// dateString := time.Now().Format("2006-01-02")
		// konsumen.TglLahir = &dateString
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

	data := entity.MasterAsuransiReal{
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
	fmt.Println("ini asuransi ", data)
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

func (lR *asuransiRepository) UpdateAsuransi(dataUpdate entity.MasterAsuransiReal) entity.MasterAsuransiReal {
	fmt.Println("ini data update yaa ", dataUpdate)
	lR.connG.Save(&dataUpdate)
	return dataUpdate
}

// func (lR *asuransiRepository) UpdateAsuransi(dataUpdate entity.MasterAsuransi) entity.MasterAsuransi {
// 	KdDlr := ""
// 	NmDlr := ""
// 	Kelurahan := ""
// 	Kecamatan := ""
// 	Kodepos := ""
// 	JnsBrg := ""
// 	stsBayar := ""
// 	var tglBayar string
// 	if dataUpdate.KdDlr != nil {
// 		KdDlr = *dataUpdate.KdDlr
// 	}
// 	if dataUpdate.NmDlr != nil {
// 		NmDlr = *dataUpdate.NmDlr
// 	}
// 	if dataUpdate.Kelurahan != nil {
// 		Kelurahan = *dataUpdate.Kelurahan
// 	}
// 	if dataUpdate.Kecamatan != nil {
// 		Kecamatan = *dataUpdate.Kecamatan
// 	}
// 	if dataUpdate.Kodepos != nil {
// 		Kodepos = *dataUpdate.Kodepos
// 	}
// 	if dataUpdate.JnsBrg != nil {
// 		JnsBrg = *dataUpdate.JnsBrg
// 	}
// 	if dataUpdate.StatusBayar != nil {
// 		stsBayar = *dataUpdate.StatusBayar
// 	}
// 	if dataUpdate.TglBayar != nil {
// 		tglBayar = *dataUpdate.TglBayar
// 	}
// 	ctx := context.Background()
// 	_, err := lR.conn.ExecContext(ctx, "UPDATE asuransi set tgl_bayar=?, sts_bayar=?, sts_asuransi=?, alasan_pending=?, alasan_tdk_berminat=?, kd_dlr=?, nm_dlr=?, kelurahan=?, kecamatan=?, kodepos=?, jns_brg=?, harga=?, kd_user=?, tgl_update=? where no_msn=? ", NewNullString(tglBayar), stsBayar, dataUpdate.Status, dataUpdate.AlasanPending, dataUpdate.AlasanTdkBerminat, KdDlr, NmDlr, Kelurahan, Kecamatan, Kodepos, JnsBrg, dataUpdate.Harga, dataUpdate.KdUser, time.Now().Format("2006-01-02"), dataUpdate.NoMsn)
// 	if err != nil {
// 		fmt.Println("ini error update ", err)
// 	}
// 	return dataUpdate
// }

func (lR *asuransiRepository) UpdateAmbilAsuransi(no_msn string, kd_user string) {
	ctx := context.Background()
	_, err := lR.conn.ExecContext(ctx, "UPDATE asuransi set sts_asuransi='P', tgl_update=?, kd_user=? where no_msn=?", time.Now().Format("2006-01-02"), kd_user, no_msn)
	if err != nil {
		fmt.Println("ini error update ", err)
	}
}

func (lR *asuransiRepository) MasterDataGorm() {
	data := entity.MasterAsuransiGorm{
		NoMsn: "KF71E1815004",
		Nik:   "ininikaku",
	}
	lR.connG.Save(&data)
	fmt.Println(" Ini data", data)
}
