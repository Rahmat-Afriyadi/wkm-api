package entity

type User struct {
	ID          uint32
	Name        string
	Username    string
	Password    string
	Group       string
	Permissions []string
}
