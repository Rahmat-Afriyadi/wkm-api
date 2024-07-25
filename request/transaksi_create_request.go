package request

type TransaksiCreateRequest struct {
	Nik        string  `json:"nik"`
	NmKonsumen string  `json:"nm_konsumen"`
	Email      string  `json:"email"`
	NoHp       string  `json:"no_hp"`
	Alamat     string  `json:"alamat"`
	TglLahir   string  `json:"tgl_lahir"`
	Admin      uint64  `json:"admin"`
	Amount     float64 `json:"amount"`
	IdProduk   string  `json:"id_produk"`
	KdMdl      string  `json:"kd_mdl"`
	NmMtr      string  `json:"nm_mtr"`
	Warna      string  `json:"warna"`
	Tahun      string  `json:"tahun"`
	Otr        int     `json:"otr"`
	Kodepos    string  `json:"kodepos"`
	Kelurahan  string  `json:"kelurahan"`
	Kecamatan  string  `json:"kecamatan"`
	Kota       string  `json:"kota"`
	NoMsn      string  `json:"no_msn"`
	NoRgk      string  `json:"no_rgk"`
	NoPlat     string  `json:"no_plat"`
}
