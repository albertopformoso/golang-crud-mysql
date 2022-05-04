package app

import (
	"fmt"
	"log"
	"github.com/joho/godotenv"
)

func LoadDotEnv(file string) {
	err := godotenv.Load(file)

	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	fmt.Println(".env file loaded successfuly")
}