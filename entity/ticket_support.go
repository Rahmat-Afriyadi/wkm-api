package entity

import (
	"time"
	"wkm/response"
)

type TicketSupport struct {
	NoTicket    string    `json:"no_ticket"`
	Kd_user     string    `json:"kd_user"`
	Case        string    `json:"case"`
	Status      int    	  `json:"status"`
	KdUserIt    *string    `json:"kd_user_it"`
	Created     *time.Time `json:"created"`
	Modified    *time.Time `json:"modified"`
	ModiBy      *string    `json:"modi_by"`
	AssignDate  *time.Time `json:"assign_date"`
	FinishDate  *time.Time `json:"finish_date"`
	JenisTicket string    `json:"jenis_ticket"`
	TierTicket  int       `json:"tier_ticket"`
	Solution    *string   `json:"solution"`
	Clients     []response.TicketClient    `json:"kd_user_client"`
}
