package middleware

import (
	"errors"
	"net/http"

	"github.com/Ardnh/go-todolist.git/helper"
	"github.com/Ardnh/go-todolist.git/model/web"
	"github.com/golang-jwt/jwt/v4"
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

func VerifyJWTToken(writer http.ResponseWriter, request *http.Request) bool {

	jwtKey := helper.LoadEnvFile("JWT_SECRECT_KEY")
	tokenString := request.Header.Get("X-API-KEY")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if token.Valid {
		return true
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   jwt.ErrTokenMalformed,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return false
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   jwt.ErrTokenExpired,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return false
	} else {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return false
	}

}
