package webapi

import (
	"github.com/gin-gonic/gin"
	_ "github.com/momzor/fizzbuzz/docs"
	"github.com/momzor/fizzbuzz/pkg/db"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type Server struct {
	Conf     Config
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

// @title  LBC test fizzbuzz
// @version 1.0
// @description LBC test fizzbuzz
// @contact.name Momzor
// @contact.email m.benaida.pro@gmail.com
// @license.url https://www.gnu.org/licenses/quick-guide-gplv3.html
// @host localhost:80
func (s *Server) Start() error {
	r := gin.New()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// middleware catching all requests for stats
	r.Use(s.statsMiddleware())

	r.GET(FIZZBUZZ_RESOURCE_ROUTE, s.FizzBuzzHandler)

	r.GET(STATS_RESOURCE_ROUTE, s.StatsHandler)

	r.Run(s.Conf.BaseUrl + ":" + s.Conf.Port)

	return nil
}
