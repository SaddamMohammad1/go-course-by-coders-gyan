package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("CRUD in go lang")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Getting error", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in getting response", res.Status)
		return
	}

	// // 1st way to print data only without storing in variable and without using struct
	// data, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error reading", err)
	// }
	// fmt.Println("Reponse data is", string(data))

	// 2nd way print data using struct and store in variable
	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding", err)
		return
	}
	fmt.Println("Todo data", todo) // Todo data {1 1 delectus aut autem false}
}
