package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetGinRouter() *gin.Engine {
	r := gin.Default()
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
