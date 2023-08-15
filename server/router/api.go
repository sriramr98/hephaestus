package router

import (
	"github.com/gin-gonic/gin"
	"gitub.com/sriramr98/hephaestus/controllers"
)

func GetAPIRoutes() *Router {
	router := &Router{}

	router.AddRoute(Route{
		Path:     "/health",
		Method:   GET,
		Handlers: []gin.HandlerFunc{controllers.GetHealth},
	})

	setupUserRoutes(router)

	return router
}
