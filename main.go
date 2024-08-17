package main

import "fmt"

// Add takes two integers and returns their sum.
func Add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(Add(3, 5)) // Example usage, should output 8
}
