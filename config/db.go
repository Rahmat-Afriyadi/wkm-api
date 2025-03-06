package config

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)


func GetConnection() (*gorm.DB, *sql.DB) {
	errEnv := godotenv.Load()
	if errEnv != nil {
		fmt.Println("ini errornya ", errEnv)
		panic("Failed to load env file")
	}

	dsn := os.Getenv("DB_WKM")
	time.LoadLocation("Asia/Jakarta")
	instance, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase:         true,
			IdentifierMaxLength: 30,
		},
		PrepareStmt:     true,
		CreateBatchSize: 50,
	})

	if err != nil {
		fmt.Println("Error db users ", err)
		panic(err)
	}

	db, _ := instance.DB()

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return instance, db
}

func GetConnectionECardPlus() (*gorm.DB, *sql.DB) {
	errEnv := godotenv.Load()
	if errEnv != nil {
		fmt.Println("ini errornya ", errEnv)
		panic("Failed to load env file")
	}

	dsn := os.Getenv("MS_WKM")
	time.LoadLocation("Asia/Jakarta")
	instance, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase:         true,
			IdentifierMaxLength: 30,
		},
		PrepareStmt:     true,
		CreateBatchSize: 50,
	})

	if err != nil {
		fmt.Println("Error db users ", err)
		panic(err)
	}

	db, _ := instance.DB()

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return instance, db
}

func GetConnectionTest() (*gorm.DB, *sql.DB) {
	errEnv := godotenv.Load()
	if errEnv != nil {
		fmt.Println("ini errornya ", errEnv)
		panic("Failed to load env file")
	}

	dsn := os.Getenv("DB_WKM_TEST")
	time.LoadLocation("Asia/Jakarta")
	instance, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase:         true,
			IdentifierMaxLength: 30,
		},
		PrepareStmt:     true,
		CreateBatchSize: 50,
	})


	// test


	if err != nil {
		fmt.Println("Error db users ", err)
		panic(err)
	}

	db, _ := instance.DB()

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return instance, db
}

func GetConnectionUser() (*gorm.DB, *sql.DB) {

	errEnv := godotenv.Load()
	if errEnv != nil {
		fmt.Println("ini errornya ", errEnv)
		panic("Failed to load env file")
	}
	dsn := os.Getenv("MST_USER")
	time.LoadLocation("Asia/Jakarta")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase:         true,
			IdentifierMaxLength: 30,
		},
		PrepareStmt:     true,
		CreateBatchSize: 50,
	})
	if err != nil {
		fmt.Println("Error db users ", err)
		panic(err)
	}

	sqlDB, _ := db.DB()

	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(10 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	return db, sqlDB
}

func NewAsuransiGorm() (*gorm.DB, *sql.DB) {
	errEnv := godotenv.Load()
	if errEnv != nil {
		fmt.Println("ini errornya ", errEnv)
		panic("Failed to load env file")
	}
	time.LoadLocation("Asia/Jakarta")
	dsn := os.Getenv("WANDA_ASURANSI")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase:         true,
			IdentifierMaxLength: 30,
		},
		PrepareStmt:     true,
		CreateBatchSize: 50,
	})
	if err != nil {
		fmt.Println("Masuk sini gk guys logger ", err)
		panic(err)
	}
	sqlDB, _ := db.DB()

	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(20 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	return db, sqlDB
}
