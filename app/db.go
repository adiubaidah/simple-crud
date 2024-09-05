package app

import (
	"adiubaidah/simple-crud/helper"
	"database/sql"
	"fmt"
	"time"
)

func NewDB() *sql.DB {
	config := GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.DBUser,
		config.DBPass,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)
	db, err := sql.Open("mysql", dsn)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
