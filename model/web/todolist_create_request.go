package web

type CreateTodolistRequest struct {
	Author      string `validate:"required"`
	Title       string `validate:"required"`
	Description string `validate:"required"`
	IsPublished bool   `validate:"required"`
}
