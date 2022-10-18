package web

type CreateTodolistRequest struct {
	UserId      int    `json:"user_id" validate:"required"`
	Author      string `json:"author" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	IsPublished bool   `json:"isPublished" validate:"required"`
}
