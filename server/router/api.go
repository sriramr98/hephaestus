package router

import (
	"github.com/gin-gonic/gin"
	"gitub.com/sriramr98/hephaestus/controllers"
	"net/http"
)

func GetAPIRoutes() Router {

	router := &Router{}

	router.AddRoute(Route{
		Path:     "/health",
		Method:   http.MethodGet,
		Handlers: []gin.HandlerFunc{controllers.GetHealth},
	})

	setupUserRoutes(router)

	return *router
}
