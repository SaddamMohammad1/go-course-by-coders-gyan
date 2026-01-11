package main

import "fmt"

func main() {

	/*
	   What is Variadic Function?
	   ---------------------------
	   - A variadic function can take ANY number of arguments.
	   - Syntax: func functionName(values ...type)
	   - Inside the function, values become a SLICE.
	*/

	// Calling variadic function with multiple arguments
	total := sum(10, 20, 30, 40)
	fmt.Println("Total:", total)

	// You can pass 1 value or 10 values
	fmt.Println("Sum of 5 values:", sum(1, 2, 3, 4, 5))

	// Passing a slice into variadic function
	nums := []int{5, 5, 5}
	fmt.Println("Sum of slice:", sum(nums...)) // ... expands slice

	// Variadic with strings
	printNames("Rahul", "Amit", "Sana")

	/*
	   Real-world Example:
	   -------------------
	   Using Variadic Function + Interface to log ANY type of value
	*/
	logInfo("User logged", 101, true, 99.5, "Success")

	// Variadic + any (can accept any type)
	logInfoWithAnyKeyword("User Logged In", 101, true, 99.9, "Success", []int{1, 2, 3})
}

/*
   Variadic Function Example
   -------------------------
   - 'nums ...int' means nums is a slice of integers
   - You can pass unlimited number of integers
*/
func sum(nums ...int) int {
	total := 0
	for _, value := range nums {
		total += value
	}
	return total
}

/*
   Variadic Function with string parameters
   ----------------------------------------
*/
func printNames(names ...string) {
	fmt.Println("Names List:")
	for _, name := range names {
		fmt.Println(name)
	}
}

/*
   REAL-WORLD VARIADIC + INTERFACE EXAMPLE
   ----------------------------------------
   - interface{} (any) accepts any type: int, string, bool, float
   - Useful for flexible logging, formatting, debugging
*/
func logInfo(values ...interface{}) {
	fmt.Println("\nLog Output:")
	for _, v := range values {
		fmt.Println("Value:", v, "| Type:", fmt.Sprintf("%T", v))
	}
}

/*
   Variadic + any
   --------------
   - any is alias for interface{}
   - Accepts any type: string, int, float, bool, struct, slice...
   - Useful for logging, debugging, flexible API inputs
*/
func logInfoWithAnyKeyword(values ...any) {
	fmt.Println("\nLog Output:")
	for _, v := range values {
		fmt.Println("Value:", v, "| Type:", fmt.Sprintf("%T", v))
	}
}
