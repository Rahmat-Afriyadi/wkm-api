package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"wkm/repository"
	"wkm/request"
)

type ECardplusService interface {
	WelcomeMessage(body request.WelcomeMessage) (map[string]interface{}, error)
	InputBayarEMembership(data request.InputBayarRequest) error

}

type eCardplusService struct {
	cR repository.ECardplusRepository
}

func NewECardplusService(cR repository.ECardplusRepository) ECardplusService {
	return &eCardplusService{
		cR:     cR,
	}
}



func (eC *eCardplusService) InputBayarEMembership(data request.InputBayarRequest) error  {
	customer := eC.cR.GetCustomerByNoMsn(data.NoMsn)
	tempPass := repository.RandomString(8)
	user, err := eC.cR.CreateUser(request.CreateECardplusUserRequest{NoHp: customer.NoWa, Fullname: customer.NmCustomerWkm, Password: tempPass})
	if err != nil {
		return err
	}
	token := eC.cR.CreateToken(customer.NoWa)
	_, err = eC.cR.GenerateEMembership(customer.NoMsn, user.ID)

	if err != nil {
		return err
	}
	_, err = eC.WelcomeMessage(request.WelcomeMessage{NoHp: customer.NoWa,NoMsn: data.NoMsn, Fullname: customer.NmCustomerWkm, Password: tempPass, Link: "https://www.e-cardplus.co.id/activation-by-token/"+token.Token})
	if err != nil {
		return err
	}
	return nil
}


func (eC *eCardplusService) WelcomeMessage(body request.WelcomeMessage) (map[string]interface{}, error) {
	customer := eC.cR.FindCustomer(body.NoMsn)
	var client = &http.Client{}
	var data map[string]any
	var param = url.Values{}
	param.Set("target", body.NoHp)
	param.Set("message", fmt.Sprintf("*Selamat bergabung %s*\n\nTerima kasih telah menjadi pelanggan setia Honda VIP Card.\nBerikut informasi terkait akun E-Membership Anda.\n\n*Username:* %s\n*Password:* %s\n\n_Gunakan link berikut untuk aktivasi akun_\n%s", customer.NmCustomerWkm, body.NoHp, body.Password, body.Link))
	param.Set("schedule", "0")
	param.Set("delay", "2")
	param.Set("countryCode", "62")
	var payload = bytes.NewBufferString(param.Encode())
	request, err := http.NewRequest("POST", "https://api.fonnte.com/send", payload)
	if err != nil {
		return map[string]any{}, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", "k!ph_r+apphR8kJY@+gS")
	response, err := client.Do(request)
	if err != nil {
		return map[string]any{}, err
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return map[string]any{}, err
	}
	return data, nil
}