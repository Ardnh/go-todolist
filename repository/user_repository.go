package repository

import (
	"context"
	"database/sql"

	"github.com/Ardnh/go-todolist.git/model/domain"
)

type UserRepository interface {
	FindByUsername(ctx context.Context, tx *sql.Tx, username string) (domain.User, error)
	Register(ctx context.Context, tx *sql.Tx, user domain.User)
}
