package service

import (
	"context"

	"github.com/Ardnh/go-todolist.git/model/web"
)

type UserService interface {
	Register(ctx context.Context, request web.CreateUserRequest)
	FindByUsername(ctx context.Context, username string) (web.UserResponseByUsername, error)
}
