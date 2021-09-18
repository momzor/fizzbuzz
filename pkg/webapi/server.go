package webapi

import (
	"fmt"

	"gihub.com/momzor/fizzbuzz/pkg/stats"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Conf   Config
	Engine *gin.Engine
}

type Config struct {
	BaseUrl string
	Port    string
}
type Request struct {
}

type Response struct {
}

type Body struct {
}

func (s *Server) Start() error {
	s.Engine = gin.Default()
	// middleware catching all requests for stats
	s.Engine.Use(stats.AccessiLstner())
	s.Engine.GET("/fizzbuzz", getFizzBuzz)
	fmt.Println(s.Conf.BaseUrl + ":" + s.Conf.Port)
	fmt.Println("DEBUG")
	s.Engine.Run(s.Conf.BaseUrl + ":" + s.Conf.Port)

	return nil
}
