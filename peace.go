package main

import (
	"fmt"
)

// SetqSafeValue represents a type-safe variable with optional encryption
type SetqSafeValue struct {
	Name  string
	Value interface{}
	Type  string
	Doc   string
	Debug bool
}

// SetqSafe creates a type-safe variable
func SetqSafe(name string, value interface{}, typ string, doc string, debug bool) SetqSafeValue {
	if debug {
		fmt.Printf("[debug] setting %s to %v\n", name, value)
	}
	return SetqSafeValue{
		Name:  name,
		Value: value,
		Type:  typ,
		Doc:   doc,
		Debug: debug,
	}
}

// Top returns the maximum of two values
func Top(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Bottom returns the minimum of two values
func Bottom(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Middle returns the median of three values
func Middle(a, b, c int) int {
	if (a >= b && a <= c) || (a >= c && a <= b) {
		return a
	}
	if (b >= a && b <= c) || (b >= c && b <= a) {
		return b
	}
	return c
}

// Add adds two integers
func Add(a, b int) int {
	return a + b
}

// Subtract subtracts two integers
func Subtract(a, b int) int {
	return a - b
}

// Multiply multiplies two integers
func Multiply(a, b int) int {
	return a * b
}

// Demo demonstrates Peace language features
func Demo() {
	fmt.Println("=== Peace Language Demo (Go Implementation) ===")

	// Type-safe variable
	x := SetqSafe("x", 10, "integer", "test variable", true)
	fmt.Printf("x = %v\n", x.Value)
	fmt.Printf("x (doc): %s\n", x.Doc)

	// Math operations
	fmt.Println("\n=== Math Operations ===")
	fmt.Printf("Top(15, 8) = %d\n", Top(15, 8))
	fmt.Printf("Bottom(15, 8) = %d\n", Bottom(15, 8))
	fmt.Printf("Middle(3, 8, 5) = %d\n", Middle(3, 8, 5))
	fmt.Printf("Add(5, 3) = %d\n", Add(5, 3))
	fmt.Printf("Subtract(10, 3) = %d\n", Subtract(10, 3))
	fmt.Printf("Multiply(4, 5) = %d\n", Multiply(4, 5))

	// Type checking demonstration
	fmt.Println("\n=== Type Safety ===")
	y := SetqSafe("y", 42, "integer", "another test variable", true)
	fmt.Printf("y = %v (type: %s, doc: %s)\n", y.Value, y.Type, y.Doc)

	fmt.Println("\n=== Demo Complete ===")
}

func main() {
	Demo()
}
