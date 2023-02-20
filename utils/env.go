package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadENV(key string) string {
	err := godotenv.Load(".env")
	HandleErr("EnvLoad Error", err)
	return os.Getenv(key)
}

