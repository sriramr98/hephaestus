package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	middlewares "gitub.com/sriramr98/hephaestus/middleware"
)

func GetGinRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Authenticator())
	apiRoutes := GetAPIRoutes()

	for _, route := range apiRoutes.Routes {
		switch route.Method {
		case http.MethodGet:
			r.GET(route.Path, route.Handlers...)
		case http.MethodPost:
			r.POST(route.Path, route.Handlers...)
		case http.MethodPatch:
			r.PATCH(route.Path, route.Handlers...)
		case http.MethodDelete:
			r.DELETE(route.Path, route.Handlers...)
		}
	}

	return r
}
