package main

import (
	"fmt"
	"os"

	"gihub.com/momzor/fizzbuzz/pkg/webapi"
)

func main() {

	c := webapi.Config{
		BaseUrl: os.Getenv("WEB_API_BASE_URL"),
		Port:    os.Getenv("WEB_API_PORT"),
	}
	s := webapi.Server{
		Conf: c,
	}

	err := s.Start()
	if err != nil {
		fmt.Println("ERROR", err)
	}

}
