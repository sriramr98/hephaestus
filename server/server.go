package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Opts struct {
	Port   int
	Engine *gin.Engine
}

type Server struct {
	Opts
}

func (s Server) Run() {
	err := s.Engine.Run(fmt.Sprintf("0.0.0.0:%d", s.Port))
	if err != nil {
		log.Fatal(err)
	}
}
