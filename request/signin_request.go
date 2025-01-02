package request

type SigninRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
