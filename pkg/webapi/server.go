package webapi

import (
	"gihub.com/momzor/fizzbuzz/pkg/stats"
	"github.com/gin-gonic/gin"
)

type Server struct {
	conf   Config
	engine *gin.Engine
}

type Config struct {
	addr string
}
type Request struct {
}

type Response struct {
}

type Body struct {
}

func (s Server) Start() error {
	s.engine = gin.Default()
	// middleware catching all requests for stats
	s.engine.Use(stats.AccessiLstner())
	s.engine.GET("/fizzbuzz", getFizzBuzz)
	s.engine.Run(s.conf.addr)

	return nil
}
