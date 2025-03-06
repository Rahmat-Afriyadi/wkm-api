package entity

import "time"

type DetailAsuransiPA struct {
	IdProduk string `json:"id_produk"`
	NmVendor string `json:"nm_vendor"`
	NmProduk string `json:"nm_produk"`
	AmountPa uint64 `json:"amount_asuransi_pa"`
}

type DetailAsuransiMtr struct {
	IdProduk string `json:"id_produk"`
	NmVendor string `json:"nm_vendor"`
	NmProduk string `json:"nm_produk"`
	NmMtr string 	`json:"nm_mtr"`
	NoMtr string 	`json:"no_mtr"`
	Otr uint64 		`json:"otr"`
	AmountOtr uint64 `json:"amount_otr"`
}


type Faktur3 struct {
	NoMsn               string     `form:"no_msn" json:"no_msn" gorm:"primary_key;column:no_msn"`
	NoTandaTerima       string     `form:"no_tanda_terima" json:"no_tanda_terima" gorm:"column:no_tanda_terima"`
	TglCetakTandaTerima time.Time  `form:"tgl_cetak_tanda_terima" json:"tgl_cetak_tanda_terima" gorm:"column:tgl_cetak_tanda_terima"`
	TglVerifikasi       time.Time  `form:"tgl_verifikasi" json:"tgl_verifikasi" gorm:"column:tgl_verifikasi"`
	TglExpired          *time.Time `form:"tgl_expired" json:"tgl_expired" gorm:"column:tgl_expired"`
	NmCustomer          string     `form:"nm_customer11" json:"nm_customer11" gorm:"column:nm_customer11"`
	NamaKtp             string     `form:"nama_ktp" json:"nama_ktp" gorm:"column:nama_ktp"`
	NmMtr               string     `form:"nm_mtr" json:"nm_mtr" gorm:"column:nm_mtr"`
	Telp1               string     `form:"no_telp1" json:"no_telp1" gorm:"column:no_telp1"`
	Hp1                 string     `form:"no_hp1" json:"no_hp1" gorm:"column:no_hp1"`
	KdUser              string     `form:"kd_user" json:"kd_user" gorm:"column:kd_user"`
	User                SecUser    `json:"user_2" gorm:"->;references:Username;foreignKey:kd_user"`
	KdUser2             string     `form:"kd_user2" json:"kd_user2" gorm:"column:kd_user2"`
	KdUser10            *string    `form:"kd_user10" json:"kd_user10" gorm:"column:kd_user10"`
	User10              SecUser    `json:"user_10" gorm:"->;references:Username;foreignKey:KdUser10"`
	Print               uint8      `form:"print" json:"print" gorm:"column:print"`
	StsCetak3           string     `form:"sts_cetak3" json:"sts_cetak3" gorm:"column:sts_cetak3"`
	StsJnsBayar         string     `form:"sts_jenis_bayar" json:"sts_jenis_bayar" gorm:"column:sts_jenis_bayar"`
	StsRenewal         string     `form:"sts_renewal" json:"sts_renewal" gorm:"column:sts_renewal"`
	StsKartu            string     `form:"sts_kartu" json:"sts_kartu" gorm:"column:sts_kartu"`
	StsBawaKartu        string     `form:"sts_bawa_kartu" json:"sts_bawa_kartu" gorm:"column:sts_bawa_kartu"`
	StsBayarRenewal     string     `form:"sts_bayar_renewal" json:"sts_bayar_renewal" gorm:"column:sts_bayar_renewal"`
	ALamatBantuan       string     `form:"alamat_bantuan" json:"alamat_bantuan" gorm:"column:alamat_bantuan"`
	StsKirim            string     `form:"sts_kirim" json:"sts_kirim" gorm:"column:sts_kirim"`
	KdCard              string     `form:"kd_card" json:"kd_card" gorm:"column:kd_card"`
	MstCard             MstCard    `json:"mst_card" gorm:"->;references:KdCard;foreignKey:KdCard"`
	KdKurir             string     `form:"kode_kurir" json:"kode_kurir" gorm:"column:kode_kurir"`
	Kurir               MstKurir   `form:"kurir" json:"kurir" gorm:"->;references:KdKurir;foreignKey:KdKurir"`
	NoKartu             string     `form:"no_kartu" json:"no_kartu" gorm:"column:no_kartu"`
	Kartu               StockCard  `form:"kartu" json:"kartu" gorm:"->;references:NoKartu;foreignKey:NoKartu"`
	// Tanggal tanggal bayar
	TglBayarRenewal         *time.Time `form:"tgl_bayar_renewal" json:"tgl_bayar_renewal" gorm:"column:tgl_bayar_renewal"`
	TglBayarRenewalFin      *time.Time `form:"tgl_bayar_renewal_fin" json:"tgl_bayar_renewal_fin" gorm:"column:tgl_bayar_renewal_fin"`
	TglBayarRenewalFinKeyIn *time.Time `form:"tgl_bayar_renewal_fin_key_in" json:"tgl_bayar_renewal_fin_key_in" gorm:"column:tgl_bayar_renewal_fin_key_in"`
	// O
	StsAsuransiPa string `form:"sts_asuransi_pa" json:"sts_asuransi_pa" gorm:"column:sts_asuransi_pa"`
	// S
	StsBayarAsuransiPa string `form:"sts_bayar_asuransi_pa" json:"sts_bayar_asuransi_pa" gorm:"column:sts_bayar_asuransi_pa"`

	// alamat Rumah
	Alamat  string `form:"alamat21" json:"alamat21" gorm:"column:alamat21"`
	Kota    string `form:"kota2" json:"kota2" gorm:"column:kota2"`
	Kec     string `form:"kec2" json:"kec2" gorm:"column:kec2"`
	Kel     string `form:"kel2" json:"kel2" gorm:"column:kel2"`
	Rt      string `form:"rt2" json:"rt2" gorm:"column:rt2"`
	Rw      string `form:"rw2" json:"rw2" gorm:"column:rw2"`
	Kodepos string `form:"kodepos2" json:"kodepos2" gorm:"column:kodepos2"`

	// alamat kantor
	KerjaDi    string `form:"kerja_di" json:"kerja_di" gorm:"column:kerja_di"`
	AlamatKtr  string `form:"alamat_ktr" json:"alamat_ktr" gorm:"column:alamat_ktr"`
	RtKtr      string `form:"rt_ktr" json:"rt_ktr" gorm:"column:rt_ktr"`
	RwKtr      string `form:"rw_ktr" json:"rw_ktr" gorm:"column:rw_ktr"`
	KelKtr     string `form:"kel_ktr" json:"kel_ktr" gorm:"column:kel_ktr"`
	KecKtr     string `form:"kec_ktr" json:"kec_ktr" gorm:"column:kec_ktr"`
	KodeposKtr string `form:"kodepos_ktr" json:"kodepos_ktr" gorm:"column:kodepos_ktr"`
	KotaKtr    string `form:"kota1" json:"kota1" gorm:"column:kota1"`

	// alamat AHHASS Dealer
	NamaPt      string `form:"alamat_srt12" json:"alamat_srt12" gorm:"column:alamat_srt12"`
	AlamatAD    string `form:"alamat_srt11" json:"alamat_srt11" gorm:"column:alamat_srt11"`
	KotaSrt1    string `form:"kota_srt1" json:"kota_srt1" gorm:"column:kota_srt1"`
	KecSrt1     string `form:"kec_srt1" json:"kec_srt1" gorm:"column:kec_srt1"`
	KelSrt1     string `form:"kel_srt1" json:"kel_srt1" gorm:"column:kel_srt1"`
	KodeposSrt1 string `form:"kodepos_srt1" json:"kodepos_srt1" gorm:"column:kodepos_srt1"`


	TypeKartu string `json:"type_kartu" gorm:"-;type:json"`
	AsuransiPa DetailAsuransiPA `json:"asuransi_pa" gorm:"-;type:json"`
	AsuransiMtr DetailAsuransiMtr `json:"asuransi_mtr" gorm:"-;type:json"`
}

func (Faktur3) TableName() string {
	return "tr_wms_faktur3"
}
