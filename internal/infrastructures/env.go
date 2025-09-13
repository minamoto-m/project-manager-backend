package infrastructures

import (
	"log"

	"github.com/joho/godotenv"
)

func InitEnvironment() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error: loading .env file")
		return
	}
}
