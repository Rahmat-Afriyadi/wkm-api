package service

import (
	"wkm/entity"
	"wkm/repository"
)

type MstMtrService interface {
	CreateMstMtr(data entity.MstMtr) error
	MasterData(search string, limit int, pageParams int) []entity.MstMtr
	MasterDataCount(search string) int64
	DetailMstMtr(id string) entity.MstMtr
	Update(body entity.MstMtr) error
}

type mstMtrService struct {
	trR repository.MstMtrRepository
}

func NewMstMtrService(tR repository.MstMtrRepository) MstMtrService {
	return &mstMtrService{
		trR: tR,
	}
}

func (s *mstMtrService) Update(body entity.MstMtr) error {
	return s.trR.Update(body)
}
func (s *mstMtrService) DetailMstMtr(id string) entity.MstMtr {
	return s.trR.DetailMstMtr(id)
}
func (s *mstMtrService) MasterData(search string, limit int, pageParams int) []entity.MstMtr {
	return s.trR.MasterData(search, limit, pageParams)
}
func (s *mstMtrService) MasterDataCount(search string) int64 {
	return s.trR.MasterDataCount(search)
}

func (s *mstMtrService) CreateMstMtr(data entity.MstMtr) error {
	return s.trR.CreateMstMtr(data)
}
