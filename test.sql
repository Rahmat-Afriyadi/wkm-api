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



UPDATE tr_wms_faktur3 a
inner join membership b on a.no_msn=b.no_msn and a.sts_cetak3=b.renewal_ke
inner join customer_mtr c on a.no_msn = c.no_msn
SET a.print=0, a.alamat_bantuan=c.alamat_bantuan_wkm,a.alamat_ktr=c.alamat_ktr_wkm,a.alamat21=c.alamat_wkm,a.kd_card=b.jns_membership,a.kd_aktivitas_jual_r=c.kd_aktivitas_jual_membership,a.kec_ktr=c.kec_ktr_wkm,a.kec2=c.kec_wkm,a.kel_ktr=c.kel_ktr_wkm,a.kel2=c.kel_wkm,a.kerja_di=c.kerja_di_wkm,
a.ket_alamat21=c.ket_alamat_wkm,
a.ket_no_hp1=c.ket_no_hp_fkt,a.ket_no_telp1=c.ket_no_telp_fkt,a.ket_no_telp2=c.ket_no_telp_wkm,
a.sts_kirim=b.kirim_ke,a.kodepos_ktr=c.kodepos_ktr_wkm,
a.kodepos2=c.kodepos_wkm,a.kota_ktr=c.kota_ktr_wkm,a.kota2=c.kota_wkm,a.nama_ktp=c.nama_customer_wkm,a.no_hp2=c.no_hp_wkm,a.no_telp2=c.no_telp_wkm,a.no_yg_dihub_renewal=c.no_yg_dihub_ts,a.rt_ktr=c.rt_ktr_wkm,a.rt2=c.rt_wkm,a.rw_ktr=c.rw_ktr_wkm,a.rw2=c.rw_wkm,a.sts_renewal=c.sts_membership,a.tgl_verifikasi=c.tgl_call_tele,a.tgl_bayar_renewal=b.tgl_janji_bayar,a.sts_jenis_bayar=b.jns_bayar, a.sts_asuransi_pa='O' where a.no_msn = ? and a.kd_user =?


select * from (select no_msn, nm_customer11, kd_user, no_yg_dihub_renewal, case when trim(no_wa) REGEXP '^[+62]|^[0-9]*$' and no_wa is not null and no_wa not like '021%' then no_wa when trim(sms_no) REGEXP '^[+62]|^[0-9]*$' and sms_no is not null and sms_no not like '021%' then sms_no when trim(no_telp2) REGEXP '^[+62]|^[0-9]*$' and no_telp2 is not null and no_telp2 not like '021%' then no_telp2 when trim(no_telp1) REGEXP '^[+62]|^[0-9]*$' and no_telp1 is not null and no_telp1 not like '021%' then no_telp1 when trim(no_hp2) REGEXP '^[+62]|^[0-9]*$' and no_hp2 is not null and no_hp2 not like '021%' then no_hp2 when trim(no_hp1)REGEXP '^[+62]|^[0-9]*$' and no_hp1 is not null and no_hp1 not like '021%' then no_hp1 end as no_wa, tgl_akhir_tenor from (select no_msn,nm_customer11,kd_user,case when ket_no_telp1 in ('1','1A','1B') then no_telp1 else no_hp1 end as 'no_telp1', case when ket_no_telp2 in ('1','2') then no_telp2 else no_hp1 end as 'no_telp2',case when no_hp2 is not null and no_hp2 != '' then no_hp2 else no_hp1 end as 'no_hp2', case when ket_no_hp1 in ('1','1A','1B') then no_hp1 else no_hp1 end as 'no_hp1', no_yg_dihub_renewal,sms_no,no_wa,tgl_akhir_tenor from tr_wms_faktur3 where month(tgl_faktur) = 4 and day(tgl_faktur)=29 and jns_beli = 1 ) t) u where no_wa is not null and no_wa != ''