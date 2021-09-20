package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"gihub.com/momzor/fizzbuzz/pkg/db"
	"gihub.com/momzor/fizzbuzz/pkg/webapi"
)

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
