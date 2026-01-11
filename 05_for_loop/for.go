package main

import "fmt"

// for loop -> Only construct in go for looping, in go while loop not exist.
func main() {
	// Achieve while looping style
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// Classic for loop
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	// range (Comes in go 1.22 version)
	for i := range 3 {
		fmt.Println(i)
	}
}
