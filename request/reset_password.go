package request

type ResetPassword struct {
	IdUser             uint32
	PasswordLama       string `form:"password_lama" json:"password_lama"`
	Password           string `form:"password" json:"password"`
	PasswordKonfirmasi string `form:"password_konfirmasi" json:"password_konfirmasi"`
}
