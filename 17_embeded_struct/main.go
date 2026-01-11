package main

import "fmt"

// Base struct
type Person struct {
	Name string
	Age  int
}

// Embedded struct
type Employee struct {
	Person  // Embedded (promoted fields)
	Company string
	Role    string
}

// Another base struct
type Vehicle struct {
	Brand string
	Year  int
}

// Car embeds Vehicle
type Car struct {
	Vehicle
	Type string
}

func main() {

	// Embedding example
	emp := Employee{
		Person: Person{
			Name: "Saddam",
			Age:  25,
		},
		Company: "Google",
		Role:    "Developer",
	}

	fmt.Println(emp.Name) // promoted field
	fmt.Println(emp.Company)

	// Embedded struct with methods
	car := Car{
		Vehicle: Vehicle{
			Brand: "Honda",
			Year:  2020,
		},
		Type: "Sedan",
	}

	fmt.Println(car.Brand) // promoted field
	car.showDetails()
}

// Method from embedded struct
func (v Vehicle) showDetails() {
	fmt.Println("Vehicle:", v.Brand, "| Year:", v.Year)
}
