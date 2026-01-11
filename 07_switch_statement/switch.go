package main

import "fmt"

func main() {

	day := 3

	/*
	   switch checks value of `day`
	   and runs the matching case
	*/
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	default:
		fmt.Println("Invalid day")
	}
}
