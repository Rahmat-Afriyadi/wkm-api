package query

func NewQueryUpdateAmbilData() string {
	return "update customer_mtr a inner join tr_wms_faktur3 b on a.no_msn = b.no_msn set a.kd_user_ts = b.kd_user,a.renewal_ke=b.sts_cetak3,a.tgl_call_tele=now(),a.modified=now(),a.sts_membership=b.sts_renewal where a.no_msn = ?"
}

func NewQueryUpdateAmbilDataKonfirmasi() string {
	return "update customer_mtr a inner join tr_wms_faktur3 b on a.no_msn = b.no_msn set a.kd_user_ts = b.kd_user,a.renewal_ke=b.sts_cetak3,a.tgl_call_tele=b.tgl_verifikasi,a.modified=now(),a.sts_membership=b.sts_renewal where a.no_msn = ?"
}