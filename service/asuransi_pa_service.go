package service

import (
	"wkm/entity"
	"wkm/repository"
)

type AsuransiPAService interface {
	CreateAsuransiPA(data entity.AsuransiPA) error
	UpdateAsuransiPA(id string, data entity.AsuransiPA) error
}

type asuransiPAService struct {
	aSR repository.AsuransiPARepository
}

func NewAsuransiPAService(aSP repository.AsuransiPARepository) AsuransiPAService {
	return &asuransiPAService{
		aSR: aSP,
	}
}

func (aS *asuransiPAService) CreateAsuransiPA(data entity.AsuransiPA) error {
	// Call repository to create AsuransiPA
	err := aS.aSR.CreateAsuransiPA(data)
	if err != nil {
		// Log the error if needed (optional)
		// log.Printf("Error creating AsuransiPA: %v", err)
		return err // Return the error to the caller
	}

	// Return nil if the creation was successful
	return nil
}

func (aS *asuransiPAService) UpdateAsuransiPA(id string, data entity.AsuransiPA) error {
	// Call repository to create AsuransiPA
	err := aS.aSR.UpdateAsuransiPA(id, data)
	if err != nil {
		return err
	}
	return nil
}