package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Ardnh/go-todolist.git/exception"
	"github.com/Ardnh/go-todolist.git/helper"
	"github.com/Ardnh/go-todolist.git/model/domain"
	"github.com/Ardnh/go-todolist.git/model/web"
	"github.com/Ardnh/go-todolist.git/repository"
	"github.com/go-playground/validator/v10"
)

type TodolistServiceImpl struct {
	TodolistRepository repository.TodolistRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewTodolistService(todolistRepository repository.TodolistRepository, DB *sql.DB, validate *validator.Validate) TodolistService {
	return &TodolistServiceImpl{
		TodolistRepository: todolistRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *TodolistServiceImpl) Create(ctx context.Context, request web.CreateTodolistRequest) web.TodolistResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	helper.PanicIfError(err)
	todolist := domain.Todolist{
		UserId:      request.UserId,
		Author:      request.Author,
		Title:       request.Title,
		Description: request.Description,
		IsPublished: request.IsPublished,
	}

	todolist = service.TodolistRepository.Save(ctx, tx, todolist)
	return helper.ToTodolistResponse(todolist)
}

func (service *TodolistServiceImpl) Update(ctx context.Context, request web.UpdateTodolistRequest) web.TodolistResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	fmt.Println(request)

	todolist, err := service.TodolistRepository.FindById(ctx, tx, request.Id)

	fmt.Println(todolist)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	todolist.Author = request.Author
	todolist.Title = request.Title
	todolist.Description = request.Description
	todolist.IsPublished = request.IsPublished

	todolist = service.TodolistRepository.Update(ctx, tx, todolist)
	return helper.ToTodolistResponse(todolist)
}

func (service *TodolistServiceImpl) Delete(ctx context.Context, todolistId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todolist, err := service.TodolistRepository.FindById(ctx, tx, todolistId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.TodolistRepository.Delete(ctx, tx, todolist)
}

func (service *TodolistServiceImpl) FindAll(ctx context.Context) []web.TodolistResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todolists := service.TodolistRepository.FindAll(ctx, tx)

	return helper.ToTodolistResponses(todolists)
}

func (service *TodolistServiceImpl) FindById(ctx context.Context, todolistId int) web.TodolistResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todolist, err := service.TodolistRepository.FindById(ctx, tx, todolistId)
	if err != nil {
		helper.PanicIfError(err)
		exception.NewNotFoundError(err.Error())
	}
	return web.TodolistResponse(todolist)
}
