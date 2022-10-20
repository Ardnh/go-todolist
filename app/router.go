package app

import (
	"database/sql"

	"github.com/Ardnh/go-todolist.git/controller"
	"github.com/Ardnh/go-todolist.git/exception"
	"github.com/Ardnh/go-todolist.git/middleware"
	"github.com/Ardnh/go-todolist.git/repository"
	"github.com/Ardnh/go-todolist.git/service"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(db *sql.DB, validate *validator.Validate) *httprouter.Router {
	router := httprouter.New()

	todolistRepository := repository.NewTodolistRepository()
	todolistService := service.NewTodolistService(todolistRepository, db, validate)
	todolistController := controller.NewTodolistController(todolistService)

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	router.GET("/api/todolists", todolistController.FindAll)
	router.GET("/api/todolist/:todolistId", todolistController.FindById)
	router.POST("/api/todolist", middleware.AuthCheck(todolistController.Create))
	router.PUT("/api/todolist/:todolistId", middleware.AuthCheck(todolistController.Update))
	router.DELETE("/api/todolist/:todolistId", middleware.AuthCheck(todolistController.Delete))

	router.POST("/api/login", userController.Login)
	router.POST("/api/register", userController.Register)

	router.PanicHandler = exception.ErrorHandler

	return router
}
