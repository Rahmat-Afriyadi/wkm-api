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
	UploadLogo(body entity.MasterProduk) error
	Create(body entity.MasterProduk) (entity.MasterProduk, error)
	DeleteManfaat(id string) error
	DeleteSyarat(id string) error
	DeletePaket(id string) error
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

func (s *produkService) UploadLogo(body entity.MasterProduk) error {
	return s.trR.UploadLogo(body)
}

func (s *produkService) Create(body entity.MasterProduk) (entity.MasterProduk, error) {
	return s.trR.Create(body)
}

func (s *produkService) DeleteManfaat(id string) error {
	return s.trR.DeleteManfaat(id)
}
func (s *produkService) DeleteSyarat(id string) error {
	return s.trR.DeleteSyarat(id)
}
func (s *produkService) DeletePaket(id string) error {
	return s.trR.DeletePaket(id)
}
