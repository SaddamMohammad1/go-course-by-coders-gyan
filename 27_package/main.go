// =====================================================
// main.go — MAIN PACKAGE (ENTRY POINT)
// =====================================================

// STEP 1: package main → required for executable program
package main

// STEP 2: Import built-in & custom packages
import (
	"fmt" // Standard package

	"github.com/SaddamMohammad1/packages/helper" // Custom package (folder name: helper)
)

// STEP 3: main() → program starts here
func main() {

	fmt.Println("=== Go Packages Example ===")

	// Calling function from helper package
	result := helper.Add(10, 20)
	fmt.Println("Add Result:", result)

	multi := helper.Multiply(5, 5)
	fmt.Println("Multiply Result:", multi)
}
