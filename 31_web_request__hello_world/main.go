package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Web request in go lang")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting GET response", err)
		return
	}
	defer res.Body.Close()
	fmt.Printf("Types of response %T\n", res) // Types of response *http.Response

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response", err)
		return
	}
	fmt.Println("Response", string(data))
}
