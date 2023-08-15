package router

import "github.com/gin-gonic/gin"

type Method int

const (
	GET Method = iota
	POST
	PUT
	PATCH
	DELETE
)

type Route struct {
	Path     string
	Handlers []gin.HandlerFunc
	Method   Method
}

type Router struct {
	Routes []Route
}

func (r *Router) AddRoute(route Route) {
	r.Routes = append(r.Routes, route)
}
