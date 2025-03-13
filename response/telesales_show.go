package response

import "time"

type TelesalesResponse struct {
	NoMsn                  string     `gorm:"-;type:json" json:"no_msn" `
	NmCustomerFkt          string     `gorm:"-;type:json" json:"nm_customer_fkt" `
	KdDlr                  string     `gorm:"-;type:json" json:"kd_dlr" `
	NmDlr                  string     `gorm:"-;type:json" json:"nm_dlr" `
	NoRgk                  string     `gorm:"-;type:json" json:"no_rgk" `
	TglLahirFkt            *time.Time `gorm:"-;type:json" json:"tgl_lahir_fkt" `
	JnsKlmFkt              string     `gorm:"-;type:json" json:"jns_klm_fkt" `
	NoTelpFkt              string     `gorm:"-;type:json" json:"no_telp_fkt" `
	KetNoTelpFkt           string     `gorm:"-;type:json" json:"ket_no_telp_fkt" `
	NoHpFkt                string     `gorm:"-;type:json" json:"no_hp_fkt" `
	KetNoHpFkt             string     `gorm:"-;type:json" json:"ket_no_hp_fkt" `
	AlamatFkt              string     `gorm:"-;type:json" json:"alamat_fkt" `
	KelFkt                 string     `gorm:"-;type:json" json:"kel_fkt" `
	KecFkt                 string     `gorm:"-;type:json" json:"kec_fkt" `
	KotaFkt                string     `gorm:"-;type:json" json:"kota_fkt" `
	KodeposFkt             string     `gorm:"-;type:json" json:"kodepos_fkt" `
	PicPerusahaanFkt       string     `gorm:"-;type:json" json:"pic_perusahaan_fkt" `
	NoMtr                  string     `gorm:"-;type:json" json:"no_mtr" `
	NmMtr                  string     `gorm:"-;type:json" json:"nm_mtr" `
	TglMohon               *time.Time `gorm:"-;type:json" json:"tgl_mohon" `
	TglFaktur              *time.Time `gorm:"-;type:json" json:"tgl_faktur" `
	KodeKerjaFkt           string     `gorm:"-;type:json" json:"kode_kerja_fkt" `
	KodeDidikFkt           string     `gorm:"-;type:json" json:"kode_didik_fkt" `
	KeluarBlnFkt           string     `gorm:"-;type:json" json:"keluar_bln_fkt" `
	MotorHir               string     `gorm:"-;type:json" json:"motor_hir" `
	JnsBeli                string     `gorm:"-;type:json" json:"jns_beli" `
	BdnUsaha              string     `gorm:"-;type:json" json:"bdn_usaha" `
	JnsJualFkt            string     `gorm:"-;type:json" json:"jns_jual_fkt" `
	AktifJualFkt          string     `gorm:"-;type:json;" json:"aktif_jual_fkt" `
	NoKtpnpwpFkt          string     `gorm:"-;type:json" json:"no_ktpnpwp_fkt" `
	NoNpwpFkt             string     `gorm:"-;type:json" json:"no_npwp_fkt" `
	KerjaDiFkt            string     `gorm:"-;type:json" json:"kerja_di_fkt" `
	AlamatKtrFkt          string     `gorm:"-;type:json" json:"alamat_ktr_fkt" `
	PropKtrFkt            string     `gorm:"-;type:json" json:"prop_ktr_fkt" `
	KotaKtrFkt            string     `gorm:"-;type:json" json:"kota_ktr_fkt" `
	KelKtrFkt             string     `gorm:"-;type:json" json:"kel_ktr_fkt" `
	KecKtrFkt             string     `gorm:"-;type:json" json:"kec_ktr_fkt" `
	KodeposKtrFkt         string     `gorm:"-;type:json" json:"kodepos_ktr_fkt" `
	NoPolFkt              string     `gorm:"-;type:json" json:"no_pol_fkt" `
	NoPolwKM              string     `gorm:"-;type:json" json:"no_pol_wkm" `
	HobbyFkt              string     `gorm:"-;type:json" json:"hobby_fkt" `
	StsSourceFkt          string     `gorm:"-;type:json" json:"sts_source_fkt" `
	NmSalesFkt            string     `gorm:"-;type:json" json:"nm_sales_fkt" `
	IdSalesFkt            string     `gorm:"-;type:json" json:"id_sales_fkt" `
	AgamaFkt              string     `gorm:"-;type:json" json:"agama_fkt" `
	TujuanPakaiFkt        string     `gorm:"-;type:json" json:"tujuan_pakai_fkt" `
	DpMtrFkt              int32    `gorm:"-;type:json" json:"dp_mtr_fkt" `
	CicilanMtrFkt         int32    `gorm:"-;type:json" json:"cicilan_mtr_fkt" `
	AngsuranFkt           string     `gorm:"-;type:json" json:"angsuran_fkt" `
	EmailFkt              string     `gorm:"-;type:json" json:"email_fkt" `
	NoLeasFkt             string     `gorm:"-;type:json" json:"no_leas_fkt" `
	JnsMotorFkt           string     `gorm:"-;type:json" json:"jns_motor_fkt" `
	SmDibeli              string     `gorm:"-;type:json" json:"sm_dibeli" `
	NoKk                  string     `gorm:"-;type:json" json:"no_kk" `
	TglAkhirTenor        *time.Time `gorm:"-;type:json" json:"tgl_akhir_tenor" `
	SmFacebookFkt        string     `gorm:"-;type:json" json:"sm_facebook_fkt" `
	SmInstagramFkt       string     `gorm:"-;type:json" json:"sm_instagram_fkt" `
	SmTwitterFkt         string     `gorm:"-;type:json" json:"sm_twitter_fkt" `
	SmYoutubeFkt         string     `gorm:"-;type:json" json:"sm_youtube_fkt" `
	TglLahirWkm          *time.Time `gorm:"-;type:json" json:"tgl_lahir_wkm" `
	JnsKlmWkm            string     `gorm:"-;type:json" json:"jns_klm_wkm" `
	NoTelpWkm            string     `gorm:"-;type:json" json:"no_telp_wkm" `
	KetNoTelpWkm         string     `gorm:"-;type:json" json:"ket_no_telp_wkm" `
	NoHpWkm              string     `gorm:"-;type:json" json:"no_hp_wkm" `
	NoWa                 string     `gorm:"-;type:json" json:"no_wa" `
	NoTelpKtrWkm         string     `gorm:"-;type:json" json:"no_telp_ktr_wkm" `
	EmailWkm             string     `gorm:"-;type:json" json:"email_wkm" `
	AlamatWkm            string     `gorm:"-;type:json" json:"alamat_wkm" `
	KetAlamatWkm         string     `gorm:"-;type:json" json:"ket_alamat_wkm" `
	KetNoInfo         string     `gorm:"-;type:json" json:"ket_wa_info" `
	RtWkm                string     `gorm:"-;type:json" json:"rt_wkm" `
	RwWkm                string     `gorm:"-;type:json" json:"rw_wkm" `
	KelWkm               string     `gorm:"-;type:json" json:"kel_wkm" `
	KecWkm               string     `gorm:"-;type:json" json:"kec_wkm" `
	KotaWkm              string     `gorm:"-;type:json" json:"kota_wkm" `
	KodeposWkm           string     `gorm:"-;type:json" json:"kodepos_wkm" `
	TglCallTele           *time.Time `gorm:"-;type:json" json:"tgl_call_tele" `
	JnsJualWkm            string     `gorm:"-;type:json" json:"jns_jual_wkm" `
	AktifJualWkm          string     `gorm:"-;type:json;" json:"aktif_jual_wkm" `
	KetAktifJualWkm       string     `gorm:"-;type:json" json:"ket_aktif_jual_wkm" `
	NoKtpnpwpWkm          string     `gorm:"-;type:json" json:"no_ktpnpwp_wkm" `
	KdUserTs              string     `gorm:"-;type:json" json:"kd_user_ts" `
	KerjaDiWkm            string     `gorm:"-;type:json" json:"kerja_di_wkm" `
	JabatanWkm            string     `gorm:"-;type:json" json:"jabatan_wkm" `
	AlamatKtrWkm          string     `gorm:"-;type:json" json:"alamat_ktr_wkm" `
	KotaKtrWkm            string     `gorm:"-;type:json" json:"kota_ktr_wkm" `
	RtKtrWkm              string     `gorm:"-;type:json" json:"rt_ktr_wkm" `
	RwKtrWkm              string     `gorm:"-;type:json" json:"rw_ktr_wkm" `
	KelKtrWkm             string     `gorm:"-;type:json" json:"kel_ktr_wkm" `
	KecKtrWkm             string     `gorm:"-;type:json" json:"kec_ktr_wkm" `
	KodeposKtrWkm         string     `gorm:"-;type:json" json:"kodepos_ktr_wkm" `
	NmCustomerWkm         string     `gorm:"-;type:json" json:"nm_customer_wkm" `
	KdDlrWkm              string     `gorm:"-;type:json" json:"kd_dlr_wkm" `
	StsValidWkm           string     `gorm:"-;type:json" json:"sts_valid_wkm" `
	StsSource2Wkm         string     `gorm:"-;type:jsondefault:' '" json:"sts_source2_wkm" `
	PicPerusahaanWkm      string     `gorm:"-;type:json" json:"pic_perusahaan_wkm" `
	HobbyWkmOthers        string     `gorm:"-;type:json" json:"hobby_wkm_others" `
	HobbyWkm              string     `gorm:"-;type:json" json:"hobby_wkm" `
	StsKawinWkm           string     `gorm:"-;type:json" json:"sts_kawin_wkm" `
	NmSalesWkm            string     `gorm:"-;type:json" json:"nm_sales_wkm" `
	AgamaWkm              string     `gorm:"-;type:json" json:"agama_wkm" `
	KodeDidikWkm          string     `gorm:"-;type:json" json:"kode_didik_wkm" `
	KodeKerjaWkm          string     `gorm:"-;type:json" json:"kode_kerja_wkm" `
	NmKerjaWkm            string     `gorm:"-;type:json" json:"nm_kerja_wkm" `
	TujuanPakaiWkm        string     `gorm:"-;type:json" json:"tujuan_pakai_wkm" `
	KeluarBlnWkm          string     `gorm:"-;type:json" json:"keluar_bln_wkm" `
	StsMembership         string     `gorm:"-;type:json" json:"sts_membership" `
	AlamatBantuanWkm      string     `gorm:"-;type:json" json:"alamat_bantuan_wkm" `
	AlasanTdkMembership    string     `gorm:"-;type:json" json:"alasan_tdk_membership" `
	AlasanTdkMembershipDetail string   `gorm:"-;type:json" json:"alasan_tdk_membership_detail" `
	AlasanPendingMembership string     `gorm:"-;type:json" json:"alasan_pending_membership" `
	StsAsuransiPa         string     `gorm:"-;type:json" json:"sts_asuransi_pa" `
	AlasanTdkAsuransiPa   string     `gorm:"-;type:json" json:"alasan_tdk_asuransi_pa" `
	AlasanTdkAsuransiPaDetail string   `gorm:"-;type:json" json:"alasan_tdk_asuransi_pa_detail" `
	AlasanPendingAsuransiPa string     `gorm:"-;type:json" json:"alasan_pending_asuransi_pa" `
	StsAsuransiMtr        string     `gorm:"-;type:json" json:"sts_asuransi_mtr" `
	AlasanTdkAsuransiMtr  string     `gorm:"-;type:json" json:"alasan_tdk_asuransi_mtr" `
	AlasanTdkAsuransiMtrDetail string   `gorm:"-;type:json" json:"alasan_tdk_asuransi_mtr_detail" `
	AlasanPendingAsuransiMtr string     `gorm:"-;type:json" json:"alasan_pending_asuransi_mtr" `
	NoYgDihubTs          string     `gorm:"-;type:json" json:"no_yg_dihub_ts" `
	RenewalKe             uint32     `gorm:"-;type:json" json:"renewal_ke" `
	AngsuranWkm           string     `gorm:"-;type:json" json:"angsuran_wkm" `
	NoLeasWkm             string     `gorm:"-;type:json" json:"no_leas_wkm" `
	NmLeasWkm             string     `gorm:"-;type:json" json:"nm_leas_wkm" `
	TglProspectMembership  *time.Time `gorm:"-;type:json" json:"tgl_prospect_membership" `
	TglProspectAsuransiPa  *time.Time `gorm:"-;type:json" json:"tgl_prospect_asuransi_pa" `
	TglProspectAsuransiMtr *time.Time `gorm:"-;type:json" json:"tgl_prospect_asuransi_mtr" `
	SmFacebookWkm         string     `gorm:"-;type:json" json:"sm_facebook_wkm" `
	SmInstagramWkm        string     `gorm:"-;type:json" json:"sm_instagram_wkm" `
	SmTwitterWkm          string     `gorm:"-;type:json" json:"sm_twitter_wkm" `
	SmYoutubeWkm          string     `gorm:"-;type:json" json:"sm_youtube_wkm" `
	DpMtrWkm              uint32     `gorm:"-;type:json" json:"dp_mtr_wkm" `
	KdAktivitasJualMembership string  `gorm:"-;type:json" json:"kd_aktivitas_jual_membership" `
	JmlCallMembership      uint32     `gorm:"-;type:json" json:"jml_call_membership" `
	StsStnk      string     `gorm:"-;type:json" json:"sts_stnk" `

	NoKartu2 string `gorm:"-;type:json" json:"no_kartu2" `
	JnsMembershipSebelum string `gorm:"-;type:json" json:"jns_membership_sebelum" `

	DescDidikFkt string `json:"desc_didik_fkt" gorm:"-;type:json"`
	DescKerjaFkt string `json:"desc_kerja_fkt" gorm:"-;type:json"`
	DescAgamaFkt string `json:"desc_agama_fkt" gorm:"-;type:json"`
	DescKeluarBlnFkt string `json:"desc_bln_fkt" gorm:"-;type:json"`
	DescTujuanPakaiFkt string `json:"desc_tujuan_pakai_fkt" gorm:"-;type:json"`
	DescHobbyFkt string `json:"desc_hobby_fkt" gorm:"-;type:json"`

	MembershipID                 string `json:"membership_id" gorm:"-;type:json"`
	KirimKe                      string `json:"kirim_ke" gorm:"-;type:json"`
	JnsBayar                     string `json:"jns_bayar" gorm:"-;type:json"`
	TglJanjiBayar            string `json:"tgl_janji_bayar" gorm:"-;type:json"`
	JnsMembership            string `json:"jns_membership" gorm:"-;type:json"`
	JnsMembershipName            string `json:"jns_membership_name" gorm:"-;type:json"`
	TypeKartu                string `json:"type_kartu" gorm:"-;type:json"`
	KdPromoTransfer          string `json:"kd_promo_transfer" gorm:"-;type:json"`
	DescTglFaktur            string `json:"desc_tgl_faktur" gorm:"-;type:json"`
	AsuransiMtrTahun	int `json:"asuransi_mtr_tahun" gorm:"-;type:json"`
	DescTglLahirFkt          string `json:"desc_tgl_lahir_fkt" gorm:"-;type:json"`
	AsuransiNmMtr	string `json:"asuransi_nm_mtr" gorm:"-;type:json"`
	AsuransiNoMtr	string `json:"asuransi_no_mtr" gorm:"-;type:json"`
	JnsJualFktKet            string `json:"jns_jual_fkt_ket" gorm:"-;type:json"`
	NoHub                    string `json:"no_hub" gorm:"-;type:json"`

	AsuransiPAID        string     `json:"asuransi_pa_id" gorm:"-;type:json"`
	IDProdukAsuransiPA  string     `json:"id_produk_asuransi_pa" gorm:"-;type:json"`
	NamaProdukAsuransiPA string  `json:"nm_produk_asuransi_pa" gorm:"-;type:json"`
	NamaVendorPA        string  `json:"nm_vendor_pa" gorm:"-;type:json"`
	AmountAsuransiPA    string `json:"amount_asuransi_pa" gorm:"-;type:json"`
	RatePA              float64 `json:"rate_pa" gorm:"-;type:json"`
	AdminPA             uint64 `json:"admin_pa" gorm:"-;type:json"`

	AsuransiMTRID        string     `json:"asuransi_mtr_id" gorm:"-;type:json"`
	IDProdukAsuransiMTR  string     `json:"id_produk_asuransi_mtr" gorm:"-;type:json"`
	NamaProdukAsuransiMTR string  `json:"nm_produk_asuransi_mtr" gorm:"-;type:json"`
	Warna                string  `json:"warna" gorm:"-;type:json"`
	NamaVendorMTR        string  `json:"nm_vendor_mtr" gorm:"-;type:json"`
	RateMTR              float64 `json:"rate_mtr" gorm:"-;type:json"`
	AdminMTR             uint64 `json:"admin_mtr" gorm:"-;type:json"`




}


type TelesalesBalikanResponseList struct {
	NoMsn                  string     `gorm:"-;type:json" json:"no_msn" `
	NmCustomerFkt          string     `gorm:"-;type:json" json:"nm_customer_fkt" `
	TglUpdateKartuBalikan       string  `json:"tgl_update_kartu_balikan" gorm:"-;type:json"`
	AlasanBelumBayar2 string	 `json:"alasan_belum_bayar2" gorm:"-;type:json"`

}