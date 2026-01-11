package main

import "fmt"

func main() {

	/*
	   What is Closure?
	   -----------------
	   - A closure is a function that remembers the variables
	     from the outer function, even after the outer function has finished.
	   - The inner function can access and modify outer variables.
	   - Closure = Function + its surrounding environment.
	*/

	// Getting a closure function
	counter := makeCounter()

	// Using closure multiple times
	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3

	// Creating another independent closure
	newCounter := makeCounter()
	fmt.Println(newCounter()) // 1
	fmt.Println(newCounter()) // 2

	/*
	   Closures are useful in:
	   - Counters
	   - Generators
	   - Maintaining state
	   - Avoiding global variables
	*/
}

/*
   Closure Example
   ----------------
   - makeCounter() returns an inner function
   - inner function remembers 'count'
   - count is NOT reset on each call
*/
func makeCounter() func() int {
	count := 0 // Outer variable → remembered by closure

	return func() int {
		count++
		return count
	}
}
