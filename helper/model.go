package helper

import (
	"github.com/Ardnh/go-todolist.git/model/domain"
	"github.com/Ardnh/go-todolist.git/model/web"
)

func ToTodolistResponse(todolist domain.Todolist) web.TodolistResponse {
	return web.TodolistResponse{
		Id:          todolist.Id,
		Author:      todolist.Author,
		Title:       todolist.Title,
		Description: todolist.Description,
		IsPublished: todolist.IsPublished,
	}
}

func ToTodolistResponses(todolists []domain.Todolist) []web.TodolistResponse {
	var todolistResponses []web.TodolistResponse
	for _, todolist := range todolists {
		todolistResponses = append(todolistResponses, ToTodolistResponse(todolist))
	}

	return todolistResponses
}

func ToUserResponseByUsername(user domain.User) web.UserResponseByUsername {
	return web.UserResponseByUsername{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
		Password:  user.Password,
	}
}
