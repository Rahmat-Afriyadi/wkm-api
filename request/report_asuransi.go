package request

type ReportAsuransi struct {
	AwalTanggal  string `form:"awal_tgl" json:"awal_tgl"`
	AkhirTAnggal string `form:"akhir_tgl" json:"akhir_tgl"`
}
