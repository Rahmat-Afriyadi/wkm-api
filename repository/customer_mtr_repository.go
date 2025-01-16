package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
	"wkm/entity"
	"wkm/utils/query"

	"gorm.io/gorm"
)

type CustomerMtrRepository interface {
	ListAmbilData() []entity.Faktur3
	AmbilData(no_msn string, kd_user string) error
	Show(no_msn string)entity.CustomerMtr
	Update(customer entity.CustomerMtr) (entity.CustomerMtr, error)
}

type customerMtrRepository struct {
	conn     *sql.DB
	connGorm *gorm.DB
}

func NewCustomerMtrRepository(conn *sql.DB, connGorm *gorm.DB) CustomerMtrRepository {
	return &customerMtrRepository{
		conn:     conn,
		connGorm: connGorm,
	}
}

func (r *customerMtrRepository) ListAmbilData() []entity.Faktur3 {
	data := []entity.Faktur3{}
	r.connGorm.Select("no_msn").Where("sts_renewal is null").Find(&data)
	return data
}

func (r *customerMtrRepository) AmbilData(no_msn string, kd_user string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	queryAmbilData := query.NewQueryAmbilData()
	defer cancel()
	var kdUser sql.NullString
	err := r.conn.QueryRowContext(ctx, "select kd_user from tr_wms_faktur3 where no_msn = ? and sts_renewal is null", no_msn).Scan(&kdUser)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("data tidak ditemukan")
		} else {
			return err
		}
	}
	if kdUser.String != "" {
		return fmt.Errorf("data tersebut telah di ambil oleh kd_user %s", kdUser.String)
	}
	now := time.Now()
	r.conn.QueryContext(ctx, "update tr_wms_faktur3 set kd_user = ?, tgl_verifikasi = ?, sts_renewal='P' where no_msn = ?", kd_user,now.Format("2006-01-02"),no_msn)
	r.conn.QueryContext(ctx, queryAmbilData, no_msn)
	return nil
}


func (r *customerMtrRepository) Show(no_msn string) entity.CustomerMtr {
	data := entity.CustomerMtr{NoMsn: no_msn}
	r.connGorm.Find(&data)
	return data
}

func (r *customerMtrRepository) Update(customer entity.CustomerMtr) (entity.CustomerMtr, error) {
	data := entity.CustomerMtr{NoMsn: customer.NoMsn}
	r.connGorm.Select("no_msn, nm_customer_fkt").Find(&data)
	if data.NmCustomerFkt == "" {
		return entity.CustomerMtr{}, errors.New("data gk ada maaf")
	}

	r.connGorm.Save(&customer)
	return customer, nil
}