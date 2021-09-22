package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/momzor/fizzbuzz/pkg/db"
	"github.com/momzor/fizzbuzz/pkg/webapi"
)

// Build dependecies and injects them for server creation
// Create a web server, start and exposes it
func main() {
	c := webapi.Config{
		BaseUrl: os.Getenv("WEB_API_BASE_URL"),
		Port:    os.Getenv("WEB_API_PORT"),
		Env:     os.Getenv("APP_ENV"),
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
	am := webapi.NewAPIMiddleware(dbC)
	h := webapi.NewAPIHandler(dbC)
	s := webapi.Server{
		Conf:       c,
		DBClient:   dbC,
		Middleware: am,
		Handler:    h,
	}
	// todo v1.1: inject a logger, and monitoring agent in server
	s.InitServer()
	if err != nil {
		log.Fatal(fmt.Sprintf("An error occured during the web Api init : %s", err))
	}

	err = s.Router.Run(s.Conf.BaseUrl + ":" + s.Conf.Port)
	if err != nil {
		log.Fatal(fmt.Sprintf("Can not start server : %s", err))
	}

}
