package repositories

import (
	"fmt"
	. "main/types"

	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _connection *gorm.DB

func connection() *gorm.DB {
	if _connection != nil {
		return _connection
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_ADDRESS"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Account{}, &Transaction{})

	_connection = db

	return db
}

func Connected() bool {
	return !connection().Config.DryRun
}

