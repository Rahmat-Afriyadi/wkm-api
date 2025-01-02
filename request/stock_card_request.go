package request

import "time"

type StockCardRequest struct {
	NoKartu    string    `form:"no_kartu" json:"no_kartu" `
	TglCetak   time.Time `form:"tgl_cetak" json:"tgl_cetak" `
	TglUpdate  time.Time `form:"tgl_update" json:"tgl_update" `
	TglExpired time.Time `form:"tgl_expired" json:"tgl_expired" `
	StsKartu   string    `form:"sts_kartu" json:"sts_kartu" `
}
