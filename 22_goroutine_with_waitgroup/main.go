package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // tells WaitGroup this goroutine is finished

	fmt.Println("Worker", id, "started")
	time.Sleep(1 * time.Second)
	fmt.Println("Worker", id, "finished")
}

func main() {

	var wg sync.WaitGroup // WaitGroup variable

	// Creating 3 goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)         // Add 1 goroutine to wait for
		go worker(i, &wg) // start goroutine
	}

	wg.Wait() // wait until all goroutines call Done()
	fmt.Println("All workers completed.")
}
