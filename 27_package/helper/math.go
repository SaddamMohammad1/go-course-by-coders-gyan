// =====================================================
// math.go — CUSTOM PACKAGE
// =====================================================

// STEP 1: This file belongs to package "helper"
package helper

// STEP 2: Function names start with Capital letter → Exported
// This function can now be used in main.go
func Add(a int, b int) int {
	return a + b
}

// STEP 3: Another exported function
func Multiply(a int, b int) int {
	return a * b
}

/*
IMPORTANT NOTES:
----------------
✔ package helper → defines this as a library package
✔ Functions start with Capital → public (exported)
✔ Use inside main.go using:
      helper.Add()
      helper.Multiply()
*/
