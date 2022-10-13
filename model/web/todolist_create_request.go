package web

type CreateTodolistRequest struct {
	Author      string `json:"author" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	IsPublished string `json:"isPublished" validate:"required"`
}
