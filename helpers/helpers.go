package helpers

import (
	"log"

	"github.com/joho/godotenv"
)

func LogIfError(err error, info string) {
	if err != nil {
		log.Printf("[ERROR] %s : %s", info, err)
	}
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("[ERROR] Error loading .env file")
		panic(err)
	}
}

type EmptyObj struct {
}
