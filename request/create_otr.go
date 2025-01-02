package request

type CreateOtr struct {
	CreateFrom     string `json:"create_from"`
	ProductNama    string `json:"product_nama"`
	MotorpriceKode string `json:"motorprice_kode"`
	Otr            uint64 `json:"otr"`
	Tahun          uint16 `json:"tahun"`
	ProductKode    string `json:"product_kode"`
	WrnKode        string `json:"wrn_kode"`
}
