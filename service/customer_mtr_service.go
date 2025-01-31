package service

import (
	"wkm/entity"
	"wkm/repository"
)

type CustomerMtrService interface {
	ListAmbilData() []entity.Faktur3
	AmbilData(no_msn string, kd_user string) error
	Show(no_msn string) entity.CustomerMtr
	UpdateOkeMembership(customer entity.CustomerMtr) (entity.CustomerMtr,error)
}

type customerMtrService struct {
	cR repository.CustomerMtrRepository
}

func NewCustomerMtrService(cR repository.CustomerMtrRepository) CustomerMtrService {
	return &customerMtrService{
		cR:     cR,
	}
}

func (cR *customerMtrService) ListAmbilData() []entity.Faktur3 {
	return cR.cR.ListAmbilData()
}

func (cR *customerMtrService) AmbilData(no_msn string, kd_user string) error {
	return cR.cR.AmbilData(no_msn, kd_user)
}

func (cR *customerMtrService) Show(no_msn string) entity.CustomerMtr {
	return cR.cR.Show(no_msn)
}
func (cR *customerMtrService) UpdateOkeMembership(customer entity.CustomerMtr) (entity.CustomerMtr,error) {
	return cR.cR.UpdateOkeMembership(customer)
}

