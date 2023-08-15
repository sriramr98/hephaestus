package router

import "github.com/gin-gonic/gin"

func GetGinRouter() *gin.Engine {
	r := gin.Default()
	apiRoutes := getAPIRoutes()

	for _, route := range apiRoutes.Routes {
		switch route.Method {
		case GET:
			r.GET(route.Path, route.Handlers...)
		case POST:
			r.POST(route.Path, route.Handlers...)
		case PATCH:
			r.PATCH(route.Path, route.Handlers...)
		case DELETE:
			r.DELETE(route.Path, route.Handlers...)
		}
	}

	return r
}
