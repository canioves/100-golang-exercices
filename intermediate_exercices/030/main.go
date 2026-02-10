// Exercise: Functions

// Create a variadic function that sums many numbers passed as arguments
// Variadic functions can be called with any number of arguments.

package main

import "fmt"

func sum(nums ...int) int {
	var result int
	for _, x := range nums {
		result += x
	}

	return result
}

func main() {
	// Your code goes here
	result := sum(1, 2, 3, 4, 5)
	fmt.Print(result)
}
