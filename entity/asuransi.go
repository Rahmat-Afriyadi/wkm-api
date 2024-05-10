package entity

type MasterAsuransi struct {
	NoMsn             string  `form:"no_msn" json:"no_msn"`
	NamaCustomer      string  `form:"nama_customer" json:"nama_customer"`
	NamaMotor         string  `form:"nama_motor" json:"nama_motor"`
	TglFaktur         string  `form:"tgl_faktur" json:"tgl_faktur"`
	NoTelepon         string  `form:"no_telepon" json:"no_telepon"`
	Status            string  `form:"status" json:"status"`
	AlasanPending     *string `form:"alasan_pending" json:"alasan_pending"`
	KdUser            string
	AlasanTdkBerminat *string `form:"alasan_tdk_berminat" json:"alasan_tdk_berminat"`
	KdDlr             *string `form:"kd_dlr" json:"kd_dlr"`
	NmDlr             *string `form:"nm_dlr" json:"nm_dlr"`
	Kelurahan         *string `form:"kelurahan" json:"kelurahan"`
	Kecamatan         *string `form:"kecamatan" json:"kecamatan"`
	Kodepos           *string `form:"kodepos" json:"kodepos"`
	JnsBrg            *string `form:"jns_brg" json:"jns_brg"`
	Harga             int64   `form:"harga" json:"harga"`
}
