package service

import (
	"wkm/response"
	"wkm/repository"
)

type MstService interface {
	ListClientUser() ([]response.ClientUser, error)
}

type mstService struct {
	uR repository.UserRepository
}

func NewMstService(uR repository.UserRepository) MstService {
	return &mstService{
		uR: uR,
	}
}

func (ur *mstService) ListClientUser() ([]response.ClientUser, error) {
	user:= ur.uR.All()
	
	return user, nil
}

