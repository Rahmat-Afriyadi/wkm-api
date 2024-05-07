package service

import (
	"wkm/entity"
	"wkm/repository"
)

type DlrService interface {
	MasterData(search string) []entity.MasterDlr
}

type dlrService struct {
	trR repository.DlrRepository
}

func NewDlrService(tR repository.DlrRepository) DlrService {
	return &dlrService{
		trR: tR,
	}
}

func (s *dlrService) MasterData(search string) []entity.MasterDlr {
	return s.trR.MasterData(search)
}
