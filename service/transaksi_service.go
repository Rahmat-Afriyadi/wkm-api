package service

import (
	"wkm/entity"
	"wkm/repository"
	"wkm/request"
)

type TransaksiService interface {
	MasterData(search string, limit int, pageParams int) []entity.Transaksi
	MasterDataCount(search string) int64
	Detail(id string) entity.Transaksi
	Update(body request.TransaksiRequest) error
	UploadDokumen(body entity.Transaksi) error
	Create(request.TransaksiRequest) (entity.Transaksi, error)
}

type transaksiService struct {
	trR repository.TransaksiRepository
}

func NewTransaksiService(tR repository.TransaksiRepository) TransaksiService {
	return &transaksiService{
		trR: tR,
	}
}

func (s *transaksiService) MasterData(search string, limit int, pageParams int) []entity.Transaksi {
	return s.trR.MasterData(search, limit, pageParams)
}

func (lR *transaksiService) MasterDataCount(search string) int64 {
	return lR.trR.MasterDataCount(search)
}

func (s *transaksiService) Detail(id string) entity.Transaksi {
	return s.trR.DetailTransaksi(id)
}

func (s *transaksiService) Update(body request.TransaksiRequest) error {
	return s.trR.Update(body)
}

func (s *transaksiService) UploadDokumen(body entity.Transaksi) error {
	return s.trR.UploadDokumen(body)
}

func (s *transaksiService) Create(body request.TransaksiRequest) (entity.Transaksi, error) {
	return s.trR.Create(body)
}
