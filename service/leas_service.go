package service

import (
	"wkm/entity"
	"wkm/repository"
)

type LeasService interface {
	MasterData() []entity.MasterLeas
}

type leasService struct {
	trR repository.LeasRepository
}

func NewLeasService(tR repository.LeasRepository) LeasService {
	return &leasService{
		trR: tR,
	}
}

func (s *leasService) MasterData() []entity.MasterLeas {
	return s.trR.MasterData()
}
