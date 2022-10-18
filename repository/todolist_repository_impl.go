package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Ardnh/go-todolist.git/helper"
	"github.com/Ardnh/go-todolist.git/model/domain"
)

type TodolistRepositoryImpl struct {
}

func NewTodolistRepository() TodolistRepository {
	return &TodolistRepositoryImpl{}
}

func (repository *TodolistRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, todolist domain.Todolist) domain.Todolist {
	SQL := "INSERT INTO todolist( user_id, author, title, description, isPublished ) VALUE(?,?,?,?,?) ;"
	result, err := tx.ExecContext(ctx, SQL, todolist.UserId, todolist.Author, todolist.Title, todolist.Description, todolist.IsPublished)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	todolist.Id = int(id)

	return todolist
}

func (repository *TodolistRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Todolist {
	SQL := "SELECT * FROM todolist;"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()

	var todolists []domain.Todolist
	for rows.Next() {
		todolist := domain.Todolist{}
		err := rows.Scan(&todolist.Id, &todolist.UserId, &todolist.Author, &todolist.Title, &todolist.Description, &todolist.IsPublished)
		helper.PanicIfError(err)

		todolists = append(todolists, todolist)
	}

	return todolists
}

func (repository *TodolistRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Todolist, error) {
	SQL := "SELECT id, user_id, author, title, description, isPublished FROM todolist WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer rows.Close()

	var todolist domain.Todolist
	if rows.Next() {
		err := rows.Scan(&todolist.Id, &todolist.UserId, &todolist.Author, &todolist.Title, &todolist.Description, &todolist.IsPublished)
		helper.PanicIfError(err)
		return todolist, nil
	} else {
		return todolist, errors.New("todolist not found")
	}
}

func (repository *TodolistRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, todolist domain.Todolist) domain.Todolist {
	SQL := "UPDATE todolist SET author = ?, title = ?, description = ?, isPublished = ? WHERE id = ? AND user_id = ?;"
	_, err := tx.ExecContext(ctx, SQL, todolist.Author, todolist.Title, todolist.Description, todolist.IsPublished, todolist.Id, todolist.UserId)
	helper.PanicIfError(err)

	return todolist
}

func (repository *TodolistRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, todolist domain.Todolist) {
	SQL := "DELETE FROM todolist WHERE id = ? AND user_id = ?;"
	_, err := tx.ExecContext(ctx, SQL, todolist.Id, todolist.UserId)
	helper.PanicIfError(err)
}
