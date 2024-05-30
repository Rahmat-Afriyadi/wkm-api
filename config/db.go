package config

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_wkm?parseTime=true&loc=Asia%2FJakarta")
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

func GetConnectionUser() *gorm.DB {
	dsn := "root:@tcp(localhost:3306)/users?parseTime=true&loc=Asia%2FJakarta"
	// dsn := "root2:root2@tcp(192.168.70.30:3306)/db_wkm?parseTime=true"
	// db, err := sql.Open("mysql", "root2:root2@tcp(192.168.70.30:3306)/db_wkm?parseTime=true")
	time.LoadLocation("Asia/Jakarta")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase:         true,
			IdentifierMaxLength: 30,
		},
		PrepareStmt:     false,
		CreateBatchSize: 50,
	})
	if err != nil {
		fmt.Println("Masuk sini gk guys ", err)
		panic(err)
	}

	return db
}

func GetConnectionAsuransi() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/wanda_asuransi?parseTime=true&loc=Asia%2FJakarta")
	// db, err := sql.Open("mysql", "root2:root2@tcp(192.168.70.30:3306)/asuransi?parseTime=true")
	if err != nil {
		fmt.Println(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func NewAsuransiGorm() *gorm.DB {
	time.LoadLocation("Asia/Jakarta")
	dsn := "root@tcp(localhost:3306)/wanda_asuransi?parseTime=true&loc=Asia%2FJakarta"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase:         true,
			IdentifierMaxLength: 30,
		},
		PrepareStmt:     false,
		CreateBatchSize: 50,
	})
	if err != nil {
		fmt.Println("Masuk sini gk guys logger ", err)
		panic(err)
	}
	return db
}
