package domain

type Todolist struct {
	Id          int
	UserId      int
	Author      string
	Title       string
	Description string
	IsPublished bool
}
