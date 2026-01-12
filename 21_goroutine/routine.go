package main

import (
	"fmt"
	"time"
)

// A goroutine is a lightweight thread used to run functions at the same time (concurrently) in Go.
// goroutine like cron/schedule

func main() {
	go greeter("Hello")
	greeter("World")
}

// Function to run in goroutine
func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Microsecond)
		fmt.Println(s)
	}
}
