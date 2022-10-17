package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	Login(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	Register(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
}
