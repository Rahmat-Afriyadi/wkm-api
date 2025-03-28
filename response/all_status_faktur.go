package response

type AllStatusResponse struct {
	NoMsn         string  `json:"no_msn"`
	StsMembership *string `json:"sts_membership"`
	StsBayar      *string `json:"sts_bayar"`
	NmCustomer    string  `json:"nm_customer"`
	TglVerifikasi string  `json:"tgl_verifikasi"`
	FromTable     string  `json:"from_table"`
}
