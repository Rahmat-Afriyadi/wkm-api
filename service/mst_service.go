package service

import (
	"wkm/entity"
	"wkm/repository"
	"wkm/response"
)

type MstService interface {
	ListClientUser() ([]response.ClientUser, error)
	MasterAgama() []entity.MstAgama
	MasterTujuPak() []entity.MstTujuPak
	MasterPendidikan() []entity.MstPendidikan
	MasterKeluarBln() []entity.MstKeluarBln
	MasterAktivitasJual() []entity.MstAktivitasJual
}

type mstService struct {
	uR repository.UserRepository
	mR repository.MstRepository
}

func NewMstService(uR repository.UserRepository, mR repository.MstRepository) MstService {
	return &mstService{
		uR: uR,
		mR: mR,
	}
}

func (ur *mstService) ListClientUser() ([]response.ClientUser, error) {
	user:= ur.uR.All()
	
	return user, nil
}

func (ur *mstService) MasterAgama() []entity.MstAgama {
	data:= ur.mR.MasterAgama()
	return data 
}
func (ur *mstService) MasterTujuPak() []entity.MstTujuPak {
	data:= ur.mR.MasterTujuPak()
	return data 
}
func (ur *mstService) MasterPendidikan() []entity.MstPendidikan {
	data:= ur.mR.MasterPendidikan()
	return data 
}
func (ur *mstService) MasterKeluarBln() []entity.MstKeluarBln {
	data:= ur.mR.MasterKeluarBln()
	return data 
}
func (ur *mstService) MasterAktivitasJual() []entity.MstAktivitasJual {
	data:= ur.mR.MasterAktivitasJual()
	return data 
}

