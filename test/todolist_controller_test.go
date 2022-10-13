package test

import (
	"database/sql"
	"time"

	"github.com/Ardnh/go-todolist.git/helper"
)

func SetUpTestDB() *sql.DB {
	db, err := sql.Open("mysql", "admin:admin@tcp(localhost:3306)/belajar_golang_restful_api_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
