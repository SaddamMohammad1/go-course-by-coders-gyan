package main

import "fmt"

func main() {
	age := 16

	if age >= 18 {
		fmt.Println("Person is an adult")
	} else if age >= 12 {
		fmt.Println("Person is teenager")
	} else {
		fmt.Println("Person is a kid")
	}

	// Logical operator in if else
	var role = "admin"
	var hasPermission = false

	if role == "admin" && hasPermission {
		fmt.Println("Valid access")
	}

	/*
		Note:
			Go does not have ternary operator, you will have to use normal if else.
	*/
}
