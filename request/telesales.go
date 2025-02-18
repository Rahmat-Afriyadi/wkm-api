package request

import "time"

type CustomerMtr struct {
	NoMsn                  string     ` json:"no_msn" form:"no_msn"`
	NmCustomerFkt          string     ` json:"nm_customer_fkt" form:"nm_customer_fkt"`
	KdDlr                  string     ` json:"kd_dlr" form:"kd_dlr"`
	NmDlr                  string     ` json:"nm_dlr" form:"nm_dlr"`
	NoRgk                  string     ` json:"no_rgk" form:"no_rgk"`
	TglLahirFkt            *time.Time ` json:"tgl_lahir_fkt" form:"tgl_lahir_fkt"`
	JnsKlmFkt              string     ` json:"jns_klm_fkt" form:"jns_klm_fkt"`
	NoTelpFkt              string     ` json:"no_telp_fkt" form:"no_telp_fkt"`
	KetNoTelpFkt           string     ` json:"ket_no_telp_fkt" form:"ket_no_telp_fkt"`
	NoHpFkt                string     ` json:"no_hp_fkt" form:"no_hp_fkt"`

	KetNoHpFkt             string     ` json:"ket_no_hp_fkt" form:"ket_no_hp_fkt"`
	AlamatFkt              string     ` json:"alamat_fkt" form:"alamat_fkt"`
	KelFkt                 string     ` json:"kel_fkt" form:"kel_fkt"`
	KecFkt                 string     ` json:"kec_fkt" form:"kec_fkt"`
	KotaFkt                string     ` json:"kota_fkt" form:"kota_fkt"`
	KodeposFkt             string     ` json:"kodepos_fkt" form:"kodepos_fkt"`
	PicPerusahaanFkt       string     ` json:"pic_perusahaan_fkt" form:"pic_perusahaan_fkt"`
	NoMtr                  string     ` json:"no_mtr" form:"no_mtr"`
	NmMtr                  string     ` json:"nm_mtr" form:"nm_mtr"`
	TglMohon               *time.Time ` json:"tgl_mohon" form:"tgl_mohon"`
	TglFaktur              *time.Time ` json:"tgl_faktur" form:"tgl_faktur"`
	KodeKerjaFkt           string     ` json:"kode_kerja_fkt" form:"kode_kerja_fkt"`
	KodeDidikFkt           string     ` json:"kode_didik_fkt" form:"kode_didik_fkt"`
	KeluarBlnFkt           string     ` json:"keluar_bln_fkt" form:"keluar_bln_fkt"`
	MotorHir               string     ` json:"motor_hir" form:"motor_hir"`
	JnsBeli                string     ` json:"jns_beli" form:"jns_beli"`
	BdnUsaha              string     ` json:"bdn_usaha" form:"bdn_usaha"`
	JnsJualFkt            string     ` json:"jns_jual_fkt" form:"jns_jual_fkt"`
	AktifJualFkt          string     `json:"aktif_jual_fkt" form:"aktif_jual_fkt"`
	NoKtpnpwpFkt          string     ` json:"no_ktpnpwp_fkt" form:"no_ktpnpwp_fkt"`
	NoNpwpFkt             string     ` json:"no_npwp_fkt" form:"no_npwp_fkt"`
	KerjaDiFkt            string     ` json:"kerja_di_fkt" form:"kerja_di_fkt"`
	AlamatKtrFkt          string     ` json:"alamat_ktr_fkt" form:"alamat_ktr_fkt"`
	PropKtrFkt            string     ` json:"prop_ktr_fkt" form:"prop_ktr_fkt"`
	KotaKtrFkt            string     ` json:"kota_ktr_fkt" form:"kota_ktr_fkt"`
	KelKtrFkt             string     ` json:"kel_ktr_fkt" form:"kel_ktr_fkt"`
	KecKtrFkt             string     ` json:"kec_ktr_fkt" form:"kec_ktr_fkt"`
	KodeposKtrFkt         string     ` json:"kodepos_ktr_fkt" form:"kodepos_ktr_fkt"`
	NoPolFkt              string     ` json:"no_pol_fkt" form:"no_pol_fkt"`
	NoPolwKM              string     `json:"no_pol_wkm" form:"no_pol_wkm"`
	HobbyFkt              string     ` json:"hobby_fkt" form:"hobby_fkt"`
	StsSourceFkt          string     ` json:"sts_source_fkt" form:"sts_source_fkt"`
	NmSalesFkt            string     ` json:"nm_sales_fkt" form:"nm_sales_fkt"`
	IdSalesFkt            string     ` json:"id_sales_fkt" form:"id_sales_fkt"`
	AgamaFkt              string     ` json:"agama_fkt" form:"agama_fkt"`
	TujuanPakaiFkt        string     ` json:"tujuan_pakai_fkt" form:"tujuan_pakai_fkt"`
	DpMtrFkt              int32    ` json:"dp_mtr_fkt" form:"dp_mtr_fkt"`
	CicilanMtrFkt         int32    ` json:"cicilan_mtr_fkt" form:"cicilan_mtr_fkt"`
	AngsuranFkt           string     ` json:"angsuran_fkt" form:"angsuran_fkt"`
	EmailFkt              string     ` json:"email_fkt" form:"email_fkt"`
	NoLeasFkt             string     ` json:"no_leas_fkt" form:"no_leas_fkt"`
	JnsMotorFkt           string     ` json:"jns_motor_fkt" form:"jns_motor_fkt"`
	SmDibeli              string     ` json:"sm_dibeli" form:"sm_dibeli"`
	NoKk                  string     ` json:"no_kk" form:"no_kk"`
	TglAkhirTenor        *time.Time ` json:"tgl_akhir_tenor" form:"tgl_akhir_tenor"`
	SmFacebookFkt        string     ` json:"sm_facebook_fkt" form:"sm_facebook_fkt"`
	SmInstagramFkt       string     ` json:"sm_instagram_fkt" form:"sm_instagram_fkt"`
	SmTwitterFkt         string     ` json:"sm_twitter_fkt" form:"sm_twitter_fkt"`
	SmYoutubeFkt         string     ` json:"sm_youtube_fkt" form:"sm_youtube_fkt"`
	TglLahirWkm          *time.Time ` json:"tgl_lahir_wkm" form:"tgl_lahir_wkm"`
	JnsKlmWkm            string     ` json:"jns_klm_wkm" form:"jns_klm_wkm"`
	NoTelpWkm            string     ` json:"no_telp_wkm" form:"no_telp_wkm"`
	KetNoTelpWkm         string     ` json:"ket_no_telp_wkm" form:"ket_no_telp_wkm"`
	NoHpWkm              string     ` json:"no_hp_wkm" form:"no_hp_wkm"`
	NoWa                 string     ` json:"no_wa" form:"no_wa"`
	NoTelpKtrWkm         string     ` json:"no_telp_ktr_wkm" form:"no_telp_ktr_wkm"`
	EmailWkm             string     ` json:"email_wkm" form:"email_wkm"`
	AlamatWkm            string     ` json:"alamat_wkm" form:"alamat_wkm"`
	KetAlamatWkm         string     ` json:"ket_alamat_wkm" form:"ket_alamat_wkm"`
	RtWkm                string     ` json:"rt_wkm" form:"rt_wkm"`
	RwWkm                string     ` json:"rw_wkm" form:"rw_wkm"`
	KelWkm               string     ` json:"kel_wkm" form:"kel_wkm"`
	KecWkm               string     ` json:"kec_wkm" form:"kec_wkm"`
	KotaWkm              string     ` json:"kota_wkm" form:"kota_wkm"`
	KodeposWkm           string     ` json:"kodepos_wkm" form:"kodepos_wkm"`
	TglCallTele           *time.Time ` json:"tgl_call_tele" form:"tgl_call_tele"`
	JnsJualWkm            string     ` json:"jns_jual_wkm" form:"jns_jual_wkm"`
	AktifJualWkm          string     `json:"aktif_jual_wkm" form:"aktif_jual_wkm"`
	KetAktifJualWkm       string     ` json:"ket_aktif_jual_wkm" form:"ket_aktif_jual_wkm"`
	NoKtpnpwpWkm          string     ` json:"no_ktpnpwp_wkm" form:"no_ktpnpwp_wkm"`
	KdUserTs              string     ` json:"kd_user_ts" form:"kd_user_ts"`
	KerjaDiWkm            string     ` json:"kerja_di_wkm" form:"kerja_di_wkm"`
	JabatanWkm            string     ` json:"jabatan_wkm" form:"jabatan_wkm"`
	AlamatKtrWkm          string     ` json:"alamat_ktr_wkm" form:"alamat_ktr_wkm"`
	KotaKtrWkm            string     ` json:"kota_ktr_wkm" form:"kota_ktr_wkm"`
	RtKtrWkm              string     ` json:"rt_ktr_wkm" form:"rt_ktr_wkm"`
	RwKtrWkm              string     ` json:"rw_ktr_wkm" form:"rw_ktr_wkm"`
	KelKtrWkm             string     ` json:"kel_ktr_wkm" form:"kel_ktr_wkm"`
	KecKtrWkm             string     ` json:"kec_ktr_wkm" form:"kec_ktr_wkm"`
	KodeposKtrWkm         string     ` json:"kodepos_ktr_wkm" form:"kodepos_ktr_wkm"`
	NmCustomerWkm         string     ` json:"nm_customer_wkm" form:"nm_customer_wkm"`
	KdDlrWkm              string     `json:"kd_dlr_wkm" form:"kd_dlr_wkm"`
	StsValidWkm           string     ` json:"sts_valid_wkm" form:"sts_valid_wkm"`
	StsSource2Wkm         string     ` json:"sts_source2_wkm" form:"sts_source2_wkm"`
	PicPerusahaanWkm      string     ` json:"pic_perusahaan_wkm" form:"pic_perusahaan_wkm"`
	HobbyWkmOthers        string     ` json:"hobby_wkm_others" form:"hobby_wkm_others"`
	HobbyWkm              string     ` json:"hobby_wkm" form:"hobby_wkm"`
	StsKawinWkm           string     ` json:"sts_kawin_wkm" form:"sts_kawin_wkm"`
	NmSalesWkm            string     ` json:"nm_sales_wkm" form:"nm_sales_wkm"`
	AgamaWkm              string     ` json:"agama_wkm" form:"agama_wkm"`
	KodeDidikWkm          string     ` json:"kode_didik_wkm" form:"kode_didik_wkm"`
	KodeKerjaWkm          string     ` json:"kode_kerja_wkm" form:"kode_kerja_wkm"`
	NmKerjaWkm            string     ` json:"nm_kerja_wkm" form:"nm_kerja_wkm"`
	TujuanPakaiWkm        string     ` json:"tujuan_pakai_wkm" form:"tujuan_pakai_wkm"`
	KeluarBlnWkm          string     ` json:"keluar_bln_wkm" form:"keluar_bln_wkm"`
	StsMembership         string     ` json:"sts_membership" form:"sts_membership"`
	AlamatBantuanWkm      string     ` json:"alamat_bantuan_wkm" form:"alamat_bantuan_wkm"`
	AlasanTdkMembership    string     ` json:"alasan_tdk_membership" form:"alasan_tdk_membership"`
	AlasanTdkMembershipDetail string   ` json:"alasan_tdk_membership_detail" form:"alasan_tdk_membership_detail"`
	AlasanPendingMembership string     ` json:"alasan_pending_membership" form:"alasan_pending_membership"`
	StsAsuransiPa         string     ` json:"sts_asuransi_pa" form:"sts_asuransi_pa"`
	AlasanTdkAsuransiPa   string     ` json:"alasan_tdk_asuransi_pa" form:"alasan_tdk_asuransi_pa"`
	AlasanTdkAsuransiPaDetail string   ` json:"alasan_tdk_asuransi_pa_detail" form:"alasan_tdk_asuransi_pa_detail"`
	AlasanPendingAsuransiPa string     ` json:"alasan_pending_asuransi_pa" form:"alasan_pending_asuransi_pa"`
	StsAsuransiMtr        string     ` json:"sts_asuransi_mtr" form:"sts_asuransi_mtr"`
	AlasanTdkAsuransiMtr  string     ` json:"alasan_tdk_asuransi_mtr" form:"alasan_tdk_asuransi_mtr"`
	AlasanTdkAsuransiMtrDetail string   ` json:"alasan_tdk_asuransi_mtr_detail" form:"alasan_tdk_asuransi_mtr_detail"`
	AlasanPendingAsuransiMtr string     ` json:"alasan_pending_asuransi_mtr" form:"alasan_pending_asuransi_mtr"`
	NoYgDihubTs          string     ` json:"no_yg_dihub_ts" form:"no_yg_dihub_ts"`
	KetNoInfo         string     ` json:"ket_wa_info" form:"ket_wa_info"`

	KetHubTs          string     ` json:"ket_hub_ts" form:"ket_hub_ts"`
	RenewalKe             uint32     ` json:"renewal_ke" form:"renewal_ke"`
	AngsuranWkm           string     ` json:"angsuran_wkm" form:"angsuran_wkm"`
	NoLeasWkm             string     ` json:"no_leas_wkm" form:"no_leas_wkm"`
	NmLeasWkm             string     ` json:"nm_leas_wkm" form:"nm_leas_wkm"`
	TglProspectMembership  *time.Time ` json:"tgl_prospect_membership" form:"tgl_prospect_membership"`
	TglProspectAsuransiPa  *time.Time `gorm:"column:tgl_prospect_asuransi_pa;" json:"tgl_prospect_asuransi_pa" form:"tgl_prospect_asuransi_pa"`
	TglProspectAsuransiMtr *time.Time `gorm:"column:tgl_prospect_asuransi_mtr;" json:"tgl_prospect_asuransi_mtr" form:"tgl_prospect_asuransi_mtr"`

	SmFacebookWkm         string     ` json:"sm_facebook_wkm" form:"sm_facebook_wkm"`
	SmInstagramWkm        string     ` json:"sm_instagram_wkm" form:"sm_instagram_wkm"`
	SmTwitterWkm          string     ` json:"sm_twitter_wkm" form:"sm_twitter_wkm"`
	SmYoutubeWkm          string     ` json:"sm_youtube_wkm" form:"sm_youtube_wkm"`
	DpMtrWkm              uint32     ` json:"dp_mtr_wkm" form:"dp_mtr_wkm"`
	KdAktivitasJualMembership string  ` json:"kd_aktivitas_jual_membership" form:"kd_aktivitas_jual_membership"`
	JmlCallMembership      uint32     ` json:"jml_call_membership" form:"jml_call_membership"`
	Modified               *time.Time ` json:"modified" form:"modified"`


	// khusus membership
	JnsMembership          string     `json:"jns_membership" form:"jns_membership"`
	JnsBayar                 string    `json:"jns_bayar" form:"jns_bayar"`
	KirimKe                string     `json:"kirim_ke" form:"kirim_ke"`
	TglJanjiBayar          *time.Time `json:"tgl_janji_bayar" form:"tgl_janji_bayar"`
	KdPromoTransfer          string     `json:"kd_promo_transfer" form:"kd_promo_transfer"`
	TypeKartu                string    `json:"type_kartu" gorm:"column:type_kartu" form:"type_kartu"`


	// khusus Asuransi Motor
	Warna          string     `json:"warna" form:"warna"`
	IdProdukAsuransIMotor     string     `json:"id_produk_asuransi_mtr" form:"id_produk_asuransi_mtr"`
	OTR                       uint64     `json:"asuransi_mtr_otr" form:"asuransi_mtr_otr"`
	Amount                    uint64     `json:"asuransi_mtr_amount" form:"asuransi_mtr_amount"`
	ThnMtr                    uint32     `json:"asuransi_mtr_tahun" form:"asuransi_mtr_tahun"`
	
	// khusus Asuransi Pa
	IdProdukAsuransIPa          string     `json:"id_produk_asuransi_pa" form:"id_produk_asuransi_pa"`
	AmountPa                    uint64     `json:"amount_asuransi_pa" form:"amount_asuransi_pa"`
}