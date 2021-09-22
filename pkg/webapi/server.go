package webapi

import (
	"github.com/gin-gonic/gin"
	"github.com/momzor/fizzbuzz/pkg/db"
	_ "github.com/momzor/fizzbuzz/pkg/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// Server representing a web server and its dependeincies
type Server struct {
	Conf       Config
	DBClient   db.Client
	Router     *gin.Engine
	Middleware APIMiddleware
	Handler    APIHandler
}

// Config Configuration for web server
type Config struct {
	Env     string
	BaseUrl string
	Port    string
}

// ErrorResponse
type ErrorResponse struct {
	Status  int
	Errors  []string `json:"errors,omitempty"`
	Message string   `json:"message"`
}

// @title LeBonCoin test fizzbuzz
// @version 1.0
// @description LeBonCoin test fizzbuzz
// @contact.name Momzor
// @contact.email m.benaida.pro@gmail.com
// @license.url https://www.gnu.org/licenses/quick-guide-gplv3.html
func (s *Server) InitServer() {

	gin.SetMode(s.Conf.Env)
	r := gin.New()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// middleware catching all requests for stats
	r.Use(s.Middleware.StatsMiddleware())
	//handlers
	r.GET(FIZZBUZZ_RESOURCE_ROUTE, FizzBuzzHandler)
	r.GET(STATS_RESOURCE_ROUTE, s.Handler.StatsHandler())
	s.Router = r

}
