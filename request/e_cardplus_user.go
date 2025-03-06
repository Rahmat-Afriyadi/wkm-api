package request

type CreateECardplusUserRequest struct {
	NoHp     string `json:"no_hp"`
	Fullname string `json:"name"`
	Password string `json:"password"`
}
