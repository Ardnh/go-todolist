package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWTKey(id string) (string, error) {

	jwtKey := LoadEnvFile("JWT_SECRECT_KEY")
	signingKey := []byte(jwtKey)
	expirationTime := time.Now().Add(60 * time.Minute)

	type MyCustomClaims struct {
		Id string `json:"id"`
		jwt.RegisteredClaims
	}

	// Create the claims
	claims := MyCustomClaims{
		id,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	// Create claims while leaving out some of the optional fields
	claims = MyCustomClaims{
		id,
		jwt.RegisteredClaims{
			// Also fixed dates can be used for the NumericDate
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signingKey)

	PanicIfError(err)

	return ss, err
}
