package main

import "fmt"

func main() {

	/*
	   What is Struct?
	   ----------------
	   - Struct is a USER-DEFINED data type.
	   - It groups multiple fields into a single type.
	   - Helps to represent REAL-WORLD objects:
	     Example:
	     - Employee (name, age, salary)
	     - Product (title, price, stock)
	     - Student (roll, class, marks)

	   Why Struct?
	   ------------
	   - Combine different data types in one group.
	   - Better data organization.
	   - Pass structured data to functions.
	   - Modify struct using pointers.
	*/

	// -------------------------------------
	// 1. Creating a Struct Value
	// -------------------------------------
	emp1 := Employee{
		Name:   "Saddam",
		Age:    25,
		Salary: 50000,
	}

	fmt.Println("Employee 1:", emp1)
	fmt.Println("Name:", emp1.Name)
	fmt.Println("Age:", emp1.Age)

	// -------------------------------------
	// 2. Updating Struct Field
	// -------------------------------------
	emp1.Age = 26
	fmt.Println("Updated Age:", emp1.Age)

	// -------------------------------------
	// 3. Using Struct with new() → returns pointer
	// -------------------------------------
	emp2 := new(Employee) // returns *Employee
	emp2.Name = "Aatiya"
	emp2.Age = 22
	emp2.Salary = 45000

	fmt.Println("\nEmployee 2 (Pointer Struct):", *emp2)

	// -------------------------------------
	// 4. Passing Struct to a Function
	// -------------------------------------
	printEmployee(emp1)

	// -------------------------------------
	// 5. Pointer with Struct → modify original struct
	// -------------------------------------
	updateSalary(&emp1, 65000) // passing address
	fmt.Println("\nAfter Pointer Update:", emp1)

	// -------------------------------------
	// 6. Pointer to struct example (Manual)
	// -------------------------------------
	empPointer := &Employee{Name: "Rahul", Age: 28, Salary: 70000}

	fmt.Println("\nPointer Struct Example:", empPointer)
	fmt.Println("Access actual data using *:", *empPointer)

	empPointer.Age = 30 // direct modification
	fmt.Println("Updated Age using pointer:", empPointer.Age)
}

/*
   Defining a Struct
   -----------------
*/
type Employee struct {
	Name   string
	Age    int
	Salary int
}

/*
   Function taking struct (copy value)
   -----------------------------------
   - Struct is passed by VALUE (copy)
*/
func printEmployee(e Employee) {
	fmt.Println("\n--- Employee Details ---")
	fmt.Println("Name:", e.Name)
	fmt.Println("Age:", e.Age)
	fmt.Println("Salary:", e.Salary)
}

/*
   Function using pointer to struct
   ---------------------------------
   - Pointer allows modifying original struct.
   - *Employee means function receives address.
*/
func updateSalary(emp *Employee, newSalary int) {
	emp.Salary = newSalary // updates original struct
}
