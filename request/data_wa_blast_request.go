package request

type DataWaBlastRequest struct {
	AwalTenor           string   `form:"awal_tenor" json:"awal_tenor"`
	AkhirTenor          string   `form:"akhir_tenor" json:"akhir_tenor"`
	NoLeas              string   `form:"no_leas" json:"no_leas"`
	KodeKerjaFilterType string   `form:"kode_kerja_filter_type" json:"kode_kerja_filter_type"`
	KodeKerja           []string `form:"kode_kerja" json:"kode_kerja"`
}
