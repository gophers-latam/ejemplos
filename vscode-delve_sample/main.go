package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting program...")
	result := sum(5, 3)
	fmt.Printf("Result: %d\n", result)
}

func sum(a, b int) int {
	time.Sleep(2 * time.Second) // Simulate some work
	sum := a + b
	fmt.Printf("Sum of %d and %d is %d\n", a, b, sum)
	return sum
}
