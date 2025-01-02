package request

import "time"

type InputBayarRequest struct {
	TglBayar time.Time `json:"tgl_bayar"`
	NoMsn    string    `json:"no_msn"`
	KdUserFa string    `json:"kd_user_fa"`
}
