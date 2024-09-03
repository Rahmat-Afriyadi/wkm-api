select * from (
    select no_msn, nm_customer11, case when trim(no_wa) REGEXP '^[+62]|^[0-9]*$' and no_wa is not null and no_wa not like '021%' then no_wa when trim(sms_no) REGEXP '^[+62]|^[0-9]*$' and sms_no is not null and sms_no not like '021%' then sms_no when trim(no_telp2) REGEXP '^[+62]|^[0-9]*$' and no_telp2 is not null and no_telp2 not like '021%' then no_telp2 when trim(no_telp1) REGEXP '^[+62]|^[0-9]*$' and no_telp1 is not null and no_telp1 not like '021%' then no_telp1 when trim(no_hp2) REGEXP '^[+62]|^[0-9]*$' and no_hp2 is not null and no_hp2 not like '021%' then no_hp2 when trim(no_hp1)REGEXP '^[+62]|^[0-9]*$' and no_hp1 is not null and no_hp1 not like '021%' then no_hp1 end as no_wa, tgl_akhir_tenor from (select no_msn,nm_customer11, case when ket_no_telp1 in ('1','1A','1B') then no_telp1 else null end as 'no_telp1', case when ket_no_telp2 in ('1','2') then no_telp2 else null end as 'no_telp2', case when no_hp2 is not null and no_hp2 != '' then no_hp2 else null end as 'no_hp2', case when ket_no_hp1 in ('1','1A','1B') then no_hp1 else null end as 'no_hp1', no_yg_dihub_renewal,sms_no,no_wa,tgl_akhir_tenor from tr_wms_faktur3))