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

	engine := router.GetRouterEngine()

	server := Server{
		Opts{
			Port:   8080,
			Engine: engine,
		},
	}

	server.Run()
}
