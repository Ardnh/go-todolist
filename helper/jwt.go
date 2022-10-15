package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Id string
	jwt.RegisteredClaims
}

func GenerateJWTKey(id string) (string, error) {

	jwtKey := LoadEnvFile("JWT_SECRECT_KEY")
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

func VerifyJWTToken(token string) {

}
