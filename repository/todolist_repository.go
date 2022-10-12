package repository

import (
	"context"
	"database/sql"

	"github.com/Ardnh/go-todolist.git/model/domain"
)

type TodolistRepository interface {
	Save(ctx context.Context, tx *sql.Tx, todolist domain.Todolist) domain.Todolist
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Todolist
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Todolist, error)
	Update(ctx context.Context, tx *sql.Tx, todolist domain.Todolist) domain.Todolist
	Delete(ctx context.Context, tx *sql.Tx, todolist domain.Todolist)
}
