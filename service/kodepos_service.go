package service

import (
	"wkm/entity"
	"wkm/repository"
)

type KodeposService interface {
	MasterData(search string) []entity.MasterKodepos
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
