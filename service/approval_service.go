package service

import (
	"wkm/entity"
	"wkm/repository"
)

type ApprovalService interface {
	Update(data entity.DetailApproval)
	MokitaToken() entity.MstToken
	MokitaUpdateToken(token string)
}

type approvalService struct {
	trR repository.ApprovalRepository
}

func NewApprovalService(tR repository.ApprovalRepository) ApprovalService {
	return &approvalService{
		trR: tR,
	}
}

func (s *approvalService) Update(data entity.DetailApproval) {
	s.trR.Update(data)
}
func (s *approvalService) MokitaToken() entity.MstToken {
	return s.trR.MokitaToken()
}
func (s *approvalService) MokitaUpdateToken(token string) {
	s.trR.MokitaUpdateToken(token)
}
