package repository

import (
	"context"
	"database/sql"
	"errors"

	// "fmt"

	"github.com/Ardnh/go-todolist.git/helper"
	"github.com/Ardnh/go-todolist.git/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, name string) (domain.User, error) {
	SQL := "SELECT id, firstname, lastname, username, password FROM user WHERE username = ?;"
	// rows, err := tx.QueryContext(ctx, SQL, fmt.Sprintf("%s", name))
	rows, err := tx.QueryContext(ctx, SQL, name)
	helper.PanicIfError(err)
	defer rows.Close()

	var user domain.User
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.UserName, &user.Password)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user not found")
	}

}

func (repository *UserRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, user domain.User) {
	SQL := "INSERT INTO user( firstname, lastname, username, password ) VALUES ( ?,?,?,? )"
	_, err := tx.ExecContext(ctx, SQL, user.FirstName, user.LastName, user.UserName, user.Password)
	helper.PanicIfError(err)
}
