package request

import "time"

type ExtendBayarRequest struct {
	Id             string    `json:"id"`
	NoMsn          string    `json:"no_msn"`
	KdUserFa       string    `json:"kd_user"`
	KdUserLf       string    `json:"kd_user_lf"`
	TglActualBayar time.Time `json:"tgl_actual_bayar"`
	StsApproval    string    `json:"sts_approval"`
	Deskripsi      string    `json:"deskripsi"`
	RenewalKe      string    `json:"sts_cetak3"`
}
