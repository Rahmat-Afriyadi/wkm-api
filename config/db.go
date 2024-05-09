package config

import (
	"database/sql"
	"fmt"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_wkm?parseTime=true")
	// db, err := sql.Open("mysql", "root2:root2@tcp(192.168.70.30:3306)/db_wkm?parseTime=true")
	if err != nil {
		fmt.Println(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
