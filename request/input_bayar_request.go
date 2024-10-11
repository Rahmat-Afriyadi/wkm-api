package request

type InputBayarRequest struct {
	TglBayar string `json:"tgl_bayar"`
	NoMsn    string `json:"no_msn"`
}
