package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	fmt.Println("We are learning json")

	person := Person{Name: "Test", Age: 28, IsAdult: true}
	fmt.Println("Person data is", person) // Person data is {Test 28 true}

	// convert person into JSON encoding (Marshalling)
	jsonData, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}
	fmt.Println("Now Person data is", string(jsonData)) // Now Person data is {"name":"Test","age":28,"is_adult":true}

	// Decodeing (Unmarshalling)
	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		panic(err)
	}
	fmt.Println("Now after decoding person data is", personData) // Now after decoding person data is {Test 28 true}
}
