package main

import (
	"fmt"
	"time"

	"github.com/morelmiles/ugx_rates/client"
)

func main() {
	done := make(chan bool)
	go keepRunning()
	<-done

	client.Config()
}

func keepRunning() {
	for {
		fmt.Printf("%v+\n", time.Now())
		time.Sleep(time.Hour.Round(5))
	}
}
