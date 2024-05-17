package entity

type User struct {
	ID          uint32
	Name        string
	Username    string
	Password    string
	Group       string
	Permissions []string
}

type UserAsuransi struct {
	ID         uint32
	Nama       string
	Username   string
	Password   string
	DataSource string
}
