package request

type TransaksiCreateRequest struct {
	Admin     uint64  `json:"admin"`
	Alamat    string  `json:"alamat"`
	Amount    float64 `json:"amount"`
	IdProduk  string  `json:"id_produk"`
	KdMdl     string  `json:"kd_mdl"`
	NmMtr     string  `json:"nm_mtr"`
	Warna     string  `json:"warna"`
	Tahun     string  `json:"tahun"`
	Otr       string  `json:"otr"`
	Kodepos   string  `json:"kodepos"`
	Kelurahan string  `json:"kelurahan"`
	Kecamatan string  `json:"kecamatan"`
	Kota      string  `json:"kota"`
	NoMsn     string  `json:"no_msn"`
	NoRgk     string  `json:"no_rgk"`
	NoPlat    string  `json:"no_plat"`
}
