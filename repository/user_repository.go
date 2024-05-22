package repository

import (
	"context"
	"database/sql"
	"fmt"
	"wkm/entity"
	"wkm/request"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	FindById(id uint32) entity.User
	FindByIdAsuransi(id uint32) entity.UserAsuransi
	FindByUsername(username string) entity.User
	FindByUsernameAsuransi(username string) entity.UserAsuransi
	MasterData() []entity.User
	GeneratePassword()
	ResetPassword(data request.ResetPassword)
}

type userRepository struct {
	conn *sql.DB
}

func NewUserRepository(conn *sql.DB) UserRepository {
	return &userRepository{
		conn: conn,
	}
}

func (lR *userRepository) FindById(id uint32) entity.User {
	var data entity.User
	ctx := context.Background()
	query := "select id, name, username, password2, 'group' from users_wkms where id=?"
	statement, err := lR.conn.PrepareContext(ctx, query)
	if err != nil {
		fmt.Println("errror disini")
		fmt.Println(err)
	}
	err = statement.QueryRow(id).Scan(&data.ID, &data.Name, &data.Username, &data.Password, &data.Group)
	if err != nil {
		fmt.Println("errornya di roww user", err)
		fmt.Println(err)
	}

	return data
}

func (lR *userRepository) FindByIdAsuransi(id uint32) entity.UserAsuransi {
	var data entity.UserAsuransi
	ctx := context.Background()
	query := "select id, name, username, password2, data_source from users where id=?"
	statement, err := lR.conn.PrepareContext(ctx, query)
	if err != nil {
		fmt.Println("errror disini")
		fmt.Println(err)
	}
	err = statement.QueryRow(id).Scan(&data.ID, &data.Nama, &data.Username, &data.Password, &data.DataSource)
	if err != nil {
		fmt.Println("errornya di roww ", err)
		fmt.Println(err)
	}

	return data
}

func (lR *userRepository) MasterData() []entity.User {
	var datas []entity.User
	ctx := context.Background()
	query := "select * from mst_users"
	statement, err := lR.conn.PrepareContext(ctx, query)
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
		var a string
		if err := rows.Scan(&data.ID, &data.Name, &data.Username, &data.Password, &data.Group, &a); err != nil {
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
		lR.conn.Exec(query, password, v.ID)
	}
}

func (lR *userRepository) ResetPassword(data request.ResetPassword) {
	query := "UPDATE users SET password2 = ? WHERE id = ?"
	lR.conn.Exec(query, data.Password, data.IdUser)
}

func (lR *userRepository) FindByUsername(username string) entity.User {
	var data entity.User
	ctx := context.Background()
	query := "select id, name, username, password2, 'group' from users_wkms where username=?"
	statement, err := lR.conn.PrepareContext(ctx, query)
	if err != nil {
		fmt.Println(err)
	}
	row := statement.QueryRow(username)
	err = row.Scan(&data.ID, &data.Name, &data.Username, &data.Password, &data.Group)
	if err != nil {
		fmt.Println(err)
	}

	// query_permissions := "select permission_type from wkms_permissions where user_id=?"
	// statement_permission, err := lR.conn.PrepareContext(ctx, query_permissions)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// rows, err := statement_permission.QueryContext(ctx, data.ID)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for rows.Next() {
	// 	var permission string
	// 	if err := rows.Scan(&permission); err != nil {
	// 		fmt.Println("Error scanning row:", err)
	// 		continue
	// 	}
	// 	data.Permissions = append(data.Permissions, permission)
	// }

	return data
}

func (lR *userRepository) FindByUsernameAsuransi(username string) entity.UserAsuransi {
	var data entity.UserAsuransi
	ctx := context.Background()
	query := "select id, name, username, password2, data_source from users where username=?"
	statement, err := lR.conn.PrepareContext(ctx, query)
	if err != nil {
		fmt.Println(err)
	}
	row := statement.QueryRow(username)
	err = row.Scan(&data.ID, &data.Nama, &data.Username, &data.Password, &data.DataSource)
	if err != nil {
		fmt.Println(err)
	}

	return data
}
