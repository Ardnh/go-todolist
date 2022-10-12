package service

import (
	"context"

	"github.com/Ardnh/go-todolist.git/model/web"
)

type TodolistService interface {
	Create(ctx context.Context, request web.CreateTodolistRequest) web.TodolistResponse
	Update(ctx context.Context, request web.UpdateTodolistRequest) web.TodolistResponse
	Delete(ctx context.Context, todolistId int)
	FindById(ctx context.Context, todolistid int) web.TodolistResponse
	FindAll(ctx context.Context) []web.TodolistResponse
}
