package main

import "fmt"

func main() {

	/*
	   Slice Definition
	   ----------------
	   - Slice is a dynamic, flexible view of an underlying array.
	   - It stores: pointer → length → capacity.
	   - Slices are used more than arrays because they grow using append().
	*/

	// Basic slice
	fruits := []string{"Apple", "Mango", "Banana"}
	fmt.Println("Fruits:", fruits)
	fmt.Println("Len:", len(fruits)) // 3
	fmt.Println("Cap:", cap(fruits)) // 3

	/*
	   Capacity (cap) Example with append()
	   ------------------------------------
	   - cap() = total available space before Go reallocates memory.
	   - When capacity is full, append() grows the underlying array automatically.
	*/

	fmt.Println("\nCapacity Growth Example using append():")

	numbers := []int{} // empty slice

	for i := 1; i <= 10; i++ {
		numbers = append(numbers, i)
		fmt.Println("Added:", i, "Length:", len(numbers), "Capacity:", cap(numbers))
	}

	/*
	   make() Function Example
	   ------------------------
	   make([]type, length, capacity)
	   - length: current number of elements
	   - capacity: total available space
	   - Useful when we know required size so Go can pre-allocate memory.
	*/

	fmt.Println("\nUsing make() to create a slice:")

	data := make([]int, 3, 5) // len = 3, cap = 5
	fmt.Println("Slice created using make():", data)
	fmt.Println("Length:", len(data))   // 3
	fmt.Println("Capacity:", cap(data)) // 5

	/*
	   make() + append example
	   ------------------------
	   - When we append within capacity → same underlying array.
	   - When capacity is full → new array is created.
	*/

	fmt.Println("\nAppending into make() slice:")

	for i := 1; i <= 4; i++ {
		data = append(data, i)
		fmt.Println("After append:", data, "Len:", len(data), "Cap:", cap(data))
	}

	/*
	   Slice from Array
	   -----------------
	   - Slicing an array creates a slice view.
	*/

	arr := [5]int{10, 20, 30, 40, 50}
	sliceArr := arr[1:4] // 20, 30, 40
	fmt.Println("\nSlice from Array:", sliceArr)

	/*
	   Shared Memory Example
	   ---------------------
	   - Slices share underlying array memory.
	   - Updating slice updates array.
	*/

	original := [3]int{1, 2, 3}
	ref := original[:] // full slice

	ref[1] = 99

	fmt.Println("\nOriginal Array:", original)
	fmt.Println("Slice View:", ref)

	/*
	   Summary Notes
	   -------------
	   - Slice = reference type (lightweight).
	   - len() → number of used elements.
	   - cap() → total free + used memory space.
	   - append() grows slice.
	   - make() defines initial length + capacity.
	   - Slice is used everywhere in Go instead of arrays.
	*/
}
