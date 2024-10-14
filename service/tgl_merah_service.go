package service

import (
	"wkm/entity"
	"wkm/repository"
	"wkm/request"
)

type TglMerahService interface {
	MasterData(search string, limit int, pageParams int) []entity.TglMerah
	MasterDataCount(search string) int64
	Detail(id uint64) entity.TglMerah
	Update(body request.TglMerahRequest) error
	UploadDokumen(body entity.TglMerah) error
	Create(request.TglMerahRequest) (entity.TglMerah, error)
	CreateFromFile(datas []entity.TglMerah) error
	Delete(id uint64) error
}

type tglMerahService struct {
	trR repository.TglMerahRepository
}

func NewTglMerahService(tR repository.TglMerahRepository) TglMerahService {
	return &tglMerahService{
		trR: tR,
	}
}

func (s *tglMerahService) MasterData(search string, limit int, pageParams int) []entity.TglMerah {
	return s.trR.MasterData(search, limit, pageParams)
}

func (lR *tglMerahService) MasterDataCount(search string) int64 {
	return lR.trR.MasterDataCount(search)
}

func (s *tglMerahService) Detail(id uint64) entity.TglMerah {
	return s.trR.DetailTglMerah(id)
}

func (s *tglMerahService) Update(body request.TglMerahRequest) error {
	return s.trR.Update(body)
}
func (s *tglMerahService) Delete(id uint64) error {
	return s.trR.Delete(id)
}

func (s *tglMerahService) UploadDokumen(body entity.TglMerah) error {
	return s.trR.UploadDokumen(body)
}

func (s *tglMerahService) Create(body request.TglMerahRequest) (entity.TglMerah, error) {
	return s.trR.Create(body)
}

func (s *tglMerahService) CreateFromFile(datas []entity.TglMerah) error {
	return s.trR.BulkCreate(datas)
}
