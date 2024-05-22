package service

import (
	"wkm/entity"
	"wkm/repository"
)

type AsuransiService interface {
	MasterData(dataSource string) []entity.MasterAsuransi
	MasterDataPending(search string, dataSource string) []entity.MasterAsuransi
	MasterDataOke(dataSource string) []entity.MasterAsuransi
	FindAsuransiByNoMsn(no_msn string) entity.MasterAsuransiReal
	UpdateAsuransi(data entity.MasterAsuransiReal) entity.MasterAsuransiReal
	UpdateAsuransiBerminat(no_msn string)
	UpdateAsuransiBatalBayar(no_msn string)
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

func (s *asuransiService) MasterData(dataSource string) []entity.MasterAsuransi {
	return s.trR.MasterData(dataSource)
}

func (s *asuransiService) MasterDataPending(search string, dataSource string) []entity.MasterAsuransi {
	return s.trR.MasterDataPending(search, dataSource)
}

func (s *asuransiService) MasterDataOke(dataSource string) []entity.MasterAsuransi {
	return s.trR.MasterDataOke(dataSource)
}

func (s *asuransiService) FindAsuransiByNoMsn(no_msn string) entity.MasterAsuransiReal {
	return s.trR.FindAsuransiByNoMsn(no_msn)
}

func (s *asuransiService) UpdateAsuransi(data entity.MasterAsuransiReal) entity.MasterAsuransiReal {
	return s.trR.UpdateAsuransi(data)
}

func (s *asuransiService) UpdateAsuransiBerminat(no_msn string) {
	s.trR.UpdateAsuransiBerminat(no_msn)
}

func (s *asuransiService) UpdateAsuransiBatalBayar(no_msn string) {
	s.trR.UpdateAsuransiBatalBayar(no_msn)
}

func (s *asuransiService) UpdateAmbilAsuransi(no_msn string, kd_user string) {
	s.trR.UpdateAmbilAsuransi(no_msn, kd_user)
}
