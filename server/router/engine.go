package router

import (
	"github.com/gin-gonic/gin"
)

func GetRouterEngine() *gin.Engine {
	r := gin.Default()
	apiRouter := GetAPIRoutes()
	for _, route := range apiRouter.Routes {
		r.Handle(route.Method, route.Path, route.Handlers...)
	}

	return r
}
