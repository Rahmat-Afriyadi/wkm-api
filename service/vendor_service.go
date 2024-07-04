package service

import (
	"wkm/entity"
	"wkm/repository"
)

type VendorService interface {
	MasterData(search string, limit int, pageParams int) []entity.MasterVendor
	MasterDataCount(search string) int64
	Detail(id string) entity.MasterVendor
	Update(body entity.MasterVendor) error
	Create(body entity.MasterVendor) error
}

type vendorService struct {
	trR repository.VendorRepository
}

func NewVendorService(tR repository.VendorRepository) VendorService {
	return &vendorService{
		trR: tR,
	}
}

func (s *vendorService) MasterData(search string, limit int, pageParams int) []entity.MasterVendor {
	return s.trR.MasterData(search, limit, pageParams)
}

func (lR *vendorService) MasterDataCount(search string) int64 {
	return lR.trR.MasterDataCount(search)
}

func (s *vendorService) Detail(id string) entity.MasterVendor {
	return s.trR.DetailVendor(id)
}

func (s *vendorService) Update(body entity.MasterVendor) error {
	return s.trR.Update(body)
}

func (s *vendorService) Create(body entity.MasterVendor) error {
	return s.trR.Create(body)
}
