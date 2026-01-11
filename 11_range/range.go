package main

import "fmt"

func main() {

	/*
	   What is range?
	   --------------
	   - range is a keyword used in for-loops.
	   - It helps to iterate over:
	     * slice
	     * array
	     * map
	     * string
	     * channel (advanced)
	   - It returns index + value (for slices/arrays)
	   - It returns key + value (for maps)
	*/

	// -----------------------------
	// 1. range with Slice
	// -----------------------------
	fruits := []string{"Apple", "Mango", "Banana"}

	fmt.Println("Using range on Slice:")
	for index, value := range fruits {
		fmt.Println("Index:", index, "Value:", value)
	}

	/*
	   Note:
	   If you want only value → use _ (blank identifier)
	   Example: for _, v := range fruits
	*/

	// -----------------------------
	// 2. range with Array
	// -----------------------------
	numbers := [3]int{10, 20, 30}

	fmt.Println("\nUsing range on Array:")
	for i, num := range numbers {
		fmt.Println("Index:", i, "Number:", num)
	}

	// -----------------------------
	// 3. range with Map
	// -----------------------------
	student := map[string]int{
		"Rahul": 21,
		"Amit":  25,
	}

	fmt.Println("\nUsing range on Map:")
	for key, value := range student {
		fmt.Println("Name:", key, "Age:", value)
	}

	// -----------------------------
	// 4. range with String
	// -----------------------------
	/*
	   - range returns Unicode code points (runes)
	*/

	text := "GoLang"

	fmt.Println("\nUsing range on String:")
	for index, char := range text {
		fmt.Println("Index:", index, "Character:", string(char))
	}

	// -----------------------------
	// Summary Notes
	// -----------------------------
	/*
	   - range is used in for loops to simplify iteration.
	   - For slice/array → gives index + value.
	   - For map → gives key + value.
	   - For string → gives index + rune.
	   - Very useful and clean syntax for looping.
	*/
}
