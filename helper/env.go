package helper

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvFile(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		PanicIfError(err)
	}

	return os.Getenv(key)
}
