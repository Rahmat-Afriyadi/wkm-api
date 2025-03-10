package entity

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
	"wkm/config"
)

var	(
	_, conn = config.GetConnection()
	_, connECardplus = config.GetConnectionECardPlus()
)



func GenerateNoTT() (string, error) {
	var lastNoTT string
	nextNoTT := 0
	err := conn.QueryRow("select VAL_CHAR from mst_runnum WHERE VAL_ID = 'tr_wms_faktur3' and VAL_TAG = 'no_tanda_terima'").Scan(&lastNoTT)
	if err != nil {
		return "", err
	}
	fmt.Sscanf(lastNoTT[1:], "%d", &nextNoTT)
	nextNoTTF := fmt.Sprintf("T%09d", nextNoTT+1)
	_, err = conn.Exec("UPDATE mst_runnum SET VAL_CHAR = ? WHERE VAL_ID = 'tr_wms_faktur3' and VAL_TAG = 'no_tanda_terima'", nextNoTTF)
	if err != nil {
		return "", err
	}
	return nextNoTTF, nil
}

func GenerateEcardNumber(jnsCard string) (string, error) {
	now := time.Now()
	jnsMembership := ""
	if strings.Contains(jnsCard, "BASIC") {
		jnsMembership = "01"
	}else if strings.Contains(jnsCard, "GOLD") {
		jnsMembership = "02"
	}else if strings.Contains(jnsCard, "PLATINUM") {
		jnsMembership = "03"
	}else if strings.Contains(jnsCard, "PLATINUM PLUS") {
		jnsMembership = "23"
	}

	formatKartuSekarang := fmt.Sprintf("11%s %s", jnsMembership, now.Format("0601"))
	var nomorKartuSekarang string
	nomorUrutKartuSekarang :=0
	err := connECardplus.QueryRow("select no_kartu from member where no_kartu like ? order by no_kartu desc", formatKartuSekarang+"%").Scan(&nomorKartuSekarang)
	if err != nil {
		if err == sql.ErrNoRows {
			return  formatKartuSekarang + " 0000 0001", nil
		}else {
			return "", err
		}
	}
	if nomorKartuSekarang != "" {
		fmt.Sscanf(strings.ReplaceAll(strings.TrimSpace(nomorKartuSekarang[10:]), " ", ""), "%d", &nomorUrutKartuSekarang)
		formatUrut := fmt.Sprintf("%08d", nomorUrutKartuSekarang+1) 
		return fmt.Sprintf("%s %s %s", formatKartuSekarang, formatUrut[:4], formatUrut[4:]), nil
	}
	return "", nil
}