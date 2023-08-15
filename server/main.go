package main

import (
	"log"

	"github.com/joho/godotenv"
	"gitub.com/sriramr98/hephaestus/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	r := router.GetGinRouter()

	r.Run()
}
