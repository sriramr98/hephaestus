package main

import (
	"github.com/joho/godotenv"
	"gitub.com/sriramr98/hephaestus/router"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	r := router.GetGinRouter()
	err := r.Run()
	if err != nil {
		log.Fatalf("%v", err)
	}
}
