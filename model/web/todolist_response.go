package web

type TodolistResponse struct {
	Id          int    `json:"id"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsPublished bool   `json:"isPublished"`
}
