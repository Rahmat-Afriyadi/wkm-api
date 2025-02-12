package entity

import "time"

type CustomerMtr struct {
	NoMsn                  string     `gorm:"primary_key;column:no_msn;" json:"no_msn" form:"no_msn"`
	NmCustomerFkt          string     `gorm:"column:nm_customer_fkt;" json:"nm_customer_fkt" form:"nm_customer_fkt"`
	KdDlr                  string     `gorm:"column:kd_dlr;" json:"kd_dlr" form:"kd_dlr"`
	NmDlr                  string     `gorm:"column:nm_dlr;" json:"nm_dlr" form:"nm_dlr"`
	NoRgk                  string     `gorm:"column:no_rgk;" json:"no_rgk" form:"no_rgk"`
	TglLahirFkt            *time.Time `gorm:"column:tgl_lahir_fkt;" json:"tgl_lahir_fkt" form:"tgl_lahir_fkt"`
	JnsKlmFkt              string     `gorm:"column:jns_klm_fkt;" json:"jns_klm_fkt" form:"jns_klm_fkt"`
	NoTelpFkt              string     `gorm:"column:no_telp_fkt;" json:"no_telp_fkt" form:"no_telp_fkt"`
	KetNoTelpFkt           string     `gorm:"column:ket_no_telp_fkt;" json:"ket_no_telp_fkt" form:"ket_no_telp_fkt"`
	NoHpFkt                string     `gorm:"column:no_hp_fkt;" json:"no_hp_fkt" form:"no_hp_fkt"`
	
	KetNoHpFkt             string     `gorm:"column:ket_no_hp_fkt;" json:"ket_no_hp_fkt" form:"ket_no_hp_fkt"`
	AlamatFkt              string     `gorm:"column:alamat_fkt;" json:"alamat_fkt" form:"alamat_fkt"`
	KelFkt                 string     `gorm:"column:kel_fkt;" json:"kel_fkt" form:"kel_fkt"`
	KecFkt                 string     `gorm:"column:kec_fkt;" json:"kec_fkt" form:"kec_fkt"`
	KotaFkt                string     `gorm:"column:kota_fkt;" json:"kota_fkt" form:"kota_fkt"`
	KodeposFkt             string     `gorm:"column:kodepos_fkt;" json:"kodepos_fkt" form:"kodepos_fkt"`
	PicPerusahaanFkt       string     `gorm:"column:pic_perusahaan_fkt;" json:"pic_perusahaan_fkt" form:"pic_perusahaan_fkt"`
	NoMtr                  string     `gorm:"column:no_mtr;" json:"no_mtr" form:"no_mtr"`
	NmMtr                  string     `gorm:"column:nm_mtr;" json:"nm_mtr" form:"nm_mtr"`
	TglMohon               *time.Time `gorm:"column:tgl_mohon;" json:"tgl_mohon" form:"tgl_mohon"`
	TglFaktur              *time.Time `gorm:"column:tgl_faktur;" json:"tgl_faktur" form:"tgl_faktur"`
	KodeKerjaFkt           string     `gorm:"column:kode_kerja_fkt;" json:"kode_kerja_fkt" form:"kode_kerja_fkt"`
	KodeDidikFkt           string     `gorm:"column:kode_didik_fkt;" json:"kode_didik_fkt" form:"kode_didik_fkt"`
	KeluarBlnFkt           string     `gorm:"column:keluar_bln_fkt;" json:"keluar_bln_fkt" form:"keluar_bln_fkt"`
	MotorHir               string     `gorm:"column:motor_hir;" json:"motor_hir" form:"motor_hir"`
	JnsBeli                string     `gorm:"column:jns_beli;" json:"jns_beli" form:"jns_beli"`
	BdnUsaha              string     `gorm:"column:bdn_usaha;" json:"bdn_usaha" form:"bdn_usaha"`
	JnsJualFkt            string     `gorm:"column:jns_jual_fkt;" json:"jns_jual_fkt" form:"jns_jual_fkt"`
	AktifJualFkt          string     `gorm:"column:aktif_jual_fkt;;" json:"aktif_jual_fkt" form:"aktif_jual_fkt"`
	NoKtpnpwpFkt          string     `gorm:"column:no_ktpnpwp_fkt;" json:"no_ktpnpwp_fkt" form:"no_ktpnpwp_fkt"`
	NoNpwpFkt             string     `gorm:"column:no_npwp_fkt;" json:"no_npwp_fkt" form:"no_npwp_fkt"`
	KerjaDiFkt            string     `gorm:"column:kerja_di_fkt;" json:"kerja_di_fkt" form:"kerja_di_fkt"`
	AlamatKtrFkt          string     `gorm:"column:alamat_ktr_fkt;" json:"alamat_ktr_fkt" form:"alamat_ktr_fkt"`
	PropKtrFkt            string     `gorm:"column:prop_ktr_fkt;" json:"prop_ktr_fkt" form:"prop_ktr_fkt"`
	KotaKtrFkt            string     `gorm:"column:kota_ktr_fkt;" json:"kota_ktr_fkt" form:"kota_ktr_fkt"`
	KelKtrFkt             string     `gorm:"column:kel_ktr_fkt;" json:"kel_ktr_fkt" form:"kel_ktr_fkt"`
	KecKtrFkt             string     `gorm:"column:kec_ktr_fkt;" json:"kec_ktr_fkt" form:"kec_ktr_fkt"`
	KodeposKtrFkt         string     `gorm:"column:kodepos_ktr_fkt;" json:"kodepos_ktr_fkt" form:"kodepos_ktr_fkt"`
	NoPolFkt              string     `gorm:"column:no_pol_fkt;" json:"no_pol_fkt" form:"no_pol_fkt"`
	HobbyFkt              string     `gorm:"column:hobby_fkt;" json:"hobby_fkt" form:"hobby_fkt"`
	StsSourceFkt          string     `gorm:"column:sts_source_fkt;" json:"sts_source_fkt" form:"sts_source_fkt"`
	NmSalesFkt            string     `gorm:"column:nm_sales_fkt;" json:"nm_sales_fkt" form:"nm_sales_fkt"`
	IdSalesFkt            string     `gorm:"column:id_sales_fkt;" json:"id_sales_fkt" form:"id_sales_fkt"`
	AgamaFkt              string     `gorm:"column:agama_fkt;" json:"agama_fkt" form:"agama_fkt"`
	TujuanPakaiFkt        string     `gorm:"column:tujuan_pakai_fkt;" json:"tujuan_pakai_fkt" form:"tujuan_pakai_fkt"`
	DpMtrFkt              int32    `gorm:"column:dp_mtr_fkt;" json:"dp_mtr_fkt" form:"dp_mtr_fkt"`
	CicilanMtrFkt         int32    `gorm:"column:cicilan_mtr_fkt;" json:"cicilan_mtr_fkt" form:"cicilan_mtr_fkt"`
	AngsuranFkt           string     `gorm:"column:angsuran_fkt;" json:"angsuran_fkt" form:"angsuran_fkt"`
	EmailFkt              string     `gorm:"column:email_fkt;" json:"email_fkt" form:"email_fkt"`
	NoLeasFkt             string     `gorm:"column:no_leas_fkt;" json:"no_leas_fkt" form:"no_leas_fkt"`
	JnsMotorFkt           string     `gorm:"column:jns_motor_fkt;" json:"jns_motor_fkt" form:"jns_motor_fkt"`
	SmDibeli              string     `gorm:"column:sm_dibeli;" json:"sm_dibeli" form:"sm_dibeli"`
	NoKk                  string     `gorm:"column:no_kk;" json:"no_kk" form:"no_kk"`
	TglAkhirTenor        *time.Time `gorm:"column:tgl_akhir_tenor;" json:"tgl_akhir_tenor" form:"tgl_akhir_tenor"`
	SmFacebookFkt        string     `gorm:"column:sm_facebook_fkt;" json:"sm_facebook_fkt" form:"sm_facebook_fkt"`
	SmInstagramFkt       string     `gorm:"column:sm_instagram_fkt;" json:"sm_instagram_fkt" form:"sm_instagram_fkt"`
	SmTwitterFkt         string     `gorm:"column:sm_twitter_fkt;" json:"sm_twitter_fkt" form:"sm_twitter_fkt"`
	SmYoutubeFkt         string     `gorm:"column:sm_youtube_fkt;" json:"sm_youtube_fkt" form:"sm_youtube_fkt"`
	TglLahirWkm          *time.Time `gorm:"column:tgl_lahir_wkm;" json:"tgl_lahir_wkm" form:"tgl_lahir_wkm"`
	JnsKlmWkm            string     `gorm:"column:jns_klm_wkm;" json:"jns_klm_wkm" form:"jns_klm_wkm"`
	NoTelpWkm            string     `gorm:"column:no_telp_wkm;" json:"no_telp_wkm" form:"no_telp_wkm"`
	KetNoTelpWkm         string     `gorm:"column:ket_no_telp_wkm;" json:"ket_no_telp_wkm" form:"ket_no_telp_wkm"`
	NoHpWkm              string     `gorm:"column:no_hp_wkm;" json:"no_hp_wkm" form:"no_hp_wkm"`
	NoWa                 string     `gorm:"column:no_wa;" json:"no_wa" form:"no_wa"`
	NoTelpKtrWkm         string     `gorm:"column:no_telp_ktr_wkm;" json:"no_telp_ktr_wkm" form:"no_telp_ktr_wkm"`
	EmailWkm             string     `gorm:"column:email_wkm;" json:"email_wkm" form:"email_wkm"`
	AlamatWkm            string     `gorm:"column:alamat_wkm;" json:"alamat_wkm" form:"alamat_wkm"`
	KetAlamatWkm         string     `gorm:"column:ket_alamat_wkm;" json:"ket_alamat_wkm" form:"ket_alamat_wkm"`
	KetNoInfo         string     `gorm:"column:ket_wa_info;" json:"ket_wa_info" form:"ket_wa_info"`
	RtWkm                string     `gorm:"column:rt_wkm;" json:"rt_wkm" form:"rt_wkm"`
	RwWkm                string     `gorm:"column:rw_wkm;" json:"rw_wkm" form:"rw_wkm"`
	KelWkm               string     `gorm:"column:kel_wkm;" json:"kel_wkm" form:"kel_wkm"`
	KecWkm               string     `gorm:"column:kec_wkm;" json:"kec_wkm" form:"kec_wkm"`
	KotaWkm              string     `gorm:"column:kota_wkm;" json:"kota_wkm" form:"kota_wkm"`
	KodeposWkm           string     `gorm:"column:kodepos_wkm;" json:"kodepos_wkm" form:"kodepos_wkm"`
	TglCallTele           *time.Time `gorm:"column:tgl_call_tele;" json:"tgl_call_tele" form:"tgl_call_tele"`
	JnsJualWkm            string     `gorm:"column:jns_jual_wkm;" json:"jns_jual_wkm" form:"jns_jual_wkm"`
	AktifJualWkm          string     `gorm:"column:aktif_jual_wkm;;" json:"aktif_jual_wkm" form:"aktif_jual_wkm"`
	KetAktifJualWkm       string     `gorm:"column:ket_aktif_jual_wkm;" json:"ket_aktif_jual_wkm" form:"ket_aktif_jual_wkm"`
	NoKtpnpwpWkm          string     `gorm:"column:no_ktpnpwp_wkm;" json:"no_ktpnpwp_wkm" form:"no_ktpnpwp_wkm"`
	KdUserTs              string     `gorm:"column:kd_user_ts;" json:"kd_user_ts" form:"kd_user_ts"`
	KerjaDiWkm            string     `gorm:"column:kerja_di_wkm;" json:"kerja_di_wkm" form:"kerja_di_wkm"`
	JabatanWkm            string     `gorm:"column:jabatan_wkm;" json:"jabatan_wkm" form:"jabatan_wkm"`
	AlamatKtrWkm          string     `gorm:"column:alamat_ktr_wkm;" json:"alamat_ktr_wkm" form:"alamat_ktr_wkm"`
	KotaKtrWkm            string     `gorm:"column:kota_ktr_wkm;" json:"kota_ktr_wkm" form:"kota_ktr_wkm"`
	RtKtrWkm              string     `gorm:"column:rt_ktr_wkm;" json:"rt_ktr_wkm" form:"rt_ktr_wkm"`
	RwKtrWkm              string     `gorm:"column:rw_ktr_wkm;" json:"rw_ktr_wkm" form:"rw_ktr_wkm"`
	KelKtrWkm             string     `gorm:"column:kel_ktr_wkm;" json:"kel_ktr_wkm" form:"kel_ktr_wkm"`
	KecKtrWkm             string     `gorm:"column:kec_ktr_wkm;" json:"kec_ktr_wkm" form:"kec_ktr_wkm"`
	KodeposKtrWkm         string     `gorm:"column:kodepos_ktr_wkm;" json:"kodepos_ktr_wkm" form:"kodepos_ktr_wkm"`
	NmCustomerWkm         string     `gorm:"column:nm_customer_wkm;" json:"nm_customer_wkm" form:"nm_customer_wkm"`
	KdDlrWkm              string     `gorm:"column:kd_dlr_wkm;;" json:"kd_dlr_wkm" form:"kd_dlr_wkm"`
	StsValidWkm           string     `gorm:"column:sts_valid_wkm;" json:"sts_valid_wkm" form:"sts_valid_wkm"`
	StsSource2Wkm         string     `gorm:"column:sts_source2_wkm;default:' '" json:"sts_source2_wkm" form:"sts_source2_wkm"`
	PicPerusahaanWkm      string     `gorm:"column:pic_perusahaan_wkm;" json:"pic_perusahaan_wkm" form:"pic_perusahaan_wkm"`
	HobbyWkmOthers        string     `gorm:"column:hobby_wkm_others;" json:"hobby_wkm_others" form:"hobby_wkm_others"`
	HobbyWkm              string     `gorm:"column:hobby_wkm;" json:"hobby_wkm" form:"hobby_wkm"`
	StsKawinWkm           string     `gorm:"column:sts_kawin_wkm;" json:"sts_kawin_wkm" form:"sts_kawin_wkm"`
	NmSalesWkm            string     `gorm:"column:nm_sales_wkm;" json:"nm_sales_wkm" form:"nm_sales_wkm"`
	AgamaWkm              string     `gorm:"column:agama_wkm;" json:"agama_wkm" form:"agama_wkm"`
	KodeDidikWkm          string     `gorm:"column:kode_didik_wkm;" json:"kode_didik_wkm" form:"kode_didik_wkm"`
	KodeKerjaWkm          string     `gorm:"column:kode_kerja_wkm;" json:"kode_kerja_wkm" form:"kode_kerja_wkm"`
	NmKerjaWkm            string     `gorm:"column:nm_kerja_wkm;" json:"nm_kerja_wkm" form:"nm_kerja_wkm"`
	TujuanPakaiWkm        string     `gorm:"column:tujuan_pakai_wkm;" json:"tujuan_pakai_wkm" form:"tujuan_pakai_wkm"`
	KeluarBlnWkm          string     `gorm:"column:keluar_bln_wkm;" json:"keluar_bln_wkm" form:"keluar_bln_wkm"`
	StsMembership         string     `gorm:"column:sts_membership;" json:"sts_membership" form:"sts_membership"`
	AlamatBantuanWkm      string     `gorm:"column:alamat_bantuan_wkm;" json:"alamat_bantuan_wkm" form:"alamat_bantuan_wkm"`
	AlasanTdkMembership    string     `gorm:"column:alasan_tdk_membership;" json:"alasan_tdk_membership" form:"alasan_tdk_membership"`
	AlasanTdkMembershipDetail string   `gorm:"column:alasan_tdk_membership_detail;" json:"alasan_tdk_membership_detail" form:"alasan_tdk_membership_detail"`
	AlasanPendingMembership string     `gorm:"column:alasan_pending_membership;" json:"alasan_pending_membership" form:"alasan_pending_membership"`
	StsAsuransiPa         string     `gorm:"column:sts_asuransi_pa;" json:"sts_asuransi_pa" form:"sts_asuransi_pa"`
	AlasanTdkAsuransiPa   string     `gorm:"column:alasan_tdk_asuransi_pa;" json:"alasan_tdk_asuransi_pa" form:"alasan_tdk_asuransi_pa"`
	AlasanTdkAsuransiPaDetail string   `gorm:"column:alasan_tdk_asuransi_pa_detail;" json:"alasan_tdk_asuransi_pa_detail" form:"alasan_tdk_asuransi_pa_detail"`
	AlasanPendingAsuransiPa string     `gorm:"column:alasan_pending_asuransi_pa;" json:"alasan_pending_asuransi_pa" form:"alasan_pending_asuransi_pa"`
	StsAsuransiMtr        string     `gorm:"column:sts_asuransi_mtr;" json:"sts_asuransi_mtr" form:"sts_asuransi_mtr"`
	AlasanTdkAsuransiMtr  string     `gorm:"column:alasan_tdk_asuransi_mtr;" json:"alasan_tdk_asuransi_mtr" form:"alasan_tdk_asuransi_mtr"`
	AlasanTdkAsuransiMtrDetail string   `gorm:"column:alasan_tdk_asuransi_mtr_detail;" json:"alasan_tdk_asuransi_mtr_detail" form:"alasan_tdk_asuransi_mtr_detail"`
	AlasanPendingAsuransiMtr string     `gorm:"column:alasan_pending_asuransi_mtr;" json:"alasan_pending_asuransi_mtr" form:"alasan_pending_asuransi_mtr"`
	NoYgDihubTs          string     `gorm:"column:no_yg_dihub_ts;" json:"no_yg_dihub_ts" form:"no_yg_dihub_ts"`
	RenewalKe             uint32     `gorm:"column:renewal_ke;" json:"renewal_ke" form:"renewal_ke"`
	AngsuranWkm           string     `gorm:"column:angsuran_wkm;" json:"angsuran_wkm" form:"angsuran_wkm"`
	NoLeasWkm             string     `gorm:"column:no_leas_wkm;" json:"no_leas_wkm" form:"no_leas_wkm"`
	NmLeasWkm             string     `gorm:"column:nm_leas_wkm;" json:"nm_leas_wkm" form:"nm_leas_wkm"`
	TglProspectMembership  *time.Time `gorm:"column:tgl_prospect_membership;" json:"tgl_prospect_membership" form:"tgl_prospect_membership"`
	TglProspectAsuransiPa  *time.Time `gorm:"column:tgl_prospect_asuransi_pa;" json:"tgl_prospect_asuransi_pa" form:"tgl_prospect_asuransi_pa"`
	TglProspectAsuransiMtr *time.Time `gorm:"column:tgl_prospect_asuransi_mtr;" json:"tgl_prospect_asuransi_mtr" form:"tgl_prospect_asuransi_mtr"`
	SmFacebookWkm         string     `gorm:"column:sm_facebook_wkm;" json:"sm_facebook_wkm" form:"sm_facebook_wkm"`
	SmInstagramWkm        string     `gorm:"column:sm_instagram_wkm;" json:"sm_instagram_wkm" form:"sm_instagram_wkm"`
	SmTwitterWkm          string     `gorm:"column:sm_twitter_wkm;" json:"sm_twitter_wkm" form:"sm_twitter_wkm"`
	SmYoutubeWkm          string     `gorm:"column:sm_youtube_wkm;" json:"sm_youtube_wkm" form:"sm_youtube_wkm"`
	DpMtrWkm              uint32     `gorm:"column:dp_mtr_wkm;" json:"dp_mtr_wkm" form:"dp_mtr_wkm"`
	KdAktivitasJualMembership string  `gorm:"column:kd_aktivitas_jual_membership;" json:"kd_aktivitas_jual_membership" form:"kd_aktivitas_jual_membership"`
	JmlCallMembership      uint32     `gorm:"column:jml_call_membership;" json:"jml_call_membership" form:"jml_call_membership"`
	Modified               *time.Time `gorm:"column:modified;" json:"modified" form:"modified"`


	Memberships           []Membership    `form:"memberships" json:"memberships" gorm:"foreignKey:NoMSN"`
	AsuransiPa            AsuransiPA    `form:"asuransi_pa" json:"asuransi_pa" gorm:"foreignKey:NoMSN"`
	AsuransiMtr           AsuransiMtr    `form:"asuransi_mtr" json:"asuransi_mtr" gorm:"foreignKey:NoMSN"`


}

func (CustomerMtr) TableName() string {
	return "customer_mtr"
}