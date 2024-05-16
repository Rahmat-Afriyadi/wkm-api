package service

import (
	"wkm/entity"
	"wkm/repository"
)

type AsuransiService interface {
	MasterData() []entity.MasterAsuransi
	MasterDataPending(search string) []entity.MasterAsuransi
	MasterDataOke() []entity.MasterAsuransi
	FindAsuransiByNoMsn(no_msn string) entity.MasterAsuransi
	UpdateAsuransi(data entity.MasterAsuransi) entity.MasterAsuransi
	UpdateAmbilAsuransi(no_msn string, kd_user string)
}

type asuransiService struct {
	trR repository.AsuransiRepository
}

func NewAsuransiService(tR repository.AsuransiRepository) AsuransiService {
	return &asuransiService{
		trR: tR,
	}
}

func (s *asuransiService) MasterData() []entity.MasterAsuransi {
	return s.trR.MasterData()
}

func (s *asuransiService) MasterDataPending(search string) []entity.MasterAsuransi {
	return s.trR.MasterDataPending(search)
}

func (s *asuransiService) MasterDataOke() []entity.MasterAsuransi {
	return s.trR.MasterDataOke()
}

func (s *asuransiService) FindAsuransiByNoMsn(no_msn string) entity.MasterAsuransi {
	return s.trR.FindAsuransiByNoMsn(no_msn)
}

func (s *asuransiService) UpdateAsuransi(data entity.MasterAsuransi) entity.MasterAsuransi {
	return s.trR.UpdateAsuransi(data)
}

func (s *asuransiService) UpdateAmbilAsuransi(no_msn string, kd_user string) {
	s.trR.UpdateAmbilAsuransi(no_msn, kd_user)
}
