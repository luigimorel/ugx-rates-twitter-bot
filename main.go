package main

import (
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	runCronJob()
}

func runCronJob() {
	c := gocron.NewScheduler(time.UTC)
	c.At("13:00:00").Do(func() {
		Config()
	})
}
