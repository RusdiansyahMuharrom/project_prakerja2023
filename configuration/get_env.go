package configuration

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func Get_env(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error Loading .env")
	}
	return os.Getenv(key)
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return Get_env("APP_PORT")
	}
	return port
}

func GetHostScheme() (string, []string) {
	app_host := Get_env("APP_HOST")
	temp := strings.Split(app_host, "://")
	schema := []string{temp[0]}

	return temp[1], schema
}
