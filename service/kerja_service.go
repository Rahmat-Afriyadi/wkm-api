package service

import (
	"wkm/entity"
	"wkm/repository"
)

type KerjaService interface {
	MasterData() []entity.MasterKerja
}

type kerjaService struct {
	trR repository.KerjaRepository
}

func NewKerjaService(tR repository.KerjaRepository) KerjaService {
	return &kerjaService{
		trR: tR,
	}
}

func (s *kerjaService) MasterData() []entity.MasterKerja {
	return s.trR.MasterData()
}
