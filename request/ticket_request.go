package request

import "wkm/entity"

type TicketRequest struct {
	KdUser      string        `json:"kd_user"`
	Problem        string        `json:"problem"`
	Status      int           `json:"status"`
	JenisTicket string        `json:"jenis_ticket"`
	Solution    string        `json:"solution"`
	Clients     []entity.User `json:"kd_user_clients"`
	KdUserIt	string 		  `json:"kd_user_it"`
	Month 		int 			`json:"month"`
	Year		int 			`json:"year"`
}
