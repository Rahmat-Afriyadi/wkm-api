package service

import (
	"wkm/entity"
	"wkm/repository"
	"wkm/request"
)

type OtrService interface {
	DetailOtrNa(search string, tahun uint16) entity.Otr
	OtrNaList() []entity.Otr
	OtrMstProduk(search string) []entity.MstMtr
	OtrMstNa(search string) []entity.OtrNa
	CreateOtr(data request.CreateOtr)
}

type otrService struct {
	trR repository.OtrRepository
}

func NewOtrService(tR repository.OtrRepository) OtrService {
	return &otrService{
		trR: tR,
	}
}

func (s *otrService) DetailOtrNa(search string, tahun uint16) entity.Otr {
	return s.trR.DetailOtrNa(search, tahun)
}

func (s *otrService) OtrNaList() []entity.Otr {
	return s.trR.OtrNaList()
}

func (s *otrService) OtrMstProduk(search string) []entity.MstMtr {
	return s.trR.OtrMstProduk(search)
}
func (s *otrService) OtrMstNa(search string) []entity.OtrNa {
	return s.trR.OtrMstNa(search)
}
func (s *otrService) CreateOtr(data request.CreateOtr) {
	s.trR.CreateOtr(data)
}
