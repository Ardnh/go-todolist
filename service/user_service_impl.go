package service

import (
	"context"
	"database/sql"

	"github.com/Ardnh/go-todolist.git/helper"
	"github.com/Ardnh/go-todolist.git/model/domain"
	"github.com/Ardnh/go-todolist.git/model/web"
	"github.com/Ardnh/go-todolist.git/repository"
	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, db *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Register(ctx context.Context, request web.CreateUserRequest) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	hashPassword, err := helper.GenerateHashPassword(request.Password)

	if err != nil {
		helper.PanicIfError(err)
	}

	user := domain.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		UserName:  request.UserName,
		Password:  hashPassword,
	}

	service.userRepository.Register(ctx, tx, user)
}

func (service *UserServiceImpl) FindByUsername(ctx context.Context, request string) (web.UserResponseByUsername, error) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.userRepository.FindByUsername(ctx, tx, request)
	if err != nil {
		helper.PanicIfError(err)

		return helper.ToUserResponseByUsername(user), err
	}

	return helper.ToUserResponseByUsername(user), nil
}
