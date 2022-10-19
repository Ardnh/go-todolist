package helper

import (
	"errors"
	"fmt"
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

func VerifyJWTToken(tokenString string) {

	jwtKey := LoadEnvFile("JWT_SECRECT_KEY")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if token.Valid {
		fmt.Println("You look nice today")
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		fmt.Println("That's not even a token")
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		// Token is either expired or not active yet
		fmt.Println("Timing is everything")
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}

}

// func GetJWT(key string, id string) {
// 	mySigningKey := []byte(key)

// 	type MyCustomClaims struct {
// 		Id string `json:"id"`
// 		jwt.RegisteredClaims
// 	}

// 	// Create the claims
// 	claims := MyCustomClaims{
// 		id,
// 		jwt.RegisteredClaims{
// 			// A usual scenario is to set the expiration time relative to the current time
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
// 		},
// 	}

// 	// Create claims while leaving out some of the optional fields
// 	claims = MyCustomClaims{
// 		id,
// 		jwt.RegisteredClaims{
// 			// Also fixed dates can be used for the NumericDate
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	ss, err := token.SignedString(mySigningKey)
// 	fmt.Printf("%v %v", ss, err)
// }
