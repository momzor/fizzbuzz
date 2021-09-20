package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/momzor/fizzbuzz/pkg/db"
	"github.com/momzor/fizzbuzz/pkg/webapi"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}
func main() {

	c := webapi.Config{
		BaseUrl: os.Getenv("WEB_API_BASE_URL"),
		Port:    os.Getenv("WEB_API_PORT"),
	}
	dbctx := context.Background()
	dbC, err := db.NewClient(dbctx, db.Config{
		URI: os.Getenv("MONGO_URI"),
		DB:  os.Getenv("MONGO_DB"),
	})
	defer dbC.Disconnect(dbctx)
	if err != nil {
		log.Fatal(fmt.Sprintf("Can not get Configuration value from ENV : %s", err))
	}

	if err := dbC.Health(context.Background()); err != nil {
		log.Fatal(fmt.Sprintf("An error occured while instanciating db dependency : %s", err))
	}
	// todo v1.1: inject a logger, and monitoring agent
	s := webapi.Server{
		Conf:     c,
		DBClient: dbC,
	}

	err = s.Start()
	if err != nil {
		log.Fatal(fmt.Sprintf("An error occured while instanciating web Api: %s", err))
	}

}
