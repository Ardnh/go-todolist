package app

import (
	"github.com/Ardnh/go-todolist.git/controller"
	"github.com/Ardnh/go-todolist.git/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(todolistController controller.TodolistController, userController controller.UserController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/todolists", todolistController.FindAll)
	router.GET("/api/todolist/:todolistId", todolistController.FindById)
	router.POST("/api/todolist", todolistController.Create)
	router.PUT("/api/todolist/:todolistId", todolistController.Update)
	router.DELETE("/api/todolist/:todolistId", todolistController.Delete)

	router.POST("/api/login", userController.Login)
	router.POST("/api/register", userController.Register)

	router.PanicHandler = exception.ErrorHandler

	return router
}
