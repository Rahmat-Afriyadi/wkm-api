package response

type DataRenewalResponse struct {
	// Year  int `json:"year"`
    // Month int `json:"month"`
	JnsCard string `json:"jns_card"`
	TotalJumlahData int `json:"total_jumlah_data"`
}