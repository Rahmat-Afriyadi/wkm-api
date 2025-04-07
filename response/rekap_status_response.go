package response

type RekapStatus struct {
	KdUser                string `json:"kd_user"`
	JmlData               int    `json:"jml_data"`
	SudahTerima           int    `json:"sudah_terima"`
	BelumTerima           int    `json:"belum_terima"`
	RenewalOkCashUpdate   int    `json:"renewal_ok_cash_update"`
	RenewalOkCash         int    `json:"renewal_ok_cash"`
	RenewalOkTransfer     int    `json:"renewal_ok_transfer"`
	PikirRagu             int    `json:"pikir2_ragu2"`
	TelpKembali           int    `json:"telp_kembali"`
	TidakDiangkat         int    `json:"tdk_diangkat"`
	BelumRegist           int    `json:"blm_regist"`
	Prospek               int    `json:"prospek"`
	Basic                 int    `json:"basic"`
	Gold                  int    `json:"gold"`
	Platinum              int    `json:"platinum"`
	AlasanTidakRenewal    map[string]int `json:"alasan_tidak_renewal"`
}
