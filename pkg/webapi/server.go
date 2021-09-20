package webapi

import (
	"gihub.com/momzor/fizzbuzz/pkg/db"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Conf     Config
	Engine   *gin.Engine
	DBClient db.Client
}

type Config struct {
	BaseUrl string
	Port    string
}

type ErrorResponse struct {
	Status  int
	Errors  []string `json:"errors,omitempty"`
	Message string   `json:"message"`
}

func (s *Server) Start() error {
	s.Engine = gin.Default()
	// middleware catching all requests for stats
	s.Engine.Use(s.statsMiddleware())
	s.Engine.GET(FIZZBUZZ_RESOURCE_ROUTE, s.FizzBuzzHandler)
	s.Engine.GET(STATS_RESOURCE_ROUTE, s.StatsHandler)
	s.Engine.Run(s.Conf.BaseUrl + ":" + s.Conf.Port)

	return nil
}
