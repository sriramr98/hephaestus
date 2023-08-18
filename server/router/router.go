package router

import "github.com/gin-gonic/gin"

type Method int

type Route struct {
	Path     string
	Handlers []gin.HandlerFunc
	Method   string
}

type Router struct {
	Routes []Route
}

func (r *Router) AddRoute(route Route) {
	r.Routes = append(r.Routes, route)
}
