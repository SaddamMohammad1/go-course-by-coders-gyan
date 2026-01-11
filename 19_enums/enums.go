package main

import "fmt"

// Enum type
type OrderStatus int

// Enum values using iota
const (
	Pending    OrderStatus = iota // 0
	Processing                    // 1
	Shipped                       // 2
	Delivered                     // 3
	Cancelled                     // 4
)

// Convert enum to readable text
func (s OrderStatus) GetStatus() string {
	switch s {
	case Pending:
		return "Pending"
	case Processing:
		return "Processing"
	case Shipped:
		return "Shipped"
	case Delivered:
		return "Delivered"
	case Cancelled:
		return "Cancelled"
	default:
		return "Unknown Status"
	}
}

// Function using enum
func updateOrderStatus(status OrderStatus) {
	fmt.Println(status)                                                  // 2 or 3
	fmt.Println("Order Status:", status.GetStatus(), "| Value:", status) // Order Status: Shipped | Value: 2
}

func main() {

	// Set order status (Shipped = 2)
	status := Shipped
	updateOrderStatus(status)

	// Change order status (Delivered = 3)
	status = Delivered
	updateOrderStatus(status)
}
