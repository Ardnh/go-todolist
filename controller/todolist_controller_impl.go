package controller

import (
	"net/http"
	"strconv"

	"github.com/Ardnh/go-todolist.git/helper"
	"github.com/Ardnh/go-todolist.git/model/web"
	"github.com/Ardnh/go-todolist.git/service"
	"github.com/julienschmidt/httprouter"
)

type TodolistControllerImpl struct {
	service service.TodolistService
}

func NewTodolistController(todolistService service.TodolistService) TodolistController {
	return &TodolistControllerImpl{
		service: todolistService,
	}
}

func (controller *TodolistControllerImpl) Create(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	todolistCreateRequest := web.CreateTodolistRequest{}
	helper.ReadFromRequestBody(request, &todolistCreateRequest)

	todolistResponse := controller.service.Create(request.Context(), todolistCreateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todolistResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TodolistControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todolistUpdateRequest := web.UpdateTodolistRequest{}
	helper.ReadFromRequestBody(request, &todolistUpdateRequest)

	todolistId := params.ByName("todolistId")
	id, err := strconv.Atoi(todolistId)
	helper.PanicIfError(err)

	todolistUpdateRequest.Id = id

	todolistResponse := controller.service.Update(request.Context(), todolistUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todolistResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TodolistControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	todolistId := params.ByName("todolistId")
	id, err := strconv.Atoi(todolistId)
	helper.PanicIfError(err)

	controller.service.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TodolistControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	todolistId := params.ByName("todolistId")
	id, err := strconv.Atoi(todolistId)
	helper.PanicIfError(err)

	todolistResponse := controller.service.FindById(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todolistResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TodolistControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todolistResponse := controller.service.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todolistResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
