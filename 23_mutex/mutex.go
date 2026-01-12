package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

/*
WaitGroup
---------
- Used to wait until all goroutines finish.
- wg.Add(1) increases counter, wg.Done() decreases it.
- wg.Wait() blocks main until counter becomes 0.
*/
var wg sync.WaitGroup

/*
Mutex
-----
- Used to lock shared data when multiple goroutines try to write at same time.
- Prevents race conditions.
- Lock() → only one goroutine allowed
- Unlock() → others can continue
*/
var mut sync.Mutex

func main() {

	// List of websites to check
	websiteList := []string{
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://githube.com",
		"https://instagram.com",
	}

	// Loop creates goroutines to check each website
	for _, web := range websiteList {
		go getStatusCode(web) // goroutine → runs function asynchronously
		wg.Add(1)             // tell WaitGroup that 1 goroutine started
	}

	wg.Wait() // wait until all goroutines are done
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {

	defer wg.Done() // mark this goroutine as finished

	res, err := http.Get(endpoint) // make HTTP request

	if err != nil {
		fmt.Println("Oops in endpoint")
	} else {

		/*
			Lock shared slice before writing
			--------------------------------
			- Many goroutines will try to append at same time
			- Mutex ensures only one goroutine writes at a time
		*/
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d status code for  %s\n", res.StatusCode, endpoint)
	}
}
