package main

import (
	"fmt"
	"time"
)

// A goroutine is a lightweight thread used to run functions at the same time (concurrently) in Go.
// goroutine like cron/schedule

// Function to run in goroutine
func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {

	// Running function normally
	fmt.Println("Running normally:")
	printNumbers()

	// Running function as goroutine
	fmt.Println("\nRunning as Goroutine:")
	go printNumbers() // runs in background

	// Prevent main from exiting too early
	time.Sleep(2 * time.Second)
}
