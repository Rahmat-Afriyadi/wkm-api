package service

import (
	"wkm/entity"
	"wkm/repository"
)

type KodeposService interface {
	MasterData(search string) []entity.MasterKodepos
	MasterData1(search string) []entity.MasterKodepos1
}

type kodeposService struct {
	trR repository.KodeposRepository
}

func NewKodeposService(tR repository.KodeposRepository) KodeposService {
	return &kodeposService{
		trR: tR,
	}
}

func (s *kodeposService) MasterData(search string) []entity.MasterKodepos {
	return s.trR.MasterData(search)
}

func (s *kodeposService) MasterData1(search string) []entity.MasterKodepos1 {
	return s.trR.MasterData1(search)
}
