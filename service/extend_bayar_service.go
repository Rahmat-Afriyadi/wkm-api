package service

import (
	"wkm/entity"
	"wkm/repository"
	"wkm/request"
)

type ExtendBayarService interface {
	MasterData(search string, tgl1 string, tgl2 string, limit int, pageParams int) []entity.ExtendBayar
	MasterDataCount(search string, tgl1 string, tgl2 string) int64
	MasterDataLf(search string, tgl1 string, tgl2 string, limit int, pageParams int) []entity.ExtendBayar
	MasterDataLfCount(search string, tgl1 string, tgl2 string) int64
	Detail(id string) entity.ExtendBayar
	UpdateFa(body request.ExtendBayarRequest) error
	UpdateLf(body request.ExtendBayarRequest) error
	Create(request.ExtendBayarRequest) (entity.ExtendBayar, error)
	Delete(id string, kdUserFa string) error
	UpdateApprovalLf(body request.ExtendBayarApprovalRequest) error
}

type extendBayarService struct {
	trR repository.ExtendBayarRepository
}

func NewExtendBayarService(tR repository.ExtendBayarRepository) ExtendBayarService {
	return &extendBayarService{
		trR: tR,
	}
}

func (s *extendBayarService) MasterData(search string, tgl1 string, tgl2 string, limit int, pageParams int) []entity.ExtendBayar {
	return s.trR.MasterData(search, tgl1, tgl2, limit, pageParams)
}

func (lR *extendBayarService) MasterDataCount(search string, tgl1 string, tgl2 string) int64 {
	return lR.trR.MasterDataCount(search, tgl1, tgl2)
}
func (s *extendBayarService) MasterDataLf(search string, tgl1 string, tgl2 string, limit int, pageParams int) []entity.ExtendBayar {
	return s.trR.MasterDataLf(search, tgl1, tgl2, limit, pageParams)
}

func (lR *extendBayarService) MasterDataLfCount(search string, tgl1 string, tgl2 string) int64 {
	return lR.trR.MasterDataLfCount(search, tgl1, tgl2)
}

func (s *extendBayarService) Detail(id string) entity.ExtendBayar {
	return s.trR.DetailExtendBayar(id)
}

func (s *extendBayarService) UpdateFa(body request.ExtendBayarRequest) error {
	return s.trR.UpdateFa(body)
}
func (s *extendBayarService) UpdateLf(body request.ExtendBayarRequest) error {
	return s.trR.UpdateLf(body)
}
func (s *extendBayarService) Delete(id string, kdUserFa string) error {
	return s.trR.Delete(id, kdUserFa)
}

func (s *extendBayarService) Create(body request.ExtendBayarRequest) (entity.ExtendBayar, error) {
	return s.trR.Create(body)
}

func (s *extendBayarService) UpdateApprovalLf(body request.ExtendBayarApprovalRequest) error {
	return s.trR.UpdateApprovalLf(body)
}
