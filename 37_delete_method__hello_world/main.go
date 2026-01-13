package main

import (
	"fmt"
	"net/http"
)

func performDeleteRequest() {
	const myUrl = "https://jsonplaceholder.typicode.com/todos/1"

	// creating Delete request
	req, err := http.NewRequest(http.MethodDelete, myUrl, nil)
	if err != nil {
		fmt.Println("Request err", err)
		return
	}

	// Send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Response error", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("Response data", res.Status)
}

func main() {
	fmt.Println("Delete method in go lang")

	// perform delete request
	performDeleteRequest()
}
