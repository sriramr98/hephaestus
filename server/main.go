package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gitub.com/sriramr98/hephaestus/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	r := gin.Default()
	apiRoutes := router.GetAPIRoutes()

	for _, route := range apiRoutes.Routes {
		switch route.Method {
		case router.GET:
			r.GET(route.Path, route.Handlers...)
		case router.POST:
			r.POST(route.Path, route.Handlers...)
		case router.PATCH:
			r.PATCH(route.Path, route.Handlers...)
		case router.DELETE:
			r.DELETE(route.Path, route.Handlers...)
		}
	}

	r.Run()
}
