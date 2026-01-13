package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performPostRequest() {
	todo := Todo{
		UserID:    23,
		Title:     "Testing",
		Completed: true,
	}

	// Convert todo struct to JSON, for post request
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error in marshalling", err)
		return
	}

	// convert json data to string
	jsonString := string(jsonData)

	// Convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	myUrl := "https://jsonplaceholder.typicode.com/todos"
	// Send post request
	// http.Post() sends a POST request to the given URL.
	// Syntax: http.Post(URL, Content-Type, Body)
	//
	// myUrl        → The API endpoint where data will be sent
	// "application/json" → Tells server that request body contains JSON
	// jsonReader   → io.Reader containing JSON data (e.g., bytes.NewBuffer(jsonData))
	res, err := http.Post(myUrl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error sending request", err)
		return
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error in response", err)
	}
	fmt.Println("Response data is", string(data))
}

func main() {
	fmt.Println("Post request in go lang")

	// Call performPostRequest
	performPostRequest()
}
