package repository

import (
	"context"
	"database/sql"
	"fmt"
	"wkm/entity"

	"golang.org/x/crypto/bcrypt"
)

type UserAsuransiRepository interface {
	FindByIdAsuransi(id uint32) entity.User
	FindByUsernameAsuransi(username string) entity.UserAsuransi
	MasterData() []entity.User
	GeneratePassword()
}

type userAsuransiRepository struct {
	conn *sql.DB
}

func NewUserAsuransiRepository(conn *sql.DB) UserAsuransiRepository {
	return &userAsuransiRepository{
		conn: conn,
	}
}

func (lR *userAsuransiRepository) FindByIdAsuransi(id uint32) entity.User {
	var data entity.User
	ctx := context.Background()
	query := "select id, name, username, password2, data_source from users where id=?"
	statement, err := lR.conn.PrepareContext(ctx, query)
	if err != nil {
		fmt.Println("errror disini")
		fmt.Println(err)
	}
	err = statement.QueryRow(id).Scan(&data.ID, &data.Name, &data.Username, &data.Password, &data.Group)
	if err != nil {
		fmt.Println("errornya di roww ", err)
		fmt.Println(err)
	}

	return data
}

func (lR *userAsuransiRepository) MasterData() []entity.User {
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

func (lR *userAsuransiRepository) GeneratePassword() {
	users := lR.MasterData()
	for _, v := range users {
		query := "UPDATE mst_users SET password2 = ? WHERE id = ?"
		password, _ := bcrypt.GenerateFromPassword([]byte(v.Password), 8)
		// Execute the SQL query
		lR.conn.Exec(query, password, v.ID)
	}
}

func (lR *userAsuransiRepository) FindByUsernameAsuransi(username string) entity.UserAsuransi {
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
