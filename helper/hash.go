package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), 5)

	if err != nil {
		return string(bytes), err
	}

	return string(bytes), nil
}

func CheckPassword(password string, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))

	if err != nil {
		return false
	} else {
		return true
	}

}
