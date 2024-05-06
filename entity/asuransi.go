package entity

type MasterAsuransi struct {
	NoMsn        string `form:"no_msn" json:"no_msn"`
	NamaCustomer string `form:"nama_customer" json:"nama_customer"`
	NoTelepon    string `form:"no_telepon" json:"no_telepon"`
	Status       string `form:"status" json:"status"`
	Alasan       string `form:"alasan" json:"alasan"`
}
