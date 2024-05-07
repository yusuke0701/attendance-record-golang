package model

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// DB接続とテーブルを作成する
func DBConnection() *sql.DB {
	var err error

	db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("DB Error: %w", err))
	}

	CreateTable(db)

	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("DB Error: %w", err))
	}

	return sqlDB
}

// Task型のテーブルを作成する
func CreateTable(db *gorm.DB) {
	db.AutoMigrate(&Task{})
}
