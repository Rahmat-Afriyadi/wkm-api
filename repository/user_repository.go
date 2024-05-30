package repository

import (
	"context"
	"database/sql"
	"fmt"
	"wkm/entity"
	"wkm/request"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindById(id uint32) entity.User
	FindByUsername(username string) entity.User
	MasterData() []entity.User
	GeneratePassword()
	ResetPassword(data request.ResetPassword)
}

type userRepository struct {
	connUser     *gorm.DB
	connAsuransi *sql.DB
}

func NewUserRepository(connUser *gorm.DB, connAsuransi *sql.DB) UserRepository {
	return &userRepository{
		connUser:     connUser,
		connAsuransi: connAsuransi,
	}
}

func (lR *userRepository) FindById(id uint32) entity.User {
	user := entity.User{ID: id}
	lR.connUser.Find(&user)

	var permissions []entity.Permission
	lR.connUser.Where("role_id", user.RoleId).Find(&permissions)
	for _, v := range permissions {
		user.Permissions = append(user.Permissions, v.Name)
	}

	return user
}

func (lR *userRepository) MasterData() []entity.User {
	var datas []entity.User
	ctx := context.Background()
	query := "select * from mst_users"
	statement, err := lR.connAsuransi.PrepareContext(ctx, query)
	if err != nil {
		fmt.Println(err)
	}
	rows, err := statement.QueryContext(ctx)
	if err != nil {
		fmt.Println("errornya di rows ", err)
		fmt.Println(err)
	}

	for rows.Next() {
		var data entity.User
		if err := rows.Scan(&data.ID, &data.Name, &data.Username, &data.Password, &data.RoleId, &data.DataSource); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		datas = append(datas, data)
	}
	return datas
}

func (lR *userRepository) GeneratePassword() {
	users := lR.MasterData()
	for _, v := range users {
		query := "UPDATE mst_users SET password2 = ? WHERE id = ?"
		password, _ := bcrypt.GenerateFromPassword([]byte(v.Password), 8)
		// Execute the SQL query
		lR.connAsuransi.Exec(query, password, v.ID)
	}
}

func (lR *userRepository) ResetPassword(data request.ResetPassword) {
	query := "UPDATE users SET password2 = ? WHERE id = ?"
	lR.connAsuransi.Exec(query, data.Password, data.IdUser)
}

func (lR *userRepository) FindByUsername(username string) entity.User {
	var user entity.User
	lR.connUser.Where("username", username).First(&user)

	var permissions []entity.Permission
	lR.connUser.Where("role_id", user.RoleId).Find(&permissions)
	for _, v := range permissions {
		user.Permissions = append(user.Permissions, v.Name)
	}

	return user
}
