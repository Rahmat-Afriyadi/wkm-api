package service

import (
	"wkm/entity"
	"wkm/repository"
)

type MerkService interface {
	MasterData(jenis_kendaraan int) []entity.Merk
}

type merkService struct {
	trR repository.MerkRepository
}

func NewMerkService(tR repository.MerkRepository) MerkService {
	return &merkService{
		trR: tR,
	}
}

func (s *merkService) MasterData(jenis_kendaraan int) []entity.Merk {
	return s.trR.MasterData(jenis_kendaraan)
}
