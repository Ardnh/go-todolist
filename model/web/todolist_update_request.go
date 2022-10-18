package web

type UpdateTodolistRequest struct {
	Id          int    `json:"id" validate:"required"`
	UserId      int    `json:"user_id" validate:"required"`
	Author      string `json:"author" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	IsPublished bool   `json:"isPublished" validate:"required"`
}
