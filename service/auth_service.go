package service

import (
	"wkm/entity"
	"wkm/repository"
	"wkm/request"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	SignInUser(r request.SigninRequest) (entity.User, error)
	SignInUserAsuransi(r request.SigninRequest) (entity.User, error)
	RefreshToken(r uint32) (entity.User, error)
	RefreshTokenAsuransi(r uint32) (entity.User, error)
	ResetPassword(data request.ResetPassword) request.Response
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

func (s *authService) SignInUserAsuransi(r request.SigninRequest) (entity.User, error) {
	return s.uR.FindByUsername(r.Username), nil
}

func (s *authService) RefreshToken(r uint32) (entity.User, error) {
	return s.uR.FindById(r), nil
}

func (s *authService) RefreshTokenAsuransi(r uint32) (entity.User, error) {
	return s.uR.FindById(r), nil
}

func (s *authService) ResetPassword(data request.ResetPassword) request.Response {
	user := s.uR.FindById(data.IdUser)
	if user.ID == 0 {
		return request.Response{Status: 400, Message: "User tidak ditemukan"}
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.PasswordLama))
	if err != nil {
		return request.Response{Status: 400, Message: "Password salah"}
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 8)
	data.Password = string(password)
	s.uR.ResetPassword(data)
	return request.Response{Status: 201, Message: "Data berhasil diupdate"}
}
