package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
	"wkm/entity"
	"wkm/request"

	"gorm.io/gorm"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

type ECardplusRepository interface {
	CreateToken(no_hp string) entity.Token
	FindCustomer(no_hp string) entity.CustomerMtr
	GetFakturByNoMsn(no_msn string) entity.Faktur3
	GetCustomerByNoMsn(no_msn string) entity.CustomerMtr
	CreateUser(data request.CreateECardplusUserRequest) (entity.UserECardPlus, error)
	GenerateEMembership(data string, user_id string) (entity.Membership, error)
}

type eCardplusRepository struct {
	conn     *sql.DB
	connGorm *gorm.DB
	connECardplus     *sql.DB
	gormECardplus *gorm.DB
}

func NewECardplusRepository(conn *sql.DB, connGorm *gorm.DB,connECardplus *sql.DB, gormECardplus *gorm.DB) ECardplusRepository {
	return &eCardplusRepository{
		conn:     conn,
		connGorm: connGorm,
		connECardplus:connECardplus,
		gormECardplus:gormECardplus,
	}
}

func (ec *eCardplusRepository) CreateToken(no_hp string) entity.Token {
	token := entity.Token{NoHp: no_hp, Token:RandomString(20)}
	ec.gormECardplus.Save(&token)
	return token
}

func (ec *eCardplusRepository) GetFakturByNoMsn(no_msn string) entity.Faktur3 {
	faktur3 := entity.Faktur3{NoMsn: no_msn}
	ec.connGorm.Find(&faktur3)
	return faktur3
}

func (ec *eCardplusRepository) GetCustomerByNoMsn(no_msn string) entity.CustomerMtr {
	customeerMtr := entity.CustomerMtr{NoMsn: no_msn}
	ec.connGorm.Find(&customeerMtr)
	return customeerMtr
}

func (ec *eCardplusRepository)	GenerateEMembership(no_msn string, user_id string) (entity.Membership, error) {
	var membership entity.Membership
	var jnsMembership string
	faktur3 := entity.Faktur3{NoMsn: no_msn}
	customerMtr := entity.CustomerMtr{NoMsn: no_msn}
	ec.connGorm.Find(&faktur3)
	ec.connGorm.Find(&customerMtr)
	now := time.Now()
	tglExpired := now.AddDate(1,0,0)
	if faktur3.NmCustomer == "" {
		return  membership, errors.New("data faktur tidak ditemukan")
	}
	if customerMtr.NmCustomerFkt != "" {
		ec.connGorm.Where("no_msn = ? and renewal_ke = ?", faktur3.NoMsn, faktur3.StsCetak3).Preload("MstCard").First(&membership)
	}
	if strings.Contains(membership.MstCard.JnsCard, "BASIC") {
		jnsMembership = "01"
	}else if strings.Contains(membership.MstCard.JnsCard, "GOLD") {
		jnsMembership = "02"
	}else if strings.Contains(membership.MstCard.JnsCard, "PLATINUM") {
		jnsMembership = "03"
	}else if strings.Contains(membership.MstCard.JnsCard, "PLATINUM PLUS") {
		jnsMembership = "23"
	}

	if membership.StsBayar == "S" && membership.TypeKartu == "E" {
		var lastNoTT string
		nextNoTT := 0
		err := ec.conn.QueryRow("select VAL_CHAR from mst_runnum WHERE VAL_ID = 'tr_wms_faktur3' and VAL_TAG = 'no_tanda_terima'").Scan(&lastNoTT)
		if err != nil {
			return membership, err
		}
		fmt.Sscanf(lastNoTT[1:], "%d", &nextNoTT)
		nextNoTTF := fmt.Sprintf("T%09d", nextNoTT+1)
		_, err = ec.conn.Exec("UPDATE mst_runnum SET VAL_CHAR = ? WHERE VAL_ID = 'tr_wms_faktur3' and VAL_TAG = 'no_tanda_terima'", nextNoTTF)
		if err != nil {
			return membership, err
		}
		membership.NoTandaTerima = nextNoTTF
		membership.TglCetakTandaTerima = &now
		
		formatKartuSekarang := fmt.Sprintf("11%s %s", jnsMembership, now.Format("0601"))
		var nomorKartuSekarang string
		nomorUrutKartuSekarang :=0
		err = ec.connECardplus.QueryRow("select no_kartu from member where no_kartu like ? order by no_kartu desc", formatKartuSekarang+"%").Scan(&nomorKartuSekarang)
		if err != nil {
			if err == sql.ErrNoRows {
				membership.NoKartu =  formatKartuSekarang + " 0000 0001"
				membership.TglExpired =  &tglExpired
			}else {
				return membership, err
			}
		}
		if nomorKartuSekarang != "" {
			fmt.Sscanf(strings.ReplaceAll(strings.TrimSpace(nomorKartuSekarang[10:]), " ", ""), "%d", &nomorUrutKartuSekarang)
			formatUrut := fmt.Sprintf("%08d", nomorUrutKartuSekarang+1) 
			membership.NoKartu = fmt.Sprintf("%s %s %s", formatKartuSekarang, formatUrut[:4], formatUrut[4:]) 
			membership.TglExpired = &tglExpired
		}
		ec.connGorm.Save(&membership)
		ec.gormECardplus.Save(&entity.ECardPlusMember{NoMsn: membership.NoMSN, NmCustomer: customerMtr.NmCustomerWkm, NoKartu: membership.NoKartu, TglExpired: now.AddDate(1,0,0), UserId: &user_id})
	}
	return membership, nil
}

func (ec *eCardplusRepository) FindCustomer(no_hp string) entity.CustomerMtr {
	var customer entity.CustomerMtr
	ec.connGorm.Select("no_msn, nm_customer_wkm, no_wa").Where("no_wa = ? ", no_hp).First(&customer)
	return customer
}

func (ec *eCardplusRepository) CreateUser(data request.CreateECardplusUserRequest) (entity.UserECardPlus, error) {
	user := entity.UserECardPlus{NoHp: data.NoHp}
	ec.gormECardplus.Where("no_hp", user.NoHp).First(&user)
	if user.ID != "" {
		return user, nil
	}
	user.Name = data.Fullname
	user.Password = data.Password
	result := ec.gormECardplus.Create(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}
