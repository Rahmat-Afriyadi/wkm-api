package entity

type MstToken struct {
	Id    string `json:"id" gorm:"column:idtoken"`
	Name  string `json:"name" gorm:"column:nm_user"`
	Token string `json:"token" gorm:"column:token"`
}

func (MstToken) TableName() string {
	return "token"
}
