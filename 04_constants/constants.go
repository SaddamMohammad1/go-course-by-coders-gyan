package main

import "fmt"

const age = 30

func main() {
	const name string = "Constant name"

	// Constant grouping
	const (
		port = 5000
		host = "local"
	)

	fmt.Println(port, host)

}
