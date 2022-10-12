package app

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Ardnh/go-todolist.git/helper"
)

func NewDb() *sql.DB {
	db_username := helper.LoadEnvFile("DB_USERNAME")
	db_password := helper.LoadEnvFile("DB_PASSWORD")
	db_name := helper.LoadEnvFile("DB_NAME")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", db_username, db_password, db_name))
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
