package middleware

import (
	"net/http"

	"github.com/Ardnh/go-todolist.git/helper"
	"github.com/Ardnh/go-todolist.git/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Header.Get("X-API-KEY") == "RAHASIA" {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
	}
}
