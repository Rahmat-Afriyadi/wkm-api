package response

type RekapLeaderTs struct {
	JumlahDataMembership        int `json:"jumlah_data_membership"`
	DataBerminatMembership      int `json:"data_berminat_membership"`
	DataBerminatMembershipCash      int `json:"data_berminat_membership_cash"`
	DataBerminatMembershipTransfer      int `json:"data_berminat_membership_transfer"`
	DataSuksesMembership      int `json:"data_sukses_membership"`
	DataProspectMembership      int `json:"data_prospect_membership"`
	DataTidakBerminatMembership int `json:"data_tidak_berminat_membership"`
	DataPendingMembership       int `json:"data_pending_membership"`
	JumlahDataMembershipBefore         int `json:"jumlah_data_membership_before"`
	DataBerminatMembershipBefore      int `json:"data_berminat_membership_before"`
	DataSuksesMembershipBefore       int `json:"data_sukses_membership_before"`
	DataProspectMembershipBefore      int `json:"data_prospect_membership_before"`
	DataTidakBerminatMembershipBefore  int `json:"data_tidak_berminat_membership_before"`
	DataPendingMembershipBefore        int `json:"data_pending_membership_before"`
	DataMemberBasicPerBulan      map[int]int    `json:"data_member_basic_per_bulan"`
	DataMemberGoldPerBulan      map[int]int    `json:"data_member_gold_per_bulan"`
	DataMemberPlatPerBulan      map[int]int    `json:"data_member_plat_per_bulan"`
	DataMemberPlatPlusPerBulan      map[int]int    `json:"data_member_plat_plus_per_bulan"`
	DataBerminatPA      int `json:"data_berminat_pa"`
	DataSuksesPA      int `json:"data_sukses_pa"`
	DataProspectPA      int `json:"data_prospect_pa"`
	DataTidakBerminatPA int `json:"data_tidak_berminat_pa"`
	DataPendingPA       int `json:"data_pending_pa"`
	DataBerminatPABefore      int `json:"data_berminat_pa_before"`
	DataSuksesPABefore      int `json:"data_sukses_pa_before"`
	DataProspectPABefore      int `json:"data_prospect_pa_before"`
	DataTidakBerminatPABefore int `json:"data_tidak_berminat_pa_before"`
	DataPendingPABefore       int `json:"data_pending_pa_before"`
	DataPanda1PerBulan      map[int]int    `json:"data_panda1_per_bulan"`
	DataPanda2PerBulan      map[int]int    `json:"data_panda2_per_bulan"`
	DataPanda3PerBulan      map[int]int    `json:"data_panda3_per_bulan"`
	DataBerminatMtr      int `json:"data_berminat_mtr"`
	DataSuksesMtr     int `json:"data_sukses_mtr"`
	DataProspectMtr      int `json:"data_prospect_mtr"`
	DataTidakBerminatMtr int `json:"data_tidak_berminat_mtr"`
	DataPendingMtr       int `json:"data_pending_mtr"`
	DataBerminatMtrBefore      int `json:"data_berminat_mtr_before"`
	DataSuksesMtrBefore     int `json:"data_sukses_mtr_before"`
	DataProspectMtrBefore      int `json:"data_prospect_mtr_before"`
	DataTidakBerminatMtrBefore int `json:"data_tidak_berminat_mtr_before"`
	DataPendingMtrBefore       int `json:"data_pending_mtr_before"`
	DataTloPerBulan      map[int]int    `json:"data_tlo_per_bulan"`
	DataTloPlusPerBulan      map[int]int    `json:"data_tlo_plus_per_bulan"`
	DataKomersialPerBulan      map[int]int    `json:"data_komersial_per_bulan"`
}