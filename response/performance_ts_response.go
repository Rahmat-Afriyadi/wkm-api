package response

type ListPerformanceTs struct {
	NamaUser     string  `json:"nama_user"`
	JumlahSukses int     `json:"jumlah_sukses"`
	Contribution float64 `json:"contribution"`
}

// Struktur utama untuk response API
type PerformanceTs struct {
	TopUsers       []ListPerformanceTs `json:"top_users"`
	LowPerformance []ListPerformanceTs `json:"low_performance"`
}