package request

import "time"

type TglMerahRequest struct {
	KdUser    string    `json:"kd_user"`
	TglAwal   time.Time `json:"tgl_awal"`
	TglAkhir  time.Time `json:"tgl_akhir"`
	Deskripsi string    `json:"deskripsi"`
}
