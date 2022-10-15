package controller

import "github.com/Ardnh/go-todolist.git/service"

type UserControllerImpl struct {
	sevice service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		sevice: userService,
	}
}

func (controller *UserControllerImpl) Login() {

}

func (controller *UserControllerImpl) Register() {

}
