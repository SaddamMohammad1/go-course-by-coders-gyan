package main

import "fmt"

func main() {

	/*
	   What is Map?
	   ------------
	   - Map is a built-in data structure in Go.
	   - It stores data in KEY → VALUE format (like a dictionary).
	   - Key must be unique.
	   - Fast to read, write, and search.
	   - Syntax: map[keyType]valueType
	*/

	// Creating a map using map literal
	student := map[string]int{
		"Rahul": 21,
		"Amit":  22,
		"Neha":  20,
	}

	fmt.Println("Students Map:", student)

	/*
	   Accessing map values
	   ---------------------
	   - Use key to get value.
	*/

	fmt.Println("Rahul Age:", student["Rahul"])

	/*
	   Adding new key-value pair
	   --------------------------
	*/

	student["Sana"] = 23
	fmt.Println("After Adding Sana:", student)

	/*
	   Updating a value
	   -----------------
	*/

	student["Amit"] = 25
	fmt.Println("After Updating Amit:", student)

	/*
	   Checking if key exists
	   -----------------------
	   - value, exists := map[key]
	*/

	age, ok := student["Rohan"]                   // Rohan is not present
	fmt.Println("Rohan Age:", age, "Exists?", ok) // 0 false

	/*
	   Deleting a key
	   --------------
	*/

	delete(student, "Neha")
	fmt.Println("After Deleting Neha:", student)

	/*
	   Looping through map
	   --------------------
	*/

	fmt.Println("\nLooping Map:")
	for name, age := range student {
		fmt.Println(name, "→", age)
	}

	/*
	   Creating map using make()
	   --------------------------
	   - make() is used when we want to create an empty map first.
	   - This avoids runtime panic if map is nil.
	*/

	country := make(map[string]string)
	country["India"] = "Delhi"
	country["USA"] = "Washington"
	country["UK"] = "London"

	fmt.Println("\nCountries Map:", country)

	/*
	   Deleting from country map
	*/

	delete(country, "UK")
	fmt.Println("After Deleting UK:", country)

	/*
	   Important Notes
	   ---------------
	   - Maps are reference types (like slices).
	   - Cannot take address of map element.
	   - Reading non-existing key returns default zero value.
	   - Order is NOT guaranteed in iteration.
	*/
}
