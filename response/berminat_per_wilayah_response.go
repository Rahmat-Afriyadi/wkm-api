package response

type RekapBerminatPerWilayah struct {
	Kota      string `json:"kota"`
	Kecamatan string `json:"kecamatan"`
	Jumlah    int    `json:"jumlah"`
}