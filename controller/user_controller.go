package controller

import (
	"net/http"
)

type UserController interface {
	Login(writer http.ResponseWriter, request *http.Request)
	Register(writer http.ResponseWriter, request *http.Request)
}
