package service

import (
	"time"
	"wkm/entity"
	"wkm/repository"
	"wkm/request"
)

type TglMerahService interface {
	MasterData(search string, tgl1 string, tgl2 string, limit int, pageParams int) []entity.TglMerah
	MasterDataCount(search string, tgl1 string, tgl2 string) int64
	Detail(id uint64) entity.TglMerah
	Update(body request.TglMerahRequest) error
	UploadDokumen(body entity.TglMerah) error
	Create(request.TglMerahRequest) (entity.TglMerah, error)
	CreateFromFile(datas []entity.TglMerah) error
	Delete(id uint64) error
	MinTglBayar() time.Time
}

type tglMerahService struct {
	trR repository.TglMerahRepository
}

func NewTglMerahService(tR repository.TglMerahRepository) TglMerahService {
	return &tglMerahService{
		trR: tR,
	}
}

func (s *tglMerahService) MasterData(search string, tgl1 string, tgl2 string, limit int, pageParams int) []entity.TglMerah {
	return s.trR.MasterData(search, tgl1, tgl2, limit, pageParams)
}

func (lR *tglMerahService) MasterDataCount(search string, tgl1 string, tgl2 string) int64 {
	return lR.trR.MasterDataCount(search, tgl1, tgl2)
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

func (s *tglMerahService) MinTglBayar() time.Time {
	return s.trR.GetMinTglBayar()
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
