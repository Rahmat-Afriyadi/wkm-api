package response

type RekapTele struct {
	JumlahDataMembership        int `json:"jumlah_data_membership"`
	DataBerminatMembership      int `json:"data_berminat_membership"`
	DataSuksesMembership      int `json:"data_sukses_membership"`
	DataProspectMembership      int `json:"data_prospect_membership"`
	DataTidakBerminatMembership int `json:"data_tidak_berminat_membership"`
	DataPendingMembership       int `json:"data_pending_membership"`
	DataBerminatMembershipPerBulan      map[int]int    `json:"data_berminat_membership_per_bulan"`
	DataBerminatPA      int `json:"data_berminat_pa"`
	DataSuksesPA      int `json:"data_sukses_pa"`
	DataProspectPA      int `json:"data_prospect_pa"`
	DataTidakBerminatPA int `json:"data_tidak_berminat_pa"`
	DataPendingPA       int `json:"data_pending_pa"`
	DataBerminatPAPerBulan      map[int]int    `json:"data_berminat_pa_per_bulan"`
	DataBerminatMtr      int `json:"data_berminat_mtr"`
	DataSuksesMtr     int `json:"data_sukses_mtr"`
	DataProspectMtr      int `json:"data_prospect_mtr"`
	DataTidakBerminatMtr int `json:"data_tidak_berminat_mtr"`
	DataPendingMtr       int `json:"data_pending_mtr"`
	DataBerminatMtrPerBulan      map[int]int    `json:"data_berminat_mtr_per_bulan"`
}