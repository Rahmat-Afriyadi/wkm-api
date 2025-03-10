package repository

import (
	"database/sql"
	"errors"
	"math/rand"
	"time"
	"wkm/entity"
	"wkm/request"

	"golang.org/x/crypto/bcrypt"
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
	FindCustomer(no_msn string) entity.CustomerMtr
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

	noKartu, err := entity.GenerateEcardNumber(membership.MstCard.JnsCard)
	if err != nil {
		return entity.Membership{}, err
	}
	
	if membership.StsBayar == "S" && membership.TypeKartu == "E" {
		nextNoTTF, err := entity.GenerateNoTT()
		if err != nil {
			return entity.Membership{}, err
		}
		membership.NoTandaTerima = nextNoTTF
		membership.TglCetakTandaTerima = &now
		membership.NoKartu = noKartu
		membership.TglExpired = &tglExpired
		ec.connGorm.Save(&membership)
		ec.gormECardplus.Save(&entity.ECardPlusMember{NoMsn: membership.NoMSN, NmCustomer: customerMtr.NmCustomerWkm, NoKartu: membership.NoKartu, TglExpired: now.AddDate(1,0,0), UserId: &user_id})
	}
	return membership, nil
}

func (ec *eCardplusRepository) FindCustomer(no_msn string) entity.CustomerMtr {
	var customer entity.CustomerMtr
	ec.connGorm.Select("no_msn, nm_customer_wkm, no_wa").Where("no_msn = ? ", no_msn).First(&customer)
	return customer
}

func (ec *eCardplusRepository) CreateUser(data request.CreateECardplusUserRequest) (entity.UserECardPlus, error) {
	user := entity.UserECardPlus{NoHp: data.NoHp}
	ec.gormECardplus.Where("no_hp", user.NoHp).First(&user)
	if user.ID != "" {
		return user, nil
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 8)
	user.Name = data.Fullname
	user.Password = string(password)
	result := ec.gormECardplus.Save(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}
