package service

import (
	"wkm/entity"
	"wkm/repository"
	"wkm/request"
)

type CustomerMtrService interface {
	MasterData(search string, sts string, jns string, username string, limit int, pageParams int) []entity.CustomerMtr
	MasterDataCount(search string, sts string, jns string, username string) int64
	ListAmbilData() []entity.Faktur3
	AmbilData(no_msn string, kd_user string) error
	Show(no_msn string) entity.CustomerMtr
	UpdateOkeMembership(customer request.CustomerMtr) (entity.CustomerMtr,error)
}

type customerMtrService struct {
	cR repository.CustomerMtrRepository
}

func NewCustomerMtrService(cR repository.CustomerMtrRepository) CustomerMtrService {
	return &customerMtrService{
		cR:     cR,
	}
}

func (cS *customerMtrService) MasterData(search string, sts string, jns string, username string, limit int, pageParams int) []entity.CustomerMtr {
	return 	cS.cR.MasterData(search, sts, jns, username, limit, pageParams)
}
func (cS *customerMtrService) MasterDataCount(search string, sts string, jns string, username string) int64 {
	return 	cS.cR.MasterDataCount(search, sts, jns, username)
}

func (cS *customerMtrService) ListAmbilData() []entity.Faktur3 {
	return cS.cR.ListAmbilData()
}

func (cS *customerMtrService) AmbilData(no_msn string, kd_user string) error {
	return cS.cR.AmbilData(no_msn, kd_user)
}

func (cS *customerMtrService) Show(no_msn string) entity.CustomerMtr {
	return cS.cR.Show(no_msn)
}
func (cS *customerMtrService) UpdateOkeMembership(customer request.CustomerMtr) (entity.CustomerMtr,error) {
	return cS.cR.UpdateOkeMembership(customer)
}

