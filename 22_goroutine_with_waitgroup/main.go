package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // pointer

func main() {
	websiteList := []string{
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://githube.com",
		"https://instagram.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
}

func getStatusCode(endpoint string) {

	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Oops in endpoint")
	} else {
		fmt.Printf("%d status code for  %s\n", res.StatusCode, endpoint)
	}
}
