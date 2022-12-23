package main

import (
	"time"

	"github.com/morelmiles/ugx_rates/client"
)

func main() {
	done := make(chan bool)

	go keepRunning()

	<-done

}

func keepRunning() {
	for {
		client.Config()

		time.Sleep(time.Duration(time.Now().Day()))
	}
}
