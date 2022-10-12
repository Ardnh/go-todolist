package main

import (
	"net/http"

	"github.com/Ardnh/go-todolist.git/app"
	"github.com/Ardnh/go-todolist.git/controller"
	"github.com/Ardnh/go-todolist.git/helper"
	"github.com/Ardnh/go-todolist.git/middleware"
	"github.com/Ardnh/go-todolist.git/repository"
	"github.com/Ardnh/go-todolist.git/service"
	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDb()
	validate := validator.New()
	todolistRepository := repository.NewTodolistRepository()
	todolistService := service.NewTodolistService(todolistRepository, db, validate)
	todolistController := controller.NewTodolistController(todolistService)

	router := app.NewRouter(todolistController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
