package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Ardnh/go-todolist.git/helper"
	"github.com/Ardnh/go-todolist.git/model/web"
	"github.com/golang-jwt/jwt/v4"
	"github.com/julienschmidt/httprouter"
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

func verifyJWTToken(tokenString string) (bool, error) {

	jwtKey := helper.LoadEnvFile("JWT_SECRECT_KEY")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if token.Valid {
		return true, nil
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		return false, jwt.ErrTokenMalformed
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		return false, jwt.ErrTokenExpired
	} else {
		return false, jwt.ErrTokenNotValidYet
	}
}

func AuthCheck(h httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		tokenString := request.Header.Get("X-API-KEY")

		if tokenString != "" {
			ok, err := verifyJWTToken(tokenString)

			fmt.Println(ok)
			fmt.Println(err)

			if !ok && err != nil {
				webResponse := web.WebResponse{
					Code:   http.StatusInternalServerError,
					Status: "INTERNAL SERVER ERROR",
					Data:   err,
				}

				helper.WriteToResponseBody(writer, webResponse)
			} else {
				h(writer, request, params)
			}

		} else {
			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
			}

			helper.WriteToResponseBody(writer, webResponse)
		}
	}
}
