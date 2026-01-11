package main

import "fmt"

/*
   - Array is FIXED size
   - Same type values only
   - Index starts from 0
   - len(array) gives size
   - [...] lets compiler set size
   - 2D array = array of arrays
*/
func main() {

	/*
	   ==========================
	   1. BASIC ARRAY IN GO
	   ==========================
	   Array → Fixed-size collection of same data type.
	   Once size is set, it cannot change.
	*/
	var fruitList [4]string

	// Assign values using index
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[2] = "Mango"
	fruitList[3] = "Peach"

	fmt.Println("Fruit List:", fruitList)        // Print full array
	fmt.Println("First Fruit:", fruitList[0])    // Access single
	fmt.Println("Array Length:", len(fruitList)) // Size of array

	/*
	   =====================================
	   2. ARRAY WITH INITIALIZATION
	   =====================================
	   Size is fixed, values assigned directly.
	*/
	numbers := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Numbers:", numbers)

	/*
	   =====================================
	   3. ARRAY WITHOUT SPECIFYING SIZE
	   =====================================
	   `...` → Compiler decides size automatically.
	*/
	colors := [...]string{"Red", "Green", "Blue"}
	fmt.Println("Colors:", colors)

	/*
	   =====================================
	   4. LOOP THROUGH ARRAY
	   =====================================
	   `range` gives index + value
	*/
	fmt.Println("\nLooping through colors array:")
	for i, v := range colors {
		fmt.Println(i, v)
	}

	/*
	   ============================
	   5. 2D ARRAY (MATRIX)
	   ============================
	   Declaration:
	   [rows][columns]type
	*/
	var matrix [3][3]int

	// Assign values
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3

	matrix[1][0] = 4
	matrix[1][1] = 5
	matrix[1][2] = 6

	matrix[2][0] = 7
	matrix[2][1] = 8
	matrix[2][2] = 9

	fmt.Println("\nMatrix (3x3):")
	fmt.Println(matrix)

	/*
	   =====================================
	   6. 2D ARRAY WITH INITIAL VALUES
	   =====================================
	*/
	matrix2 := [2][3]int{
		{1, 2, 3}, // Row 1
		{4, 5, 6}, // Row 2
	}

	fmt.Println("\nMatrix2 (2x3):")

	fmt.Println(matrix2)

	/*
	   =====================================
	   7. LOOP THROUGH 2D ARRAY
	   =====================================
	   Nested loop → iterate rows & columns
	*/
	fmt.Println("\nLooping Matrix2:")
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(matrix2[i][j], " ")
		}
		fmt.Println()
	}
}
