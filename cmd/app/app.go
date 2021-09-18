package main

import (
	"fmt"
	"time"

	"gihub.com/momzor/fizzbuzz/pkg/webapi"
)

func main() {
	s := webapi.Server{}

	err := s.Start()
	if err != nil {

		fmt.Println("ERROR", err)
	}

	fmt.Println("run baby run....")

	for {
		time.Sleep(time.Nanosecond * 2000)
		fmt.Println("runings....")
	}

}
