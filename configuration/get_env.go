package configuration

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Get_env(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error Loading .env")
		panic(err)
	}
	return os.Getenv(key)
}
