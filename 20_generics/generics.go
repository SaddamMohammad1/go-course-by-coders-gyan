package main

import "fmt"

//
// ===============================
//     GENERICS IN GO – ALL TYPES
// ===============================
//

// 1. Basic Generic Function (works with int + float)
func Add[T int | float64](a T, b T) T {
	return a + b
}

// 2. Generic Function that accepts ANY type
func PrintValue[T any](value T) {
	fmt.Println("Value:", value)
}

// 3. Generic Function for slices
func FirstElement[T any](items []T) T {
	return items[0]
}

// 4. Generic Map Function (key + value types)
func GetValue[K comparable, V any](m map[K]V, key K) V {
	return m[key]
}

// 5. Generic Struct
type Box[T any] struct {
	Value T
}

// 6. Generic Struct with multiple type parameters
type Pair[A any, B any] struct {
	First  A
	Second B
}

// 7. Generic with custom constraints
type Number interface {
	int | float64
}

func Multiply[T Number](a T, b T) T {
	return a * b
}

func main() {

	// ---------------------------------------
	// 1. Basic Generic Function
	// ---------------------------------------
	fmt.Println("Add int:", Add(10, 20))
	fmt.Println("Add float:", Add(3.5, 2.5))

	// ---------------------------------------
	// 2. ANY Type Function
	// ---------------------------------------
	PrintValue("Hello Go")
	PrintValue(101)
	PrintValue(true)

	// ---------------------------------------
	// 3. Generic Slice Function
	// ---------------------------------------
	fmt.Println("First element (int):", FirstElement([]int{10, 20, 30}))
	fmt.Println("First element (string):", FirstElement([]string{"Go", "Lang"}))

	// ---------------------------------------
	// 4. Generic Map Function
	// ---------------------------------------
	users := map[int]string{1: "Saddam", 2: "Aatiya"}
	fmt.Println("Map value:", GetValue(users, 2))

	// ---------------------------------------
	// 5. Generic Struct
	// ---------------------------------------
	intBox := Box[int]{Value: 100}
	stringBox := Box[string]{Value: "Generics"}

	fmt.Println("Box int:", intBox.Value)
	fmt.Println("Box string:", stringBox.Value)

	// ---------------------------------------
	// 6. Struct with 2 type parameters
	// ---------------------------------------
	p := Pair[string, int]{First: "Age", Second: 25}
	fmt.Println("Pair:", p.First, p.Second)

	// ---------------------------------------
	// 7. Constraint-based Generic Function
	// ---------------------------------------
	fmt.Println("Multiply int:", Multiply(4, 5))
	fmt.Println("Multiply float:", Multiply(2.5, 3.0))
}
