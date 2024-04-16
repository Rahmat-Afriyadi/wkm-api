package service

import (
	"wkm/entity"
	"wkm/repository"
	"wkm/request"
)

type AuthService interface {
	SignInUser(r request.SigninRequest) (entity.User, error)
	RefreshToken(r uint32) (entity.User, error)
	GeneratePassword()
}

type authService struct {
	uR repository.UserRepository
}

func NewAuthService(uR repository.UserRepository) AuthService {
	return &authService{
		uR,
	}
}

func (s *authService) SignInUser(r request.SigninRequest) (entity.User, error) {
	return s.uR.FindByUsername(r.Username), nil
}
func (s *authService) RefreshToken(r uint32) (entity.User, error) {
	return s.uR.FindById(r), nil
}

func (s *authService) GeneratePassword() {
	s.uR.GeneratePassword()
}
