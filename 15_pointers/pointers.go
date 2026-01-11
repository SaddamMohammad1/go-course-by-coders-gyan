package main

import "fmt"

func main() {

	/*
	   What is Pointer?
	   -----------------
	   - Pointer stores the MEMORY ADDRESS of a variable.
	   - Instead of storing the value, it stores *where* the value is located.
	   - '&'  → gives address of variable
	   - '*'  → dereference (get value stored at that address)
	   - Useful for:
	     * Passing data without copying
	     * Updating values inside functions
	     * Efficient memory usage
	*/

	// Normal variable
	num := 10
	fmt.Println("Value of num:", num)
	fmt.Println("Address of num:", &num) // memory address

	// Pointer variable: stores address of num
	var ptr *int = &num
	fmt.Println("\nPointer ptr:", ptr)     // address
	fmt.Println("Value using *ptr:", *ptr) // dereference → value 10

	/*
	   Changing value using pointer
	   ----------------------------
	   - If we modify *ptr, it updates the original variable
	*/

	*ptr = 50
	fmt.Println("\nUpdated value of num using pointer:", num) // now 50

	/*
	   Pointer with function
	   ---------------------
	   - Passing pointer to modify original variable
	*/

	value := 5
	fmt.Println("\nBefore function call:", value)

	increaseValue(&value) // passing address
	fmt.Println("After function call:", value)

	/*
	   Nil Pointer
	   -----------
	   - Pointer with no address
	   - Default value of pointer = nil
	*/

	var emptyPtr *int
	fmt.Println("\nNil Pointer:", emptyPtr)
}

// Function using pointer to update value
func increaseValue(num *int) {
	/*
	   num *int → pointer parameter
	   *num     → access/modify actual value
	*/
	*num = *num + 10
}
