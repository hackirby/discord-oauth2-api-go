package main

import (
	"discord-oauth2/internal/app"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
