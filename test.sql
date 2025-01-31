select * from (
    select no_msn, nm_customer11, case when trim(no_wa) REGEXP '^[+62]|^[0-9]*$' and no_wa is not null and no_wa not like '021%' then no_wa when trim(sms_no) REGEXP '^[+62]|^[0-9]*$' and sms_no is not null and sms_no not like '021%' then sms_no when trim(no_telp2) REGEXP '^[+62]|^[0-9]*$' and no_telp2 is not null and no_telp2 not like '021%' then no_telp2 when trim(no_telp1) REGEXP '^[+62]|^[0-9]*$' and no_telp1 is not null and no_telp1 not like '021%' then no_telp1 when trim(no_hp2) REGEXP '^[+62]|^[0-9]*$' and no_hp2 is not null and no_hp2 not like '021%' then no_hp2 when trim(no_hp1)REGEXP '^[+62]|^[0-9]*$' and no_hp1 is not null and no_hp1 not like '021%' then no_hp1 end as no_wa, tgl_akhir_tenor from (select no_msn,nm_customer11, case when ket_no_telp1 in ('1','1A','1B') then no_telp1 else null end as 'no_telp1', case when ket_no_telp2 in ('1','2') then no_telp2 else null end as 'no_telp2', case when no_hp2 is not null and no_hp2 != '' then no_hp2 else null end as 'no_hp2', case when ket_no_hp1 in ('1','1A','1B') then no_hp1 else null end as 'no_hp1', no_yg_dihub_renewal,sms_no,no_wa,tgl_akhir_tenor from tr_wms_faktur3))





UPDATE tr_wms_faktur3 SET 
alamat21=?,
ket_alamat21=?,
rt2=?,
rw2=?,
kodepos2=?,
kel2=?,
kec2=?,
kota2=?,
ket_no_telp1=?,
no_telp2=?,
ket_no_telp2=?,
ket_no_hp1=?,
no_hp2=?,
no_telp_ktr2=?,
sts_sms=?,
no_yg_dihub_renewal=?,
email2=?,
alamat_bantuan=?,
jns_klm2=?,
sts_kawin=?,
agama2=?,
kode_didik2=?,
keluar_bln2=?,
tujuan_pakai2=?,
terima_kartu=?,
sts_renewal=?,
kd_card=?,
sts_jenis_bayar=?,
tambahan=?,
kd_aktivitas_jual_r=?,
alasan_tdk_renewal=?,
alasan_tdk_renewal2=?,
alasan_pending_renewal=?,
sts_kirim=?,
sts_asuransi=?,
kerja_di=?,
kode_kerja2=?,
alamat_ktr=?,
rt_ktr=?,
rw_ktr=?,
kodepos_ktr=?,
kel_ktr=?,
kec_ktr=?,
kota_ktr=?,
hobby2=?,
tgl_verifikasi=?,
tgl_bayar_renewal=?,
nama_ktp=?,
tgl_prospect=NULL,jml_call_renewal=jml_call_renewal+1,tgl_cetak_tanda_terima=NULL 
WHERE no_msn=? 
AND kd_user=?




