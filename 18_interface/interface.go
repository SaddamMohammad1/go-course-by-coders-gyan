package main

import "fmt"

// Interface: defines behavior (what function must exist)
type Shape interface {
	Area() float64
}

// Rectangle struct implementing Shape interface
type Rectangle struct {
	Width  float64
	Height float64
}

// Area method (HOW Rectangle calculates area)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle struct implementing Shape interface
type Circle struct {
	Radius float64
}

// Area method (HOW Circle calculates area)
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Function accepting interface (works for ANY type with Area())
func printArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

func main() {

	// Rectangle implements Shape
	r := Rectangle{Width: 10, Height: 5}
	printArea(r)

	// Circle implements Shape
	c := Circle{Radius: 7}
	printArea(c)
}
