package helper

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvFile(key string) string {
	err := godotenv.Load()
	if err != nil {
		PanicIfError(err)
	}

	return os.Getenv(key)
}
