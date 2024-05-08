package service

import (
	"wkm/entity"
	"wkm/repository"
)

type ProdukService interface {
	MasterData(search string) []entity.MasterProduk
}

type produkService struct {
	trR repository.ProdukRepository
}

func NewProdukService(tR repository.ProdukRepository) ProdukService {
	return &produkService{
		trR: tR,
	}
}

func (s *produkService) MasterData(search string) []entity.MasterProduk {
	return s.trR.MasterData(search)
}
