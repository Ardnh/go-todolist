package controller

import (
	"fmt"
	"net/http"

	"github.com/Ardnh/go-todolist.git/exception"
	"github.com/Ardnh/go-todolist.git/helper"
	"github.com/Ardnh/go-todolist.git/model/web"
	"github.com/Ardnh/go-todolist.git/service"
	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	service service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		service: userService,
	}
}

func (controller *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {

	userLoginRequest := web.UserLoginRequest{}
	helper.ReadFromRequestBody(request, &userLoginRequest)
	user, err := controller.service.FindByUsername(request.Context(), userLoginRequest.Username)
	if err != nil {
		exception.NewNotFoundError(err.Error())
	}

	fmt.Println(user)

	if helper.CheckPassword(userLoginRequest.Password, user.Password) {
		id := fmt.Sprintf("%d", user.Id)
		token, err := helper.GenerateJWTKey(id)
		if err != nil {
			exception.InternalServerError(writer, request, err)
		}

		webResponse := web.UserResponseWithToken{
			Code:   200,
			Status: "OK",
			Data: web.ResponseWithToken{
				Id:        user.Id,
				FirstName: user.FirstName,
				LastName:  user.LastName,
				UserName:  user.UserName,
				Token:     token,
			},
		}

		helper.WriteToResponseBody(writer, webResponse)
	} else {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Wrong Password",
		}
		helper.WriteToResponseBody(writer, webResponse)
	}

}

func (controller *UserControllerImpl) Register(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {

	userRegisterController := web.CreateUserRequest{}
	helper.ReadFromRequestBody(request, &userRegisterController)

	controller.service.Register(request.Context(), userRegisterController)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
