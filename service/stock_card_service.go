package service

import (
	"wkm/entity"
	"wkm/repository"
	"wkm/request"
)

type StockCardService interface {
	MasterData(search string, limit int, pageParams int) []entity.StockCard
	MasterDataCount(search string) int64
	Detail(no_kartu string) entity.StockCard
	Update(body request.StockCardRequest) error
	UploadDokumen(body entity.StockCard) error
	Create(request.StockCardRequest) (entity.StockCard, error)
}

type stockCardService struct {
	trR repository.StockCardRepository
}

func NewStockCardService(tR repository.StockCardRepository) StockCardService {
	return &stockCardService{
		trR: tR,
	}
}

func (s *stockCardService) MasterData(search string, limit int, pageParams int) []entity.StockCard {
	return s.trR.MasterData(search, limit, pageParams)
}

func (lR *stockCardService) MasterDataCount(search string) int64 {
	return lR.trR.MasterDataCount(search)
}

func (s *stockCardService) Detail(no_kartu string) entity.StockCard {
	return s.trR.DetailStockCard(no_kartu)
}

func (s *stockCardService) Update(body request.StockCardRequest) error {
	return s.trR.Update(body)
}

func (s *stockCardService) UploadDokumen(body entity.StockCard) error {
	return s.trR.UploadDokumen(body)
}

func (s *stockCardService) Create(body request.StockCardRequest) (entity.StockCard, error) {
	return s.trR.Create(body)
}
