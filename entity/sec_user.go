package entity

type SecUser struct {
	Username string `form:"username" json:"username" gorm:"column:username"`
	FullName string `form:"fullname" json:"fullname" gorm:"column:fullname"`
}

func (SecUser) TableName() string {
	return "sec_user"
}
