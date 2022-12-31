package main

import "fmt"

// This can be optimized by injecting version at build time. Requires a server
// https://blog.alexellis.io/inject-build-time-vars-golang/
func LogVersion() string {
	return fmt.Sprintln("v0.10")
}
