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

func performUpdateRequest() {
	todo := Todo{
		UserID:    323,
		Title:     "Go lang update Testing",
		Completed: false,
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

	myUrl := "https://jsonplaceholder.typicode.com/todos/1"

	// creating PUT request
	req, err := http.NewRequest(http.MethodPut, myUrl, jsonReader)
	if err != nil {
		fmt.Println("Request err", err)
		return
	}
	req.Header.Set("Content-type", "application/json")

	// Send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Response error", err)
		return
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error in reading", err)
		return
	}
	fmt.Println("Response data", string(data))

}

func main() {
	fmt.Println("Update method in go lang")

	// Put request
	performUpdateRequest()
}
