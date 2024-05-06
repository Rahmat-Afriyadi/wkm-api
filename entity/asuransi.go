package entity

type MasterAsuransi struct {
	NoMsn        string `json:"no_msn"`
	NamaCustomer string `json:"nama_customer"`
	NoTelepon    string `json:"no_telepon"`
	Status       string `json:"status"`
	Alasan       string `json:"alasan"`
}
