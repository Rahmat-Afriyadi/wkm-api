package service

import (
	"wkm/entity"
	"wkm/repository"
)

type ProdukService interface {
	MasterData(search string, jenis_asuransi int, limit int, pageParams int) []entity.MasterProduk
	MasterDataCount(search string, jenis_asuransi int) int64
	Detail(id string) entity.MasterProduk
	Update(body entity.MasterProduk) error
	Create(body entity.MasterProduk) error
}

type produkService struct {
	trR repository.ProdukRepository
}

func NewProdukService(tR repository.ProdukRepository) ProdukService {
	return &produkService{
		trR: tR,
	}
}

func (s *produkService) MasterData(search string, jenis_asuransi int, limit int, pageParams int) []entity.MasterProduk {
	return s.trR.MasterData(search, jenis_asuransi, limit, pageParams)
}

func (lR *produkService) MasterDataCount(search string, jenis_asuransi int) int64 {
	return lR.trR.MasterDataCount(search, jenis_asuransi)
}

func (s *produkService) Detail(id string) entity.MasterProduk {
	return s.trR.DetailProduk(id)
}

func (s *produkService) Update(body entity.MasterProduk) error {
	return s.trR.Update(body)
}

func (s *produkService) Create(body entity.MasterProduk) error {
	return s.trR.Create(body)
}
