package configuration

import (
	"os"

	"github.com/joho/godotenv"
)

func Get_env(key string) string {
	godotenv.Load(".env")

	return os.Getenv(key)
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return Get_env("APP_PORT")
	}
	return port
}
