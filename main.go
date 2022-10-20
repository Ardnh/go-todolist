package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Ardnh/go-todolist.git/app"
	"github.com/Ardnh/go-todolist.git/helper"
	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDb()
	validate := validator.New()

	router := app.NewRouter(db, validate)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
