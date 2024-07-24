package service

import (
	"wkm/entity"
	"wkm/repository"
	"wkm/request"
)

type TransaksiService interface {
	MasterData(search string, jenis_asuransi int, limit int, pageParams int) []entity.MasterProduk
	MasterDataCount(search string, jenis_asuransi int) int64
	Detail(id string) entity.MasterProduk
	Update(body entity.MasterProduk) error
	UploadLogo(body entity.MasterProduk) error
	Create(request.TransaksiCreateRequest) (entity.Transaksi, error)
	DeleteManfaat(id string) error
	DeleteSyarat(id string) error
	DeletePaket(id string) error
}

type transaksiService struct {
	trR repository.TransaksiRepository
}

func NewTransaksiService(tR repository.TransaksiRepository) TransaksiService {
	return &transaksiService{
		trR: tR,
	}
}

func (s *transaksiService) MasterData(search string, jenis_asuransi int, limit int, pageParams int) []entity.MasterProduk {
	return s.trR.MasterData(search, jenis_asuransi, limit, pageParams)
}

func (lR *transaksiService) MasterDataCount(search string, jenis_asuransi int) int64 {
	return lR.trR.MasterDataCount(search, jenis_asuransi)
}

func (s *transaksiService) Detail(id string) entity.MasterProduk {
	return s.trR.DetailTransaksi(id)
}

func (s *transaksiService) Update(body entity.MasterProduk) error {
	return s.trR.Update(body)
}

func (s *transaksiService) UploadLogo(body entity.MasterProduk) error {
	return s.trR.UploadLogo(body)
}

func (s *transaksiService) Create(body request.TransaksiCreateRequest) (entity.Transaksi, error) {
	return s.trR.Create(body)
}

func (s *transaksiService) DeleteManfaat(id string) error {
	return s.trR.DeleteManfaat(id)
}
func (s *transaksiService) DeleteSyarat(id string) error {
	return s.trR.DeleteSyarat(id)
}
func (s *transaksiService) DeletePaket(id string) error {
	return s.trR.DeletePaket(id)
}
