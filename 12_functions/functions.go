package main

import "fmt"

func main() {

	/*
	   What is a Function?
	   --------------------
	   - Function is a block of reusable code.
	   - It performs a specific task.
	   - In Go:
	      * function starts with 'func'
	      * can take parameters
	      * can return values
	      * main() is the entry function
	*/

	fmt.Println("Welcome to Functions in Go")

	// Calling simple function
	sayHello()

	// Calling function with parameter
	printName("Saddam")

	// Calling function with two parameters + return value
	sum := addNumbers(10, 20)
	fmt.Println("Sum is:", sum)

	// Calling multi-return function
	result, status := divide(20, 5)
	fmt.Println("Divide Result:", result, "Status:", status)

	// Named return values example
	x := multiply(5, 6)
	fmt.Println("Multiply:", x)
}

/*
   Simple function (no parameter, no return)
*/
func sayHello() {
	fmt.Println("Hello from function!")
}

/*
   Function with 1 parameter
   -------------------------
   - Takes name as input
*/
func printName(name string) {
	fmt.Println("Your name is:", name)
}

/*
   Function with parameters + return type
   ---------------------------------------
   - Takes two integers
   - Returns integer result
*/
func addNumbers(a int, b int) int {
	return a + b
}

/*
   Function with multiple return values
   -------------------------------------
   - Go supports returning many values
*/
func divide(x int, y int) (int, string) {
	if y == 0 {
		return 0, "Cannot divide (y is zero)"
	}
	return x / y, "Success"
}

/*
   Named return function
   ----------------------
   - result is a named return variable
   - simply return → returns that value
*/
func multiply(a int, b int) (result int) {
	result = a * b
	return
}
