package service

import (
	"wkm/entity"
	"wkm/repository"

	// "wkm/request"
	"wkm/response"
)

type MstService interface {
	ListClientUser() ([]response.ClientUser, error)
	MasterAgama() []entity.MstAgama
	MasterTujuPak() []entity.MstTujuPak
	MasterPendidikan() []entity.MstPendidikan
	MasterKeluarBln() []entity.MstKeluarBln
	MasterAktivitasJual() []entity.MstAktivitasJual
	CreateScript(data entity.MstScript, username string) error
	UpdateScript(id string, data entity.MstScript, username string) error
	MasterScript() []entity.MstScript
	ListAllScript() []entity.MstScript
	ViewScript(id string) (entity.MstScript, error)
}

type mstService struct {
	uR repository.UserRepository
	mR repository.MstRepository
}

func NewMstService(uR repository.UserRepository, mR repository.MstRepository) MstService {
	return &mstService{
		uR: uR,
		mR: mR,
	}
}

func (ur *mstService) ListClientUser() ([]response.ClientUser, error) {
	user := ur.uR.All()

	return user, nil
}

func (ur *mstService) MasterAgama() []entity.MstAgama {
	data := ur.mR.MasterAgama()
	return data
}
func (ur *mstService) MasterTujuPak() []entity.MstTujuPak {
	data := ur.mR.MasterTujuPak()
	return data
}
func (ur *mstService) MasterPendidikan() []entity.MstPendidikan {
	data := ur.mR.MasterPendidikan()
	return data
}
func (ur *mstService) MasterKeluarBln() []entity.MstKeluarBln {
	data := ur.mR.MasterKeluarBln()
	return data
}
func (ur *mstService) MasterAktivitasJual() []entity.MstAktivitasJual {
	data := ur.mR.MasterAktivitasJual()
	return data
}
func (ur *mstService) MasterScript() []entity.MstScript {
	data := ur.mR.MasterScript()
	return data
}
func (ur *mstService) ViewScript(id string) (entity.MstScript, error) {
	data, err := ur.mR.ViewScript(id)
	if err != nil {
		return entity.MstScript{}, err // Mengembalikan error jika terjadi kesalahan
	}

	res := entity.MstScript{
		Id: data.Id,
		Title:    data.Title,
		Script:   data.Script,
		IsActive: data.IsActive,
	}

	return res, nil
}

func (ur *mstService) ListAllScript() []entity.MstScript {

	return ur.mR.ListAllScript()
}
func (ur *mstService) CreateScript(data entity.MstScript, username string) error {
	err := ur.mR.CreateScript(data, username)
	if err != nil {
		return err
	}

	return nil
}

func (ur *mstService) UpdateScript(id string, data entity.MstScript, username string) error {
	err := ur.mR.UpdateScript(id, data, username)
	if err != nil {
		return err
	}

	return nil
}
