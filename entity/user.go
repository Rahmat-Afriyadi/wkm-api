package entity

type User struct {
	ID          uint32   `form:"id" json:"id" gorm:"primary_key;column:id"`
	Name        string   `form:"name" json:"name" gorm:"column:name"`
	Username    string   `form:"username" json:"username" gorm:"column:username"`
	Password    string   `form:"password" json:"password" gorm:"column:password2"`
	DataSource  string   `form:"data_source" json:"data_source" gorm:"column:data_source"`
	Permissions []string `gorm:"type:text;->"`
	RoleId      uint32   `form:"role_id" json:"role_id" gorm:"column:role_id"`
	Role        Role     `form:"role_id" json:"role_id" gorm:"->;references:ID;foreignKey:ID`
	Tier        uint32 	 `form:"tier" json:"tier" gorm:"column:tier"`		 
}

func (User) TableName() string {
	return "mst_users"
}

type Role struct {
	ID   string `form:"id" json:"id" gorm:"primary_key;column:id"`
	Name string `form:"name" json:"name" gorm:"column:name"`
}

func (Role) TableName() string {
	return "mst_roles"
}

type Permission struct {
	ID     string `form:"id" json:"id" gorm:"primary_key;column:id"`
	Name   string `form:"name" json:"name" gorm:"column:name"`
	RoleId string `form:"role_id" json:"role_id" gorm:"column:role_id"`
}

func (Permission) TableName() string {
	return "mst_permissions"
}
