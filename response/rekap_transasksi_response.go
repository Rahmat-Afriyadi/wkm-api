package response

type RekapTransaksi struct {
	NamaUser          string
	JmlData           int
	RenewalOkCash     int
	RenewalOkTransfer int
	RenewalOkDigital  int
	Basic             int
	Gold              int
	Platinum          int
	PlatinumP         int
}