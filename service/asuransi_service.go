package service

import (
	"wkm/entity"
	"wkm/repository"
)

type AsuransiService interface {
	MasterData() []entity.MasterAsuransi
	FindAsuransiByNoMsn(no_msn string) entity.MasterAsuransi
	UpdateAsuransi(data entity.MasterAsuransi) entity.MasterAsuransi
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

func (s *asuransiService) FindAsuransiByNoMsn(no_msn string) entity.MasterAsuransi {
	return s.trR.FindAsuransiByNoMsn(no_msn)
}

func (s *asuransiService) UpdateAsuransi(data entity.MasterAsuransi) entity.MasterAsuransi {
	return s.trR.UpdateAsuransi(data)
}
